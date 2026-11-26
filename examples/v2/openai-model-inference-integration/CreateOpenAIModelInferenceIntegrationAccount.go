// Create an OpenAI Model Inference account returns "Created" response

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
	body := datadogV2.OpenAIModelInferenceIntegrationAccountCreateRequest{
		Data: datadogV2.OpenAIModelInferenceIntegrationAccountCreateData{
			Attributes: datadogV2.OpenAIModelInferenceIntegrationAccountCreateAttributes{
				Authentication: datadogV2.OpenAIModelInferenceIntegrationAccountAuthenticationRequest{
					OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest: &datadogV2.OpenAIModelInferenceIntegrationAccountBearerTokenAuthRequest{
						AuthType: datadogV2.OPENAIMODELINFERENCEINTEGRATIONACCOUNTBEARERTOKENAUTHTYPE_BEARER_TOKEN,
						Token:    "your-token",
					}},
				Dataflows: &datadogV2.OpenAIModelInferenceIntegrationDataflowsRequest{
					OpenaiLlmCompletions: &datadogV2.OpenAIModelInferenceLLMCompletionsIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
					OpenaiProjectCloudCostData: &datadogV2.OpenAIModelInferenceProjectCloudCostDataIntegrationDataflowRequest{
						Enabled: datadog.PtrBool(true),
					},
				},
				Name: "openai-model-inference-acme",
				Settings: datadogV2.OpenAIModelInferenceIntegrationAccountSettingsRequest{
					CustomTags: datadog.PtrString("team:ai,env:prod"),
				},
			},
			Type: datadogV2.INTEGRATIONACCOUNTTYPE_INTEGRATION_ACCOUNT,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateOpenAIModelInferenceIntegrationAccount", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOpenAIModelInferenceIntegrationApi(apiClient)
	resp, r, err := api.CreateOpenAIModelInferenceIntegrationAccount(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAIModelInferenceIntegrationApi.CreateOpenAIModelInferenceIntegrationAccount`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpenAIModelInferenceIntegrationApi.CreateOpenAIModelInferenceIntegrationAccount`:\n%s\n", responseContent)
}
