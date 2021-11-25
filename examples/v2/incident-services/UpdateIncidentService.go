// Update an existing incident service returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "service" in the system
	SERVICE_DATA_ID := os.Getenv("SERVICE_DATA_ID")

	body := datadog.IncidentServiceUpdateRequest{
		Data: datadog.IncidentServiceUpdateData{
			Type: datadog.IncidentServiceType("services"),
			Attributes: &datadog.IncidentServiceUpdateAttributes{
				Name: "service name-updated",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("UpdateIncidentService", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentServicesApi.UpdateIncidentService(ctx, SERVICE_DATA_ID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.UpdateIncidentService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.UpdateIncidentService`:\n%s\n", responseContent)
}
