// Validate a detection rule with detection method 'sequence_detection' returns "OK" response

package main


import (
	"context"
	"encoding/json"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.SecurityMonitoringRuleValidatePayload{
SecurityMonitoringStandardRulePayload: &datadogV2.SecurityMonitoringStandardRulePayload{
Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
{
Name: datadog.PtrString(""),
Status: datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
Notifications: []string{
},
Condition: datadog.PtrString("step_b > 0"),
},
},
HasExtendedTitle: datadog.PtrBool(true),
IsEnabled: true,
Message: "My security monitoring rule",
Name: "My security monitoring rule",
Options: datadogV2.SecurityMonitoringRuleOptions{
EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ZERO_MINUTES.Ptr(),
KeepAlive: datadogV2.SECURITYMONITORINGRULEKEEPALIVE_FIVE_MINUTES.Ptr(),
MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_TEN_MINUTES.Ptr(),
DetectionMethod: datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_SEQUENCE_DETECTION.Ptr(),
SequenceDetectionOptions: &datadogV2.SecurityMonitoringRuleSequenceDetectionOptions{
StepTransitions: []datadogV2.SecurityMonitoringRuleSequenceDetectionStepTransition{
{
Child: datadog.PtrString("step_b"),
EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
Parent: datadog.PtrString("step_a"),
},
},
Steps: []datadogV2.SecurityMonitoringRuleSequenceDetectionStep{
{
Condition: datadog.PtrString("a > 0"),
EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ONE_MINUTE.Ptr(),
Name: datadog.PtrString("step_a"),
},
{
Condition: datadog.PtrString("b > 0"),
EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_ONE_MINUTE.Ptr(),
Name: datadog.PtrString("step_b"),
},
},
},
},
Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{
{
Query: datadog.PtrString("source:source_here"),
GroupByFields: []string{
"@userIdentity.assumed_role",
},
DistinctFields: []string{
},
Aggregation: datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
Name: datadog.PtrString(""),
},
{
Query: datadog.PtrString("source:source_here2"),
GroupByFields: []string{
},
DistinctFields: []string{
},
Aggregation: datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
Name: datadog.PtrString(""),
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
	r, err := api.ValidateSecurityMonitoringRule(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ValidateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}