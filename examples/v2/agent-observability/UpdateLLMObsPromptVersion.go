// Update an Agent Observability prompt version returns "OK" response

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
	body := datadogV2.LLMObsUpdatePromptVersionRequest{
		Data: datadogV2.LLMObsUpdatePromptVersionData{
			Attributes: datadogV2.LLMObsUpdatePromptVersionDataAttributes{
				EnvIds: []string{},
				Labels: []datadogV2.LLMObsPromptVersionLabel{
					datadogV2.LLMOBSPROMPTVERSIONLABEL_PRODUCTION,
				},
			},
			Type: datadogV2.LLMOBSPROMPTVERSIONTYPE_PROMPT_TEMPLATE_VERSIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLLMObsPromptVersion", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentObservabilityApi(apiClient)
	resp, r, err := api.UpdateLLMObsPromptVersion(ctx, "prompt_id", 9223372036854775807, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentObservabilityApi.UpdateLLMObsPromptVersion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentObservabilityApi.UpdateLLMObsPromptVersion`:\n%s\n", responseContent)
}
