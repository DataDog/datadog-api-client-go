// Create an incident Microsoft Teams configuration returns "Created" response

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
	body := datadogV2.IncidentMicrosoftTeamsConfigurationRequest{
		Data: datadogV2.IncidentMicrosoftTeamsConfigurationDataRequest{
			Attributes: datadogV2.IncidentMicrosoftTeamsConfigurationDataAttributesRequest{
				ManualMeetingCreation: datadog.PtrBool(false),
				PostMeetingSummary:    datadog.PtrBool(true),
			},
			Type: datadogV2.INCIDENTMICROSOFTTEAMSCONFIGURATIONTYPE_MICROSOFT_TEAMS_CONFIGURATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentMicrosoftTeamsConfiguration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentsApi(apiClient)
	resp, r, err := api.CreateIncidentMicrosoftTeamsConfiguration(ctx, body, *datadogV2.NewCreateIncidentMicrosoftTeamsConfigurationOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.CreateIncidentMicrosoftTeamsConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.CreateIncidentMicrosoftTeamsConfiguration`:\n%s\n", responseContent)
}
