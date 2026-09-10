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

// writeGcloudShim installs a fake `gcloud` executable that records its argv to
// argvFile and prints token on stdout, returning a directory to prepend to
// PATH. This exercises the real mint path (mintIdentityToken, argument
// construction, output handling) without requiring the gcloud CLI.
func writeGcloudShim(t *testing.T, argvFile, token string) string {
	t.Helper()
	dir := t.TempDir()
	script := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> " + argvFile + "\n" +
		"echo " + token + "\n"
	path := dir + "/gcloud"
	if err := os.WriteFile(path, []byte(script), 0o755); err != nil {
		t.Fatalf("write gcloud shim: %v", err)
	}
	return dir
}

// TestGCPMintIdentityTokenViaShim drives the real mint path through a fake
// gcloud and pins the invocation that builds the GCP audience contract: the
// identity token must carry the datadog/<org-uuid> audience, the
// impersonated service account, and --include-email so the email claim
// identifies the SA.
func TestGCPMintIdentityTokenViaShim(t *testing.T) {
	argvFile := t.TempDir() + "/argv"
	const token = "shim.jwt.sig"
	shimDir := writeGcloudShim(t, argvFile, token)
	t.Setenv("PATH", shimDir)

	auth := datadog.GCPAuth{ImpersonateServiceAccount: "sa@project.iam.gserviceaccount.com"}
	got, err := auth.GetIdentityToken(context.Background(), "org-9")
	if err != nil {
		t.Fatalf("GetIdentityToken: %v", err)
	}
	if got != token {
		t.Errorf("token = %q, want %q", got, token)
	}

	argv, err := os.ReadFile(argvFile)
	if err != nil {
		t.Fatalf("read shim argv: %v", err)
	}
	wantArgs := []string{
		"auth",
		"print-identity-token",
		"--audiences=datadog/org-9",
		"--impersonate-service-account=sa@project.iam.gserviceaccount.com",
		"--include-email",
	}
	gotArgs := strings.Split(strings.TrimSpace(string(argv)), "\n")
	if len(gotArgs) != len(wantArgs) {
		t.Fatalf("gcloud argv = %v, want %v", gotArgs, wantArgs)
	}
	for i := range wantArgs {
		if gotArgs[i] != wantArgs[i] {
			t.Errorf("gcloud argv[%d] = %q, want %q", i, gotArgs[i], wantArgs[i])
		}
	}
}

// TestGCPMintIdentityTokenShimError pins the error path: a failing gcloud
// surfaces its stderr in the wrapped error.
func TestGCPMintIdentityTokenShimError(t *testing.T) {
	dir := t.TempDir()
	script := "#!/bin/sh\necho 'permission denied' >&2\nexit 1\n"
	if err := os.WriteFile(dir+"/gcloud", []byte(script), 0o755); err != nil {
		t.Fatalf("write gcloud shim: %v", err)
	}
	t.Setenv("PATH", dir)

	auth := datadog.GCPAuth{ImpersonateServiceAccount: "sa@project.iam.gserviceaccount.com"}
	_, err := auth.GetIdentityToken(context.Background(), "org-9")
	if err == nil {
		t.Fatal("expected error from failing gcloud shim")
	}
	if !strings.Contains(err.Error(), "permission denied") {
		t.Errorf("err = %v, want it to carry the CLI stderr", err)
	}
}
