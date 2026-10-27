// Validate an observability pipeline with spike protection processor threshold overrides returns "OK" response

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
								"datadog-agent-source",
							},
							Processors: []datadogV2.ObservabilityPipelineConfigProcessorItem{
								datadogV2.ObservabilityPipelineConfigProcessorItem{
									ObservabilityPipelineSpikeProtectionProcessor: &datadogV2.ObservabilityPipelineSpikeProtectionProcessor{
										Enabled:     true,
										Id:          "spike-protection-processor",
										Include:     "*",
										Type:        datadogV2.OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORTYPE_SPIKE_PROTECTION,
										Name:        "MyPartitionedSpikeProtection",
										SpikeAction: datadogV2.OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORSPIKEACTION_NO_ACTION,
										PartitionFields: []string{
											"service",
										},
										ActionWhenMissingPartitions: datadogV2.OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORACTIONWHENMISSINGPARTITIONS_DEFAULT_BUCKET.Ptr(),
										MinimalThresholdOverrides: []datadogV2.ObservabilityPipelineSpikeProtectionProcessorThresholdOverride{
											{
												Fields: []datadogV2.ObservabilityPipelineFieldValue{
													{
														Name:  "service",
														Value: "checkout",
													},
												},
												MinimalThreshold: datadogV2.ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold{
													ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes: &datadogV2.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdBytes{
														Type:  datadogV2.OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORMINIMALTHRESHOLDBYTESTYPE_THRESHOLD,
														Bytes: 53687091200,
													}},
											},
											{
												Fields: []datadogV2.ObservabilityPipelineFieldValue{
													{
														Name:  "service",
														Value: "payments",
													},
												},
												MinimalThreshold: datadogV2.ObservabilityPipelineSpikeProtectionProcessorMinimalThreshold{
													ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep: &datadogV2.ObservabilityPipelineSpikeProtectionProcessorMinimalThresholdAlwaysKeep{
														Type: datadogV2.OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORMINIMALTHRESHOLDALWAYSKEEPTYPE_ALWAYS_KEEP,
													}},
											},
										},
									}},
							},
						},
					},
					Sources: []datadogV2.ObservabilityPipelineConfigSourceItem{
						datadogV2.ObservabilityPipelineConfigSourceItem{
							ObservabilityPipelineDatadogAgentSource: &datadogV2.ObservabilityPipelineDatadogAgentSource{
								Id:   "datadog-agent-source",
								Type: datadogV2.OBSERVABILITYPIPELINEDATADOGAGENTSOURCETYPE_DATADOG_AGENT,
							}},
					},
				},
				Name: "Pipeline with Spike Protection Threshold Overrides",
			},
			Type: "pipelines",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
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
