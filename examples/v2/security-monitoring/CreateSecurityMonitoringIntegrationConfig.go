// Create an entity context sync configuration returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SecurityMonitoringIntegrationConfigCreateRequest{
		Data: datadogV2.SecurityMonitoringIntegrationConfigCreateData{
			Attributes: datadogV2.SecurityMonitoringIntegrationConfigCreateAttributes{
				SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes: &datadogV2.SecurityMonitoringGoogleWorkspaceIntegrationConfigCreateAttributes{
					Domain:          "siem-test.com",
					IntegrationType: datadogV2.SECURITYMONITORINGINTEGRATIONTYPEGOOGLEWORKSPACE_GOOGLE_WORKSPACE,
					Name:            "My GWS Integration",
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
					Settings: map[string]interface{}{
						"setting1": "value1",
					},
				}},
			Type: datadogV2.SECURITYMONITORINGINTEGRATIONCONFIGRESOURCETYPE_INTEGRATION_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityMonitoringIntegrationConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringIntegrationConfig(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringIntegrationConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringIntegrationConfig`:\n%s\n", responseContent)
}
