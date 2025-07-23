// Create a pipeline with Decoder Processor returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.LogsPipeline{
		Filter: &datadogV1.LogsFilter{
			Query: datadog.PtrString("source:python"),
		},
		Name: "testDecoderProcessor",
		Processors: []datadogV1.LogsProcessor{
			datadogV1.LogsProcessor{
				LogsDecoderProcessor: &datadogV1.LogsDecoderProcessor{
					Type:                 datadogV1.LOGSDECODERPROCESSORTYPE_DECODER_PROCESSOR,
					IsEnabled:            datadog.PtrBool(true),
					Name:                 datadog.PtrString("test_decoder"),
					Source:               "encoded.field",
					Target:               "decoded.field",
					BinaryToTextEncoding: datadogV1.LOGSDECODERPROCESSORBINARYTOTEXTENCODING_BASE16,
					InputRepresentation:  datadogV1.LOGSDECODERPROCESSORINPUTREPRESENTATION_UTF_8,
				}},
		},
		Tags: []string{},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsPipelinesApi(apiClient)
	resp, r, err := api.CreateLogsPipeline(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.CreateLogsPipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.CreateLogsPipeline`:\n%s\n", responseContent)
}
