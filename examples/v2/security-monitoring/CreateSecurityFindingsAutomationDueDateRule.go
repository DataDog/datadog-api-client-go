// Create a due date rule returns "Successfully created the due date rule" response

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
	body := datadogV2.DueDateRuleCreateRequest{
		Data: datadogV2.DueDateRuleDataCreate{
			Attributes: datadogV2.DueDateRuleAttributesCreate{
				Action: datadogV2.DueDateRuleAction{
					DueDaysPerSeverity: []datadogV2.DueDatePerSeverityItem{
						{
							DueInDays: 7,
							Severity:  datadogV2.DUEDATESEVERITY_CRITICAL,
						},
					},
					DueFrom: datadogV2.DUEDATEFROM_FIRST_SEEN,
				},
				Enabled: datadog.PtrBool(true),
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
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityFindingsAutomationDueDateRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityFindingsAutomationDueDateRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityFindingsAutomationDueDateRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityFindingsAutomationDueDateRule`:\n%s\n", responseContent)
}
