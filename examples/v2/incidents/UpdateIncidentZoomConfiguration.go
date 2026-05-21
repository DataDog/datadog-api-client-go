// Update an incident Zoom configuration returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentZoomConfiguration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncidentZoomConfiguration(ctx, uuid.MustParse("9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d"), body, *datadogV2.NewUpdateIncidentZoomConfigurationOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncidentZoomConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncidentZoomConfiguration`:\n%s\n", responseContent)
}
