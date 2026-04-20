// Create or update a custom evaluator configuration returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.LLMObsCustomEvalConfigUpdateRequest{
		Data: datadogV2.LLMObsCustomEvalConfigUpdateData{
			Attributes: datadogV2.LLMObsCustomEvalConfigUpdateAttributes{
				Category: datadog.PtrString("Custom"),
				EvalName: datadog.PtrString("my-custom-evaluator"),
				LlmJudgeConfig: &datadogV2.LLMObsCustomEvalConfigLLMJudgeConfig{
					AssessmentCriteria: &datadogV2.LLMObsCustomEvalConfigAssessmentCriteria{
						MaxThreshold: *datadog.NewNullableFloat64(datadog.PtrFloat64(1.0)),
						MinThreshold: *datadog.NewNullableFloat64(datadog.PtrFloat64(0.7)),
						PassValues: *datadog.NewNullableList(&[]string{
							"pass",
							"yes",
						}),
						PassWhen: *datadog.NewNullableBool(datadog.PtrBool(true)),
					},
					InferenceParams: datadogV2.LLMObsCustomEvalConfigInferenceParams{
						FrequencyPenalty: datadog.PtrFloat64(0.0),
						MaxTokens:        datadog.PtrInt64(1024),
						PresencePenalty:  datadog.PtrFloat64(0.0),
						Temperature:      datadog.PtrFloat64(0.7),
						TopK:             datadog.PtrInt64(50),
						TopP:             datadog.PtrFloat64(1.0),
					},
					LastUsedLibraryPromptTemplateName: *datadog.NewNullableString(datadog.PtrString("sentiment-analysis-v1")),
					ModifiedLibraryPromptTemplate:     *datadog.NewNullableBool(datadog.PtrBool(false)),
					OutputSchema:                      nil,
					ParsingType:                       datadogV2.LLMOBSCUSTOMEVALCONFIGPARSINGTYPE_STRUCTURED_OUTPUT.Ptr(),
					PromptTemplate: []datadogV2.LLMObsCustomEvalConfigPromptMessage{
						{
							Content: datadog.PtrString("Rate the quality of the following response:"),
							Contents: []datadogV2.LLMObsCustomEvalConfigPromptContent{
								{
									Type: "text",
									Value: datadogV2.LLMObsCustomEvalConfigPromptContentValue{
										Text: datadog.PtrString("What is the sentiment of this review?"),
										ToolCall: &datadogV2.LLMObsCustomEvalConfigPromptToolCall{
											Arguments: datadog.PtrString(`{"location": "San Francisco"}`),
											Id:        datadog.PtrString("call_abc123"),
											Name:      datadog.PtrString("get_weather"),
											Type:      datadog.PtrString("function"),
										},
										ToolCallResult: &datadogV2.LLMObsCustomEvalConfigPromptToolResult{
											Name:   datadog.PtrString("get_weather"),
											Result: datadog.PtrString("sunny, 72F"),
											ToolId: datadog.PtrString("call_abc123"),
											Type:   datadog.PtrString("function"),
										},
									},
								},
							},
							Role: "user",
						},
					},
				},
				LlmProvider: &datadogV2.LLMObsCustomEvalConfigLLMProvider{
					Bedrock: &datadogV2.LLMObsCustomEvalConfigBedrockOptions{
						Region: datadog.PtrString("us-east-1"),
					},
					IntegrationAccountId: datadog.PtrString("my-account-id"),
					IntegrationProvider:  datadogV2.LLMOBSCUSTOMEVALCONFIGINTEGRATIONPROVIDER_OPENAI.Ptr(),
					ModelName:            datadog.PtrString("gpt-4o"),
					VertexAi: &datadogV2.LLMObsCustomEvalConfigVertexAIOptions{
						Location: datadog.PtrString("us-central1"),
						Project:  datadog.PtrString("my-gcp-project"),
					},
				},
				Target: datadogV2.LLMObsCustomEvalConfigTarget{
					ApplicationName:    "my-llm-app",
					Enabled:            true,
					EvalScope:          datadogV2.LLMOBSCUSTOMEVALCONFIGEVALSCOPE_SPAN.Ptr(),
					Filter:             *datadog.NewNullableString(datadog.PtrString("@service:my-service")),
					RootSpansOnly:      *datadog.NewNullableBool(datadog.PtrBool(true)),
					SamplingPercentage: *datadog.NewNullableFloat64(datadog.PtrFloat64(50.0)),
				},
			},
			Id:   datadog.PtrString("my-custom-evaluator"),
			Type: datadogV2.LLMOBSCUSTOMEVALCONFIGTYPE_EVALUATOR_CONFIG,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLLMObsCustomEvalConfig", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	r, err := api.UpdateLLMObsCustomEvalConfig(ctx, "eval_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpdateLLMObsCustomEvalConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
