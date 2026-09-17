package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	awsauth "github.com/DataDog/datadog-api-client-go/v2/auth/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
)

func isolateAWSAuthEnvironment(t *testing.T) {
	t.Helper()
	for _, key := range []string{
		"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN", "AWS_PROFILE", "AWS_DEFAULT_PROFILE",
		"AWS_REGION", "AWS_DEFAULT_REGION", "AWS_WEB_IDENTITY_TOKEN_FILE", "AWS_ROLE_ARN", "AWS_ROLE_SESSION_NAME",
		"AWS_CONTAINER_CREDENTIALS_FULL_URI", "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI", "AWS_CONTAINER_AUTHORIZATION_TOKEN",
		"AWS_CONTAINER_AUTHORIZATION_TOKEN_FILE", "AWS_ENDPOINT_URL", "AWS_ENDPOINT_URL_STS", "AWS_CA_BUNDLE",
		"AWS_USE_FIPS_ENDPOINT", "AWS_USE_DUALSTACK_ENDPOINT",
	} {
		t.Setenv(key, "")
	}
	t.Setenv("AWS_EC2_METADATA_DISABLED", "true")
	t.Setenv("AWS_CONFIG_FILE", filepath.Join(t.TempDir(), "config"))
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", filepath.Join(t.TempDir(), "credentials"))
}

func TestAWSProviderRefreshThroughAPIClient(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	var retrievals, exchanges, apiCalls atomic.Int32
	// The first source credential is already expired. A second Retrieve must
	// replace it when the Datadog token expires; no sleeps or live AWS calls.
	cache := aws.NewCredentialsCache(aws.CredentialsProviderFunc(func(context.Context) (aws.Credentials, error) {
		n := retrievals.Add(1)
		expiry := time.Now().Add(time.Hour)
		if n == 1 {
			expiry = time.Now().Add(-time.Hour)
		}
		return aws.Credentials{AccessKeyID: fmt.Sprintf("key-%d", n), SecretAccessKey: "secret", SessionToken: fmt.Sprintf("session-%d", n), CanExpire: true, Expires: expiry}, nil
	}))
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v2/delegated-token":
			n := exchanges.Add(1)
			headers, _ := decodeProof(t, strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated "))
			if !strings.Contains(firstHeader(headers, "Authorization"), fmt.Sprintf("Credential=key-%d/", n)) {
				t.Error("delegated proof did not use refreshed AWS credentials")
			}
			if firstHeader(headers, "X-Amz-Security-Token") != fmt.Sprintf("session-%d", n) {
				t.Error("stale AWS session token")
			}
			fmt.Fprintf(w, `{"data":{"attributes":{"access_token":"token-%d","expires":"%d"}}}`, n, time.Now().Add(time.Hour).Unix())
		case "/api/v2/current_user":
			apiCalls.Add(1)
			if r.Header.Get("Authorization") != fmt.Sprintf("Bearer token-%d", exchanges.Load()) {
				t.Error("API request did not use current delegated token")
			}
			io.WriteString(w, `{"data":{"id":"test-user","type":"users","attributes":{"name":"Test User"}}}`)
		default:
			t.Errorf("unexpected request path %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	provider, err := awsauth.New(awsauth.WithConfigOptions(awsconfig.WithCredentialsProvider(cache)))
	if err != nil {
		t.Fatal(err)
	}
	config := datadog.NewConfiguration()
	config.DelegatedTokenConfig = delegatedTokenConfig()
	config.DelegatedTokenConfig.ProviderAuth = provider
	creds := &datadog.DelegatedTokenCredentials{}
	ctx := context.WithValue(delegatedTokenContext(server.URL), datadog.ContextDelegatedToken, creds)
	users := datadogV2.NewUsersApi(datadog.NewAPIClient(config))
	for i := 0; i < 2; i++ {
		if _, _, err := users.GetCurrentUser(ctx); err != nil {
			t.Fatal(err)
		}
	}
	if exchanges.Load() != 1 || retrievals.Load() != 1 {
		t.Fatal("unexpired Datadog token was not reused")
	}
	creds.Expiration = time.Now().Add(-time.Second)
	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatal(err)
	}
	if exchanges.Load() != 2 || retrievals.Load() != 2 || apiCalls.Load() != 3 {
		t.Fatalf("retrievals=%d exchanges=%d calls=%d", retrievals.Load(), exchanges.Load(), apiCalls.Load())
	}
}

type awsAuthRoundTripFunc func(*http.Request) (*http.Response, error)

func (f awsAuthRoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestAWSProviderUsesCustomHTTPClientForCredentialsAndExchange(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	t.Setenv("AWS_CONTAINER_CREDENTIALS_FULL_URI", "http://127.0.0.1:12345/credentials")
	var credentialRequests, exchanges atomic.Int32
	client := &http.Client{Transport: awsAuthRoundTripFunc(func(r *http.Request) (*http.Response, error) {
		var body string
		switch r.URL.Path {
		case "/credentials":
			credentialRequests.Add(1)
			body = fmt.Sprintf(`{"AccessKeyId":"container-key","SecretAccessKey":"secret","Token":"session","Expiration":"%s"}`, time.Now().Add(time.Hour).UTC().Format(time.RFC3339))
		case "/api/v2/delegated-token":
			exchanges.Add(1)
			body = `{"data":{"attributes":{"access_token":"test-token"}}}`
		default:
			return nil, fmt.Errorf("unexpected request: %s", r.URL.Path)
		}
		return &http.Response{StatusCode: 200, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body)), Request: r}, nil
	})}
	provider, err := awsauth.New(awsauth.WithHTTPClient(client))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := provider.Authenticate(delegatedTokenContext("https://token.invalid"), delegatedTokenConfig()); err != nil {
		t.Fatal(err)
	}
	if credentialRequests.Load() != 1 || exchanges.Load() != 1 {
		t.Fatal("custom client was not used for both credential retrieval and token exchange")
	}
}

func TestAWSProviderExchangeFailuresDoNotLeakTokens(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	for _, tc := range []struct {
		name   string
		status int
		body   string
	}{
		{"rejected proof", 403, `{"error":"sensitive-proof"}`},
		{"malformed JSON", 200, `{"access_token":"sensitive-token"`},
		{"missing data", 200, `{"access_token":"sensitive-token"}`},
		{"missing attributes", 200, `{"data":{"access_token":"sensitive-token"}}`},
		{"invalid token type", 200, `{"data":{"attributes":{"access_token":["sensitive-token"]}}}`},
		{"empty token", 200, `{"data":{"attributes":{"access_token":""}}}`},
		{"oversize response", 200, strings.Repeat("x", (1<<20)+1)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(tc.status); io.WriteString(w, tc.body) }))
			defer server.Close()
			provider, err := awsauth.New(awsauth.WithStaticCredentials("key", "secret", ""))
			if err != nil {
				t.Fatal(err)
			}
			token, err := provider.Authenticate(delegatedTokenContext(server.URL), delegatedTokenConfig())
			if err == nil || token != nil {
				t.Fatal("invalid response accepted")
			}
			if strings.Contains(err.Error(), "sensitive-") {
				t.Fatal("token material leaked into error")
			}
		})
	}
}

func TestAWSProviderHonorsCancellation(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	started := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { close(started); <-r.Context().Done() }))
	defer server.Close()
	provider, err := awsauth.New(awsauth.WithStaticCredentials("key", "secret", ""))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(delegatedTokenContext(server.URL))
	defer cancel()
	done := make(chan error, 1)
	go func() { _, err := provider.Authenticate(ctx, delegatedTokenConfig()); done <- err }()
	select {
	case <-started:
	case <-time.After(5 * time.Second):
		t.Fatal("request never reached server")
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("expected cancellation, got %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("token exchange ignored cancellation")
	}
}

func TestAWSProviderCredentialPrecedenceAndDefaultRegion(t *testing.T) {
	isolateAWSAuthEnvironment(t)
	t.Setenv("AWS_ACCESS_KEY_ID", "environment-key")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "environment-secret")
	for _, explicit := range []bool{false, true} {
		t.Run(fmt.Sprintf("explicit=%t", explicit), func(t *testing.T) {
			var proof string
			server := newDelegatedTokenServer(t, func(r *http.Request) { proof = strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated ") })
			defer server.Close()
			var options []awsauth.Option
			wantKey := "environment-key"
			if explicit {
				options = append(options, awsauth.WithStaticCredentials("explicit-key", "explicit-secret", ""))
				wantKey = "explicit-key"
			}
			provider, err := awsauth.New(options...)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := provider.Authenticate(delegatedTokenContext(server.URL), delegatedTokenConfig()); err != nil {
				t.Fatal(err)
			}
			headers, endpoint := decodeProof(t, proof)
			if !strings.Contains(firstHeader(headers, "Authorization"), "Credential="+wantKey+"/") {
				t.Error("incorrect credential precedence")
			}
			if endpoint != "https://sts.us-east-1.amazonaws.com" {
				t.Errorf("unexpected default endpoint %s", endpoint)
			}
		})
	}
}
