// Create an incident role assignment returns "Created" response

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
	body := datadogV2.IncidentRoleAssignmentRequest{
		Data: datadogV2.IncidentRoleAssignmentDataRequest{
			Attributes: &datadogV2.IncidentRoleAssignmentDataAttributesRequest{
				Role: *datadog.NewNullableString(datadog.PtrString("commander")),
			},
			Relationships: datadogV2.IncidentRoleAssignmentRelationshipsRequest{
				ReservedRole: &datadogV2.IncidentRoleAssignmentRoleRelationship{
					Data: &datadogV2.IncidentRoleAssignmentRoleRelationshipData{
						Id:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
						Type: "incident_reserved_roles",
					},
				},
				Responder: datadogV2.IncidentRoleAssignmentResponderRelationship{
					Data: datadogV2.IncidentRoleAssignmentResponderRelationshipData{
						Id:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
						Type: "users",
					},
				},
				UserDefinedRole: &datadogV2.IncidentRoleAssignmentRoleRelationship{
					Data: &datadogV2.IncidentRoleAssignmentRoleRelationshipData{
						Id:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
						Type: "incident_reserved_roles",
					},
				},
			},
			Type: datadogV2.INCIDENTROLEASSIGNMENTTYPE_INCIDENT_ROLE_ASSIGNMENTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentRoleAssignment", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentRoleAssignment(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentRoleAssignment`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentRoleAssignment`:\n%s\n", responseContent)
}
