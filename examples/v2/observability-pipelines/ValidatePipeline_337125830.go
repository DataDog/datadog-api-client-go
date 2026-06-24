// Validate an observability pipeline with websocket source bearer auth returns "OK" response

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
					Destinations: []datadogV2.ObservabilityPipelineConfigDestinationItem{
						datadogV2.ObservabilityPipelineConfigDestinationItem{
							ObservabilityPipelineDatadogLogsDestination: &datadogV2.ObservabilityPipelineDatadogLogsDestination{
								Id: "datadog-logs-destination",
								Inputs: []string{
									"my-processor-group",
								},
								Type: datadogV2.OBSERVABILITYPIPELINEDATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS,
							}},
					},
					ProcessorGroups: []datadogV2.ObservabilityPipelineConfigProcessorGroup{
						{
							Enabled: true,
							Id:      "my-processor-group",
							Include: "service:my-service",
							Inputs: []string{
								"websocket-source",
							},
							Processors: []datadogV2.ObservabilityPipelineConfigProcessorItem{
								datadogV2.ObservabilityPipelineConfigProcessorItem{
									ObservabilityPipelineFilterProcessor: &datadogV2.ObservabilityPipelineFilterProcessor{
										Enabled: true,
										Id:      "filter-processor",
										Include: "status:error",
										Type:    datadogV2.OBSERVABILITYPIPELINEFILTERPROCESSORTYPE_FILTER,
									}},
							},
						},
					},
					Sources: []datadogV2.ObservabilityPipelineConfigSourceItem{
						datadogV2.ObservabilityPipelineConfigSourceItem{
							ObservabilityPipelineWebsocketSource: &datadogV2.ObservabilityPipelineWebsocketSource{
								Id:           "websocket-source",
								Type:         datadogV2.OBSERVABILITYPIPELINEWEBSOCKETSOURCETYPE_WEBSOCKET,
								Decoding:     datadogV2.OBSERVABILITYPIPELINEDECODING_DECODE_JSON,
								AuthStrategy: datadogV2.OBSERVABILITYPIPELINEWEBSOCKETSOURCEAUTHSTRATEGY_BEARER,
								TokenKey:     datadog.PtrString("WS_BEARER_TOKEN"),
								UriKey:       datadog.PtrString("WS_URI"),
								Tls: &datadogV2.ObservabilityPipelineWebsocketSourceTls{
									ObservabilityPipelineWebsocketSourceTlsEnabled: &datadogV2.ObservabilityPipelineWebsocketSourceTlsEnabled{
										Mode: datadogV2.OBSERVABILITYPIPELINEWEBSOCKETSOURCETLSENABLEDMODE_ENABLED,
									}},
							}},
					},
				},
				Name: "Pipeline with WebSocket Source",
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
