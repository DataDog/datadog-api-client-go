// Update a specific LLM Observability prompt version returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "prompt" in the system
	PromptDataAttributesPromptID := os.Getenv("PROMPT_DATA_ATTRIBUTES_PROMPT_ID")

	// there is a valid "prompt_version" in the system
	PromptVersionDataAttributesVersion, _ := strconv.ParseInt(os.Getenv("PROMPT_VERSION_DATA_ATTRIBUTES_VERSION"), 10, 64)

	body := datadogV2.LLMObsUpdatePromptVersionRequest{
		Data: datadogV2.LLMObsUpdatePromptVersionData{
			Attributes: datadogV2.LLMObsUpdatePromptVersionDataAttributes{
				Description: datadog.PtrString("Give concise answers and cite relevant help-center articles."),
			},
			Type: datadogV2.LLMOBSPROMPTVERSIONTYPE_PROMPT_TEMPLATE_VERSIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLLMObsPromptVersion", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.UpdateLLMObsPromptVersion(ctx, PromptDataAttributesPromptID, PromptVersionDataAttributesVersion, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpdateLLMObsPromptVersion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.UpdateLLMObsPromptVersion`:\n%s\n", responseContent)
}
