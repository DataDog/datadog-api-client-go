// Create a pipeline with Array Map Processor using category sub-processor returns "OK" response

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
		Name: "testPipelineArrayMapCategory",
		Processors: []datadogV1.LogsProcessor{
			datadogV1.LogsProcessor{
				LogsArrayMapProcessor: &datadogV1.LogsArrayMapProcessor{
					Type:      datadogV1.LOGSARRAYMAPPROCESSORTYPE_ARRAY_MAP_PROCESSOR,
					IsEnabled: datadog.PtrBool(true),
					Name:      datadog.PtrString("categorize items"),
					Source:    "items",
					Target:    "out",
					Processors: []datadogV1.LogsArrayMapSubProcessor{
						datadogV1.LogsArrayMapSubProcessor{
							LogsArrayMapCategorySubProcessor: &datadogV1.LogsArrayMapCategorySubProcessor{
								Type:   datadogV1.LOGSCATEGORYPROCESSORTYPE_CATEGORY_PROCESSOR,
								Target: "$targetElem.level",
								Categories: []datadogV1.LogsCategoryProcessorCategory{
									{
										Filter: &datadogV1.LogsFilter{
											Query: datadog.PtrString("@$sourceElem.status:error"),
										},
										Name: datadog.PtrString("error"),
									},
									{
										Filter: &datadogV1.LogsFilter{
											Query: datadog.PtrString("*"),
										},
										Name: datadog.PtrString("info"),
									},
								},
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
