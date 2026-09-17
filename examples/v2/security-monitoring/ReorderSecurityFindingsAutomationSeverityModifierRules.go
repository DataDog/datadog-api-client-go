// Reorder severity modifier rules returns "Successfully reordered the severity modifier rules" response

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

	body := datadogV2.SeverityModifierRuleReorderRequest{
		Data: []datadogV2.SeverityModifierRuleReorderItem{
			{
				Id:   ValidSeverityModifierRuleDataID,
				Type: datadogV2.SEVERITYMODIFIERRULETYPE_SEVERITY_MODIFIER_RULES,
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ReorderSecurityFindingsAutomationSeverityModifierRules", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ReorderSecurityFindingsAutomationSeverityModifierRules(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ReorderSecurityFindingsAutomationSeverityModifierRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ReorderSecurityFindingsAutomationSeverityModifierRules`:\n%s\n", responseContent)
}
