// Create an incident returns "CREATED" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
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
							Value: *common.NewNullableString(common.PtrString("resolved")),
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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncident", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncident(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncident`:\n%s\n", responseContent)
}
