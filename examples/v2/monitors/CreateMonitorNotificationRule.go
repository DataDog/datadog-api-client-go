// Create a monitor notification rule returns "OK" response

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
	body := datadogV2.MonitorNotificationRuleCreateRequest{
Data: datadogV2.MonitorNotificationRuleCreateRequestData{
Attributes: datadogV2.MonitorNotificationRuleAttributes{
Filter: &datadogV2.MonitorNotificationRuleFilter{
MonitorNotificationRuleFilterTags: &datadogV2.MonitorNotificationRuleFilterTags{
Tags: []string{
"test:example-monitor",
},
}},
Name: "test rule",
Recipients: []string{
"slack-test-channel",
"jira-test",
},
},
Type: datadogV2.MONITORNOTIFICATIONRULERESOURCETYPE_MONITOR_NOTIFICATION_RULE.Ptr(),
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	resp, r, err := api.CreateMonitorNotificationRule(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.CreateMonitorNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MonitorsApi.CreateMonitorNotificationRule`:\n%s\n", responseContent)
}