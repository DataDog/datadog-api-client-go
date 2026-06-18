// Update a due date rule returns "Successfully updated the due date rule" response

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
	// there is a valid "valid_due_date_rule" in the system
	ValidDueDateRuleDataID := uuid.MustParse(os.Getenv("VALID_DUE_DATE_RULE_DATA_ID"))

	body := datadogV2.DueDateRuleUpdateRequest{
		Data: datadogV2.DueDateRuleDataCreate{
			Attributes: datadogV2.DueDateRuleAttributesCreate{
				Action: datadogV2.DueDateRuleAction{
					DueDaysPerSeverity: []datadogV2.DueDatePerSeverityItem{
						{
							DueInDays: 14,
							Severity:  datadogV2.DUEDATESEVERITY_CRITICAL,
						},
					},
					DueFrom: datadogV2.DUEDATEFROM_FIRST_SEEN,
				},
				Enabled: datadog.PtrBool(false),
				Name:    "Example-Security-Monitoring",
				Rule: datadogV2.AutomationRuleScope{
					FindingTypes: []datadogV2.SecurityFindingType{
						datadogV2.SECURITYFINDINGTYPE_MISCONFIGURATION,
					},
					Query: datadog.PtrString("env:staging"),
				},
			},
			Type: datadogV2.DUEDATERULETYPE_DUE_DATE_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateSecurityFindingsAutomationDueDateRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.UpdateSecurityFindingsAutomationDueDateRule(ctx, ValidDueDateRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.UpdateSecurityFindingsAutomationDueDateRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.UpdateSecurityFindingsAutomationDueDateRule`:\n%s\n", responseContent)
}
