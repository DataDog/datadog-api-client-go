// Create an incident communication returns "Created" response

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
	body := datadogV2.IncidentCommunicationRequest{
		Data: datadogV2.IncidentCommunicationDataRequest{
			Attributes: datadogV2.IncidentCommunicationDataAttributesRequest{
				CommunicationType: datadogV2.INCIDENTCOMMUNICATIONKIND_MANUAL,
				Content: datadogV2.IncidentCommunicationContent{
					GroupingKey: datadog.PtrString("update-1"),
					Handles: []datadogV2.IncidentCommunicationContentHandle{
						{
							CreatedAt:   datadog.PtrString("2024-01-01T00:00:00.000Z"),
							DisplayName: datadog.PtrString("#incidents-channel"),
							Handle:      "@slack-incidents-channel",
						},
					},
					Message: "Incident update for INC-123.",
					Status:  datadog.PtrInt32(0),
					Subject: datadog.PtrString("Incident INC-123: Update"),
				},
			},
			Type: datadogV2.INCIDENTCOMMUNICATIONTYPE_COMMUNICATION,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentCommunication", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentCommunication(ctx, "incident_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentCommunication`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentCommunication`:\n%s\n", responseContent)
}
