// Create a new pipeline returns "Pipeline created" response

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
	body := datadogV2.Pipeline{
		Data: datadogV2.PipelineData{
			Attributes: datadogV2.PipelineDataAttributes{
				Config: datadogV2.PipelineDataAttributesConfig{
					Destinations: []interface{}{},
					Processors:   []interface{}{},
					Sources:      []interface{}{},
				},
				Name: "",
			},
			Id:   datadog.PtrString("e8890e15-053e-4d34-9404-1b41b9e403e2"),
			Type: "pipeline",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewObservabilityPipelinesApi(apiClient)
	resp, r, err := api.CreatePipeline(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityPipelinesApi.CreatePipeline`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ObservabilityPipelinesApi.CreatePipeline`:\n%s\n", responseContent)
}
