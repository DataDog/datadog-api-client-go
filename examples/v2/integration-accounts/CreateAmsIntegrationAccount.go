// Create integration account returns "Created: The account was successfully created." response

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
	body := datadogV2.AmsIntegrationAccountCreateRequest{
		Data: datadogV2.AmsIntegrationAccountCreateRequestData{
			Attributes: datadogV2.AmsIntegrationAccountCreateRequestAttributes{
				Name: "My Production Account",
				Secrets: map[string]interface{}{
					"api_key_token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
				},
				Settings: map[string]interface{}{
					"account_sid":    "ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					"alerts":         true,
					"api_key":        "SKxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					"call_summaries": true,
					"ccm_enabled":    true,
					"censor_logs":    true,
					"events":         true,
					"messages":       true,
				},
			},
			Type: datadogV2.AMSINTEGRATIONACCOUNTTYPE_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIntegrationAccountsApi(apiClient)
	resp, r, err := api.CreateAmsIntegrationAccount(ctx, "integration_name", "interface_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAccountsApi.CreateAmsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAccountsApi.CreateAmsIntegrationAccount`:\n%s\n", responseContent)
}
