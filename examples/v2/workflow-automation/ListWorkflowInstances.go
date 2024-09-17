// List workflow instances returns "OK" response

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
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewWorkflowAutomationApi(apiClient)
	resp, r, err := api.ListWorkflowInstances(ctx, "ccf73164-1998-4785-a7a3-8d06c7e5f558", *datadogV2.NewListWorkflowInstancesOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowAutomationApi.ListWorkflowInstances`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `WorkflowAutomationApi.ListWorkflowInstances`:\n%s\n", responseContent)
}
