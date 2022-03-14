// Update an existing incident returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "incident" in the system
	IncidentDataID := os.Getenv("INCIDENT_DATA_ID")

	body := datadog.IncidentUpdateRequest{
		Data: datadog.IncidentUpdateData{
			Id:   IncidentDataID,
			Type: datadog.INCIDENTTYPE_INCIDENTS,
			Attributes: &datadog.IncidentUpdateAttributes{
				Fields: map[string]datadog.IncidentFieldAttributes{
					"state": datadog.IncidentFieldAttributes{
						IncidentFieldAttributesSingleValue: &datadog.IncidentFieldAttributesSingleValue{
							Type:  datadog.INCIDENTFIELDATTRIBUTESSINGLEVALUETYPE_DROPDOWN.Ptr(),
							Value: *datadog.NewNullableString(datadog.PtrString("resolved")),
						}},
				},
				Title: datadog.PtrString("A test incident title-updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("UpdateIncident", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsApi.UpdateIncident(ctx, IncidentDataID, body, *datadog.NewUpdateIncidentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncident`:\n%s\n", responseContent)
}
