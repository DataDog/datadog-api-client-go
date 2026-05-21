// Validate an observability pipeline with HTTP server source valid_tokens returns "OK" response

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
								"http-server-source",
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
							ObservabilityPipelineHttpServerSource: &datadogV2.ObservabilityPipelineHttpServerSource{
								Id:           "http-server-source",
								Type:         datadogV2.OBSERVABILITYPIPELINEHTTPSERVERSOURCETYPE_HTTP_SERVER,
								AuthStrategy: datadogV2.OBSERVABILITYPIPELINEHTTPSERVERSOURCEAUTHSTRATEGY_NONE,
								Decoding:     datadogV2.OBSERVABILITYPIPELINEDECODING_DECODE_JSON,
								ValidTokens: []datadogV2.ObservabilityPipelineHttpServerSourceValidToken{
									{
										TokenKey: "HTTP_SERVER_TOKEN",
										Enabled:  datadog.PtrBool(true),
										PathToToken: &datadogV2.ObservabilityPipelineHttpServerSourceValidTokenPathToToken{
											ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader: &datadogV2.ObservabilityPipelineHttpServerSourceValidTokenPathToTokenHeader{
												Header: "X-Token",
											}},
										FieldToAdd: &datadogV2.ObservabilityPipelineSourceValidTokenFieldToAdd{
											Key:   "token_name",
											Value: "primary_token",
										},
									},
									{
										TokenKey: "HTTP_SERVER_TOKEN_BACKUP",
										Enabled:  datadog.PtrBool(true),
										PathToToken: &datadogV2.ObservabilityPipelineHttpServerSourceValidTokenPathToToken{
											ObservabilityPipelineHttpServerSourceValidTokenPathToTokenLocation: datadogV2.OBSERVABILITYPIPELINEHTTPSERVERSOURCEVALIDTOKENPATHTOTOKENLOCATION_PATH.Ptr()},
									},
								},
							}},
					},
				},
				Name: "Pipeline with HTTP server valid_tokens",
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
