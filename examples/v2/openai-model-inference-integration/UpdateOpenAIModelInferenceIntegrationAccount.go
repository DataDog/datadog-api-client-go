// Update an OpenAI Model Inference account returns "OK" response

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
	body := datadogV2.OpenAIModelInferenceIntegrationAccountUpdateRequest{
		Data: datadogV2.OpenAIModelInferenceIntegrationAccountUpdateData{
			Attributes: datadogV2.OpenAIModelInferenceIntegrationAccountUpdateAttributes{
				Authentication: &datadogV2.OpenAIModelInferenceIntegrationAccountAuthenticationUpdate{
					OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate: &datadogV2.OpenAIModelInferenceIntegrationAccountBearerTokenAuthUpdate{
						AuthType: datadogV2.OPENAIMODELINFERENCEINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    datadog.PtrString("your-token"),
					}},
				Dataflows: &datadogV2.OpenAIModelInferenceIntegrationDataflowsRequest{
					OpenaiLlmCompletions: &datadogV2.OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					OpenaiProjectCloudCostData: &datadogV2.OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: datadog.PtrString("openai-model-inference-acme"),
				Settings: &datadogV2.OpenAIModelInferenceIntegrationAccountSettingsUpdate{
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
	configuration.SetUnstableOperationEnabled("v2.UpdateOpenAIModelInferenceIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpenAIModelInferenceIntegrationApi(apiClient)
	resp, r, err := api.UpdateOpenAIModelInferenceIntegrationAccount(ctx, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAIModelInferenceIntegrationApi.UpdateOpenAIModelInferenceIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpenAIModelInferenceIntegrationApi.UpdateOpenAIModelInferenceIntegrationAccount`:\n%s\n", responseContent)
}
