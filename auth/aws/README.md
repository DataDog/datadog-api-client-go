# Datadog API client AWS authentication

This optional module adds AWS Workload Identity Federation authentication to
`datadog-api-client-go` without adding the AWS SDK to the core client's module
graph.

It uses the AWS SDK for Go v2 default configuration and credential chain,
including environment credentials, shared configuration and credential files,
web identity, ECS/EKS container credentials, EC2 instance metadata, SSO, assume
role profiles, and process credentials supported by the SDK.

```go
import (
	awsauth "github.com/DataDog/datadog-api-client-go/auth/aws"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

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
```

The default chain honors standard AWS settings such as `AWS_PROFILE`,
`AWS_CONFIG_FILE`, and `AWS_SHARED_CREDENTIALS_FILE`. Library callers that do
not want to use environment selection can pass
`awsauth.WithSharedConfigProfile("sandbox")`. Explicit credentials remain
available through `awsauth.WithStaticCredentials`.
