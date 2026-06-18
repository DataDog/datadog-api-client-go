// Reorder due date rules returns "Successfully reordered the due date rules" response

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

	body := datadogV2.DueDateRuleReorderRequest{
		Data: []datadogV2.DueDateRuleReorderItem{
			{
				Id:   ValidDueDateRuleDataID,
				Type: datadogV2.DUEDATERULETYPE_DUE_DATE_RULES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ReorderSecurityFindingsAutomationDueDateRules", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ReorderSecurityFindingsAutomationDueDateRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ReorderSecurityFindingsAutomationDueDateRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ReorderSecurityFindingsAutomationDueDateRules`:\n%s\n", responseContent)
}
