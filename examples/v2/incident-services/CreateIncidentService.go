// Create a new incident service returns "CREATED" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.IncidentServiceCreateRequest{
		Data: datadog.IncidentServiceCreateData{
			Type: datadog.INCIDENTSERVICETYPE_SERVICES,
			Attributes: &datadog.IncidentServiceCreateAttributes{
				Name: "Example-Create_a_new_incident_service_returns_CREATED_response",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("CreateIncidentService", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentServicesApi.CreateIncidentService(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.CreateIncidentService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.CreateIncidentService`:\n%s\n", responseContent)
}
