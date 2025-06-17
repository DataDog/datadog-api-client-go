// Add a member team returns "Added" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.AddMemberTeamRequest{
		Data: datadogV2.MemberTeam{
			Id:   "aeadc05e-98a8-11ec-ac2c-da7ad0900001",
			Type: datadogV2.MEMBERTEAMTYPE_MEMBER_TEAMS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.AddMemberTeam", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	r, err := api.AddMemberTeam(ctx, "super_team_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.AddMemberTeam`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
