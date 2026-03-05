// Delete LLM Observability experiments returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.LLMObsDeleteExperimentsRequest{
		Data: datadogV2.LLMObsDeleteExperimentsDataRequest{
			Attributes: datadogV2.LLMObsDeleteExperimentsDataAttributesRequest{
				ExperimentIds: []string{
					"3fd6b5e0-8910-4b1c-a7d0-5b84de329012",
				},
			},
			Type: datadogV2.LLMOBSEXPERIMENTTYPE_EXPERIMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteLLMObsExperiments", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	r, err := api.DeleteLLMObsExperiments(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.DeleteLLMObsExperiments`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
