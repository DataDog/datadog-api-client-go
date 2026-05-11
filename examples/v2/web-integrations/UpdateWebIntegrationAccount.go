// Update a web integration account returns "OK" response

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
	body := datadogV2.WebIntegrationAccountUpdateRequest{
		Data: datadogV2.WebIntegrationAccountUpdateRequestData{
			Attributes: datadogV2.WebIntegrationAccountUpdateRequestAttributes{
				Name: datadog.PtrString("my-databricks-account"),
				Secrets: map[string]interface{}{
					"client_secret": "my-client-secret",
				},
				Settings: map[string]interface{}{
					"workspace_url": "https://example.azuredatabricks.net",
				},
			},
			Type: datadogV2.WEBINTEGRATIONACCOUNTTYPE_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateWebIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWebIntegrationsApi(apiClient)
	resp, r, err := api.UpdateWebIntegrationAccount(ctx, "integration_name", "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebIntegrationsApi.UpdateWebIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebIntegrationsApi.UpdateWebIntegrationAccount`:\n%s\n", responseContent)
}
