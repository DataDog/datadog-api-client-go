package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

func TestGCPAuthenticate(t *testing.T) {
	type testCase struct {
		name           string
		auth           datadog.GCPAuth
		envToken       string
		expectedProof  string
		expectErr      bool
		expectErrParts []string
	}

	testCases := []testCase{
		{
			name:          "Pre-minted token from field",
			auth:          datadog.GCPAuth{IdentityToken: "field.jwt.sig"},
			expectedProof: "field.jwt.sig",
		},
		{
			name:          "Pre-minted token from env",
			auth:          datadog.GCPAuth{},
			envToken:      "env.jwt.sig",
			expectedProof: "env.jwt.sig",
		},
		{
			name:           "No token and no impersonation fails",
			auth:           datadog.GCPAuth{},
			expectErr:      true,
			expectErrParts: []string{"missing service account to impersonate"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envToken != "" {
				t.Setenv(datadog.GCPIdentityTokenName, tc.envToken)
			} else {
				t.Setenv(datadog.GCPIdentityTokenName, "")
				os.Unsetenv(datadog.GCPIdentityTokenName)
			}

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

			ctx, cfg := delegatedTokenTestContext(t, server, datadog.ProviderGCP)
			cfg.ProviderAuth = &tc.auth

			creds, err := cfg.ProviderAuth.Authenticate(ctx, cfg)
			if tc.expectErr {
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

func TestGCPMintIdentityTokenRequiresImpersonation(t *testing.T) {
	auth := datadog.GCPAuth{}
	_, err := auth.GetIdentityToken(context.Background(), "org-uuid")
	if err == nil || !strings.Contains(err.Error(), "missing service account to impersonate") {
		t.Errorf("err = %v, want missing-impersonation error", err)
	}
}

func TestGCPIdentityTokenPrefersFieldOverEnv(t *testing.T) {
	t.Setenv(datadog.GCPIdentityTokenName, "env.jwt.sig")
	auth := datadog.GCPAuth{IdentityToken: "field.jwt.sig"}
	token, err := auth.GetIdentityToken(context.Background(), "org-uuid")
	if err != nil {
		t.Fatalf("GetIdentityToken: %v", err)
	}
	if token != "field.jwt.sig" {
		t.Errorf("token = %q, want field.jwt.sig", token)
	}
}

func TestGCPContextOverride(t *testing.T) {
	t.Setenv(datadog.GCPIdentityTokenName, "env.jwt.sig")
	ctx := context.WithValue(context.Background(), datadog.ContextGCPVariables, map[string]string{
		datadog.GCPIdentityTokenName: "ctx.jwt.sig",
	})
	auth := datadog.GCPAuth{}
	token, err := auth.GetIdentityToken(ctx, "org-uuid")
	if err != nil {
		t.Fatalf("GetIdentityToken: %v", err)
	}
	if token != "ctx.jwt.sig" {
		t.Errorf("token = %q, want ctx.jwt.sig (context override wins over env)", token)
	}
}
