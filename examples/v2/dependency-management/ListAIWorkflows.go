// List AI workflows returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.ListAIWorkflows", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDependencyManagementApi(apiClient)
	resp, r, err := api.ListAIWorkflows(ctx, *datadogV2.NewListAIWorkflowsOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependencyManagementApi.ListAIWorkflows`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DependencyManagementApi.ListAIWorkflows`:\n%s\n", responseContent)
}
