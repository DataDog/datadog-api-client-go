// Get a severity modifier rule returns "Successfully retrieved the severity modifier rule" response

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

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetSecurityFindingsAutomationSeverityModifierRule", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.GetSecurityFindingsAutomationSeverityModifierRule(ctx, ValidSeverityModifierRuleDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.GetSecurityFindingsAutomationSeverityModifierRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.GetSecurityFindingsAutomationSeverityModifierRule`:\n%s\n", responseContent)
}
