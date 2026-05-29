// Restore an LLM Observability dataset version returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.LLMObsDatasetRestoreVersionRequest{
		Data: datadogV2.LLMObsDatasetRestoreVersionDataRequest{
			Attributes: datadogV2.LLMObsDatasetRestoreVersionDataAttributesRequest{
				DatasetVersion: 1,
			},
			Id:   "9f64e5c7-dc5a-45c8-a17c-1b85f0bec97d",
			Type: datadogV2.LLMOBSDATASETTYPE_DATASETS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RestoreLLMObsDatasetVersion", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	r, err := api.RestoreLLMObsDatasetVersion(ctx, "project_id", "dataset_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.RestoreLLMObsDatasetVersion`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
