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
				Domain:          "siem-test.com",
				IntegrationType: datadogV2.SECURITYMONITORINGINTEGRATIONTYPE_GOOGLE_WORKSPACE,
				Name:            "My GWS Integration",
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
