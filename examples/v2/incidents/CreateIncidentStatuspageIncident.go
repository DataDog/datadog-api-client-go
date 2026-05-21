// Create a Statuspage incident for an incident returns "OK" response

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
	body := datadogV2.IncidentStatuspageIncidentRequest{
		Data: datadogV2.IncidentStatuspageIncidentDataRequest{
			Attributes: datadogV2.IncidentStatuspageIncidentDataAttributesRequest{
				Body:                 *datadog.NewNullableString(datadog.PtrString("We are investigating the issue.")),
				DeliverNotifications: *datadog.NewNullableBool(datadog.PtrBool(true)),
				Impact:               *datadog.NewNullableString(datadog.PtrString("major")),
				Name:                 *datadog.NewNullableString(datadog.PtrString("Service Outage")),
				PageId:               datadog.PtrString("abc123"),
				Status:               *datadog.NewNullableString(datadog.PtrString("investigating")),
			},
			Type: datadogV2.INCIDENTSTATUSPAGEINCIDENTTYPE_INCIDENT_INTEGRATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentStatuspageIncident", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentStatuspageIncident(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentStatuspageIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentStatuspageIncident`:\n%s\n", responseContent)
}
