// Validate an observability pipeline returns "OK" response

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
	body := datadogV2.ObservabilityPipelineCreateRequest{
		Data: datadogV2.ObservabilityPipelineCreateRequestData{
			Attributes: datadogV2.ObservabilityPipelineDataAttributes{
				Config: datadogV2.ObservabilityPipelineConfig{
					Destinations: []datadogV2.ObservabilityPipelineConfigDestinationItem{
						datadogV2.ObservabilityPipelineConfigDestinationItem{
							ObservabilityPipelineDatadogLogsDestination: &datadogV2.ObservabilityPipelineDatadogLogsDestination{
								Id: "datadog-logs-destination",
								Inputs: []string{
									"filter-processor",
								},
								Type: datadogV2.OBSERVABILITYPIPELINEDATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS,
							}},
					},
					Processors: []datadogV2.ObservabilityPipelineConfigProcessorItem{
						datadogV2.ObservabilityPipelineConfigProcessorItem{
							ObservabilityPipelineFilterProcessor: &datadogV2.ObservabilityPipelineFilterProcessor{
								Id:      "filter-processor",
								Include: "service:my-service",
								Inputs: []string{
									"datadog-agent-source",
								},
								Type: datadogV2.OBSERVABILITYPIPELINEFILTERPROCESSORTYPE_FILTER,
							}},
					},
					Sources: []datadogV2.ObservabilityPipelineConfigSourceItem{
						datadogV2.ObservabilityPipelineConfigSourceItem{
							ObservabilityPipelineDatadogAgentSource: &datadogV2.ObservabilityPipelineDatadogAgentSource{
								Id:   "datadog-agent-source",
								Type: datadogV2.OBSERVABILITYPIPELINEDATADOGAGENTSOURCETYPE_DATADOG_AGENT,
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
	configuration.SetUnstableOperationEnabled("v2.ValidatePipeline", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewObservabilityPipelinesApi(apiClient)
	resp, r, err := api.ValidatePipeline(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityPipelinesApi.ValidatePipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ObservabilityPipelinesApi.ValidatePipeline`:\n%s\n", responseContent)
}
