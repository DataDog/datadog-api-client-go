// Update an incident impact returns "OK" response

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
	body := datadogV2.IncidentImpactCreateRequest{
		Data: datadogV2.IncidentImpactCreateData{
			Attributes: datadogV2.IncidentImpactCreateAttributes{
				Description: "Service was unavailable for external users",
				EndAt:       *datadog.NewNullableTime(datadog.PtrTime(time.Date(2025, 8, 29, 13, 17, 0, 0, time.UTC))),
				Fields: map[string]interface{}{
					"customers_impacted": "all",
					"products_impacted":  "['shopping', 'marketing']",
				},
				StartAt: time.Date(2025, 8, 28, 13, 17, 0, 0, time.UTC),
			},
			Type: datadogV2.INCIDENTIMPACTTYPE_INCIDENT_IMPACTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentImpact", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentImpact(ctx, "incident_id", "impact_id", body, *datadogV2.NewUpdateIncidentImpactOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentImpact`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentImpact`:\n%s\n", responseContent)
}
