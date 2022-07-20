// Update an existing rule returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityMonitoringRuleUpdatePayload{
		Cases: []datadog.SecurityMonitoringRuleCase{
			{
				Notifications: []string{},
				Status:        datadog.SECURITYMONITORINGRULESEVERITY_CRITICAL.Ptr(),
			},
		},
		Filters: []datadog.SecurityMonitoringFilter{
			{
				Action: datadog.SECURITYMONITORINGFILTERACTION_REQUIRE.Ptr(),
			},
		},
		HasExtendedTitle: datadog.PtrBool(true),
		Options: &datadog.SecurityMonitoringRuleOptions{
			DecreaseCriticalityBasedOnEnv: datadog.PtrBool(false),
			DetectionMethod:               datadog.SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD.Ptr(),
			EvaluationWindow:              datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
			HardcodedEvaluatorType:        datadog.SECURITYMONITORINGRULEHARDCODEDEVALUATORTYPE_LOG4SHELL.Ptr(),
			ImpossibleTravelOptions: &datadog.SecurityMonitoringRuleImpossibleTravelOptions{
				BaselineUserLocations: datadog.PtrBool(true),
			},
			KeepAlive:         datadog.SECURITYMONITORINGRULEKEEPALIVE_ZERO_MINUTES.Ptr(),
			MaxSignalDuration: datadog.SECURITYMONITORINGRULEMAXSIGNALDURATION_ZERO_MINUTES.Ptr(),
			NewValueOptions: &datadog.SecurityMonitoringRuleNewValueOptions{
				ForgetAfter:       datadog.SECURITYMONITORINGRULENEWVALUEOPTIONSFORGETAFTER_ONE_DAY.Ptr(),
				LearningDuration:  datadog.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ZERO_DAYS.Ptr(),
				LearningMethod:    datadog.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGMETHOD_DURATION.Ptr(),
				LearningThreshold: datadog.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ZERO_OCCURRENCES.Ptr(),
			},
		},
		Queries: []datadog.SecurityMonitoringRuleQuery{
			{
				Aggregation:    datadog.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
				DistinctFields: []string{},
				GroupByFields:  []string{},
				Metrics:        []string{},
			},
		},
		Tags:    []string{},
		Version: datadog.PtrInt32(1),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityMonitoringApi.UpdateSecurityMonitoringRule(ctx, "rule_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringRule`:\n%s\n", responseContent)
}
