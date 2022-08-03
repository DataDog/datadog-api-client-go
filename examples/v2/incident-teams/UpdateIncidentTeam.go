// Update an existing incident team returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncidentTeam", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewIncidentTeamsApi(apiClient)
	resp, r, err := api.UpdateIncidentTeam(ctx, TeamDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.UpdateIncidentTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentTeamsApi.UpdateIncidentTeam`:\n%s\n", responseContent)
}
