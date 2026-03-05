// Delete LLM Observability projects returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.LLMObsDeleteProjectsRequest{
		Data: datadogV2.LLMObsDeleteProjectsDataRequest{
			Attributes: datadogV2.LLMObsDeleteProjectsDataAttributesRequest{
				ProjectIds: []string{
					"a33671aa-24fd-4dcd-9b33-a8ec7dde7751",
				},
			},
			Type: datadogV2.LLMOBSPROJECTTYPE_PROJECTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteLLMObsProjects", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	r, err := api.DeleteLLMObsProjects(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.DeleteLLMObsProjects`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
