// Patch a signal-based notification rule returns "Notification rule successfully patched." response

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
	// there is a valid "valid_signal_notification_rule" in the system
	ValidSignalNotificationRuleDataID := os.Getenv("VALID_SIGNAL_NOTIFICATION_RULE_DATA_ID")

	body := datadogV2.PatchNotificationRuleParameters{
		Data: &datadogV2.PatchNotificationRuleParametersData{
			Attributes: datadogV2.PatchNotificationRuleParametersDataAttributes{
				Enabled: datadog.PtrBool(true),
				Name:    datadog.PtrString("Rule 1"),
				Selectors: &datadogV2.Selectors{
					Query: datadog.PtrString("(source:production_service OR env:prod)"),
					RuleTypes: []datadogV2.RuleTypesItems{
						datadogV2.RULETYPESITEMS_MISCONFIGURATION,
						datadogV2.RULETYPESITEMS_ATTACK_PATH,
					},
					Severities: []datadogV2.RuleSeverity{
						datadogV2.RULESEVERITY_CRITICAL,
					},
					TriggerSource: datadogV2.TRIGGERSOURCE_SECURITY_FINDINGS,
				},
				Targets: []string{
					"@john.doe@email.com",
				},
				TimeAggregation: datadog.PtrInt64(86400),
				Version:         datadog.PtrInt64(1),
			},
			Id:   ValidSignalNotificationRuleDataID,
			Type: datadogV2.NOTIFICATIONRULESTYPE_NOTIFICATION_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.PatchSignalNotificationRule(ctx, ValidSignalNotificationRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.PatchSignalNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.PatchSignalNotificationRule`:\n%s\n", responseContent)
}
