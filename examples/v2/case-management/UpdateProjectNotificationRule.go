// Update a notification rule returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.CaseNotificationRuleUpdateRequest{
		Data: datadogV2.CaseNotificationRuleUpdate{
			Attributes: &datadogV2.CaseNotificationRuleAttributes{
				Recipients: []datadogV2.CaseNotificationRuleRecipient{
					{
						Data: &datadogV2.CaseNotificationRuleRecipientData{},
						Type: datadog.PtrString("EMAIL"),
					},
				},
				Triggers: []datadogV2.CaseNotificationRuleTrigger{
					{
						Data: &datadogV2.CaseNotificationRuleTriggerData{},
						Type: datadog.PtrString("CASE_CREATED"),
					},
				},
			},
			Type: datadogV2.CASENOTIFICATIONRULERESOURCETYPE_NOTIFICATION_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	r, err := api.UpdateProjectNotificationRule(ctx, "project_id", "notification_rule_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.UpdateProjectNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
