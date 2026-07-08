// Activate an entity context sync integration returns "OK" response

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
	body := datadogV2.SecurityMonitoringIntegrationActivateRequest{
		Data: &datadogV2.SecurityMonitoringIntegrationActivateData{
			Attributes: &datadogV2.SecurityMonitoringIntegrationActivateAttributes{
				Domain: datadog.PtrString("default"),
				Name:   datadog.PtrString("My Entra ID Integration"),
				Settings: map[string]interface{}{
					"setting1": "value1",
				},
			},
			Type: datadogV2.SECURITYMONITORINGINTEGRATIONACTIVATERESOURCETYPE_ACTIVATE_ENTRA_ID_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ActivateIntegration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ActivateIntegration(ctx, "entra_id", *datadogV2.NewActivateIntegrationOptionalParameters().WithBody(body))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ActivateIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ActivateIntegration`:\n%s\n", responseContent)
}
