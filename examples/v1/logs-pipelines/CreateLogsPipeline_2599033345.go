// Create a pipeline with nested pipeline processor returns "OK" response

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
		Name: "testPipelineWithNested",
		Processors: []datadogV1.LogsProcessor{
			datadogV1.LogsProcessor{
				LogsPipelineProcessor: &datadogV1.LogsPipelineProcessor{
					Type:      datadogV1.LOGSPIPELINEPROCESSORTYPE_PIPELINE,
					IsEnabled: datadog.PtrBool(true),
					Name:      datadog.PtrString("nested_pipeline_with_metadata"),
					Filter: &datadogV1.LogsFilter{
						Query: datadog.PtrString("env:production"),
					},
					Tags: []string{
						"env:prod",
						"type:nested",
					},
					Description: datadog.PtrString("This is a nested pipeline for production logs"),
				}},
		},
		Tags: []string{
			"team:test",
		},
		Description: datadog.PtrString("Pipeline containing nested processor with tags and description"),
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
