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
		Cases: []datadog.SecurityMonitoringRuleCaseCreate{},
		Filters: []datadog.SecurityMonitoringFilter{
			{
				Action: datadog.SECURITYMONITORINGFILTERACTION_REQUIRE.Ptr(),
			},
		},
		HasExtendedTitle: datadog.PtrBool(true),
		IsEnabled:        true,
		Message:          "",
		Name:             "My security monitoring rule.",
		Options: datadog.SecurityMonitoringRuleOptions{
			DetectionMethod:        datadog.SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD.Ptr(),
			EvaluationWindow:       datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
			HardcodedEvaluatorType: datadog.SECURITYMONITORINGRULEHARDCODEDEVALUATORTYPE_LOG4SHELL.Ptr(),
			ImpossibleTravelOptions: &datadog.SecurityMonitoringRuleImpossibleTravelOptions{
				BaselineUserLocations: datadog.PtrBool(true),
			},
			KeepAlive:         datadog.SECURITYMONITORINGRULEKEEPALIVE_ZERO_MINUTES.Ptr(),
			MaxSignalDuration: datadog.SECURITYMONITORINGRULEMAXSIGNALDURATION_ZERO_MINUTES.Ptr(),
			NewValueOptions: &datadog.SecurityMonitoringRuleNewValueOptions{
				ForgetAfter:      datadog.SECURITYMONITORINGRULENEWVALUEOPTIONSFORGETAFTER_ONE_DAY.Ptr(),
				LearningDuration: datadog.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ZERO_DAYS.Ptr(),
			},
		},
		Queries: []datadog.SecurityMonitoringRuleQueryCreate{},
		Tags: []string{
			"env:prod",
			"team:security",
		},
		Type: datadog.SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityPlatformApi.CreateSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityPlatformApi.CreateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityPlatformApi.CreateSecurityMonitoringRule`:\n%s\n", responseContent)
}
