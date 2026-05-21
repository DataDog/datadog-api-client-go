// Create an incident Zoom configuration returns "Created" response

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
	body := datadogV2.IncidentZoomConfigurationRequest{
		Data: datadogV2.IncidentZoomConfigurationDataRequest{
			Attributes: datadogV2.IncidentZoomConfigurationDataAttributesRequest{
				ManualMeetingCreation:   datadog.PtrBool(false),
				MeetingChatTimelineSync: datadog.PtrBool(false),
				PostMeetingSummary:      datadog.PtrBool(true),
			},
			Type: datadogV2.INCIDENTZOOMCONFIGURATIONTYPE_ZOOM_CONFIGURATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentZoomConfiguration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentZoomConfiguration(ctx, body, *datadogV2.NewCreateIncidentZoomConfigurationOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentZoomConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentZoomConfiguration`:\n%s\n", responseContent)
}
