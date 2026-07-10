// Update postmortem for an incident returns "OK" response

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
	// there is a valid "postmortem" in the system
	PostmortemDataID := os.Getenv("POSTMORTEM_DATA_ID")
	PostmortemDataRelationshipsIncidentDataID := os.Getenv("POSTMORTEM_DATA_RELATIONSHIPS_INCIDENT_DATA_ID")

	body := datadogV2.IncidentPostmortemUpdateRequest{
		Data: datadogV2.IncidentPostmortemUpdateData{
			Attributes: datadogV2.IncidentPostmortemUpdateAttributes{
				Status: datadogV2.POSTMORTEMSTATUS_IN_REVIEW.Ptr(),
			},
			Id:   PostmortemDataID,
			Type: datadogV2.INCIDENTPOSTMORTEMTYPE_INCIDENT_POSTMORTEMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentPostmortem", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentPostmortem(ctx, PostmortemDataRelationshipsIncidentDataID, body, *datadogV2.NewUpdateIncidentPostmortemOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentPostmortem`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentPostmortem`:\n%s\n", responseContent)
}
