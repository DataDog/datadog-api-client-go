// Create a detection rule with detection method "third_party" returns "OK" response

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
		Name: "Example-Create_a_detection_rule_with_detection_method_third_party_returns_OK_response",
		Queries: []datadog.SecurityMonitoringRuleQueryCreate{
			{
				Query:          "@test:true",
				Aggregation:    datadog.SECURITYMONITORINGRULEQUERYAGGREGATION_NONE.Ptr(),
				GroupByFields:  &[]string{},
				DistinctFields: &[]string{},
			},
		},
		Filters: &[]datadog.SecurityMonitoringFilter{},
		Cases: []datadog.SecurityMonitoringRuleCaseCreate{
			{
				Name:          datadog.PtrString(""),
				Status:        datadog.SECURITYMONITORINGRULESEVERITY_INFO,
				Notifications: &[]string{},
			},
		},
		Options: datadog.SecurityMonitoringRuleOptions{
			DetectionMethod:   datadog.SECURITYMONITORINGRULEDETECTIONMETHOD_THIRD_PARTY.Ptr(),
			EvaluationWindow:  datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
			KeepAlive:         datadog.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
			MaxSignalDuration: datadog.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
			ThirdPartyRuleOptions: &datadog.SecurityMonitoringRuleThirdPartyOptions{
				RootQuery:     datadog.PtrString("@pop"),
				DefaultStatus: datadog.SECURITYMONITORINGRULESEVERITY_LOW.Ptr(),
			},
		},
		Message:   "Example-Create_a_detection_rule_with_detection_method_third_party_returns_OK_response message",
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
