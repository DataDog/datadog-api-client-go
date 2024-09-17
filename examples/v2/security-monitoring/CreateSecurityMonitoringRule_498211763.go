// Create a detection rule with type 'workload_security' returns "OK" response

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
	body := datadogV2.SecurityMonitoringRuleCreatePayload{
		SecurityMonitoringStandardRuleCreatePayload: &datadogV2.SecurityMonitoringStandardRuleCreatePayload{
			Name: "Example-Security-Monitoring",
			Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{
				{
					Query:          datadog.PtrString("@test:true"),
					Aggregation:    datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
					GroupByFields:  []string{},
					DistinctFields: []string{},
					Metric:         datadog.PtrString(""),
				},
			},
			Filters: []datadogV2.SecurityMonitoringFilter{},
			Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
				{
					Name:          datadog.PtrString(""),
					Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
					Condition:     datadog.PtrString("a > 0"),
					Notifications: []string{},
				},
			},
			Options: datadogV2.SecurityMonitoringRuleOptions{
				EvaluationWindow:  datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
				KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
				MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
			},
			Message:   "Test rule",
			Tags:      []string{},
			IsEnabled: true,
			Type:      datadogV2.SECURITYMONITORINGRULETYPECREATE_WORKLOAD_SECURITY.Ptr(),
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
