// Create a web integration account returns "CREATED" response

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
	body := datadogV2.WebIntegrationAccountCreateRequest{
		Data: datadogV2.WebIntegrationAccountCreateRequestData{
			Attributes: datadogV2.WebIntegrationAccountCreateRequestAttributes{
				Name: "my-databricks-account",
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
	configuration.SetUnstableOperationEnabled("v2.CreateWebIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWebIntegrationsApi(apiClient)
	resp, r, err := api.CreateWebIntegrationAccount(ctx, "integration_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebIntegrationsApi.CreateWebIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WebIntegrationsApi.CreateWebIntegrationAccount`:\n%s\n", responseContent)
}
