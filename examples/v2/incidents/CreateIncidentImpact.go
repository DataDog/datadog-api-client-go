// Create an incident impact returns "CREATED" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "incident" in the system
	IncidentDataID := os.Getenv("INCIDENT_DATA_ID")

	body := datadogV2.IncidentImpactCreateRequest{
		Data: datadogV2.IncidentImpactCreateData{
			Type: datadogV2.INCIDENTIMPACTTYPE_INCIDENT_IMPACTS,
			Attributes: datadogV2.IncidentImpactCreateAttributes{
				StartAt:     time.Date(2025, 9, 12, 13, 50, 0, 0, time.UTC),
				EndAt:       *datadog.NewNullableTime(datadog.PtrTime(time.Date(2025, 9, 12, 14, 50, 0, 0, time.UTC))),
				Description: "Outage in the us-east-1 region",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentImpact", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentImpact(ctx, IncidentDataID, body, *datadogV2.NewCreateIncidentImpactOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentImpact`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentImpact`:\n%s\n", responseContent)
}
