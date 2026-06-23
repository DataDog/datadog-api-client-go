// Create a pipeline with Array Map Processor using arithmetic sub-processor returns "OK" response

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
		Name: "testPipelineArrayMapArithmetic",
		Processors: []datadogV1.LogsProcessor{
			datadogV1.LogsProcessor{
				LogsArrayMapProcessor: &datadogV1.LogsArrayMapProcessor{
					Type:      datadogV1.LOGSARRAYMAPPROCESSORTYPE_ARRAY_MAP_PROCESSOR,
					IsEnabled: datadog.PtrBool(true),
					Name:      datadog.PtrString("double counts"),
					Source:    "items",
					Target:    "out",
					Processors: []datadogV1.LogsArrayMapSubProcessor{
						datadogV1.LogsArrayMapSubProcessor{
							LogsArrayMapArithmeticSubProcessor: &datadogV1.LogsArrayMapArithmeticSubProcessor{
								Type:       datadogV1.LOGSARITHMETICPROCESSORTYPE_ARITHMETIC_PROCESSOR,
								Expression: "$sourceElem.count * 2",
								Target:     "$targetElem.doubled",
							}},
					},
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
