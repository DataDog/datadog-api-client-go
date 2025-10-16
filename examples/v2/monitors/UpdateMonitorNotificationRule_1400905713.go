// Update a monitor notification rule with conditional_recipients returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	// there is a valid "monitor_notification_rule" in the system
	MonitorNotificationRuleDataID := os.Getenv("MONITOR_NOTIFICATION_RULE_DATA_ID")


	body := datadogV2.MonitorNotificationRuleUpdateRequest{
Data: datadogV2.MonitorNotificationRuleUpdateRequestData{
Attributes: datadogV2.MonitorNotificationRuleAttributes{
Filter: &datadogV2.MonitorNotificationRuleFilter{
MonitorNotificationRuleFilterTags: &datadogV2.MonitorNotificationRuleFilterTags{
Tags: []string{
"test:example-monitor",
"host:abc",
},
}},
Name: "updated rule",
ConditionalRecipients: &datadogV2.MonitorNotificationRuleConditionalRecipients{
Conditions: []datadogV2.MonitorNotificationRuleCondition{
{
Scope: "transition_type:is_alert",
Recipients: []string{
"slack-test-channel",
"jira-test",
},
},
},
},
},
Id: MonitorNotificationRuleDataID,
Type: datadogV2.MONITORNOTIFICATIONRULERESOURCETYPE_MONITOR_NOTIFICATION_RULE.Ptr(),
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	resp, r, err := api.UpdateMonitorNotificationRule(ctx, MonitorNotificationRuleDataID, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.UpdateMonitorNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.UpdateMonitorNotificationRule`:\n%s\n", responseContent)
}