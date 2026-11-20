// Update an Anthropic Usage and Costs account returns "OK" response

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
	body := datadogV2.AnthropicUsageAndCostsIntegrationAccountUpdateRequest{
		Data: datadogV2.AnthropicUsageAndCostsIntegrationAccountUpdateData{
			Attributes: datadogV2.AnthropicUsageAndCostsIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.AnthropicUsageAndCostsIntegrationAccountAuthenticationUpdate{
					AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate: &datadogV2.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.ANTHROPICUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Dataflows: &datadogV2.AnthropicUsageAndCostsIntegrationDataflowsRequest{
					AnthropicCloudCostMetrics: &datadogV2.AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("anthropic-acme"),
			},
			Id:   "9b4c7e2a-1d6f-4a3b-8e5c-0f2d7a9b6c18",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateAnthropicUsageAndCostsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAnthropicUsageAndCostsIntegrationApi(apiClient)
	resp, r, err := api.UpdateAnthropicUsageAndCostsIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnthropicUsageAndCostsIntegrationApi.UpdateAnthropicUsageAndCostsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AnthropicUsageAndCostsIntegrationApi.UpdateAnthropicUsageAndCostsIntegrationAccount`:\n%s\n", responseContent)
}
