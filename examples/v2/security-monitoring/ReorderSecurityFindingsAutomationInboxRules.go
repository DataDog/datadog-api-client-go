// Reorder inbox rules returns "Successfully reordered the inbox rules" response

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

	body := datadogV2.InboxRuleReorderRequest{
		Data: []datadogV2.InboxRuleReorderItem{
			{
				Id:   ValidInboxRuleDataID,
				Type: datadogV2.INBOXRULETYPE_INBOX_RULES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ReorderSecurityFindingsAutomationInboxRules", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ReorderSecurityFindingsAutomationInboxRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ReorderSecurityFindingsAutomationInboxRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ReorderSecurityFindingsAutomationInboxRules`:\n%s\n", responseContent)
}
