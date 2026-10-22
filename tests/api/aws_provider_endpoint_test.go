package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	awsauth "github.com/DataDog/datadog-api-client-go/v2/auth/aws"
)

func TestAWSProviderProofPreservesSignedHost(t *testing.T) {
	for _, tc := range []struct{ endpoint, host string }{
		{"https://sts.us-east-1.amazonaws.com:443", "sts.us-east-1.amazonaws.com"},
		{"http://localhost:80", "localhost"},
		{"https://localhost:8443", "localhost:8443"},
		{"https://sts.us-east-1.amazonaws.com", "sts.us-east-1.amazonaws.com"},
	} {
		t.Run(tc.endpoint, func(t *testing.T) {
			isolateAWSAuthEnvironment(t)
			t.Setenv("AWS_ENDPOINT_URL_STS", tc.endpoint)
			server := newDelegatedTokenServer(t, func(r *http.Request) {
				headers, endpoint := decodeProof(t, strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated "))
				if got := firstHeader(headers, "Host"); got != tc.host {
					t.Errorf("proof Host = %q, want signed Host %q", got, tc.host)
				}
				if endpoint != tc.endpoint {
					t.Errorf("endpoint = %q, want %q", endpoint, tc.endpoint)
				}
			})
			defer server.Close()
			provider, err := awsauth.New(awsauth.WithStaticCredentials("key", "secret", ""))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := provider.Authenticate(delegatedTokenContext(server.URL), delegatedTokenConfig()); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestAWSProviderRoleCredentialsUseRegionFallback(t *testing.T) {
	for _, webIdentity := range []bool{false, true} {
		for _, configuredRegion := range []string{"", "eu-west-1"} {
			t.Run(fmt.Sprintf("web_identity=%t/region=%s", webIdentity, configuredRegion), func(t *testing.T) {
				isolateAWSAuthEnvironment(t)
				t.Setenv("AWS_REGION", configuredRegion)
				wantRegion := configuredRegion
				if wantRegion == "" {
					wantRegion = "us-east-1"
				}
				action := "AssumeRole"
				if webIdentity {
					action = "AssumeRoleWithWebIdentity"
					path := filepath.Join(t.TempDir(), "token")
					if err := os.WriteFile(path, []byte("test-web-identity-token"), 0o600); err != nil {
						t.Fatal(err)
					}
					t.Setenv("AWS_WEB_IDENTITY_TOKEN_FILE", path)
					t.Setenv("AWS_ROLE_ARN", "arn:aws:iam::123456789012:role/test")
				} else {
					config := "[profile role]\nrole_arn = arn:aws:iam::123456789012:role/test\nsource_profile = source\n[profile source]\naws_access_key_id = source-key\naws_secret_access_key = source-secret\n"
					if err := os.WriteFile(os.Getenv("AWS_CONFIG_FILE"), []byte(config), 0o600); err != nil {
						t.Fatal(err)
					}
					t.Setenv("AWS_PROFILE", "role")
				}
				credentialRequests, exchanges := 0, 0
				client := &http.Client{Transport: awsAuthRoundTripFunc(func(r *http.Request) (*http.Response, error) {
					var body, contentType string
					if r.URL.Host == "sts."+wantRegion+".amazonaws.com" {
						credentialRequests++
						if err := r.ParseForm(); err != nil {
							return nil, err
						}
						if got := r.Form.Get("Action"); got != action {
							t.Errorf("STS action = %q, want %q", got, action)
						}
						contentType = "text/xml"
						body = fmt.Sprintf(`<%sResponse xmlns="https://sts.amazonaws.com/doc/2011-06-15/"><%sResult><Credentials><AccessKeyId>role-key</AccessKeyId><SecretAccessKey>role-secret</SecretAccessKey><SessionToken>role-token</SessionToken><Expiration>2099-01-01T00:00:00Z</Expiration></Credentials></%sResult></%sResponse>`, action, action, action, action)
					} else if r.URL.Host == "token.invalid" {
						exchanges++
						headers, _ := decodeProof(t, strings.TrimPrefix(r.Header.Get("Authorization"), "Delegated "))
						authorization := firstHeader(headers, "Authorization")
						if !strings.Contains(authorization, "Credential=role-key/") || !strings.Contains(authorization, "/"+wantRegion+"/sts/aws4_request") {
							t.Error("proof did not use the assumed role and expected signing region")
						}
						contentType = "application/json"
						body = `{"data":{"attributes":{"access_token":"test-token"}}}`
					} else {
						return nil, fmt.Errorf("unexpected HTTP request: %s", r.URL)
					}
					return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": {contentType}}, Body: io.NopCloser(strings.NewReader(body)), Request: r}, nil
				})}
				provider, err := awsauth.New(awsauth.WithHTTPClient(client))
				if err != nil {
					t.Fatal(err)
				}
				if _, err := provider.Authenticate(delegatedTokenContext("https://token.invalid"), delegatedTokenConfig()); err != nil {
					t.Fatal(err)
				}
				if credentialRequests != 1 || exchanges != 1 {
					t.Fatalf("credential requests = %d, exchanges = %d, want 1 each", credentialRequests, exchanges)
				}
			})
		}
	}
}
