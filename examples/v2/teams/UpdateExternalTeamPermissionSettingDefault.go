// Update an external team permission default returns "OK" response

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
	body := datadogV2.ExternalTeamPermissionSettingDefaultUpdateRequest{
		Data: datadogV2.ExternalTeamPermissionSettingDefaultUpdateData{
			Attributes: datadogV2.ExternalTeamPermissionSettingDefaultUpdateAttributes{
				Value: "user_access_manage",
			},
			Id:   "scim:add_and_remove_subteam",
			Type: datadogV2.EXTERNALTEAMPERMISSIONSETTINGDEFAULTSTYPE_EXTERNAL_TEAM_PERMISSION_SETTING_DEFAULTS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.UpdateExternalTeamPermissionSettingDefault(ctx, datadogV2.EXTERNALTEAMPERMISSIONSETTINGDEFAULTSSOURCE_SCIM, "add_and_remove_subteam", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateExternalTeamPermissionSettingDefault`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateExternalTeamPermissionSettingDefault`:\n%s\n", responseContent)
}
