// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const ProviderGCP = "gcp"

// GCPIdentityTokenName is the environment variable carrying a pre-minted GCP
// identity token, used instead of invoking `gcloud`.
const GCPIdentityTokenName = "DD_GCP_IDENTITY_TOKEN"

// gcpAudiencePrefix is the audience convention the delegated-token servicer
// pins every GCP exchange to: the identity token must target
// `datadog/<org-uuid>` for the org being routed.
const gcpAudiencePrefix = "datadog/"

// GCPAuth exchanges a Google Cloud identity token for a Datadog delegated token.
type GCPAuth struct {
	// ImpersonateServiceAccount is the service-account email used with
	// `gcloud auth print-identity-token --impersonate-service-account`.
	// Required for minting: user identity tokens cannot set the custom
	// `datadog/<org-uuid>` audience the servicer expects.
	ImpersonateServiceAccount string

	// IdentityToken is an optional pre-minted identity token used verbatim as
	// the exchange proof, bypassing the `gcloud` invocation. When empty, the
	// token is sourced from the DD_GCP_IDENTITY_TOKEN environment variable
	// (or its context override) before falling back to `gcloud`.
	IdentityToken string
}

func (g *GCPAuth) Authenticate(ctx context.Context, config *DelegatedTokenConfig) (*DelegatedTokenCredentials, error) {
	if config == nil || config.OrgUUID == "" {
		return nil, fmt.Errorf("missing org UUID in config")
	}

	proof, err := g.GetIdentityToken(ctx, config.OrgUUID)
	if err != nil {
		return nil, err
	}

	return GetDelegatedToken(ctx, config.OrgUUID, proof)
}

// GetIdentityToken resolves the GCP identity token serving as the exchange
// proof: the IdentityToken field, then the DD_GCP_IDENTITY_TOKEN environment
// variable (with its context override), then `gcloud`.
func (g *GCPAuth) GetIdentityToken(ctx context.Context, orgUUID string) (string, error) {
	if g.IdentityToken != "" {
		return g.IdentityToken, nil
	}
	if token := gcpTokenFromContextOrEnv(ctx); token != "" {
		return token, nil
	}
	return g.mintIdentityToken(ctx, orgUUID)
}

// mintIdentityToken invokes `gcloud auth print-identity-token` with the
// Datadog org audience. Impersonation of a service account is required:
// `gcloud` rejects `--audiences` for user credentials.
func (g *GCPAuth) mintIdentityToken(ctx context.Context, orgUUID string) (string, error) {
	if g.ImpersonateServiceAccount == "" {
		return "", fmt.Errorf("missing service account to impersonate: set GCPAuth.ImpersonateServiceAccount or provide a pre-minted identity token via GCPAuth.IdentityToken or %s", GCPIdentityTokenName)
	}
	if _, err := exec.LookPath("gcloud"); err != nil {
		return "", fmt.Errorf("`gcloud` not found on PATH (install the Google Cloud CLI and run `gcloud auth login`), or provide a pre-minted identity token")
	}
	args := gcloudIdentityTokenArgs(g.ImpersonateServiceAccount, gcpAudiencePrefix+orgUUID)
	out, err := exec.CommandContext(ctx, "gcloud", args...).Output()
	if err != nil {
		return "", fmt.Errorf("gcloud auth print-identity-token: %w (run `gcloud auth login` first)", err)
	}
	token := strings.TrimSpace(string(out))
	if token == "" {
		return "", fmt.Errorf("gcloud auth print-identity-token returned an empty token")
	}
	return token, nil
}

// gcloudIdentityTokenArgs builds the `gcloud` invocation for minting an
// identity token with the Datadog org audience and the service-account email
// in the `email` claim. Factored out for testing.
func gcloudIdentityTokenArgs(impersonate, audience string) []string {
	return []string{
		"auth", "print-identity-token",
		"--audiences=" + audience,
		"--impersonate-service-account=" + impersonate,
		"--include-email",
	}
}

// gcpTokenFromContextOrEnv reads a pre-minted identity token from the context
// override (mirroring ContextAWSVariables), falling back to the environment.
func gcpTokenFromContextOrEnv(ctx context.Context) string {
	if keys, ok := ctx.Value(ContextGCPVariables).(map[string]string); ok {
		if token, ok := keys[GCPIdentityTokenName]; ok {
			return token
		}
	}
	return os.Getenv(GCPIdentityTokenName)
}
