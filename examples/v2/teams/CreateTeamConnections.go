// Create team connections returns "Created" response

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
	// there is a valid "dd_team" in the system
	DdTeamDataID := os.Getenv("DD_TEAM_DATA_ID")

	body := datadogV2.TeamConnectionCreateRequest{
		Data: []datadogV2.TeamConnectionCreateData{
			{
				Type: datadogV2.TEAMCONNECTIONTYPE_TEAM_CONNECTION,
				Attributes: &datadogV2.TeamConnectionAttributes{
					Source:    datadog.PtrString("github"),
					ManagedBy: datadog.PtrString("datadog"),
				},
				Relationships: &datadogV2.TeamConnectionRelationships{
					Team: &datadogV2.TeamRef{
						Data: &datadogV2.TeamRefData{
							Id:   DdTeamDataID,
							Type: datadogV2.TEAMREFDATATYPE_TEAM,
						},
					},
					ConnectedTeam: &datadogV2.ConnectedTeamRef{
						Data: &datadogV2.ConnectedTeamRefData{
							Id:   "@MyGitHubAccount/my-team-name",
							Type: datadogV2.CONNECTEDTEAMREFDATATYPE_GITHUB_TEAM,
						},
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateTeamConnections", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.CreateTeamConnections(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateTeamConnections`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateTeamConnections`:\n%s\n", responseContent)
}
