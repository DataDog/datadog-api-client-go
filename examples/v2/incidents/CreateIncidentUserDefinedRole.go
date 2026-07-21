// Create an incident user-defined role returns "Created" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.IncidentUserDefinedRoleRequest{
		Data: datadogV2.IncidentUserDefinedRoleDataRequest{
			Attributes: datadogV2.IncidentUserDefinedRoleDataAttributesRequest{
				Description: *datadog.NewNullableString(datadog.PtrString("The technical lead for the incident.")),
				Name:        "Tech Lead",
				Policy: &datadogV2.IncidentUserDefinedRolePolicy{
					IsSingle: true,
				},
			},
			Relationships: datadogV2.IncidentUserDefinedRoleRelationshipsRequest{
				IncidentType: datadogV2.IncidentUserDefinedRoleIncidentTypeRelationship{
					Data: datadogV2.IncidentUserDefinedRoleIncidentTypeRelationshipData{
						Id:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
						Type: "incident_types",
					},
				},
			},
			Type: datadogV2.INCIDENTUSERDEFINEDROLETYPE_INCIDENT_USER_DEFINED_ROLES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentUserDefinedRole", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentUserDefinedRole(ctx, body, *datadogV2.NewCreateIncidentUserDefinedRoleOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentUserDefinedRole`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentUserDefinedRole`:\n%s\n", responseContent)
}
