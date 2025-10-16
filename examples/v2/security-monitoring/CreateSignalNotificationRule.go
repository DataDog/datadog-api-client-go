// Create a new signal-based notification rule returns "Successfully created the notification rule." response

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
	body := datadogV2.CreateNotificationRuleParameters{
Data: &datadogV2.CreateNotificationRuleParametersData{
Attributes: datadogV2.CreateNotificationRuleParametersDataAttributes{
Enabled: datadog.PtrBool(true),
Name: "Rule 1",
Selectors: datadogV2.Selectors{
Query: datadog.PtrString("(source:production_service OR env:prod)"),
RuleTypes: []datadogV2.RuleTypesItems{
datadogV2.RULETYPESITEMS_MISCONFIGURATION,
datadogV2.RULETYPESITEMS_ATTACK_PATH,
},
Severities: []datadogV2.RuleSeverity{
datadogV2.RULESEVERITY_CRITICAL,
},
TriggerSource: datadogV2.TRIGGERSOURCE_SECURITY_FINDINGS,
},
Targets: []string{
"@john.doe@email.com",
},
TimeAggregation: datadog.PtrInt64(86400),
},
Type: datadogV2.NOTIFICATIONRULESTYPE_NOTIFICATION_RULES,
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSignalNotificationRule(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSignalNotificationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSignalNotificationRule`:\n%s\n", responseContent)
}