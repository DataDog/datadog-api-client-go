// Create an LLM Observability experiment returns "OK" response

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
	body := datadogV2.LLMObsExperimentRequest{
		Data: datadogV2.LLMObsExperimentDataRequest{
			Attributes: datadogV2.LLMObsExperimentDataAttributesRequest{
				DatasetId: "9f64e5c7-dc5a-45c8-a17c-1b85f0bec97d",
				Name:      "My Experiment v1",
				ProjectId: "a33671aa-24fd-4dcd-9b33-a8ec7dde7751",
			},
			Type: datadogV2.LLMOBSEXPERIMENTTYPE_EXPERIMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsExperiment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsExperiment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsExperiment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CreateLLMObsExperiment`:\n%s\n", responseContent)
}
