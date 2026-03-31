// Add a user to a team returns "Represents a user's association to a team" response

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

	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadogV2.UserTeamRequest{
		Data: datadogV2.UserTeamCreate{
			Attributes: &datadogV2.UserTeamAttributes{
				Role: *datadogV2.NewNullableUserTeamRole(datadogV2.USERTEAMROLE_ADMIN.Ptr()),
			},
			Relationships: &datadogV2.UserTeamRelationships{
				User: &datadogV2.RelationshipToUserTeamUser{
					Data: datadogV2.RelationshipToUserTeamUserData{
						Id:   UserDataID,
						Type: datadogV2.USERTEAMUSERTYPE_USERS,
					},
				},
			},
			Type: datadogV2.USERTEAMTYPE_TEAM_MEMBERSHIPS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.CreateTeamMembership(ctx, DdTeamDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateTeamMembership`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateTeamMembership`:\n%s\n", responseContent)
}
