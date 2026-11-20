// Create an OpenAI LLM Evaluations account returns "Created" response

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
	body := datadogV2.OpenAILLMEvaluationsIntegrationAccountCreateRequest{
		Data: datadogV2.OpenAILLMEvaluationsIntegrationAccountCreateData{
			Attributes: datadogV2.OpenAILLMEvaluationsIntegrationAccountCreateAttributes{
				Authentication: datadogV2.OpenAILLMEvaluationsIntegrationAccountAuthenticationRequest{
					OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest: &datadogV2.OpenAILLMEvaluationsIntegrationAccountBearerTokenAuthRequest{
						AuthType: datadogV2.OPENAILLMEVALUATIONSINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    "your-token",
					}},
				Dataflows: &datadogV2.OpenAILLMEvaluationsIntegrationDataflowsRequest{
					OpenaiLlmCompletions: &datadogV2.OpenAILLMEvaluationsLLMCompletionsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					OpenaiProjectCloudCostData: &datadogV2.OpenAILLMEvaluationsProjectCloudCostDataIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: "openai-llm-evaluations-acme",
				Settings: datadogV2.OpenAILLMEvaluationsIntegrationAccountSettingsRequest{
					CustomTags: datadog.PtrString("team:ai,env:prod"),
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOpenAILLMEvaluationsIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpenAILLMEvaluationsIntegrationApi(apiClient)
	resp, r, err := api.CreateOpenAILLMEvaluationsIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAILLMEvaluationsIntegrationApi.CreateOpenAILLMEvaluationsIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpenAILLMEvaluationsIntegrationApi.CreateOpenAILLMEvaluationsIntegrationAccount`:\n%s\n", responseContent)
}
