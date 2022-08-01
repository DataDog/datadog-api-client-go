// Create a new incident team returns "CREATED" response

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
	body := datadog.IncidentTeamCreateRequest{
		Data: datadog.IncidentTeamCreateData{
			Type: datadog.INCIDENTTEAMTYPE_TEAMS,
			Attributes: &datadog.IncidentTeamCreateAttributes{
				Name: "Example-Create_a_new_incident_team_returns_CREATED_response",
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateIncidentTeam", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewIncidentTeamsApi(apiClient)
	resp, r, err := api.CreateIncidentTeam(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentTeamsApi.CreateIncidentTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentTeamsApi.CreateIncidentTeam`:\n%s\n", responseContent)
}
