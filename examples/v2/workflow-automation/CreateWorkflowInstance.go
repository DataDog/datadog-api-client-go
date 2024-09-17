// Execute a workflow returns "Created" response

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
	body := datadogV2.WorkflowInstanceCreateRequest{
		Meta: &datadogV2.WorkflowInstanceCreateMeta{
			Payload: map[string]interface{}{
				"input": "value",
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	resp, r, err := api.CreateWorkflowInstance(ctx, "ccf73164-1998-4785-a7a3-8d06c7e5f558", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.CreateWorkflowInstance`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WorkflowAutomationApi.CreateWorkflowInstance`:\n%s\n", responseContent)
}
