// Create a new incident service returns "CREATED" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentService", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.IncidentServicesApi(apiClient)
	resp, r, err := api.CreateIncidentService(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.CreateIncidentService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.CreateIncidentService`:\n%s\n", responseContent)
}
