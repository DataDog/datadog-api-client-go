// Update LLM Observability dataset records returns "OK" response

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
	body := datadogV2.LLMObsDatasetRecordsUpdateRequest{
		Data: datadogV2.LLMObsDatasetRecordsUpdateDataRequest{
			Attributes: datadogV2.LLMObsDatasetRecordsUpdateDataAttributesRequest{
				Records: []datadogV2.LLMObsDatasetRecordUpdateItem{
					{
						ExpectedOutput: *datadogV2.NewNullableAnyValue(nil),
						Id:             "rec-7c3f5a1b-9e2d-4f8a-b1c6-3d7e9f0a2b4c",
						Input:          *datadogV2.NewNullableAnyValue(nil),
					},
				},
			},
			Type: datadogV2.LLMOBSRECORDTYPE_RECORDS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateLLMObsDatasetRecords", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.UpdateLLMObsDatasetRecords(ctx, "project_id", "dataset_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.UpdateLLMObsDatasetRecords`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.UpdateLLMObsDatasetRecords`:\n%s\n", responseContent)
}
