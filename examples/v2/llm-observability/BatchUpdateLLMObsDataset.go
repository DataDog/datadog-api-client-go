// Batch update LLM Observability dataset records returns "OK" response

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
	body := datadogV2.LLMObsDatasetBatchUpdateRequest{
		Data: datadogV2.LLMObsDatasetBatchUpdateDataRequest{
			Attributes: datadogV2.LLMObsDatasetBatchUpdateDataAttributesRequest{
				CreateNewVersion: datadog.PtrBool(true),
				DeleteRecords:    []string{},
				InsertRecords: []datadogV2.LLMObsDatasetBatchUpdateInsertRecord{
					{
						ExpectedOutput: *datadogV2.NewNullableAnyValue(nil),
						Id:             datadog.PtrString("rec-7c3f5a1b-9e2d-4f8a-b1c6-3d7e9f0a2b4c"),
						Input:          *datadogV2.NewNullableAnyValue(nil),
						TagOperations: &datadogV2.LLMObsDatasetRecordTagOperations{
							Add:    []string{},
							Remove: []string{},
							Set:    []string{},
						},
						Tags: []string{},
					},
				},
				Tags: []string{},
				UpdateRecords: []datadogV2.LLMObsDatasetBatchUpdateUpdateRecord{
					{
						ExpectedOutput: *datadogV2.NewNullableAnyValue(nil),
						Id:             "rec-7c3f5a1b-9e2d-4f8a-b1c6-3d7e9f0a2b4c",
						Input:          *datadogV2.NewNullableAnyValue(nil),
						TagOperations: &datadogV2.LLMObsDatasetRecordTagOperations{
							Add:    []string{},
							Remove: []string{},
							Set:    []string{},
						},
					},
				},
			},
			Id:   "9f64e5c7-dc5a-45c8-a17c-1b85f0bec97d",
			Type: datadogV2.LLMOBSDATASETTYPE_DATASETS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.BatchUpdateLLMObsDataset", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.BatchUpdateLLMObsDataset(ctx, "project_id", "dataset_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.BatchUpdateLLMObsDataset`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.BatchUpdateLLMObsDataset`:\n%s\n", responseContent)
}
