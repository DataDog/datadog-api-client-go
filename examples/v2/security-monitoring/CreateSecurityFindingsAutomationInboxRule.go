// Create an inbox rule returns "Successfully created the inbox rule" response

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
	body := datadogV2.InboxRuleCreateRequest{
		Data: datadogV2.InboxRuleDataCreate{
			Attributes: datadogV2.InboxRuleAttributesCreate{
				Action: datadogV2.InboxRuleAction{
					Description: datadog.PtrString("Needs triage"),
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
			Type: datadogV2.INBOXRULETYPE_INBOX_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateSecurityFindingsAutomationInboxRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.CreateSecurityFindingsAutomationInboxRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.CreateSecurityFindingsAutomationInboxRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.CreateSecurityFindingsAutomationInboxRule`:\n%s\n", responseContent)
}
