// Update an LLM Observability prompt returns "OK" response

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

	body := datadogV2.LLMObsUpdatePromptRequest{
		Data: datadogV2.LLMObsUpdatePromptData{
			Attributes: datadogV2.LLMObsUpdatePromptDataAttributes{
				Title: datadog.PtrString("Customer Support Assistant"),
			},
			Type: datadogV2.LLMOBSPROMPTTYPE_PROMPT_TEMPLATES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLLMObsPrompt", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.UpdateLLMObsPrompt(ctx, PromptDataAttributesPromptID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpdateLLMObsPrompt`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.UpdateLLMObsPrompt`:\n%s\n", responseContent)
}
