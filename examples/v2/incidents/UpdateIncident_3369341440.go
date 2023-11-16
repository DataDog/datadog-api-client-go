// Add commander to an incident returns "OK" response

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

	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadogV2.IncidentUpdateRequest{
		Data: datadogV2.IncidentUpdateData{
			Id:   IncidentDataID,
			Type: datadogV2.INCIDENTTYPE_INCIDENTS,
			Relationships: &datadogV2.IncidentUpdateRelationships{
				CommanderUser: *datadogV2.NewNullableNullableRelationshipToUser(&datadogV2.NullableRelationshipToUser{
					Data: *datadogV2.NewNullableNullableRelationshipToUserData(&datadogV2.NullableRelationshipToUserData{
						Id:   UserDataID,
						Type: datadogV2.USERSTYPE_USERS,
					}),
				}),
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
