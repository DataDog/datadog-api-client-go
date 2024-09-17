// Update permission setting for team returns "OK" response

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

	body := datadogV2.TeamPermissionSettingUpdateRequest{
		Data: datadogV2.TeamPermissionSettingUpdate{
			Attributes: &datadogV2.TeamPermissionSettingUpdateAttributes{
				Value: datadogV2.TEAMPERMISSIONSETTINGVALUE_ADMINS.Ptr(),
			},
			Type: datadogV2.TEAMPERMISSIONSETTINGTYPE_TEAM_PERMISSION_SETTINGS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.UpdateTeamPermissionSetting(ctx, DdTeamDataID, "manage_membership", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeamPermissionSetting`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeamPermissionSetting`:\n%s\n", responseContent)
}
