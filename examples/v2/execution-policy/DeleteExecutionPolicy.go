// Delete an execution policy returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "execution_policy" in the system
	ExecutionPolicyDataID := os.Getenv("EXECUTION_POLICY_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.DeleteExecutionPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExecutionPolicyApi(apiClient)
	r, err := api.DeleteExecutionPolicy(ctx, ExecutionPolicyDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionPolicyApi.DeleteExecutionPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
