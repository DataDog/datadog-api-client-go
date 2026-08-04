// Update a severity modifier rule returns "Successfully updated the severity modifier rule" response

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
	// there is a valid "valid_severity_modifier_rule" in the system
	ValidSeverityModifierRuleDataID := uuid.MustParse(os.Getenv("VALID_SEVERITY_MODIFIER_RULE_DATA_ID"))

	body := datadogV2.SeverityModifierRuleUpdateRequest{
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
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityFindingsAutomationSeverityModifierRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityFindingsAutomationSeverityModifierRule(ctx, ValidSeverityModifierRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityFindingsAutomationSeverityModifierRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityFindingsAutomationSeverityModifierRule`:\n%s\n", responseContent)
}
