// Create a severity modifier rule returns "Successfully created the severity modifier rule" response

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
	body := datadogV2.SeverityModifierRuleCreateRequest{
		Data: datadogV2.SeverityModifierRuleDataCreate{
			Attributes: datadogV2.SeverityModifierRuleAttributesCreate{
				Action: datadogV2.SeverityModifierRuleAction{
					SeverityModifierRuleSetAction: &datadogV2.SeverityModifierRuleSetAction{
						Description: datadog.PtrString("Lower severity for dev environment noise"),
						Severity:    datadogV2.SEVERITYMODIFIERSEVERITY_LOW,
						Type:        datadogV2.SEVERITYMODIFIERRULESETACTIONTYPE_SET,
					}},
				Enabled: datadog.PtrBool(true),
				Name:    "Downgrade misconfigurations in dev",
				Rule: datadogV2.AutomationRuleScope{
					FindingTypes: []datadogV2.SecurityFindingType{
						datadogV2.SECURITYFINDINGTYPE_MISCONFIGURATION,
					},
					Query: datadog.PtrString("env:prod team:platform"),
				},
			},
			Type: datadogV2.SEVERITYMODIFIERRULETYPE_SEVERITY_MODIFIER_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityFindingsAutomationSeverityModifierRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityFindingsAutomationSeverityModifierRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityFindingsAutomationSeverityModifierRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityFindingsAutomationSeverityModifierRule`:\n%s\n", responseContent)
}
