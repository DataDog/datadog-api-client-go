# Datadog API client AWS authentication

This opt-in package adds AWS Workload Identity Federation authentication to
`datadog-api-client-go`. It ships with the client's normal v2 releases. AWS SDK
dependencies are declared in the shared module, but applications importing only
`api/datadog` do not compile or link this package or the AWS SDK.

It uses the AWS SDK for Go v2 default configuration and credential chain,
including environment credentials, shared configuration and credential files,
web identity, ECS/EKS container credentials, EC2 instance metadata, SSO, assume
role profiles, and process credentials supported by the SDK.

```go
import (
	"context"

	awsauth "github.com/DataDog/datadog-api-client-go/v2/auth/aws"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Keep authentication values in the client's standard auth context. Empty AWS
// fields select the default credential chain (for example, an AWS_PROFILE).
auth := context.WithValue(context.Background(), datadog.ContextAWSVariables, map[string]string{
	datadog.AWSAccessKeyIdName:     awsAccessKeyID,
	datadog.AWSSecretAccessKeyName: awsSecretAccessKey,
	datadog.AWSSessionTokenName:    awsSessionToken,
})
auth = context.WithValue(auth, datadog.ContextDelegatedToken, &datadog.DelegatedTokenCredentials{})

// Select the AWS implementation in the existing delegated-token config hook.
provider, err := awsauth.New(awsauth.WithRegion("us-east-1"))
if err != nil {
	return err
}
configuration := datadog.NewConfiguration()
configuration.DelegatedTokenConfig = &datadog.DelegatedTokenConfig{
	OrgUUID:      orgUUID,
	Provider:     datadog.ProviderAWS,
	ProviderAuth: provider,
}
// Use auth for API calls made with this configuration. The client caches and
// renews the Datadog delegated token through ContextDelegatedToken as usual.
```

The default chain honors standard AWS settings such as `AWS_PROFILE`,
`AWS_CONFIG_FILE`, and `AWS_SHARED_CREDENTIALS_FILE`. Library callers that do
not want to use environment selection can pass AWS SDK load options through
`awsauth.WithConfigOptions`.

Explicit credentials in `datadog.ContextAWSVariables` take precedence over
constructor credentials and the SDK chain. Supply both the access key and secret;
the session token is optional and is never filled from another source. Partial
credentials or a context value other than `map[string]string` return an error.
An absent, nil, or all-empty credential map uses the configured fallback:
`awsauth.WithStaticCredentials`, if supplied, otherwise the AWS SDK chain.
`WithStaticCredentials` remains available for callers that configure explicit
credentials when constructing the adapter.

The adapter reads context credentials on each authentication attempt, without
caching them on the provider. SDK-chain credentials retain the SDK's normal
caching and refresh behavior. Changing AWS context credentials does not invalidate
an already cached Datadog token; they are used on its next authentication or renewal.
