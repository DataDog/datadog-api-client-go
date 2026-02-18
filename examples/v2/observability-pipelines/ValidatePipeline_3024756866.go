// Validate an observability pipeline with OCSF mapper custom mapping returns "OK" response

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
									ObservabilityPipelineOcsfMapperProcessor: &datadogV2.ObservabilityPipelineOcsfMapperProcessor{
										Enabled: true,
										Id:      "ocsf-mapper-processor",
										Include: "service:my-service",
										Mappings: []datadogV2.ObservabilityPipelineOcsfMapperProcessorMapping{
											{
												Include: "source:custom",
												Mapping: datadogV2.ObservabilityPipelineOcsfMapperProcessorMappingMapping{
													ObservabilityPipelineOcsfMappingCustom: &datadogV2.ObservabilityPipelineOcsfMappingCustom{
														Mapping: []datadogV2.ObservabilityPipelineOcsfMappingCustomFieldMapping{
															{
																Default: "",
																Dest:    "time",
																Source:  "timestamp",
															},
															{
																Default: "",
																Dest:    "severity",
																Source:  "level",
															},
															{
																Default: "",
																Dest:    "device.type",
																Lookup: &datadogV2.ObservabilityPipelineOcsfMappingCustomLookup{
																	Table: []datadogV2.ObservabilityPipelineOcsfMappingCustomLookupTableEntry{
																		{
																			Contains: datadog.PtrString("Desktop"),
																			Value:    "desktop",
																		},
																	},
																},
																Source: "host.type",
															},
														},
														Metadata: datadogV2.ObservabilityPipelineOcsfMappingCustomMetadata{
															Class: "Device Inventory Info",
															Profiles: []string{
																"container",
															},
															Version: "1.3.0",
														},
														Version: 1,
													}},
											},
										},
										Type: datadogV2.OBSERVABILITYPIPELINEOCSFMAPPERPROCESSORTYPE_OCSF_MAPPER,
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
				Name: "OCSF Custom Mapper Pipeline",
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
