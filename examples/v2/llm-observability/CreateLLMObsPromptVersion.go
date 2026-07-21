// Create a new LLM Observability prompt version returns "OK" response

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
	// there is a valid "prompt" in the system
	PromptDataAttributesPromptID := os.Getenv("PROMPT_DATA_ATTRIBUTES_PROMPT_ID")

	body := datadogV2.LLMObsCreatePromptVersionRequest{
		Data: datadogV2.LLMObsCreatePromptVersionData{
			Attributes: datadogV2.LLMObsCreatePromptVersionDataAttributes{
				Template: datadogV2.LLMObsPromptTemplate{
					LLMObsPromptChatTemplate: &datadogV2.LLMObsPromptChatTemplate{Items: []datadogV2.LLMObsPromptChatMessage{
						{
							Content: "You are a concise customer support assistant for {{company_name}}.",
							Role:    "system",
						},
						{
							Content: "Answer {{customer_name}}'s question: {{question}}",
							Role:    "user",
						},
					}}},
			},
			Type: datadogV2.LLMOBSPROMPTVERSIONTYPE_PROMPT_TEMPLATE_VERSIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsPromptVersion", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsPromptVersion(ctx, PromptDataAttributesPromptID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsPromptVersion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CreateLLMObsPromptVersion`:\n%s\n", responseContent)
}
