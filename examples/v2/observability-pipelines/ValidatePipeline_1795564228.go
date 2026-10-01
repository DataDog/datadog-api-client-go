// Validate a metrics pipeline with enrichment table processor file lookup returns "OK" response

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
								"datadog-agent-source",
							},
							Processors: []datadogV2.ObservabilityPipelineConfigProcessorItem{
								datadogV2.ObservabilityPipelineConfigProcessorItem{
									ObservabilityPipelineMetricEnrichmentTableProcessor: &datadogV2.ObservabilityPipelineMetricEnrichmentTableProcessor{
										ObservabilityPipelineMetricEnrichmentTableFileProcessor: &datadogV2.ObservabilityPipelineMetricEnrichmentTableFileProcessor{
											Enabled: true,
											Id:      "enrichment-table-processor",
											Include: "*",
											Type:    datadogV2.OBSERVABILITYPIPELINEENRICHMENTTABLEPROCESSORTYPE_ENRICHMENT_TABLE,
											File: datadogV2.ObservabilityPipelineMetricEnrichmentTableFile{
												Encoding: datadogV2.ObservabilityPipelineEnrichmentTableFileEncoding{
													Delimiter:       ",",
													Type:            datadogV2.OBSERVABILITYPIPELINEENRICHMENTTABLEFILEENCODINGTYPE_CSV,
													IncludesHeaders: true,
												},
												Key: datadogV2.ObservabilityPipelineMetricEnrichmentTableFileKey{
													Column: "service",
													Source: datadogV2.ObservabilityPipelineMetricEnrichmentTableLookupSource{
														ObservabilityPipelineMetricEnrichmentTableTagLookup: &datadogV2.ObservabilityPipelineMetricEnrichmentTableTagLookup{
															Type: datadogV2.OBSERVABILITYPIPELINEMETRICENRICHMENTTABLETAGLOOKUPTYPE_TAG,
															Name: "service",
														}},
												},
												Path: "/etc/enrichment/lookup.csv",
											},
										}}},
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
				Name: "Metrics Pipeline with Enrichment Table File Lookup",
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
