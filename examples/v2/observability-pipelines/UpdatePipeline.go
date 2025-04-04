// Update a pipeline returns "OK" response

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
	// there is a valid "pipeline" in the system
	PipelineDataID := os.Getenv("PIPELINE_DATA_ID")

	body := datadogV2.Pipeline{
		Data: datadogV2.PipelineData{
			Attributes: datadogV2.PipelineDataAttributes{
				Config: datadogV2.PipelineConfig{
					Destinations: []datadogV2.PipelineConfigDestination{
						datadogV2.PipelineConfigDestination{
							PipelineDatadogLogsDestination: &datadogV2.PipelineDatadogLogsDestination{
								Id: "updated-datadog-logs-destination-id",
								Inputs: []string{
									"filter-processor",
								},
								Type: datadogV2.PIPELINEDATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS,
							}},
					},
					Processors: []datadogV2.PipelineConfigProcessor{
						datadogV2.PipelineConfigProcessor{
							PipelineFilterProcessor: &datadogV2.PipelineFilterProcessor{
								Id:      "filter-processor",
								Include: "service:my-service",
								Inputs: []string{
									"datadog-agent-source",
								},
								Type: datadogV2.PIPELINEFILTERPROCESSORTYPE_FILTER,
							}},
					},
					Sources: []datadogV2.PipelineConfigSource{
						datadogV2.PipelineConfigSource{
							PipelineDatadogAgentSource: &datadogV2.PipelineDatadogAgentSource{
								Id:   "datadog-agent-source",
								Type: datadogV2.PIPELINEDATADOGAGENTSOURCETYPE_DATADOG_AGENT,
							}},
					},
				},
				Name: "Updated Pipeline Name",
			},
			Id:   PipelineDataID,
			Type: "pipelines",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdatePipeline", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewObservabilityPipelinesApi(apiClient)
	resp, r, err := api.UpdatePipeline(ctx, PipelineDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityPipelinesApi.UpdatePipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ObservabilityPipelinesApi.UpdatePipeline`:\n%s\n", responseContent)
}
