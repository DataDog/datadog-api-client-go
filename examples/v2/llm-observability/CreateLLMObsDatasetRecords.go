// Append records to an LLM Observability dataset returns "OK" response

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
	body := datadogV2.LLMObsDatasetRecordsRequest{
		Data: datadogV2.LLMObsDatasetRecordsDataRequest{
			Attributes: datadogV2.LLMObsDatasetRecordsDataAttributesRequest{
				Records: []datadogV2.LLMObsDatasetRecordItem{
					{
						ExpectedOutput: *datadogV2.NewNullableAnyValue(nil),
						Input:          *datadogV2.NewNullableAnyValue(nil),
					},
				},
			},
			Type: datadogV2.LLMOBSRECORDTYPE_RECORDS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateLLMObsDatasetRecords", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLLMObservabilityApi(apiClient)
	resp, r, err := api.CreateLLMObsDatasetRecords(ctx, "project_id", "dataset_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMObservabilityApi.CreateLLMObsDatasetRecords`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LLMObservabilityApi.CreateLLMObsDatasetRecords`:\n%s\n", responseContent)
}
