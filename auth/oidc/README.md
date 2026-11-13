# Datadog API client custom OIDC authentication

This opt-in package adds custom OIDC Workload Identity Federation authentication
to `datadog-api-client-go`. It exchanges a workload identity token issued by the
customer's platform (for example a Kubernetes projected service account token)
for a short-lived Datadog delegated token. It has no dependencies beyond the core
client, and applications importing only `api/datadog` do not link this package.

The token must already be addressed to Datadog: its `aud` claim must be exactly
`datadog/<org-uuid>`, as a string or a single-element array. The provider decodes
the token before contacting Datadog and refuses it, without sending it, if:

- `aud` is missing, names another service or organization, or has more than one value;
- `dd_org_uuid` is present and names another organization;
- `exp` is missing or in the past.

The provider does not verify the signature; Datadog verifies it during the
exchange. Token exchange with an identity provider (RFC 8693) is not supported.

```go
import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	oidcauth "github.com/DataDog/datadog-api-client-go/v2/auth/oidc"
)

// Read the token file on every exchange so kubelet rotation is picked up.
provider, err := oidcauth.New(oidcauth.WithTokenFile("/var/run/secrets/datadog/token"))
if err != nil {
	return err
}
configuration := datadog.NewConfiguration()
configuration.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{
	OrgUUID:      orgUUID,
	Provider:     oidcauth.ProviderName,
	ProviderAuth: provider,
}
// The client caches and renews the delegated token through ContextDelegatedToken.
ctx := context.WithValue(context.Background(), datadog.ContextDelegatedToken, &datadog.DelegatedTokenCredentials{})
```

Use `oidcauth.WithTokenEnv` instead of `WithTokenFile` to read the token from an
environment variable; an environment variable cannot change without restarting
the process. `oidcauth.WithHTTPClient` sets the client used for the exchange, for
custom proxy, TLS, or timeout settings.

## Datadog setup

The organization needs a custom OIDC provider and an identity mapping:

- the provider's issuer URL must exactly match the token's `iss`, and Datadog
  must be able to reach its discovery document and keys;
- the provider must set `org_uuid_source` to `aud` and list `datadog/<org-uuid>`
  in its audiences;
- the identity mapping selects which Datadog user or service account the
  workload acts as, and must match a narrow claim such as an exact `sub`.

For a Kubernetes projected token, mount the volume directory rather than a
single file with `subPath`; files mounted with `subPath` are not updated when
the token rotates.
