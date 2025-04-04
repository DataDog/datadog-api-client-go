// Create a new pipeline returns "OK" response

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
	body := datadogV2.PipelineCreateRequest{
		Data: datadogV2.PipelineCreateRequestData{
			Attributes: datadogV2.PipelineDataAttributes{
				Config: datadogV2.PipelineConfig{
					Destinations: []datadogV2.PipelineConfigDestination{
						datadogV2.PipelineConfigDestination{
							PipelineDatadogLogsDestination: &datadogV2.PipelineDatadogLogsDestination{
								Id: "datadog-logs-destination",
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
				Name: "Main Observability Pipeline",
			},
			Type: "pipelines",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreatePipeline", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewObservabilityPipelinesApi(apiClient)
	resp, r, err := api.CreatePipeline(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityPipelinesApi.CreatePipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ObservabilityPipelinesApi.CreatePipeline`:\n%s\n", responseContent)
}
