// Update an existing incident team returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "team" in the system
	TeamDataID := os.Getenv("TEAM_DATA_ID")

	body := datadog.IncidentTeamUpdateRequest{
		Data: datadog.IncidentTeamUpdateData{
			Type: datadog.INCIDENTTEAMTYPE_TEAMS,
			Attributes: &datadog.IncidentTeamUpdateAttributes{
				Name: "team name-updated",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("UpdateIncidentTeam", true)
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentTeamsApi.UpdateIncidentTeam(ctx, TeamDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.UpdateIncidentTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentTeamsApi.UpdateIncidentTeam`:\n%s\n", responseContent)
}
