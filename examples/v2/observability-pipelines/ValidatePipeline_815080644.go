// Validate an observability pipeline with enrichment table secret field lookup returns "OK" response

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
									ObservabilityPipelineEnrichmentTableProcessor: &datadogV2.ObservabilityPipelineEnrichmentTableProcessor{
										Enabled: true,
										Id:      "enrichment-processor",
										Include: "*",
										Target:  "enriched",
										Type:    datadogV2.OBSERVABILITYPIPELINEENRICHMENTTABLEPROCESSORTYPE_ENRICHMENT_TABLE,
										File: &datadogV2.ObservabilityPipelineEnrichmentTableFile{
											Encoding: datadogV2.ObservabilityPipelineEnrichmentTableFileEncoding{
												Delimiter:       ",",
												Type:            datadogV2.OBSERVABILITYPIPELINEENRICHMENTTABLEFILEENCODINGTYPE_CSV,
												IncludesHeaders: true,
											},
											Key: []datadogV2.ObservabilityPipelineEnrichmentTableFileKeyItems{
												{
													Column:     "user_id",
													Comparison: datadogV2.OBSERVABILITYPIPELINEENRICHMENTTABLEFILEKEYITEMSCOMPARISON_EQUALS,
													Field: datadogV2.ObservabilityPipelineEnrichmentTableFileKeyItemField{
														ObservabilityPipelineEnrichmentTableFieldSecretLookup: &datadogV2.ObservabilityPipelineEnrichmentTableFieldSecretLookup{
															Secret: "LOOKUP_KEY_SECRET",
														}},
												},
											},
											Path: "/etc/enrichment/lookup.csv",
											Schema: []datadogV2.ObservabilityPipelineEnrichmentTableFileSchemaItems{
												{
													Column: "user_id",
													Type:   datadogV2.OBSERVABILITYPIPELINEENRICHMENTTABLEFILESCHEMAITEMSTYPE_STRING,
												},
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
				Name: "Pipeline with Enrichment Table Secret Field Lookup",
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
