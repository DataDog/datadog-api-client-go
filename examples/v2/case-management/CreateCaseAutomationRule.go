// Create an automation rule returns "Created" response

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
	body := datadogV2.AutomationRuleCreateRequest{
		Data: datadogV2.AutomationRuleCreate{
			Attributes: datadogV2.AutomationRuleCreateAttributes{
				Action: datadogV2.AutomationRuleAction{
					Data: datadogV2.AutomationRuleActionData{
						Handle: datadog.PtrString("workflow-handle-123"),
					},
					Type: datadogV2.AUTOMATIONRULEACTIONTYPE_EXECUTE_WORKFLOW,
				},
				Name:  "Auto-assign workflow",
				State: datadogV2.CASEAUTOMATIONRULESTATE_ENABLED.Ptr(),
				Trigger: datadogV2.AutomationRuleTrigger{
					Data: &datadogV2.AutomationRuleTriggerData{},
					Type: datadogV2.AUTOMATIONRULETRIGGERTYPE_CASE_CREATED,
				},
			},
			Type: datadogV2.CASEAUTOMATIONRULERESOURCETYPE_RULE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCaseManagementApi(apiClient)
	resp, r, err := api.CreateCaseAutomationRule(ctx, "project_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaseManagementApi.CreateCaseAutomationRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CaseManagementApi.CreateCaseAutomationRule`:\n%s\n", responseContent)
}
