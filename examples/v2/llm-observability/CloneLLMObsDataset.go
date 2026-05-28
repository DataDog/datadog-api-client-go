// Clone an LLM Observability dataset returns "OK" response

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
	body := datadogV2.LLMObsDatasetCloneRequest{
		Data: datadogV2.LLMObsDatasetCloneDataRequest{
			Attributes: datadogV2.LLMObsDatasetCloneDataAttributesRequest{
				Description: datadog.PtrString("Clone of the original dataset for experimentation."),
				Name:        "My cloned dataset",
			},
			Id:   "9f64e5c7-dc5a-45c8-a17c-1b85f0bec97d",
			Type: datadogV2.LLMOBSDATASETTYPE_DATASETS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CloneLLMObsDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CloneLLMObsDataset(ctx, "project_id", "dataset_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CloneLLMObsDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CloneLLMObsDataset`:\n%s\n", responseContent)
}
