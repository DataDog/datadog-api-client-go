// Run an LLM inference returns "OK" response

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
	body := datadogV2.LLMObsIntegrationInferenceRequest{
		AnthropicMetadata: &datadogV2.LLMObsAnthropicMetadata{
			Effort: *datadogV2.NewNullableLLMObsAnthropicEffort(datadogV2.LLMOBSANTHROPICEFFORT_MEDIUM.Ptr()),
			Thinking: &datadogV2.LLMObsAnthropicThinkingConfig{
				BudgetTokens: *datadog.NewNullableInt64(datadog.PtrInt64(1024)),
				Type:         datadogV2.LLMOBSANTHROPICTHINKINGTYPE_ENABLED,
			},
		},
		AzureOpenaiMetadata: &datadogV2.LLMObsAzureOpenAIMetadata{
			DeploymentId: datadog.PtrString("my-gpt4-deployment"),
			ModelVersion: datadog.PtrString("0613"),
			ResourceName: datadog.PtrString("my-azure-resource"),
		},
		BedrockMetadata: &datadogV2.LLMObsBedrockMetadata{
			Region: datadog.PtrString("us-east-1"),
		},
		FrequencyPenalty:    *datadog.NewNullableFloat64(datadog.PtrFloat64(0.0)),
		JsonSchema:          *datadog.NewNullableString(datadog.PtrString(`{"type":"object","properties":{"answer":{"type":"string"}}}`)),
		MaxCompletionTokens: *datadog.NewNullableInt64(datadog.PtrInt64(1024)),
		MaxTokens:           *datadog.NewNullableInt64(datadog.PtrInt64(1024)),
		Messages: []datadogV2.LLMObsInferenceMessage{
			{
				Content: datadog.PtrString("What is the capital of France?"),
				Contents: []datadogV2.LLMObsInferenceContent{
					{
						Type: "text",
						Value: datadogV2.LLMObsInferenceContentValue{
							Text: datadog.PtrString("Hello, how can I help you?"),
							ToolCall: &datadogV2.LLMObsInferenceToolCall{
								Arguments: map[string]interface{}{
									"location": "San Francisco",
								},
								Name:   datadog.PtrString("get_weather"),
								ToolId: datadog.PtrString("call_abc123"),
								Type:   datadog.PtrString("function"),
							},
							ToolCallResult: &datadogV2.LLMObsInferenceToolResult{
								Name:   datadog.PtrString("get_weather"),
								Result: datadog.PtrString("The weather in San Francisco is 68°F and sunny."),
								ToolId: datadog.PtrString("call_abc123"),
								Type:   datadog.PtrString("function"),
							},
						},
					},
				},
				Id:   datadog.PtrString("msg_001"),
				Role: datadog.PtrString("user"),
				ToolCalls: []datadogV2.LLMObsInferenceToolCall{
					{
						Arguments: map[string]interface{}{
							"location": "San Francisco",
						},
						Name:   datadog.PtrString("get_weather"),
						ToolId: datadog.PtrString("call_abc123"),
						Type:   datadog.PtrString("function"),
					},
				},
				ToolResults: []datadogV2.LLMObsInferenceToolResult{
					{
						Name:   datadog.PtrString("get_weather"),
						Result: datadog.PtrString("The weather in San Francisco is 68°F and sunny."),
						ToolId: datadog.PtrString("call_abc123"),
						Type:   datadog.PtrString("function"),
					},
				},
			},
		},
		ModelId: "gpt-4o",
		OpenaiMetadata: &datadogV2.LLMObsOpenAIMetadata{
			ReasoningEffort:  *datadogV2.NewNullableLLMObsOpenAIReasoningEffort(datadogV2.LLMOBSOPENAIREASONINGEFFORT_MEDIUM.Ptr()),
			ReasoningSummary: *datadogV2.NewNullableLLMObsOpenAIReasoningSummary(datadogV2.LLMOBSOPENAIREASONINGSUMMARY_AUTO.Ptr()),
		},
		PresencePenalty: *datadog.NewNullableFloat64(datadog.PtrFloat64(0.0)),
		Temperature:     *datadog.NewNullableFloat64(datadog.PtrFloat64(0.7)),
		Tools: []datadogV2.LLMObsInferenceTool{
			{
				Function: datadogV2.LLMObsInferenceFunction{
					Description: datadog.PtrString("Get the current weather for a location."),
					Name:        "get_weather",
					Parameters: map[string]interface{}{
						"properties": "{'location': {'type': 'string'}}",
						"type":       "object",
					},
				},
				Type: "function",
			},
		},
		TopK: *datadog.NewNullableInt64(datadog.PtrInt64(50)),
		TopP: *datadog.NewNullableFloat64(datadog.PtrFloat64(1.0)),
		VertexAiMetadata: &datadogV2.LLMObsVertexAIMetadata{
			Location: datadog.PtrString("us-central1"),
			Project:  datadog.PtrString("my-gcp-project"),
			ProjectIds: []string{
				"my-gcp-project",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsIntegrationInference", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsIntegrationInference(ctx, datadogV2.LLMOBSINTEGRATIONNAME_OPENAI, "account_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsIntegrationInference`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CreateLLMObsIntegrationInference`:\n%s\n", responseContent)
}
