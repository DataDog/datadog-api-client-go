// Create postmortem for an incident returns "CREATED" response

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

	body := datadogV2.IncidentPostmortemCreateRequest{
		Data: datadogV2.IncidentPostmortemCreateData{
			Attributes: datadogV2.IncidentPostmortemCreateAttributes{
				DocumentUrl: "https://app.datadoghq.com/notebook/123",
				Title:       "Postmortem for IR-123",
			},
			Type: datadogV2.INCIDENTPOSTMORTEMTYPE_INCIDENT_POSTMORTEMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentPostmortem", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentPostmortem(ctx, IncidentDataID, body, *datadogV2.NewCreateIncidentPostmortemOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentPostmortem`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentPostmortem`:\n%s\n", responseContent)
}
