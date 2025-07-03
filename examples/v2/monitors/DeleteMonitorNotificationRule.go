// Delete a monitor notification rule returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "monitor_notification_rule" in the system
	MonitorNotificationRuleDataID := os.Getenv("MONITOR_NOTIFICATION_RULE_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMonitorsApi(apiClient)
	r, err := api.DeleteMonitorNotificationRule(ctx, MonitorNotificationRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.DeleteMonitorNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
