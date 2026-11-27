// Create an OpenAI Usage and Costs account returns "Created" response

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
	body := datadogV2.OpenAIUsageAndCostsIntegrationAccountCreateRequest{
		Data: datadogV2.OpenAIUsageAndCostsIntegrationAccountCreateData{
			Attributes: datadogV2.OpenAIUsageAndCostsIntegrationAccountCreateAttributes{
				Authentication: datadogV2.OpenAIUsageAndCostsIntegrationAccountAuthenticationRequest{
					OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest: &datadogV2.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthRequest{
						AuthType: datadogV2.OPENAIUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    "your-token",
					}},
				Dataflows: &datadogV2.OpenAIUsageAndCostsIntegrationDataflowsRequest{
					OpenaiCloudCostData: &datadogV2.OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: "openai-usage-and-costs-acme",
				Settings: datadogV2.OpenAIUsageAndCostsIntegrationAccountSettingsRequest{
					CustomTags: datadog.PtrString("team:ai,env:prod"),
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOpenAIUsageAndCostsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpenAIUsageAndCostsIntegrationApi(apiClient)
	resp, r, err := api.CreateOpenAIUsageAndCostsIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAIUsageAndCostsIntegrationApi.CreateOpenAIUsageAndCostsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpenAIUsageAndCostsIntegrationApi.CreateOpenAIUsageAndCostsIntegrationAccount`:\n%s\n", responseContent)
}
