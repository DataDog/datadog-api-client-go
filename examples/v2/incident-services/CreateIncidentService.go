// Create a new incident service returns "CREATED" response

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
	body := datadogV2.IncidentServiceCreateRequest{
		Data: datadogV2.IncidentServiceCreateData{
			Type: datadogV2.INCIDENTSERVICETYPE_SERVICES,
			Attributes: &datadogV2.IncidentServiceCreateAttributes{
				Name: "Example-Incident-Service",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentService", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentServicesApi(apiClient)
	resp, r, err := api.CreateIncidentService(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.CreateIncidentService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.CreateIncidentService`:\n%s\n", responseContent)
}
