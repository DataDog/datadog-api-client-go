// Update pipeline order returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.LogsPipelinesOrder{
		PipelineIds: []string{
			"tags",
			"org_ids",
			"products",
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsPipelinesApi.UpdateLogsPipelineOrder(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsPipelinesApi.UpdateLogsPipelineOrder`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsPipelinesApi.UpdateLogsPipelineOrder`:\n%s\n", responseContent)
}
