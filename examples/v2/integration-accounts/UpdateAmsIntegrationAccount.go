// Update integration account returns "OK: The account was successfully updated." response

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
	body := datadogV2.AmsIntegrationAccountUpdateRequest{
		Data: datadogV2.AmsIntegrationAccountUpdateRequestData{
			Attributes: &datadogV2.AmsIntegrationAccountUpdateRequestAttributes{
				Name: datadog.PtrString("My Production Account (Updated)"),
				Secrets: map[string]interface{}{
					"api_key_token": "new_secret_token_value",
				},
				Settings: map[string]interface{}{
					"ccm_enabled": true,
					"events":      true,
					"messages":    false,
				},
			},
			Type: datadogV2.AMSINTEGRATIONACCOUNTTYPE_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIntegrationAccountsApi(apiClient)
	resp, r, err := api.UpdateAmsIntegrationAccount(ctx, "integration_name", "interface_id", "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAccountsApi.UpdateAmsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAccountsApi.UpdateAmsIntegrationAccount`:\n%s\n", responseContent)
}
