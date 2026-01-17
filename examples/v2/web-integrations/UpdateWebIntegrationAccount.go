// Update integration account returns "OK" response

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
	// there is a valid "web_integration_account" in the system
	WebIntegrationAccountDataID := os.Getenv("WEB_INTEGRATION_ACCOUNT_DATA_ID")

	body := datadogV2.WebIntegrationAccountUpdateRequest{
		Data: datadogV2.WebIntegrationAccountUpdateRequestData{
			Type: datadogV2.WEBINTEGRATIONACCOUNTTYPE_ACCOUNT,
			Attributes: &datadogV2.WebIntegrationAccountUpdateRequestAttributes{
				Name: datadog.PtrString("Example-Web-Integration-updated"),
				Settings: map[string]interface{}{
					"events":      "False",
					"messages":    "False",
					"ccm_enabled": "False",
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWebIntegrationsApi(apiClient)
	resp, r, err := api.UpdateWebIntegrationAccount(ctx, "twilio", WebIntegrationAccountDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebIntegrationsApi.UpdateWebIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebIntegrationsApi.UpdateWebIntegrationAccount`:\n%s\n", responseContent)
}
