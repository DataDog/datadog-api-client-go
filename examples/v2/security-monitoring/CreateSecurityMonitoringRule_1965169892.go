// Create a detection rule with type 'application_security 'returns "OK" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.SecurityMonitoringRuleCreatePayload{
SecurityMonitoringStandardRuleCreatePayload: &datadogV2.SecurityMonitoringStandardRuleCreatePayload{
Type: datadogV2.SECURITYMONITORINGRULETYPECREATE_APPLICATION_SECURITY.Ptr(),
Name: "Example-Security-Monitoring_appsec_rule",
Queries: []datadogV2.SecurityMonitoringStandardRuleQuery{
{
Query: datadog.PtrString("@appsec.security_activity:business_logic.users.login.failure"),
Aggregation: datadogV2.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
GroupByFields: []string{
"service",
"@http.client_ip",
},
DistinctFields: []string{
},
},
},
Filters: []datadogV2.SecurityMonitoringFilter{
},
Cases: []datadogV2.SecurityMonitoringRuleCaseCreate{
{
Name: datadog.PtrString(""),
Status: datadogV2.SECURITYMONITORINGRULESEVERITY_INFO,
Notifications: []string{
},
Condition: datadog.PtrString("a > 100000"),
Actions: []datadogV2.SecurityMonitoringRuleCaseAction{
{
Type: datadogV2.SECURITYMONITORINGRULECASEACTIONTYPE_BLOCK_IP.Ptr(),
Options: &datadogV2.SecurityMonitoringRuleCaseActionOptions{
Duration: datadog.PtrInt64(900),
},
},
{
Type: datadogV2.SECURITYMONITORINGRULECASEACTIONTYPE_USER_BEHAVIOR.Ptr(),
Options: &datadogV2.SecurityMonitoringRuleCaseActionOptions{
UserBehaviorName: datadog.PtrString("behavior"),
},
},
{
Type: datadogV2.SECURITYMONITORINGRULECASEACTIONTYPE_FLAG_IP.Ptr(),
Options: &datadogV2.SecurityMonitoringRuleCaseActionOptions{
FlaggedIpType: datadogV2.SECURITYMONITORINGRULECASEACTIONOPTIONSFLAGGEDIPTYPE_FLAGGED.Ptr(),
},
},
},
},
},
Options: datadogV2.SecurityMonitoringRuleOptions{
KeepAlive: datadogV2.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
MaxSignalDuration: datadogV2.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
EvaluationWindow: datadogV2.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
DetectionMethod: datadogV2.SECURITYMONITORINGRULEDETECTIONMETHOD_THRESHOLD.Ptr(),
},
IsEnabled: true,
Message: "Test rule",
Tags: []string{
},
GroupSignalsBy: []string{
"service",
},
}}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityMonitoringRule(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityMonitoringRule`:\n%s\n", responseContent)
}