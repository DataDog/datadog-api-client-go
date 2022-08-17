// Update an existing incident service returns "OK" response

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
	// there is a valid "service" in the system
	ServiceDataID := os.Getenv("SERVICE_DATA_ID")

	body := datadogV2.IncidentServiceUpdateRequest{
		Data: datadogV2.IncidentServiceUpdateData{
			Type: datadogV2.INCIDENTSERVICETYPE_SERVICES,
			Attributes: &datadogV2.IncidentServiceUpdateAttributes{
				Name: "service name-updated",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentService", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentServicesApi(apiClient)
	resp, r, err := api.UpdateIncidentService(ctx, ServiceDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentServicesApi.UpdateIncidentService`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentServicesApi.UpdateIncidentService`:\n%s\n", responseContent)
}
