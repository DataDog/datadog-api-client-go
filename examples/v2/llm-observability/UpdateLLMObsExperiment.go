// Update an LLM Observability experiment returns "OK" response

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
	body := datadogV2.LLMObsExperimentUpdateRequest{
		Data: datadogV2.LLMObsExperimentUpdateDataRequest{
			Attributes: datadogV2.LLMObsExperimentUpdateDataAttributesRequest{},
			Type:       datadogV2.LLMOBSEXPERIMENTTYPE_EXPERIMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLLMObsExperiment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.UpdateLLMObsExperiment(ctx, "experiment_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpdateLLMObsExperiment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.UpdateLLMObsExperiment`:\n%s\n", responseContent)
}
