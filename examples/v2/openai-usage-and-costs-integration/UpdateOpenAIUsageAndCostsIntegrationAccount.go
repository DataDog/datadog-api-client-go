// Update an OpenAI Usage and Costs account returns "OK" response

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
	body := datadogV2.OpenAIUsageAndCostsIntegrationAccountUpdateRequest{
		Data: datadogV2.OpenAIUsageAndCostsIntegrationAccountUpdateData{
			Attributes: datadogV2.OpenAIUsageAndCostsIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.OpenAIUsageAndCostsIntegrationAccountAuthenticationUpdate{
					OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthUpdate: &datadogV2.OpenAIUsageAndCostsIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.OPENAIUSAGEANDCOSTSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Dataflows: &datadogV2.OpenAIUsageAndCostsIntegrationDataflowsRequest{
					OpenaiCloudCostData: &datadogV2.OpenAIUsageAndCostsCloudCostDataIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("openai-usage-and-costs-acme"),
				Settings: &datadogV2.OpenAIUsageAndCostsIntegrationAccountSettingsUpdate{
					CustomTags: datadog.PtrString("team:ai,env:prod"),
				},
			},
			Id:   "9c4e7a2b-1d5f-4e8a-b3c6-0f2d8a6e4b17",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOpenAIUsageAndCostsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpenAIUsageandCostsIntegrationApi(apiClient)
	resp, r, err := api.UpdateOpenAIUsageAndCostsIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAIUsageandCostsIntegrationApi.UpdateOpenAIUsageAndCostsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpenAIUsageandCostsIntegrationApi.UpdateOpenAIUsageAndCostsIntegrationAccount`:\n%s\n", responseContent)
}
