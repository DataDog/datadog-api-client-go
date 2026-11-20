// Create an Anthropic Usage and Costs account returns "Created" response

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
	body := datadogV2.AnthropicUsageAndCostsIntegrationAccountCreateRequest{
		Data: datadogV2.AnthropicUsageAndCostsIntegrationAccountCreateData{
			Attributes: datadogV2.AnthropicUsageAndCostsIntegrationAccountCreateAttributes{
				Authentication: datadogV2.AnthropicUsageAndCostsIntegrationAccountAuthenticationRequest{
					AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest: &datadogV2.AnthropicUsageAndCostsIntegrationAccountBearerTokenAuthRequest{
						AuthType: datadogV2.ANTHROPICUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    "your-token",
					}},
				Dataflows: &datadogV2.AnthropicUsageAndCostsIntegrationDataflowsRequest{
					AnthropicCloudCostMetrics: &datadogV2.AnthropicUsageAndCostsCloudCostMetricsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: "anthropic-acme",
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateAnthropicUsageAndCostsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAnthropicUsageAndCostsIntegrationApi(apiClient)
	resp, r, err := api.CreateAnthropicUsageAndCostsIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnthropicUsageAndCostsIntegrationApi.CreateAnthropicUsageAndCostsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AnthropicUsageAndCostsIntegrationApi.CreateAnthropicUsageAndCostsIntegrationAccount`:\n%s\n", responseContent)
}
