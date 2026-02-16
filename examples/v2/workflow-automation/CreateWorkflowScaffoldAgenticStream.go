// Generate workflow scaffold with agentic stream returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.WorkflowScaffoldAgenticStreamRequest{
		ChatHistory: []datadogV2.ChatMessage{
			{
				ChatId:   "chat-456",
				Content:  "Add error handling to the workflow",
				Id:       "msg-123",
				Role:     datadogV2.CHATMESSAGEROLE_USER,
				UserUuid: datadog.PtrString("550e8400-e29b-41d4-a716-446655440000"),
			},
		},
		PreviousAction: datadog.PtrString("created_initial_scaffold"),
		UserContext: &datadogV2.UserContext{
			UserInfo: datadogV2.UserInfo{
				OrgName:   "Acme Corp",
				UserEmail: "john.doe@example.com",
				UserName:  datadog.PtrString("John Doe"),
				UserUuid:  "550e8400-e29b-41d4-a716-446655440000",
			},
		},
		UserPrompt: "Create a workflow to restart a service when CPU is high",
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateWorkflowScaffoldAgenticStream", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	resp, r, err := api.CreateWorkflowScaffoldAgenticStream(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.CreateWorkflowScaffoldAgenticStream`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
