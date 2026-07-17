// Update an entity context sync configuration returns "OK" response

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
	body := datadogV2.SecurityMonitoringIntegrationConfigUpdateRequest{
		Data: datadogV2.SecurityMonitoringIntegrationConfigUpdateData{
			Attributes: datadogV2.SecurityMonitoringIntegrationConfigUpdateAttributes{
				SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes: &datadogV2.SecurityMonitoringGoogleWorkspaceIntegrationConfigUpdateAttributes{
					Domain:          datadog.PtrString("siem-test.com"),
					Enabled:         datadog.PtrBool(true),
					IntegrationType: datadogV2.SECURITYMONITORINGINTEGRATIONTYPEGOOGLEWORKSPACE_GOOGLE_WORKSPACE,
					Name:            datadog.PtrString("My GWS Integration (renamed)"),
					Secrets: &datadogV2.SecurityMonitoringIntegrationConfigGoogleWorkspaceSecrets{
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
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityMonitoringIntegrationConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringIntegrationConfig(ctx, "integration_config_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringIntegrationConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringIntegrationConfig`:\n%s\n", responseContent)
}
