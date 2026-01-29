// Create a notification rule returns "CREATED" response

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
	body := datadogV2.CaseNotificationRuleCreateRequest{
		Data: datadogV2.CaseNotificationRuleCreate{
			Attributes: datadogV2.CaseNotificationRuleCreateAttributes{
				IsEnabled: datadog.PtrBool(true),
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
	resp, r, err := api.CreateProjectNotificationRule(ctx, "project_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateProjectNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.CreateProjectNotificationRule`:\n%s\n", responseContent)
}
