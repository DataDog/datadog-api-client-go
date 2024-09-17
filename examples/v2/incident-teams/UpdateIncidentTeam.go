// Update an existing incident team returns "OK" response

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
	// there is a valid "team" in the system
	TeamDataID := os.Getenv("TEAM_DATA_ID")

	body := datadogV2.IncidentTeamUpdateRequest{
		Data: datadogV2.IncidentTeamUpdateData{
			Type: datadogV2.INCIDENTTEAMTYPE_TEAMS,
			Attributes: &datadogV2.IncidentTeamUpdateAttributes{
				Name: "team name-updated",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentTeam", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewIncidentTeamsApi(apiClient)
	resp, r, err := api.UpdateIncidentTeam(ctx, TeamDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.UpdateIncidentTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentTeamsApi.UpdateIncidentTeam`:\n%s\n", responseContent)
}
