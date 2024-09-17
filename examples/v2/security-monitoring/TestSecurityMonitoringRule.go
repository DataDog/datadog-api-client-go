// Test a rule returns "OK" response

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
	body := datadogV2.SecurityMonitoringRuleTestRequest{
		Rule: &datadogV2.SecurityMonitoringRuleTestPayload{
			SecurityMonitoringStandardRuleTestPayload: &datadogV2.SecurityMonitoringStandardRuleTestPayload{
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
				Message:          "My security monitoring rule message.",
				Name:             "My security monitoring rule.",
				Options: datadogV2.SecurityMonitoringRuleOptions{
					DecreaseCriticalityBasedOnEnv: datadog.PtrBool(false),
					DetectionMethod:               datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD.Ptr(),
					EvaluationWindow:              datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
					KeepAlive:                     datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ZERO_MINUTES.Ptr(),
					MaxSignalDuration:             datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ZERO_MINUTES.Ptr(),
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
				Type: datadogV2.SECURITYMONITORINGRULETYPETEST_LOG_DETECTION.Ptr(),
			}},
		RuleQueryPayloads: []datadogV2.SecurityMonitoringRuleQueryPayload{
			{
				ExpectedResult: datadog.PtrBool(true),
				Index:          datadog.PtrInt64(0),
				Payload: &datadogV2.SecurityMonitoringRuleQueryPayloadData{
					Ddsource: datadog.PtrString("source_here"),
					Ddtags:   datadog.PtrString("env:staging,version:5.1"),
					Hostname: datadog.PtrString("i-012345678"),
					Message:  datadog.PtrString("2019-11-19T14:37:58,995 INFO [process.name][20081] Hello World"),
					Service:  datadog.PtrString("payment"),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.TestSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.TestSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.TestSecurityMonitoringRule`:\n%s\n", responseContent)
}
