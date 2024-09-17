// Update an existing rule returns "OK" response

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

	body := datadogV2.SecurityMonitoringRuleUpdatePayload{
		Name: datadog.PtrString("Example-Security-Monitoring-Updated"),
		Queries: []datadogV2.SecurityMonitoringRuleQuery{
			datadogV2.SecurityMonitoringRuleQuery{
				SecurityMonitoringStandardRuleQuery: &datadogV2.SecurityMonitoringStandardRuleQuery{
					Query:          datadog.PtrString("@test:true"),
					Aggregation:    datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
					GroupByFields:  []string{},
					DistinctFields: []string{},
					Metrics:        []string{},
				}},
		},
		Filters: []datadogV2.SecurityMonitoringFilter{},
		Cases: []datadogV2.SecurityMonitoringRuleCase{
			{
				Name:          datadog.PtrString(""),
				Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO.Ptr(),
				Condition:     datadog.PtrString("a > 0"),
				Notifications: []string{},
			},
		},
		Options: &datadogV2.SecurityMonitoringRuleOptions{
			EvaluationWindow:  datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
			KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
			MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
		},
		Message:   datadog.PtrString("Test rule"),
		Tags:      []string{},
		IsEnabled: datadog.PtrBool(true),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringRule(ctx, SecurityRuleID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringRule`:\n%s\n", responseContent)
}
