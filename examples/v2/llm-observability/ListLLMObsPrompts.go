// List LLM Observability prompts returns "OK" response

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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListLLMObsPrompts", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.ListLLMObsPrompts(ctx, *datadogV2.NewListLLMObsPromptsOptionalParameters().WithFilterPromptId(PromptDataAttributesPromptID))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.ListLLMObsPrompts`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.ListLLMObsPrompts`:\n%s\n", responseContent)
}
