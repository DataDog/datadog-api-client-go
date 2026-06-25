// Validate an observability pipeline with ClickHouse destination with all fields set returns "OK" response

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
								Type:               datadogV2.OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONTYPE_CLICKHOUSE,
								EndpointUrlKey:     datadog.PtrString("CLICKHOUSE_ENDPOINT_URL"),
								Database:           datadog.PtrString("my_database"),
								Table:              "application_logs",
								Format:             datadogV2.OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONFORMAT_ARROW_STREAM.Ptr(),
								SkipUnknownFields:  *datadog.NewNullableBool(datadog.PtrBool(true)),
								DateTimeBestEffort: datadog.PtrBool(true),
								Compression: &datadogV2.ObservabilityPipelineClickhouseDestinationCompression{
									ObservabilityPipelineClickhouseDestinationCompressionObject: &datadogV2.ObservabilityPipelineClickhouseDestinationCompressionObject{
										Algorithm: datadogV2.OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONCOMPRESSIONALGORITHM_GZIP,
										Level:     datadog.PtrInt64(6),
									}},
								Auth: &datadogV2.ObservabilityPipelineClickhouseDestinationAuth{
									Strategy:    datadogV2.OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONAUTHSTRATEGY_BASIC,
									UsernameKey: datadog.PtrString("CLICKHOUSE_USERNAME"),
									PasswordKey: datadog.PtrString("CLICKHOUSE_PASSWORD"),
								},
								Batch: &datadogV2.ObservabilityPipelineClickhouseDestinationBatch{
									MaxEvents:   datadog.PtrInt64(1000),
									TimeoutSecs: datadog.PtrInt64(1),
								},
								BatchEncoding: &datadogV2.ObservabilityPipelineClickhouseDestinationBatchEncoding{
									Codec:               datadogV2.OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONBATCHENCODINGCODEC_ARROW_STREAM,
									AllowNullableFields: datadog.PtrBool(true),
								},
								Tls: &datadogV2.ObservabilityPipelineTls{
									CrtFile:    "/path/to/cert.crt",
									CaFile:     datadog.PtrString("/path/to/ca.crt"),
									KeyFile:    datadog.PtrString("/path/to/key.key"),
									KeyPassKey: datadog.PtrString("TLS_KEY_PASSPHRASE"),
								},
								Buffer: &datadogV2.ObservabilityPipelineBufferOptions{
									ObservabilityPipelineMemoryBufferSizeOptions: &datadogV2.ObservabilityPipelineMemoryBufferSizeOptions{
										Type:      datadogV2.OBSERVABILITYPIPELINEBUFFEROPTIONSMEMORYTYPE_MEMORY.Ptr(),
										MaxEvents: 500,
										WhenFull:  datadogV2.OBSERVABILITYPIPELINEBUFFEROPTIONSWHENFULL_BLOCK.Ptr(),
									}},
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
				Name: "Pipeline with ClickHouse Destination All Fields",
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
