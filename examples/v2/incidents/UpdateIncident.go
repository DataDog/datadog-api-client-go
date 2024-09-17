// Update an existing incident returns "OK" response

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
	// there is a valid "incident" in the system
	IncidentDataID := os.Getenv("INCIDENT_DATA_ID")

	body := datadogV2.IncidentUpdateRequest{
		Data: datadogV2.IncidentUpdateData{
			Id:   IncidentDataID,
			Type: datadogV2.INCIDENTTYPE_INCIDENTS,
			Attributes: &datadogV2.IncidentUpdateAttributes{
				Fields: map[string]datadogV2.IncidentFieldAttributes{
					"state": datadogV2.IncidentFieldAttributes{
						IncidentFieldAttributesSingleValue: &datadogV2.IncidentFieldAttributesSingleValue{
							Type:  datadogV2.INCIDENTFIELDATTRIBUTESSINGLEVALUETYPE_DROPDOWN.Ptr(),
							Value: *datadog.NewNullableString(datadog.PtrString("resolved")),
						}},
				},
				Title: datadog.PtrString("A test incident title-updated"),
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncident", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncident(ctx, IncidentDataID, body, *datadogV2.NewUpdateIncidentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncident`:\n%s\n", responseContent)
}
