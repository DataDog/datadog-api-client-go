// Validate a detection rule with detection method 'new_value' with enabled feature 'instantaneousBaseline' returns "OK"
// response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.SecurityMonitoringRuleValidatePayload{
		SecurityMonitoringStandardRulePayload: &datadogV2.SecurityMonitoringStandardRulePayload{
			Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
				{
					Name:          datadog.PtrString(""),
					Status:        datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
					Notifications: []string{},
				},
			},
			HasExtendedTitle: datadog.PtrBool(true),
			IsEnabled:        true,
			Message:          "My security monitoring rule",
			Name:             "My security monitoring rule",
			Options: datadogV2.SecurityMonitoringRuleOptions{
				EvaluationWindow:  datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
				KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_FIVE_MINUTES.Ptr(),
				MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_TEN_MINUTES.Ptr(),
				DetectionMethod:   datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_NEW_VALUE.Ptr(),
				NewValueOptions: &datadogV2.SecurityMonitoringRuleNewValueOptions{
					ForgetAfter:           datadogV2.SECURITYMONITORINGRULENEWVALUEOPTIONSFORGETAFTER_ONE_WEEK.Ptr(),
					InstantaneousBaseline: datadog.PtrBool(true),
					LearningDuration:      datadogV2.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGDURATION_ONE_DAY.Ptr(),
					LearningThreshold:     datadogV2.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGTHRESHOLD_ZERO_OCCURRENCES.Ptr(),
					LearningMethod:        datadogV2.SECURITYMONITORINGRULENEWVALUEOPTIONSLEARNINGMETHOD_DURATION.Ptr(),
				},
			},
			Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{
				{
					Query: datadog.PtrString("source:source_here"),
					GroupByFields: []string{
						"@userIdentity.assumed_role",
					},
					DistinctFields: []string{},
					Metric:         datadog.PtrString("name"),
					Metrics: []string{
						"name",
					},
					Aggregation: datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_NEW_VALUE.Ptr(),
					Name:        datadog.PtrString(""),
					DataSource:  datadogV2.SECURITYMONITORINGSTANDARDDATASOURCE_LOGS.Ptr(),
				},
			},
			Tags: []string{
				"env:prod",
				"team:security",
			},
			Type: datadogV2.SECURITYMONITORINGRULETYPECREATE_LOG_DETECTION.Ptr(),
		}}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	r, err := api.ValidateSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ValidateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
