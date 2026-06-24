// Validate an observability pipeline with ClickHouse destination returns "OK" response

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
							ObservabilityPipelineClickhouseDestination: &datadogV2.ObservabilityPipelineClickhouseDestination{
								Id: "clickhouse-destination",
								Inputs: []string{
									"my-processor-group",
								},
								Type:     datadogV2.OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONTYPE_CLICKHOUSE,
								Table:    "application_logs",
								Database: datadog.PtrString("my_database"),
								Compression: &datadogV2.ObservabilityPipelineClickhouseDestinationCompression{
									ObservabilityPipelineClickhouseDestinationCompressionAlgorithm: datadogV2.OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONCOMPRESSIONALGORITHM_GZIP.Ptr()},
								Auth: &datadogV2.ObservabilityPipelineClickhouseDestinationAuth{
									Strategy:    datadogV2.OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONAUTHSTRATEGY_BASIC,
									UsernameKey: datadog.PtrString("CLICKHOUSE_USERNAME"),
									PasswordKey: datadog.PtrString("CLICKHOUSE_PASSWORD"),
								},
								Batch: &datadogV2.ObservabilityPipelineClickhouseDestinationBatch{
									MaxEvents:   datadog.PtrInt64(1000),
									TimeoutSecs: datadog.PtrInt64(1),
								},
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
							ObservabilityPipelineDatadogAgentSource: &datadogV2.ObservabilityPipelineDatadogAgentSource{
								Id:   "datadog-agent-source",
								Type: datadogV2.OBSERVABILITYPIPELINEDATADOGAGENTSOURCETYPE_DATADOG_AGENT,
							}},
					},
				},
				Name: "Pipeline with ClickHouse Destination",
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
