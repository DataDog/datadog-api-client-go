// Validate entity context sync credentials returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SecurityMonitoringIntegrationCredentialsValidateRequest{
		Data: datadogV2.SecurityMonitoringIntegrationCredentialsValidateData{
			Attributes: datadogV2.SecurityMonitoringIntegrationCredentialsValidateAttributes{
				SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes: &datadogV2.SecurityMonitoringGoogleWorkspaceIntegrationCredentialsValidateAttributes{
					Domain:          "siem-test.com",
					IntegrationType: datadogV2.SECURITYMONITORINGINTEGRATIONTYPEGOOGLEWORKSPACE_GOOGLE_WORKSPACE,
					Secrets: datadogV2.SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets{
						AdminEmail: datadog.PtrString("admin@example.com"),
						ServiceAccountJson: datadogV2.SecurityMonitoringIntegrationConfigGoogleWorkspaceServiceAccount{
							ClientEmail: "svc@my-project.iam.gserviceaccount.com",
							PrivateKey: `-----BEGIN PRIVATE KEY-----
...
-----END PRIVATE KEY-----`,
							ProjectId: "my-project",
							Type:      "service_account",
						},
					},
				}},
			Type: datadogV2.SECURITYMONITORINGINTEGRATIONCONFIGRESOURCETYPE_INTEGRATION_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ValidateSecurityMonitoringIntegrationCredentials", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.ValidateSecurityMonitoringIntegrationCredentials(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ValidateSecurityMonitoringIntegrationCredentials`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
