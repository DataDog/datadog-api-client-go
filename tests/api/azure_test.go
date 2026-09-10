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

func TestAzureAuthenticate(t *testing.T) {
	type testCase struct {
		name           string
		auth           datadog.AzureAuth
		envToken       string
		expectedProof  string
		expectErr      bool
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
			// No token anywhere: with PATH emptied below,
			// mintAccessToken's exec.LookPath fails deterministically
			// without invoking a real `az` binary, so the case is hermetic.
			name:           "No token and no az CLI fails",
			auth:           datadog.AzureAuth{},
			expectErr:      true,
			expectErrParts: []string{"`az` not found on PATH"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envToken != "" {
				t.Setenv(datadog.AzureAccessTokenName, tc.envToken)
			} else {
				t.Setenv(datadog.AzureAccessTokenName, "")
				os.Unsetenv(datadog.AzureAccessTokenName)
			}
			if tc.expectErr {
				// Empty PATH so exec.LookPath("az") fails deterministically
				// on machines with the Azure CLI installed.
				t.Setenv("PATH", "")
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

			ctx, cfg := delegatedTokenTestContext(t, server, datadog.ProviderAzure)
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
	os.Unsetenv(datadog.AzureAccessTokenName)

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

func TestAzureAccessTokenPrefersFieldOverEnv(t *testing.T) {
	t.Setenv(datadog.AzureAccessTokenName, "env.jwt.sig")
	auth := datadog.AzureAuth{AccessToken: "field.jwt.sig"}
	token, err := auth.GetAccessToken(context.Background())
	if err != nil {
		t.Fatalf("GetAccessToken: %v", err)
	}
	if token != "field.jwt.sig" {
		t.Errorf("token = %q, want field.jwt.sig", token)
	}
}

func TestAzureContextOverride(t *testing.T) {
	t.Setenv(datadog.AzureAccessTokenName, "env.jwt.sig")
	ctx := context.WithValue(context.Background(), datadog.ContextAzureVariables, map[string]string{
		datadog.AzureAccessTokenName: "ctx.jwt.sig",
	})
	auth := datadog.AzureAuth{}
	token, err := auth.GetAccessToken(ctx)
	if err != nil {
		t.Fatalf("GetAccessToken: %v", err)
	}
	if token != "ctx.jwt.sig" {
		t.Errorf("token = %q, want ctx.jwt.sig (context override wins over env)", token)
	}
}

func TestAzureAuthenticateMissingOrgUUID(t *testing.T) {
	auth := datadog.AzureAuth{AccessToken: "token"}
	_, err := auth.Authenticate(context.Background(), &datadog.DelegatedTokenConfig{ProviderAuth: &auth})
	if err == nil || !strings.Contains(err.Error(), "missing org UUID") {
		t.Errorf("err = %v, want missing org UUID error", err)
	}
}

// writeAzShim installs a fake `az` executable that records its argv to
// argvFile and prints a JSON access-token response on stdout, returning a
// directory to prepend to PATH. This exercises the real mint path
// (mintAccessToken, argument construction, output parsing) without requiring
// the az CLI.
func writeAzShim(t *testing.T, argvFile, token string) string {
	t.Helper()
	dir := t.TempDir()
	script := "#!/bin/sh\n" +
		"printf '%s\\n' \"$@\" >> " + argvFile + "\n" +
		`echo '{"accessToken":"` + token + `"}'` + "\n"
	path := dir + "/az"
	if err := os.WriteFile(path, []byte(script), 0o755); err != nil {
		t.Fatalf("write az shim: %v", err)
	}
	return dir
}

// TestAzureMintAccessTokenViaShim drives the real mint path through a fake az
// and pins the invocation: JSON output (token extracted without depending on
// JMESPath behavior across az versions) and no --resource when Resource is
// unset.
func TestAzureMintAccessTokenViaShim(t *testing.T) {
	argvFile := t.TempDir() + "/argv"
	const token = "shim.jwt.sig"
	shimDir := writeAzShim(t, argvFile, token)
	t.Setenv("PATH", shimDir)
	t.Setenv(datadog.AzureAccessTokenName, "")
	os.Unsetenv(datadog.AzureAccessTokenName)

	auth := datadog.AzureAuth{}
	got, err := auth.GetAccessToken(context.Background())
	if err != nil {
		t.Fatalf("GetAccessToken: %v", err)
	}
	if got != token {
		t.Errorf("token = %q, want %q", got, token)
	}

	argv, err := os.ReadFile(argvFile)
	if err != nil {
		t.Fatalf("read shim argv: %v", err)
	}
	wantArgs := []string{"account", "get-access-token", "--output", "json"}
	gotArgs := strings.Split(strings.TrimSpace(string(argv)), "\n")
	if len(gotArgs) != len(wantArgs) {
		t.Fatalf("az argv = %v, want %v", gotArgs, wantArgs)
	}
	for i := range wantArgs {
		if gotArgs[i] != wantArgs[i] {
			t.Errorf("az argv[%d] = %q, want %q", i, gotArgs[i], wantArgs[i])
		}
	}
}

// TestAzureMintAccessTokenResourceArg pins the --resource passthrough.
func TestAzureMintAccessTokenResourceArg(t *testing.T) {
	argvFile := t.TempDir() + "/argv"
	shimDir := writeAzShim(t, argvFile, "shim.jwt.sig")
	t.Setenv("PATH", shimDir)
	t.Setenv(datadog.AzureAccessTokenName, "")
	os.Unsetenv(datadog.AzureAccessTokenName)

	auth := datadog.AzureAuth{Resource: "https://management.azure.com/"}
	if _, err := auth.GetAccessToken(context.Background()); err != nil {
		t.Fatalf("GetAccessToken: %v", err)
	}
	argv, err := os.ReadFile(argvFile)
	if err != nil {
		t.Fatalf("read shim argv: %v", err)
	}
	if !strings.Contains(string(argv), "--resource") || !strings.Contains(string(argv), "https://management.azure.com/") {
		t.Errorf("az argv = %q, want it to carry --resource with its value", string(argv))
	}
}

// TestAzureMintAccessTokenShimError pins the error path: a failing az
// surfaces its stderr in the wrapped error.
func TestAzureMintAccessTokenShimError(t *testing.T) {
	dir := t.TempDir()
	script := "#!/bin/sh\necho 'run az login' >&2\nexit 1\n"
	if err := os.WriteFile(dir+"/az", []byte(script), 0o755); err != nil {
		t.Fatalf("write az shim: %v", err)
	}
	t.Setenv("PATH", dir)
	t.Setenv(datadog.AzureAccessTokenName, "")
	os.Unsetenv(datadog.AzureAccessTokenName)

	auth := datadog.AzureAuth{}
	_, err := auth.GetAccessToken(context.Background())
	if err == nil {
		t.Fatal("expected error from failing az shim")
	}
	if !strings.Contains(err.Error(), "run az login") {
		t.Errorf("err = %v, want it to carry the CLI stderr", err)
	}
}
