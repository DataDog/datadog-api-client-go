// Update an OpenAI LLM Evaluations account returns "OK" response

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
	body := datadogV2.OpenAILLMEvaluationsIntegrationAccountUpdateRequest{
		Data: datadogV2.OpenAILLMEvaluationsIntegrationAccountUpdateData{
			Attributes: datadogV2.OpenAILLMEvaluationsIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.OpenAILLMEvaluationsIntegrationAccountAuthenticationUpdate{
					OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate: &datadogV2.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.OPENAILLMEVALUATIONSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Dataflows: &datadogV2.OpenAILLMEvaluationsIntegrationDataflowsRequest{
					OpenaiLlmCompletions: &datadogV2.OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					OpenaiProjectCloudCostData: &datadogV2.OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("openai-llm-evaluations-acme"),
				Settings: &datadogV2.OpenAILLMEvaluationsIntegrationAccountSettingsUpdate{
					CustomTags: datadog.PtrString("team:ai,env:prod"),
				},
			},
			Id:   "3f8b2c1d-6e4a-4b9f-8c2d-1a7e5b9c0f63",
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateOpenAILLMEvaluationsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpenAILLMEvaluationsIntegrationApi(apiClient)
	resp, r, err := api.UpdateOpenAILLMEvaluationsIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAILLMEvaluationsIntegrationApi.UpdateOpenAILLMEvaluationsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpenAILLMEvaluationsIntegrationApi.UpdateOpenAILLMEvaluationsIntegrationAccount`:\n%s\n", responseContent)
}
