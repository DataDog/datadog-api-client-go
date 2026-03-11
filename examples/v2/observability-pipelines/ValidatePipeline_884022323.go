// Validate a metrics pipeline with opentelemetry source returns "OK" response

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
	body := datadogV2.ObservabilityPipelineSpec{
		Data: datadogV2.ObservabilityPipelineSpecData{
			Attributes: datadogV2.ObservabilityPipelineDataAttributes{
				Config: datadogV2.ObservabilityPipelineConfig{
					PipelineType: datadogV2.OBSERVABILITYPIPELINECONFIGPIPELINETYPE_METRICS.Ptr(),
					Destinations: []datadogV2.ObservabilityPipelineConfigDestinationItem{
						datadogV2.ObservabilityPipelineConfigDestinationItem{
							ObservabilityPipelineDatadogMetricsDestination: &datadogV2.ObservabilityPipelineDatadogMetricsDestination{
								Id: "datadog-metrics-destination",
								Inputs: []string{
									"my-processor-group",
								},
								Type: datadogV2.OBSERVABILITYPIPELINEDATADOGMETRICSDESTINATIONTYPE_DATADOG_METRICS,
							}},
					},
					ProcessorGroups: []datadogV2.ObservabilityPipelineConfigProcessorGroup{
						{
							Enabled: true,
							Id:      "my-processor-group",
							Include: "*",
							Inputs: []string{
								"opentelemetry-source",
							},
							Processors: []datadogV2.ObservabilityPipelineConfigProcessorItem{
								datadogV2.ObservabilityPipelineConfigProcessorItem{
									ObservabilityPipelineFilterProcessor: &datadogV2.ObservabilityPipelineFilterProcessor{
										Enabled: true,
										Id:      "filter-processor",
										Include: "env:production",
										Type:    datadogV2.OBSERVABILITYPIPELINEFILTERPROCESSORTYPE_FILTER,
									}},
							},
						},
					},
					Sources: []datadogV2.ObservabilityPipelineConfigSourceItem{
						datadogV2.ObservabilityPipelineConfigSourceItem{
							ObservabilityPipelineOpentelemetrySource: &datadogV2.ObservabilityPipelineOpentelemetrySource{
								Id:   "opentelemetry-source",
								Type: datadogV2.OBSERVABILITYPIPELINEOPENTELEMETRYSOURCETYPE_OPENTELEMETRY,
							}},
					},
				},
				Name: "Metrics OTel Pipeline",
			},
			Type: "pipelines",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
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
