// Validate a detection rule returns "OK" response

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
					Condition:     datadog.PtrString("a > 0"),
				},
			},
			HasExtendedTitle: datadog.PtrBool(true),
			IsEnabled:        true,
			Message:          "My security monitoring rule",
			Name:             "My security monitoring rule",
			Options: datadogV2.SecurityMonitoringRuleOptions{
				EvaluationWindow:  datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_THIRTY_MINUTES.Ptr(),
				KeepAlive:         datadogV2.SECURITYMONITORINGRULEKEEPALIVE_THIRTY_MINUTES.Ptr(),
				MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_THIRTY_MINUTES.Ptr(),
				DetectionMethod:   datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD.Ptr(),
			},
			Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{
				{
					Query: datadog.PtrString("source:source_here"),
					GroupByFields: []string{
						"@userIdentity.assumed_role",
					},
					DistinctFields: []string{},
					Aggregation:    datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
					Name:           datadog.PtrString(""),
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
