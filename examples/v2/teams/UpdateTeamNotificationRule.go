// Update team notification rule returns "OK" response

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

	// there is a valid "team_notification_rule" in the system
	TeamNotificationRuleDataID := os.Getenv("TEAM_NOTIFICATION_RULE_DATA_ID")

	body := datadogV2.TeamNotificationRuleRequest{
		Data: datadogV2.TeamNotificationRule{
			Type: datadogV2.TEAMNOTIFICATIONRULETYPE_TEAM_NOTIFICATION_RULES,
			Id:   datadog.PtrString(TeamNotificationRuleDataID),
			Attributes: datadogV2.TeamNotificationRuleAttributes{
				Pagerduty: &datadogV2.TeamNotificationRuleAttributesPagerduty{
					ServiceName: datadog.PtrString("Datadog-prod"),
				},
				Slack: &datadogV2.TeamNotificationRuleAttributesSlack{
					Workspace: datadog.PtrString("Datadog"),
					Channel:   datadog.PtrString("aaa-governance-ops"),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.UpdateTeamNotificationRule(ctx, DdTeamDataID, TeamNotificationRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeamNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeamNotificationRule`:\n%s\n", responseContent)
}
