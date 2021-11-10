// Create a detection rule returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityMonitoringRuleCreatePayload{
		Name: "Example-Create_a_detection_rule_returns_OK_response",
		Queries: []datadog.SecurityMonitoringRuleQueryCreate{
			datadog.SecurityMonitoringRuleQueryCreate{
				Query:          "@test:true",
				Aggregation:    datadog.SecurityMonitoringRuleQueryAggregation("count").Ptr(),
				GroupByFields:  &[]string{},
				DistinctFields: &[]string{},
				Metric:         datadog.PtrString(""),
			},
		},
		Filters: &[]datadog.SecurityMonitoringFilter{},
		Cases: []datadog.SecurityMonitoringRuleCaseCreate{
			datadog.SecurityMonitoringRuleCaseCreate{
				Name:          datadog.PtrString(""),
				Status:        datadog.SecurityMonitoringRuleSeverity("info"),
				Condition:     datadog.PtrString("a > 0"),
				Notifications: &[]string{},
			},
		},
		Options: datadog.SecurityMonitoringRuleOptions{
			EvaluationWindow:  datadog.SecurityMonitoringRuleEvaluationWindow(900).Ptr(),
			KeepAlive:         datadog.SecurityMonitoringRuleKeepAlive(3600).Ptr(),
			MaxSignalDuration: datadog.SecurityMonitoringRuleMaxSignalDuration(86400).Ptr(),
		},
		Message:   "Test rule",
		Tags:      &[]string{},
		IsEnabled: true,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityMonitoringApi.CreateSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringRule`:\n%s\n", responseContent)
}
