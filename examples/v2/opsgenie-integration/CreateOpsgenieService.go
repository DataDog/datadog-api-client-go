// Create a new service object returns "CREATED" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.OpsgenieServiceCreateRequest{
		Data: datadog.OpsgenieServiceCreateData{
			Attributes: datadog.OpsgenieServiceCreateAttributes{
				Name:           "Example-Create_a_new_service_object_returns_CREATED_response",
				OpsgenieApiKey: "00000000-0000-0000-0000-000000000000",
				Region:         datadog.OPSGENIESERVICEREGIONTYPE_US,
			},
			Type: datadog.OPSGENIESERVICETYPE_OPSGENIE_SERVICE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.OpsgenieIntegrationApi.CreateOpsgenieService(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpsgenieIntegrationApi.CreateOpsgenieService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `OpsgenieIntegrationApi.CreateOpsgenieService`:\n%s\n", responseContent)
}
