// Create global incident handle returns "Created" response

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
	body := datadogV2.IncidentHandleRequest{
		Data: datadogV2.IncidentHandleDataRequest{
			Attributes: datadogV2.IncidentHandleAttributesRequest{
				Fields: &datadogV2.IncidentHandleAttributesFields{
					Severity: []string{
						"SEV-1",
					},
				},
				Name: "@incident-sev-1",
			},
			Id: datadog.PtrString("b2494081-cdf0-4205-b366-4e1dd4fdf0bf"),
			Relationships: *datadogV2.NewNullableIncidentHandleRelationshipsRequest(&datadogV2.IncidentHandleRelationshipsRequest{
				CommanderUser: &datadogV2.IncidentHandleRelationship{
					Data: datadogV2.IncidentHandleRelationshipData{
						Id:   "f7b538b1-ed7c-4e84-82de-fdf84a539d40",
						Type: "incident_types",
					},
				},
				IncidentType: datadogV2.IncidentHandleRelationship{
					Data: datadogV2.IncidentHandleRelationshipData{
						Id:   "f7b538b1-ed7c-4e84-82de-fdf84a539d40",
						Type: "incident_types",
					},
				},
			}),
			Type: datadogV2.INCIDENTHANDLETYPE_INCIDENTS_HANDLES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateGlobalIncidentHandle", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateGlobalIncidentHandle(ctx, body, *datadogV2.NewCreateGlobalIncidentHandleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateGlobalIncidentHandle`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateGlobalIncidentHandle`:\n%s\n", responseContent)
}
