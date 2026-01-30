// Execute a workflow from a template returns "OK" response

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
	body := datadogV2.WorkflowHeadlessExecutionRequest{
		Data: datadogV2.WorkflowHeadlessExecutionRequestData{
			Attributes: datadogV2.WorkflowHeadlessExecutionRequestAttributes{
				Config: datadogV2.WorkflowHeadlessExecutionConfig{
					Connections: []datadogV2.WorkflowHeadlessExecutionConnection{
						{
							ConnectionId: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
							Label:        "INTEGRATION_DATADOG",
						},
					},
					Inputs: map[string]interface{}{},
				},
				TemplateId: "template-789",
			},
			Id:   "1234",
			Type: datadogV2.WORKFLOWHEADLESSEXECUTIONREQUESTTYPE_WORKFLOW_HEADLESS_EXECUTION_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ExecuteWorkflowFromTemplate", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	resp, r, err := api.ExecuteWorkflowFromTemplate(ctx, "parent_id", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.ExecuteWorkflowFromTemplate`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WorkflowAutomationApi.ExecuteWorkflowFromTemplate`:\n%s\n", responseContent)
}
