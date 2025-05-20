// Update an existing Workflow returns "Successfully updated a workflow." response

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
	// there is a valid "workflow" in the system
	WorkflowDataID := os.Getenv("WORKFLOW_DATA_ID")

	body := datadogV2.UpdateWorkflowRequest{
		Data: datadogV2.WorkflowDataUpdate{
			Attributes: datadogV2.WorkflowDataUpdateAttributes{
				Description: datadog.PtrString("A sample workflow."),
				Name:        datadog.PtrString("Example Workflow"),
				Published:   datadog.PtrBool(true),
				Spec: &datadogV2.Spec{
					ConnectionEnvs: []datadogV2.ConnectionEnv{
						{
							Connections: []datadogV2.Connection{
								{
									ConnectionId: "11111111-1111-1111-1111-111111111111",
									Label:        "INTEGRATION_DATADOG",
								},
							},
							Env: datadogV2.CONNECTIONENVENV_DEFAULT,
						},
					},
					InputSchema: &datadogV2.InputSchema{
						Parameters: []datadogV2.InputSchemaParameters{
							{
								DefaultValue: "default",
								Name:         "input",
								Type:         datadogV2.INPUTSCHEMAPARAMETERSTYPE_STRING,
							},
						},
					},
					OutputSchema: &datadogV2.OutputSchema{
						Parameters: []datadogV2.OutputSchemaParameters{
							{
								Name:  "output",
								Type:  datadogV2.OUTPUTSCHEMAPARAMETERSTYPE_ARRAY_OBJECT,
								Value: "outputValue",
							},
						},
					},
					Steps: []datadogV2.Step{
						{
							ActionId:        "com.datadoghq.dd.monitor.listMonitors",
							ConnectionLabel: datadog.PtrString("INTEGRATION_DATADOG"),
							Name:            "Step1",
							OutboundEdges: []datadogV2.OutboundEdge{
								{
									BranchName:   "main",
									NextStepName: "Step2",
								},
							},
							Parameters: []datadogV2.Parameter{
								{
									Name:  "tags",
									Value: "service:monitoring",
								},
							},
						},
						{
							ActionId: "com.datadoghq.core.noop",
							Name:     "Step2",
						},
					},
					Triggers: []datadogV2.Trigger{
						datadogV2.Trigger{
							MonitorTriggerWrapper: &datadogV2.MonitorTriggerWrapper{
								MonitorTrigger: datadogV2.MonitorTrigger{
									RateLimit: &datadogV2.TriggerRateLimit{
										Count:    datadog.PtrInt64(1),
										Interval: datadog.PtrString("3600s"),
									},
								},
								StartStepNames: []string{
									"Step1",
								},
							}},
						datadogV2.Trigger{
							GithubWebhookTriggerWrapper: &datadogV2.GithubWebhookTriggerWrapper{
								StartStepNames: []string{
									"Step1",
								},
								GithubWebhookTrigger: datadogV2.GithubWebhookTrigger{},
							}},
					},
				},
				Tags: []string{
					"team:infra",
					"service:monitoring",
					"foo:bar",
				},
			},
			Id:   datadog.PtrString("22222222-2222-2222-2222-222222222222"),
			Type: datadogV2.WORKFLOWDATATYPE_WORKFLOWS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	resp, r, err := api.UpdateWorkflow(ctx, WorkflowDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.UpdateWorkflow`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WorkflowAutomationApi.UpdateWorkflow`:\n%s\n", responseContent)
}
