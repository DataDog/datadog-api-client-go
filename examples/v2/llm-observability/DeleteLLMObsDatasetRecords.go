// Delete LLM Observability dataset records returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.LLMObsDeleteDatasetRecordsRequest{
		Data: datadogV2.LLMObsDeleteDatasetRecordsDataRequest{
			Attributes: datadogV2.LLMObsDeleteDatasetRecordsDataAttributesRequest{
				RecordIds: []string{
					"rec-7c3f5a1b-9e2d-4f8a-b1c6-3d7e9f0a2b4c",
				},
			},
			Type: datadogV2.LLMOBSRECORDTYPE_RECORDS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteLLMObsDatasetRecords", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	r, err := api.DeleteLLMObsDatasetRecords(ctx, "project_id", "dataset_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.DeleteLLMObsDatasetRecords`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
