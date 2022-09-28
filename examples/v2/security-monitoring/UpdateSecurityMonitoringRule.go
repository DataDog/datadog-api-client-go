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
	body := datadogV2.SecurityMonitoringRuleUpdatePayload{
		Cases: []datadogV2.SecurityMonitoringRuleCase{
			{
				Notifications: []string{},
				Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_CRITICAL.Ptr(),
			},
		},
		Filters: []datadogV2.SecurityMonitoringFilter{
			{
				Action: datadogV2.SECURITYMONITORINGFILTERACTION_REQUIRE.Ptr(),
			},
		},
		HasExtendedTitle: datadog.PtrBool(true),
		Options: &datadogV2.SecurityMonitoringRuleOptions{
			DecreaseCriticalityBasedOnEnv: datadog.PtrBool(false),
			DetectionMethod:               datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD.Ptr(),
			EvaluationWindow:              datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
			HardcodedEvaluatorType:        datadogV2.SECURITYMONITORINGRULEHARDCODEDEVALUATORTYPE_LOG4SHELL.Ptr(),
			ImpossibleTravelOptions: &datadogV2.SecurityMonitoringRuleImpossibleTravelOptions{
				BaselineUserLocations: datadog.PtrBool(true),
			},
			KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ZERO_MINUTES.Ptr(),
			MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ZERO_MINUTES.Ptr(),
			NewValueOptions: &datadogV2.SecurityMonitoringRuleNewValueOptions{
				ForgetAfter:       datadogV2.SECURITYMONITORINGRULENEWVALUEOPTIONSFORGETAFTER_ONE_DAY.Ptr(),
				LearningDuration:  datadogV2.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ZERO_DAYS.Ptr(),
				LearningMethod:    datadogV2.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGMETHOD_DURATION.Ptr(),
				LearningThreshold: datadogV2.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ZERO_OCCURRENCES.Ptr(),
			},
		},
		Queries: []datadogV2.SecurityMonitoringRuleQuery{
			datadogV2.SecurityMonitoringRuleQuery{
				SecurityMonitoringStandardRuleQuery: &datadogV2.SecurityMonitoringStandardRuleQuery{
					Aggregation:    datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
					DistinctFields: []string{},
					GroupByFields:  []string{},
					Metrics:        []string{},
					Query:          "a > 3",
				}},
		},
		Tags:    []string{},
		Version: datadog.PtrInt32(1),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringRule(ctx, "rule_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringRule`:\n%s\n", responseContent)
}
