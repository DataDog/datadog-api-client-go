// Create a notification rule returns "OK" response

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
	body := datadogV2.SecurityMonitoringNotificationRuleCreateRequest{
		Data: datadogV2.SecurityMonitoringNotificationRuleCreateData{
			Attributes: datadogV2.SecurityMonitoringNotificationRuleCreateAttributes{
				Enabled: true,
				Name:    "Example-Security-Monitoring",
				Selectors: datadogV2.SecurityMonitoringNotificationRuleSelectors{
					Attributes: []string{
						"test:123",
						"env:prod",
					},
					ImplicitVmRuleMatch: false,
					RuleTags: []string{
						"test:123",
					},
					RuleTypes: []datadogV2.SecurityMonitoringRuleTypes{
						datadogV2.SECURITYMONITORINGRULETYPES_APPLICATION_SECURITY,
						datadogV2.SECURITYMONITORINGRULETYPES_LOG_DETECTION,
					},
					Severities: []datadogV2.SecurityMonitoringRuleSeverity{
						datadogV2.SECURITYMONITORINGRULESEVERITY_HIGH,
					},
					SignalTags: []string{
						"env:prod",
					},
				},
				Targets: []string{
					"@slack-test",
				},
			},
			Type: datadogV2.SECURITYMONITORINGNOTIFICATIONRULETYPE_NOTIFICATION_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringNotificationRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringNotificationRule`:\n%s\n", responseContent)
}
