// Update a notification rule returns "OK" response

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
	// there is a valid "notification_rule" in the system
	NotificationRuleDataID := os.Getenv("NOTIFICATION_RULE_DATA_ID")

	body := datadogV2.SecurityMonitoringNotificationRuleUpdateRequest{
		Data: datadogV2.SecurityMonitoringNotificationRuleUpdateData{
			Attributes: datadogV2.SecurityMonitoringNotificationRuleUpdateAttributes{
				Version: 1,
				Name:    "Example-Security-Monitoring",
				Enabled: false,
				Selectors: datadogV2.SecurityMonitoringNotificationRuleSelectors{
					Attributes: []string{
						"fim:true",
					},
					ImplicitVmRuleMatch: false,
					RuleTags: []string{
						"fim:true",
					},
					Severities: []datadogV2.SecurityMonitoringRuleSeverity{
						datadogV2.SECURITYMONITORINGRULESEVERITY_HIGH,
					},
					SignalTags: []string{},
					RuleTypes: []datadogV2.SecurityMonitoringRuleTypes{
						datadogV2.SECURITYMONITORINGRULETYPES_LOG_DETECTION,
						datadogV2.SECURITYMONITORINGRULETYPES_CLOUD_CONFIGURATION,
					},
				},
				Targets: []string{
					"test2",
				},
			},
			Type: datadogV2.SECURITYMONITORINGNOTIFICATIONRULETYPE_NOTIFICATION_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.EditSecurityMonitoringNotificationRule(ctx, NotificationRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.EditSecurityMonitoringNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.EditSecurityMonitoringNotificationRule`:\n%s\n", responseContent)
}
