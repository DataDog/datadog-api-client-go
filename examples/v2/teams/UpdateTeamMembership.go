// Update a user's membership attributes on a team returns "Represents a user's association to a team" response

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
	body := datadogV2.UserTeamUpdateRequest{
		Data: datadogV2.UserTeamUpdate{
			Attributes: &datadogV2.UserTeamAttributes{
				Role: *datadogV2.NewNullableUserTeamRole(datadogV2.USERTEAMROLE_ADMIN.Ptr()),
			},
			Type: datadogV2.USERTEAMTYPE_TEAM_MEMBERSHIPS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.UpdateTeamMembership(ctx, "team_id", "user_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeamMembership`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeamMembership`:\n%s\n", responseContent)
}
