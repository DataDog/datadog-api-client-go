// Create a pipeline with Array Processor Key Value Operation returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.LogsPipeline{
		Filter: &datadogV1.LogsFilter{
			Query: datadog.PtrString("source:python"),
		},
		Name: "testPipelineArrayKeyValue",
		Processors: []datadogV1.LogsProcessor{
			datadogV1.LogsProcessor{
				LogsArrayProcessor: &datadogV1.LogsArrayProcessor{
					Type:      datadogV1.LOGSARRAYPROCESSORTYPE_ARRAY_PROCESSOR,
					IsEnabled: datadog.PtrBool(true),
					Name:      datadog.PtrString("extract_kv"),
					Operation: datadogV1.LogsArrayProcessorOperation{
						LogsArrayProcessorOperationExtractKeyValue: &datadogV1.LogsArrayProcessorOperationExtractKeyValue{
							Type:           datadogV1.LOGSARRAYPROCESSOROPERATIONEXTRACTKEYVALUETYPE_KEY_VALUE,
							Source:         "tags",
							KeyToExtract:   "name",
							ValueToExtract: "value",
						}},
				}},
		},
		Tags: []string{},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsPipelinesApi(apiClient)
	resp, r, err := api.CreateLogsPipeline(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.CreateLogsPipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.CreateLogsPipeline`:\n%s\n", responseContent)
}
