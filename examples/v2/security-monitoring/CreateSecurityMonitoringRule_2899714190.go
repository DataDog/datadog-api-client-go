// Create a detection rule with detection method 'sequence_detection' returns "OK" response

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
			Name:      "Example-Security-Monitoring",
			Type:      datadogV2.SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION.Ptr(),
			IsEnabled: true,
			Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{
				{
					Aggregation:              datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
					DataSource:               datadogV2.SECURITYMONITORINGSTANDARDDATASOURCE_LOGS.Ptr(),
					DistinctFields:           []string{},
					GroupByFields:            []string{},
					HasOptionalGroupByFields: datadog.PtrBool(false),
					Name:                     datadog.PtrString(""),
					Query:                    datadog.PtrString("service:logs-rule-reducer source:paul test2"),
				},
				{
					Aggregation:              datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
					DataSource:               datadogV2.SECURITYMONITORINGSTANDARDDATASOURCE_LOGS.Ptr(),
					DistinctFields:           []string{},
					GroupByFields:            []string{},
					HasOptionalGroupByFields: datadog.PtrBool(false),
					Name:                     datadog.PtrString(""),
					Query:                    datadog.PtrString("service:logs-rule-reducer source:paul test1"),
				},
			},
			Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
				{
					Name:          datadog.PtrString(""),
					Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
					Notifications: []string{},
					Condition:     datadog.PtrString("step_b > 0"),
				},
			},
			Message: "Logs and signals asdf",
			Options: datadogV2.SecurityMonitoringRuleOptions{
				DetectionMethod:   datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_SEQUENCE_DETECTION.Ptr(),
				EvaluationWindow:  datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
				KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_FIVE_MINUTES.Ptr(),
				MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_TEN_MINUTES.Ptr(),
				SequenceDetectionOptions: &datadogV2.SecurityMonitoringRuleSequenceDetectionOptions{
					StepTransitions: []datadogV2.SecurityMonitoringRuleSequenceDetectionStepTransition{
						{
							Child:            datadog.PtrString("step_b"),
							EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
							Parent:           datadog.PtrString("step_a"),
						},
					},
					Steps: []datadogV2.SecurityMonitoringRuleSequenceDetectionStep{
						{
							Condition:        datadog.PtrString("a > 0"),
							EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ONE_MINUTE.Ptr(),
							Name:             datadog.PtrString("step_a"),
						},
						{
							Condition:        datadog.PtrString("b > 0"),
							EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ONE_MINUTE.Ptr(),
							Name:             datadog.PtrString("step_b"),
						},
					},
				},
			},
			Tags: []string{},
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
