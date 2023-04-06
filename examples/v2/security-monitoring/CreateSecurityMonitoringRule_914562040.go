// Create a detection rule with type 'signal_correlation' returns "OK" response

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
	// there is a valid "security_rule" in the system
	SecurityRuleID := os.Getenv("SECURITY_RULE_ID")

	// there is a valid "security_rule_bis" in the system
	SecurityRuleBisID := os.Getenv("SECURITY_RULE_BIS_ID")

	body := datadogV2.SecurityMonitoringRuleCreatePayload{
		SecurityMonitoringSignalRuleCreatePayload: &datadogV2.SecurityMonitoringSignalRuleCreatePayload{
			Name: "Example-Security-Monitoring_signal_rule",
			Queries: []datadogV2.SecurityMonitoringSignalRuleQuery{
				{
					RuleId:      SecurityRuleID,
					Aggregation: datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_EVENT_COUNT.Ptr(),
					CorrelatedByFields: []string{
						"host",
					},
					CorrelatedQueryIndex: datadog.PtrInt32(1),
				},
				{
					RuleId:      SecurityRuleBisID,
					Aggregation: datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_EVENT_COUNT.Ptr(),
					CorrelatedByFields: []string{
						"host",
					},
				},
			},
			Filters: []datadogV2.SecurityMonitoringFilter{},
			Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
				{
					Name:          datadog.PtrString(""),
					Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
					Condition:     datadog.PtrString("a > 0 && b > 0"),
					Notifications: []string{},
				},
			},
			Options: datadogV2.SecurityMonitoringRuleOptions{
				EvaluationWindow:  datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
				KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
				MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
			},
			Message:   "Test signal correlation rule",
			Tags:      []string{},
			IsEnabled: true,
			Type:      datadogV2.SECURITYMONITORINGSIGNALRULETYPE_SIGNAL_CORRELATION.Ptr(),
		}}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringRule`:\n%s\n", responseContent)
}
