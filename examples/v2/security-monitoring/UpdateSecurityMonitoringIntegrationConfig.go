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
				Domain:          datadog.PtrString("siem-test.com"),
				Enabled:         datadog.PtrBool(true),
				IntegrationType: datadogV2.SECURITYMONITORINGINTEGRATIONTYPE_GOOGLE_WORKSPACE.Ptr(),
				Name:            datadog.PtrString("My GWS Integration (renamed)"),
				Secrets: map[string]interface{}{
					"admin_email": "test@example.com",
				},
				Settings: map[string]interface{}{
					"setting1": "value1",
				},
			},
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
