// Create an LLM Observability prompt returns "OK" response

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
	body := datadogV2.LLMObsCreatePromptRequest{
		Data: datadogV2.LLMObsCreatePromptData{
			Attributes: datadogV2.LLMObsCreatePromptDataAttributes{
				PromptId: "Example-LLM-Observability",
				Title:    datadog.PtrString("Customer Support Assistant"),
				Template: datadogV2.LLMObsPromptTemplate{
					LLMObsPromptChatTemplate: &datadogV2.LLMObsPromptChatTemplate{Items: []datadogV2.LLMObsPromptChatMessage{
						{
							Content: "You are a helpful customer support assistant for {{company_name}}.",
							Role:    "system",
						},
						{
							Content: "Help {{customer_name}} with this question: {{question}}",
							Role:    "user",
						},
					}}},
			},
			Type: datadogV2.LLMOBSPROMPTTYPE_PROMPT_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsPrompt", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsPrompt(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsPrompt`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CreateLLMObsPrompt`:\n%s\n", responseContent)
}
