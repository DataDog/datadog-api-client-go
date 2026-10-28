// Create an incident responder returns "Created" response

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
	body := datadogV2.IncidentResponderRequest{
		Data: datadogV2.IncidentResponderDataRequest{
			Relationships: datadogV2.IncidentResponderRelationshipsRequest{
				User: datadogV2.IncidentResponderUserRelationship{
					Data: datadogV2.IncidentResponderUserRelationshipData{
						Id:   uuid.MustParse("00000000-0000-0000-0000-000000000000"),
						Type: "users",
					},
				},
			},
			Type: datadogV2.INCIDENTRESPONDERTYPE_INCIDENT_RESPONDERS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentResponder", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentResponder(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentResponder`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentResponder`:\n%s\n", responseContent)
}
