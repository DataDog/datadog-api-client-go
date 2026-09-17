// Get all severity modifier rules returns "Successfully retrieved the list of severity modifier rules" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListSecurityFindingsAutomationSeverityModifierRules", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSecurityMonitoringApi(apiClient)
	resp, r, err := api.ListSecurityFindingsAutomationSeverityModifierRules(ctx, *datadogV2.NewListSecurityFindingsAutomationSeverityModifierRulesOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityMonitoringApi.ListSecurityFindingsAutomationSeverityModifierRules`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SecurityMonitoringApi.ListSecurityFindingsAutomationSeverityModifierRules`:\n%s\n", responseContent)
}
