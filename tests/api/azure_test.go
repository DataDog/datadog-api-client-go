package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

func TestAzureAuthenticate(t *testing.T) {
	type testCase struct {
		name           string
		auth           datadog.AzureAuth
		envToken       string
		expectedProof  string
		expectErrParts []string
	}

	testCases := []testCase{
		{
			name:          "Pre-minted token from field",
			auth:          datadog.AzureAuth{AccessToken: "field.jwt.sig"},
			expectedProof: "field.jwt.sig:12345678-1234-1234-1234-123456789012",
		},
		{
			name:          "Pre-minted token from env",
			auth:          datadog.AzureAuth{},
			envToken:      "env.jwt.sig",
			expectedProof: "env.jwt.sig:12345678-1234-1234-1234-123456789012",
		},
		{
			name: "Token source fallback",
			auth: datadog.AzureAuth{TokenSource: func(context.Context) (string, error) {
				return "source.jwt.sig", nil
			}},
			expectedProof: "source.jwt.sig:12345678-1234-1234-1234-123456789012",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(datadog.AzureAccessTokenName, tc.envToken)

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/api/v2/delegated-token" {
					t.Errorf("got path %q, want /api/v2/delegated-token", r.URL.Path)
				}
				if got := r.Header.Get("Authorization"); got != "Delegated "+tc.expectedProof {
					t.Errorf("Authorization = %q, want %q", got, "Delegated "+tc.expectedProof)
				}
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"data":{"type":"delegated_token","attributes":{"access_token":"delegated.jwt.sig","expires":"1700000000"}}}`))
			}))
			defer server.Close()

			ctx, cfg := delegatedTokenTestContext(t, server, datadog.ProviderAzure)
			cfg.ProviderAuth = &tc.auth

			creds, err := cfg.ProviderAuth.Authenticate(ctx, cfg)
			if len(tc.expectErrParts) > 0 {
				if err == nil {
					t.Fatalf("expected error containing %v, got nil", tc.expectErrParts)
				}
				for _, part := range tc.expectErrParts {
					if !strings.Contains(err.Error(), part) {
						t.Errorf("error = %q, want it to contain %q", err.Error(), part)
					}
				}
				return
			}
			if err != nil {
				t.Fatalf("Authenticate: %v", err)
			}
			if creds.DelegatedToken != "delegated.jwt.sig" {
				t.Errorf("DelegatedToken = %q, want delegated.jwt.sig", creds.DelegatedToken)
			}
			if creds.DelegatedProof != tc.expectedProof {
				t.Errorf("DelegatedProof = %q, want %q", creds.DelegatedProof, tc.expectedProof)
			}
		})
	}
}

// TestAzureDelegatedProofSuffix pins the org-routing contract end to end: the
// delegated-token endpoint receives `Authorization: Delegated <token>:<org-uuid>`
// (the servicer splits on the last colon).
func TestAzureDelegatedProofSuffix(t *testing.T) {
	t.Setenv(datadog.AzureAccessTokenName, "")

	var gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data":{"type":"delegated_token","attributes":{"access_token":"delegated.jwt.sig"}}}`))
	}))
	defer server.Close()

	ctx, cfg := delegatedTokenTestContext(t, server, datadog.ProviderAzure)
	cfg.ProviderAuth = &datadog.AzureAuth{AccessToken: "hdr.payload.sig"}

	if _, err := cfg.ProviderAuth.Authenticate(ctx, cfg); err != nil {
		t.Fatalf("Authenticate: %v", err)
	}
	if gotAuth != "Delegated hdr.payload.sig:12345678-1234-1234-1234-123456789012" {
		t.Errorf("Authorization = %q, want suffixed org-routing proof", gotAuth)
	}
}

func TestAzureAccessTokenPrecedence(t *testing.T) {
	t.Setenv(datadog.AzureAccessTokenName, "env.jwt.sig")
	ctx := context.WithValue(context.Background(), datadog.ContextAzureVariables, map[string]string{
		datadog.AzureAccessTokenName: "ctx.jwt.sig",
	})
	tokenSourceCalled := false
	auth := datadog.AzureAuth{
		AccessToken: "field.jwt.sig",
		TokenSource: func(context.Context) (string, error) {
			tokenSourceCalled = true
			return "source.jwt.sig", nil
		},
	}

	token, err := auth.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("GetAccessToken with field: %v", err)
	}
	if token != "field.jwt.sig" {
		t.Errorf("token = %q, want field.jwt.sig", token)
	}

	auth.AccessToken = ""
	token, err = auth.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("GetAccessToken with context token: %v", err)
	}
	if token != "ctx.jwt.sig" {
		t.Errorf("token = %q, want ctx.jwt.sig", token)
	}

	token, err = auth.GetAccessToken(context.Background())
	if err != nil {
		t.Fatalf("GetAccessToken with environment token: %v", err)
	}
	if token != "env.jwt.sig" {
		t.Errorf("token = %q, want env.jwt.sig", token)
	}
	if tokenSourceCalled {
		t.Error("TokenSource called before higher-precedence sources were exhausted")
	}
}

func TestAzureAuthenticateMissingOrgUUID(t *testing.T) {
	auth := datadog.AzureAuth{AccessToken: "token"}
	_, err := auth.Authenticate(context.Background(), &datadog.DelegatedTokenConfig{ProviderAuth: &auth})
	if err == nil || !strings.Contains(err.Error(), "missing org UUID") {
		t.Errorf("err = %v, want missing org UUID error", err)
	}
}

func TestAzureTokenSourceFallback(t *testing.T) {
	t.Setenv(datadog.AzureAccessTokenName, "")
	type contextKey struct{}
	ctx := context.WithValue(context.Background(), contextKey{}, "context value")
	auth := datadog.AzureAuth{TokenSource: func(sourceCtx context.Context) (string, error) {
		if got := sourceCtx.Value(contextKey{}); got != "context value" {
			t.Errorf("TokenSource context value = %v, want context value", got)
		}
		return "source.jwt.sig", nil
	}}

	token, err := auth.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("GetAccessToken: %v", err)
	}
	if token != "source.jwt.sig" {
		t.Errorf("token = %q, want source.jwt.sig", token)
	}
}

func TestAzureTokenSourceError(t *testing.T) {
	t.Setenv(datadog.AzureAccessTokenName, "")
	sourceErr := errors.New("credential unavailable")
	auth := datadog.AzureAuth{TokenSource: func(context.Context) (string, error) {
		return "", sourceErr
	}}

	_, err := auth.GetAccessToken(context.Background())
	if !errors.Is(err, sourceErr) {
		t.Fatalf("error = %v, want wrapped token source error", err)
	}
	if !strings.Contains(err.Error(), "azure token source") {
		t.Errorf("error = %q, want TokenSource context", err.Error())
	}
}

func TestAzureTokenSourceRejectsEmptyToken(t *testing.T) {
	t.Setenv(datadog.AzureAccessTokenName, "")
	auth := datadog.AzureAuth{TokenSource: func(context.Context) (string, error) {
		return "", nil
	}}

	_, err := auth.GetAccessToken(context.Background())
	if err == nil || !strings.Contains(err.Error(), "empty access token") {
		t.Fatalf("error = %v, want empty access token error", err)
	}
}

func TestAzureAccessTokenRequiresSource(t *testing.T) {
	t.Setenv(datadog.AzureAccessTokenName, "")
	_, err := (&datadog.AzureAuth{}).GetAccessToken(context.Background())
	if err == nil {
		t.Fatal("expected missing Azure access token error")
	}
	for _, part := range []string{"AzureAuth.AccessToken", datadog.AzureAccessTokenName, "AzureAuth.TokenSource"} {
		if !strings.Contains(err.Error(), part) {
			t.Errorf("error = %q, want it to contain %q", err.Error(), part)
		}
	}
}
