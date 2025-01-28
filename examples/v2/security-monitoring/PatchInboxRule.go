// Patch an inbox rule returns "Inbox rule successfully patched" response

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
	// there is a valid "valid_inbox_rule" in the system
	ValidInboxRuleDataID := uuid.MustParse(os.Getenv("VALID_INBOX_RULE_DATA_ID"))

	body := datadogV2.PatchInboxRulesParameters{
		Data: &datadogV2.PatchInboxRulesParametersData{
			Attributes: datadogV2.PatchInboxRulesParametersDataAttributes{
				Action: &datadogV2.ActionInbox{
					ReasonDescription: datadog.PtrString("We want to focus on these items."),
				},
				Enabled: datadog.PtrBool(true),
				Name:    datadog.PtrString("Rule 1"),
				Rule: &datadogV2.AutomationRule{
					IssueType: datadogV2.ISSUETYPE_VULNERABILITY,
					Query:     datadog.PtrString("key:val"),
					RuleIds: []string{
						"rule-id-1",
					},
					RuleTypes: []datadogV2.SecurityRuleTypesItems{
						datadogV2.SECURITYRULETYPESITEMS_APPLICATION_CODE_VULNERABILITY,
					},
					Severities: []datadogV2.SecurityRuleSeverity{
						datadogV2.SECURITYRULESEVERITY_CRITICAL,
					},
				},
			},
			Id:   ValidInboxRuleDataID,
			Type: datadogV2.INBOXRULESTYPE_INBOX_RULES,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.PatchInboxRule(ctx, ValidInboxRuleDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.PatchInboxRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.PatchInboxRule`:\n%s\n", responseContent)
}