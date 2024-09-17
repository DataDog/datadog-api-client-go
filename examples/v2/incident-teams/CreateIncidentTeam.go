// Create a new incident team returns "CREATED" response

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
	body := datadogV2.IncidentTeamCreateRequest{
		Data: datadogV2.IncidentTeamCreateData{
			Type: datadogV2.INCIDENTTEAMTYPE_TEAMS,
			Attributes: &datadogV2.IncidentTeamCreateAttributes{
				Name: "Example-Incident-Team",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentTeam", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentTeamsApi(apiClient)
	resp, r, err := api.CreateIncidentTeam(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.CreateIncidentTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentTeamsApi.CreateIncidentTeam`:\n%s\n", responseContent)
}
