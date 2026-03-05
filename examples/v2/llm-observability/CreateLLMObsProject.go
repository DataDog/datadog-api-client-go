// Create an LLM Observability project returns "OK" response

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
	body := datadogV2.LLMObsProjectRequest{
		Data: datadogV2.LLMObsProjectDataRequest{
			Attributes: datadogV2.LLMObsProjectDataAttributesRequest{
				Name: "My LLM Project",
			},
			Type: datadogV2.LLMOBSPROJECTTYPE_PROJECTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsProject", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsProject(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsProject`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CreateLLMObsProject`:\n%s\n", responseContent)
}
