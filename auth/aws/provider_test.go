package awsauth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/aws/aws-sdk-go-v2/aws"
)

const testOrgUUID = "00000000-0000-0000-0000-000000000000"

func TestAuthenticateWithStaticCredentials(t *testing.T) {
	tests := []struct {
		name         string
		sessionToken string
		wantToken    bool
	}{
		{name: "long-lived credentials"},
		{name: "temporary credentials", sessionToken: "session-token", wantToken: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var proof string
			server := newDelegatedTokenServer(t, func(request *http.Request) {
				proof = strings.TrimPrefix(request.Header.Get("Authorization"), "Delegated ")
			})
			defer server.Close()

			provider, err := New(
				WithRegion("us-east-2"),
				WithStaticCredentials("test-access-key", "test-secret-key", test.sessionToken),
			)
			if err != nil {
				t.Fatalf("creating provider: %v", err)
			}

			credentials, err := provider.Authenticate(delegatedTokenContext(server.URL), delegatedTokenConfig())
			if err != nil {
				t.Fatalf("authenticating: %v", err)
			}
			if credentials.DelegatedToken != "delegated-token-1" {
				t.Fatalf("delegated token = %q, want delegated-token-1", credentials.DelegatedToken)
			}

			headers, endpoint := decodeProof(t, proof)
			if got := firstHeader(headers, orgIDHeader); got != testOrgUUID {
				t.Errorf("organization header = %q, want %q", got, testOrgUUID)
			}
			authorization := firstHeader(headers, "Authorization")
			if !strings.Contains(authorization, "Credential=test-access-key/") {
				t.Errorf("AWS authorization header does not contain the access key: %q", authorization)
			}
			if !strings.Contains(authorization, "/us-east-2/sts/aws4_request") {
				t.Errorf("AWS authorization header does not contain the signing region: %q", authorization)
			}
			for _, signedHeader := range []string{"content-length", "content-type", "host", "x-amz-date", "x-ddog-org-id"} {
				if !strings.Contains(strings.ToLower(authorization), signedHeader) {
					t.Errorf("AWS authorization header does not sign %q: %q", signedHeader, authorization)
				}
			}
			if got := firstHeader(headers, "X-Amz-Security-Token"); (got != "") != test.wantToken {
				t.Errorf("security token present = %t, want %t", got != "", test.wantToken)
			}
			if got := endpoint; got != "https://sts.us-east-2.amazonaws.com" {
				t.Errorf("STS endpoint = %q, want regional endpoint", got)
			}
		})
	}
}

func TestAuthenticateUsesSharedConfigurationProfile(t *testing.T) {
	directory := t.TempDir()
	credentialsFile := filepath.Join(directory, "credentials")
	configFile := filepath.Join(directory, "config")
	if err := os.WriteFile(credentialsFile, []byte("[cavm]\naws_access_key_id = profile-access-key\naws_secret_access_key = profile-secret-key\n"), 0o600); err != nil {
		t.Fatalf("writing credentials file: %v", err)
	}
	if err := os.WriteFile(configFile, []byte("[profile cavm]\nregion = us-west-1\n"), 0o600); err != nil {
		t.Fatalf("writing config file: %v", err)
	}
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", credentialsFile)
	t.Setenv("AWS_CONFIG_FILE", configFile)
	t.Setenv("AWS_PROFILE", "cavm")
	t.Setenv("AWS_ACCESS_KEY_ID", "")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "")
	t.Setenv("AWS_SESSION_TOKEN", "")
	t.Setenv("AWS_EC2_METADATA_DISABLED", "true")

	var proof string
	server := newDelegatedTokenServer(t, func(request *http.Request) {
		proof = strings.TrimPrefix(request.Header.Get("Authorization"), "Delegated ")
	})
	defer server.Close()

	provider, err := New()
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	if _, err := provider.Authenticate(delegatedTokenContext(server.URL), delegatedTokenConfig()); err != nil {
		t.Fatalf("authenticating from shared profile: %v", err)
	}

	headers, endpoint := decodeProof(t, proof)
	if !strings.Contains(firstHeader(headers, "Authorization"), "Credential=profile-access-key/") {
		t.Errorf("AWS authorization header does not use profile credentials: %q", firstHeader(headers, "Authorization"))
	}
	if endpoint != "https://sts.us-west-1.amazonaws.com" {
		t.Errorf("STS endpoint = %q, want profile region endpoint", endpoint)
	}
}

func TestAuthenticateAgainstDatadogIntegration(t *testing.T) {
	apiURL := os.Getenv("DD_TEST_WIF_API_URL")
	orgUUID := os.Getenv("DD_ORG_UUID")
	if apiURL == "" || orgUUID == "" {
		t.Skip("set DD_TEST_WIF_API_URL, DD_ORG_UUID, and AWS_PROFILE to test a live Datadog token exchange")
	}

	provider, err := New()
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	credentials, err := provider.Authenticate(delegatedTokenContext(apiURL), &datadog.DelegatedTokenConfig{
		OrgUUID:  orgUUID,
		Provider: datadog.ProviderAWS,
	})
	if err != nil {
		t.Fatalf("authenticating against Datadog: %v", err)
	}
	if credentials.DelegatedToken == "" {
		t.Fatal("Datadog returned an empty delegated token")
	}
	if credentials.OrgUUID != orgUUID {
		t.Fatalf("delegated token org UUID = %q, want %q", credentials.OrgUUID, orgUUID)
	}
	if !credentials.Expiration.After(time.Now()) {
		t.Fatalf("delegated token expiration = %s, want a future time", credentials.Expiration)
	}
}

func TestWithStaticCredentialsRejectsPartialCredentials(t *testing.T) {
	if _, err := New(WithStaticCredentials("access-key", "", "")); err == nil {
		t.Fatal("expected partial static credentials to be rejected")
	}
}

func TestGenerateProofResolvesAWSPartitions(t *testing.T) {
	provider, err := New()
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	proof, err := provider.generateProof(context.Background(), testOrgUUID, aws.Config{Region: "cn-north-1"}, aws.Credentials{
		AccessKeyID:     "access-key",
		SecretAccessKey: "secret-key",
	})
	if err != nil {
		t.Fatalf("generating proof: %v", err)
	}
	_, endpoint := decodeProof(t, proof)
	if endpoint != "https://sts.cn-north-1.amazonaws.com.cn" {
		t.Fatalf("STS endpoint = %q, want China partition endpoint", endpoint)
	}
}

func newDelegatedTokenServer(t *testing.T, inspect func(*http.Request)) *httptest.Server {
	t.Helper()
	var requests atomic.Int32
	return httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			t.Errorf("request method = %s, want POST", request.Method)
		}
		if !strings.HasPrefix(request.Header.Get("Authorization"), "Delegated ") {
			t.Errorf("authorization header does not contain a delegated proof")
		}
		if inspect != nil {
			inspect(request)
		}
		requestNumber := requests.Add(1)
		response.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(response, `{"data":{"attributes":{"access_token":"delegated-token-%d","expires":"%d"}}}`, requestNumber, time.Now().Add(15*time.Minute).Unix())
	}))
}

func delegatedTokenContext(serverURL string) context.Context {
	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		panic(fmt.Sprintf("parsing delegated token server URL: %v", err))
	}
	return context.WithValue(
		context.WithValue(context.Background(), datadog.ContextServerIndex, 1),
		datadog.ContextServerVariables,
		map[string]string{"protocol": parsedURL.Scheme, "name": parsedURL.Host},
	)
}

func delegatedTokenConfig() *datadog.DelegatedTokenConfig {
	return &datadog.DelegatedTokenConfig{OrgUUID: testOrgUUID, Provider: datadog.ProviderAWS}
}

func decodeProof(t *testing.T, proof string) (http.Header, string) {
	t.Helper()
	parts := strings.Split(proof, "|")
	if len(parts) != 4 {
		t.Fatalf("proof has %d components, want 4", len(parts))
	}
	headersJSON, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		t.Fatalf("decoding proof headers: %v", err)
	}
	var headers http.Header
	if err := json.Unmarshal(headersJSON, &headers); err != nil {
		t.Fatalf("unmarshalling proof headers: %v", err)
	}
	endpoint, err := base64.StdEncoding.DecodeString(parts[3])
	if err != nil {
		t.Fatalf("decoding proof endpoint: %v", err)
	}
	return headers, string(endpoint)
}

func firstHeader(headers http.Header, name string) string {
	for key, values := range headers {
		if strings.EqualFold(key, name) && len(values) > 0 {
			return values[0]
		}
	}
	return ""
}
