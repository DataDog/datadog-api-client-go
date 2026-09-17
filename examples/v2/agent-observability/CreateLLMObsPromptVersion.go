// Create a new Agent Observability prompt version returns "OK" response

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
	body := datadogV2.LLMObsCreatePromptVersionRequest{
		Data: datadogV2.LLMObsCreatePromptVersionData{
			Attributes: datadogV2.LLMObsCreatePromptVersionDataAttributes{
				EnvIds: []string{},
				Labels: []datadogV2.LLMObsPromptVersionLabel{
					datadogV2.LLMOBSPROMPTVERSIONLABEL_PRODUCTION,
				},
				Template: datadogV2.LLMObsPromptTemplate{
					LLMObsPromptTextTemplate: datadog.PtrString("You are a helpful assistant for .")},
			},
			Type: datadogV2.LLMOBSPROMPTVERSIONTYPE_PROMPT_TEMPLATE_VERSIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsPromptVersion", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAgentObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsPromptVersion(ctx, "prompt_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentObservabilityApi.CreateLLMObsPromptVersion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AgentObservabilityApi.CreateLLMObsPromptVersion`:\n%s\n", responseContent)
}
