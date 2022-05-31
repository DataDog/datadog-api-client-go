// Create an incident returns "CREATED" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadog.IncidentCreateRequest{
		Data: datadog.IncidentCreateData{
			Type: datadog.INCIDENTTYPE_INCIDENTS,
			Attributes: datadog.IncidentCreateAttributes{
				Title:            "Example-Create_an_incident_returns_CREATED_response",
				CustomerImpacted: false,
				Fields: map[string]datadog.IncidentFieldAttributes{
					"state": datadog.IncidentFieldAttributes{
						IncidentFieldAttributesSingleValue: &datadog.IncidentFieldAttributesSingleValue{
							Type:  datadog.INCIDENTFIELDATTRIBUTESSINGLEVALUETYPE_DROPDOWN.Ptr(),
							Value: *datadog.NewNullableString(datadog.PtrString("resolved")),
						}},
				},
			},
			Relationships: &datadog.IncidentCreateRelationships{
				CommanderUser: datadog.NullableRelationshipToUser{
					Data: *datadog.NewNullableNullableRelationshipToUserData(&datadog.NullableRelationshipToUserData{
						Type: datadog.USERSTYPE_USERS,
						Id:   UserDataID,
					}),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("CreateIncident", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsApi.CreateIncident(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncident`:\n%s\n", responseContent)
}
