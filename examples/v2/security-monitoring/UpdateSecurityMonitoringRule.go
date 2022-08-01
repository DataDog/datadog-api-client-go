// Update an existing rule returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
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
		HasExtendedTitle: common.PtrBool(true),
		Options: &datadog.SecurityMonitoringRuleOptions{
			DecreaseCriticalityBasedOnEnv: common.PtrBool(false),
			DetectionMethod:               datadog.SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD.Ptr(),
			EvaluationWindow:              datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
			HardcodedEvaluatorType:        datadog.SECURITYMONITORINGRULEHARDCODEDEVALUATORTYPE_LOG4SHELL.Ptr(),
			ImpossibleTravelOptions: &datadog.SecurityMonitoringRuleImpossibleTravelOptions{
				BaselineUserLocations: common.PtrBool(true),
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
		Version: common.PtrInt32(1),
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityMonitoringRule(ctx, "rule_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityMonitoringRule`:\n%s\n", responseContent)
}
