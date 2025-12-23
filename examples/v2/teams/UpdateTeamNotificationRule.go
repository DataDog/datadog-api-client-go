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
	body := datadogV2.TeamNotificationRule{
		Attributes: &datadogV2.TeamNotificationRuleAttributes{
			Email:     &datadogV2.TeamNotificationRuleAttributesEmail{},
			MsTeams:   &datadogV2.TeamNotificationRuleAttributesMsTeams{},
			Pagerduty: &datadogV2.TeamNotificationRuleAttributesPagerduty{},
			Slack:     &datadogV2.TeamNotificationRuleAttributesSlack{},
		},
		Id: datadog.PtrString("b8626d7e-cedd-11eb-abf5-da7ad0900001"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTeamsApi(apiClient)
	resp, r, err := api.UpdateTeamNotificationRule(ctx, "rule_id", "team_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.UpdateTeamNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TeamsApi.UpdateTeamNotificationRule`:\n%s\n", responseContent)
}
