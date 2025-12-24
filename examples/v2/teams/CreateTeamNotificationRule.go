// Create team notification rule returns "OK" response

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

	body := datadogV2.TeamNotificationRuleRequest{
		Data: datadogV2.TeamNotificationRule{
			Type: datadogV2.TEAMNOTIFICATIONRULETYPE_TEAM_NOTIFICATION_RULES,
			Attributes: datadogV2.TeamNotificationRuleAttributes{
				Email: &datadogV2.TeamNotificationRuleAttributesEmail{
					Enabled: datadog.PtrBool(true),
				},
				Slack: &datadogV2.TeamNotificationRuleAttributesSlack{
					Workspace: datadog.PtrString("Datadog"),
					Channel:   datadog.PtrString("aaa-omg-ops"),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.CreateTeamNotificationRule(ctx, DdTeamDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.CreateTeamNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.CreateTeamNotificationRule`:\n%s\n", responseContent)
}
