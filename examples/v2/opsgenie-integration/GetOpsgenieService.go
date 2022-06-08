// Get a single service object returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsgenieIntegrationApi.GetOpsgenieService(ctx, "integration_service_id")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.GetOpsgenieService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpsgenieIntegrationApi.GetOpsgenieService`:\n%s\n", responseContent)
}
