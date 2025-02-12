// Delete an existing Workflow returns "Successfully deleted a workflow." response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "workflow" in the system
	WorkflowDataID := os.Getenv("WORKFLOW_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	r, err := api.DeleteWorkflow(ctx, WorkflowDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.DeleteWorkflow`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
