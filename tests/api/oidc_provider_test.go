package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	oidcauth "github.com/DataDog/datadog-api-client-go/v2/auth/oidc"
)

func oidcTestToken(t *testing.T, claims map[string]any) string {
	t.Helper()
	payload, err := json.Marshal(claims)
	if err != nil {
		t.Fatalf("encoding claims: %v", err)
	}
	encoding := base64.RawURLEncoding
	return encoding.EncodeToString([]byte(`{"alg":"RS256","typ":"JWT"}`)) + "." + encoding.EncodeToString(payload) + ".signature"
}

func oidcScopedToken(t *testing.T, subject string) string {
	t.Helper()
	return oidcTestToken(t, map[string]any{
		"aud": "datadog/" + testOrgUUID,
		"sub": subject,
		"exp": time.Now().Add(10 * time.Minute).Unix(),
	})
}

func oidcWriteToken(t *testing.T, path, token string) {
	t.Helper()
	// Replace atomically, as kubelet does for projected service account tokens.
	temporary := path + ".tmp"
	if err := os.WriteFile(temporary, []byte(token+"\n"), 0o600); err != nil {
		t.Fatalf("writing token: %v", err)
	}
	if err := os.Rename(temporary, path); err != nil {
		t.Fatalf("rotating token: %v", err)
	}
}

func oidcDelegatedTokenConfig() *datadog.DelegatedTokenConfig {
	return &datadog.DelegatedTokenConfig{OrgUUID: testOrgUUID, Provider: oidcauth.ProviderName}
}

func TestOIDCAuthenticateExchangesTokenFromFile(t *testing.T) {
	token := oidcScopedToken(t, "system:serviceaccount:datadog:operator")
	path := filepath.Join(t.TempDir(), "token")
	oidcWriteToken(t, path, token)

	var authorization string
	server := newDelegatedTokenServer(t, func(request *http.Request) {
		authorization = request.Header.Get("Authorization")
	})
	defer server.Close()

	provider, err := oidcauth.New(oidcauth.WithTokenFile(path))
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	credentials, err := provider.Authenticate(delegatedTokenContext(server.URL), oidcDelegatedTokenConfig())
	if err != nil {
		t.Fatalf("authenticating: %v", err)
	}
	if want := "Delegated oidc-" + token; authorization != want {
		t.Errorf("authorization header = %q, want %q", authorization, want)
	}
	if credentials.DelegatedToken != "delegated-token-1" || credentials.OrgUUID != testOrgUUID {
		t.Errorf("credentials = %+v, want delegated-token-1 for %s", credentials, testOrgUUID)
	}
	if !credentials.Expiration.After(time.Now()) {
		t.Errorf("expiration = %s, want a future time", credentials.Expiration)
	}
}

func TestOIDCAuthenticateExchangesTokenFromEnvironment(t *testing.T) {
	token := oidcScopedToken(t, "project_path:group/project:ref_type:branch:ref:main")
	t.Setenv("DD_TEST_OIDC_TOKEN", token)

	var authorization string
	server := newDelegatedTokenServer(t, func(request *http.Request) {
		authorization = request.Header.Get("Authorization")
	})
	defer server.Close()

	provider, err := oidcauth.New(oidcauth.WithTokenEnv("DD_TEST_OIDC_TOKEN"))
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	if _, err := provider.Authenticate(delegatedTokenContext(server.URL), oidcDelegatedTokenConfig()); err != nil {
		t.Fatalf("authenticating: %v", err)
	}
	if want := "Delegated oidc-" + token; authorization != want {
		t.Errorf("authorization header = %q, want %q", authorization, want)
	}
}

func TestOIDCAuthenticateRereadsRotatedTokenFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token")
	first := oidcScopedToken(t, "first")
	second := oidcScopedToken(t, "second")
	oidcWriteToken(t, path, first)

	var proofs []string
	server := newDelegatedTokenServer(t, func(request *http.Request) {
		proofs = append(proofs, strings.TrimPrefix(request.Header.Get("Authorization"), "Delegated oidc-"))
	})
	defer server.Close()

	provider, err := oidcauth.New(oidcauth.WithTokenFile(path))
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	ctx := delegatedTokenContext(server.URL)
	if _, err := provider.Authenticate(ctx, oidcDelegatedTokenConfig()); err != nil {
		t.Fatalf("first authentication: %v", err)
	}
	oidcWriteToken(t, path, second)
	if _, err := provider.Authenticate(ctx, oidcDelegatedTokenConfig()); err != nil {
		t.Fatalf("second authentication: %v", err)
	}
	if len(proofs) != 2 || proofs[0] != first || proofs[1] != second {
		t.Fatalf("exchanged tokens did not follow file rotation")
	}
}

func TestOIDCAuthenticateRefusesTokensNotScopedToDatadog(t *testing.T) {
	future := time.Now().Add(10 * time.Minute).Unix()
	tests := []struct {
		name   string
		token  string
		reason string
	}{
		{"wrong audience", oidcTestToken(t, map[string]any{"aud": "https://internal-api.example", "exp": future}), `aud is "https://internal-api.example"`},
		{"other organization", oidcTestToken(t, map[string]any{"aud": "datadog/other-org", "exp": future}), `aud is "datadog/other-org"`},
		{"multiple audiences", oidcTestToken(t, map[string]any{"aud": []string{"datadog/" + testOrgUUID, "other-api"}, "exp": future}), "aud has 2 values"},
		{"no audience", oidcTestToken(t, map[string]any{"exp": future}), "no aud claim"},
		{"mismatched dd_org_uuid", oidcTestToken(t, map[string]any{"aud": "datadog/" + testOrgUUID, "dd_org_uuid": "other-org", "exp": future}), "dd_org_uuid"},
		{"no expiry", oidcTestToken(t, map[string]any{"aud": "datadog/" + testOrgUUID}), "no exp claim"},
		{"expired", oidcTestToken(t, map[string]any{"aud": "datadog/" + testOrgUUID, "exp": time.Now().Add(-time.Minute).Unix()}), "expired"},
		{"opaque token", "opaque-access-token", "not a JWT"},
		{"numeric audience", oidcTestToken(t, map[string]any{"aud": 7, "exp": future}), "neither a string"},
	}

	var requests atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		requests.Add(1)
	}))
	defer server.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("DD_TEST_OIDC_TOKEN", test.token)
			provider, err := oidcauth.New(oidcauth.WithTokenEnv("DD_TEST_OIDC_TOKEN"))
			if err != nil {
				t.Fatalf("creating provider: %v", err)
			}
			_, err = provider.Authenticate(delegatedTokenContext(server.URL), oidcDelegatedTokenConfig())
			var scopeError *oidcauth.TokenScopeError
			if !errors.As(err, &scopeError) || !strings.Contains(scopeError.Reason, test.reason) {
				t.Fatalf("error = %v, want TokenScopeError containing %q", err, test.reason)
			}
		})
	}
	if count := requests.Load(); count != 0 {
		t.Fatalf("refused tokens reached the delegated token endpoint %d times", count)
	}
}

func TestCheckDatadogScopedAcceptsBothAudienceForms(t *testing.T) {
	future := time.Now().Add(time.Minute).Unix()
	for name, aud := range map[string]any{
		"string":         "datadog/" + testOrgUUID,
		"single element": []string{"datadog/" + testOrgUUID},
	} {
		token := oidcTestToken(t, map[string]any{"aud": aud, "exp": future, "dd_org_uuid": testOrgUUID})
		if err := oidcauth.CheckDatadogScoped(token, testOrgUUID, time.Now()); err != nil {
			t.Errorf("%s audience: %v", name, err)
		}
	}
}

func TestOIDCNewValidatesOptions(t *testing.T) {
	tests := map[string][]oidcauth.Option{
		"no token source":   nil,
		"two token sources": {oidcauth.WithTokenFile("/token"), oidcauth.WithTokenEnv("TOKEN")},
		"empty file path":   {oidcauth.WithTokenFile("")},
		"empty env name":    {oidcauth.WithTokenEnv("")},
		"nil HTTP client":   {oidcauth.WithTokenEnv("TOKEN"), oidcauth.WithHTTPClient(nil)},
	}
	for name, options := range tests {
		if _, err := oidcauth.New(options...); err == nil {
			t.Errorf("%s: New succeeded, want an error", name)
		}
	}
}

func TestOIDCAuthenticateReportsSourceErrors(t *testing.T) {
	directory := t.TempDir()
	emptyFile := filepath.Join(directory, "empty")
	if err := os.WriteFile(emptyFile, []byte(" \n"), 0o600); err != nil {
		t.Fatal(err)
	}
	tests := map[string]oidcauth.Option{
		"missing file":  oidcauth.WithTokenFile(filepath.Join(directory, "missing")),
		"empty file":    oidcauth.WithTokenFile(emptyFile),
		"unset env var": oidcauth.WithTokenEnv("DD_TEST_OIDC_TOKEN_UNSET"),
	}
	for name, option := range tests {
		provider, err := oidcauth.New(option)
		if err != nil {
			t.Fatalf("%s: creating provider: %v", name, err)
		}
		if _, err := provider.Authenticate(context.Background(), oidcDelegatedTokenConfig()); err == nil {
			t.Errorf("%s: Authenticate succeeded, want an error", name)
		}
	}
	provider, err := oidcauth.New(oidcauth.WithTokenEnv("DD_TEST_OIDC_TOKEN_UNSET"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := provider.Authenticate(context.Background(), &datadog.DelegatedTokenConfig{}); err == nil {
		t.Error("Authenticate without an org UUID succeeded, want an error")
	}
}

type oidcCountingTransport struct {
	requests atomic.Int32
}

func (transport *oidcCountingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	transport.requests.Add(1)
	return http.DefaultTransport.RoundTrip(request)
}

func TestOIDCAuthenticateUsesConfiguredHTTPClient(t *testing.T) {
	t.Setenv("DD_TEST_OIDC_TOKEN", oidcScopedToken(t, "subject"))
	server := newDelegatedTokenServer(t, nil)
	defer server.Close()

	transport := &oidcCountingTransport{}
	provider, err := oidcauth.New(
		oidcauth.WithTokenEnv("DD_TEST_OIDC_TOKEN"),
		oidcauth.WithHTTPClient(&http.Client{Transport: transport}),
	)
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	if _, err := provider.Authenticate(delegatedTokenContext(server.URL), oidcDelegatedTokenConfig()); err != nil {
		t.Fatalf("authenticating: %v", err)
	}
	if transport.requests.Load() != 1 {
		t.Fatalf("configured HTTP client handled %d requests, want 1", transport.requests.Load())
	}
}

func TestOIDCAuthenticateDoesNotExposeMalformedResponses(t *testing.T) {
	t.Setenv("DD_TEST_OIDC_TOKEN", oidcScopedToken(t, "subject"))
	for name, handler := range map[string]http.HandlerFunc{
		"malformed body": func(response http.ResponseWriter, _ *http.Request) {
			_, _ = fmt.Fprint(response, `{"data":{"secret-token-material":true}}`)
		},
		"error status": func(response http.ResponseWriter, _ *http.Request) {
			response.WriteHeader(http.StatusUnauthorized)
			_, _ = fmt.Fprint(response, `{"errors":["secret-token-material"]}`)
		},
	} {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(handler)
			defer server.Close()
			provider, err := oidcauth.New(oidcauth.WithTokenEnv("DD_TEST_OIDC_TOKEN"))
			if err != nil {
				t.Fatal(err)
			}
			_, err = provider.Authenticate(delegatedTokenContext(server.URL), oidcDelegatedTokenConfig())
			if err == nil || strings.Contains(err.Error(), "secret-token-material") {
				t.Fatalf("error = %v, want a failure that omits the response body", err)
			}
		})
	}
}

func TestOIDCClientSendsAndRefreshesDelegatedToken(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token")
	oidcWriteToken(t, path, oidcScopedToken(t, "first"))

	var exchanges atomic.Int32
	var bearer atomic.Value
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		switch request.URL.Path {
		case "/api/v2/delegated-token":
			number := exchanges.Add(1)
			_, _ = fmt.Fprintf(response, `{"data":{"attributes":{"access_token":"delegated-token-%d","expires":"%d"}}}`, number, time.Now().Add(15*time.Minute).Unix())
		case "/api/v2/current_user":
			bearer.Store(request.Header.Get("Authorization"))
			_, _ = fmt.Fprint(response, `{"data":{"type":"users","id":"00000000-0000-0000-0000-000000000001"}}`)
		default:
			http.NotFound(response, request)
		}
	}))
	defer server.Close()

	provider, err := oidcauth.New(oidcauth.WithTokenFile(path))
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	configuration := datadog.NewConfiguration()
	configuration.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{
		OrgUUID: testOrgUUID, Provider: oidcauth.ProviderName, ProviderAuth: provider,
	}
	credentials := &datadog.DelegatedTokenCredentials{}
	ctx := context.WithValue(delegatedTokenContext(server.URL), datadog.ContextDelegatedToken, credentials)
	users := datadogV2.NewUsersApi(datadog.NewAPIClient(configuration))

	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatalf("first API call: %v", err)
	}
	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatalf("second API call: %v", err)
	}
	if exchanges.Load() != 1 || bearer.Load() != "Bearer delegated-token-1" {
		t.Fatalf("exchanges = %d, authorization = %v; want one exchange reused as a bearer token", exchanges.Load(), bearer.Load())
	}

	oidcWriteToken(t, path, oidcScopedToken(t, "second"))
	credentials.Expiration = time.Now().Add(-time.Second)
	if _, _, err := users.GetCurrentUser(ctx); err != nil {
		t.Fatalf("API call after expiry: %v", err)
	}
	if exchanges.Load() != 2 || bearer.Load() != "Bearer delegated-token-2" {
		t.Fatalf("exchanges = %d, authorization = %v; want a refreshed delegated token", exchanges.Load(), bearer.Load())
	}
}

func TestOIDCAuthenticateAgainstDatadogIntegration(t *testing.T) {
	apiURL := os.Getenv("DD_TEST_WIF_API_URL")
	orgUUID := os.Getenv("DD_ORG_UUID")
	tokenFile := os.Getenv("DD_TEST_OIDC_TOKEN_FILE")
	if apiURL == "" || orgUUID == "" || tokenFile == "" {
		t.Skip("set DD_TEST_WIF_API_URL, DD_ORG_UUID, and DD_TEST_OIDC_TOKEN_FILE to test a live Datadog token exchange")
	}
	provider, err := oidcauth.New(oidcauth.WithTokenFile(tokenFile))
	if err != nil {
		t.Fatalf("creating provider: %v", err)
	}
	credentials, err := provider.Authenticate(delegatedTokenContext(apiURL), &datadog.DelegatedTokenConfig{
		OrgUUID: orgUUID, Provider: oidcauth.ProviderName,
	})
	if err != nil {
		t.Fatalf("authenticating against Datadog: %v", err)
	}
	if credentials.DelegatedToken == "" || !credentials.Expiration.After(time.Now()) {
		t.Fatalf("Datadog returned unusable credentials (expiration %s)", credentials.Expiration)
	}
}
