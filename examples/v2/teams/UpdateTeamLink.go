// Update a team link returns "OK" response

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

	// there is a valid "team_link" in the system
	TeamLinkDataID := os.Getenv("TEAM_LINK_DATA_ID")

	body := datadogV2.TeamLinkCreateRequest{
		Data: datadogV2.TeamLinkCreate{
			Attributes: datadogV2.TeamLinkAttributes{
				Label: "New Label",
				Url:   "https://example.com",
			},
			Type: datadogV2.TEAMLINKTYPE_TEAM_LINKS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.UpdateTeamLink(ctx, DdTeamDataID, TeamLinkDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeamLink`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeamLink`:\n%s\n", responseContent)
}
