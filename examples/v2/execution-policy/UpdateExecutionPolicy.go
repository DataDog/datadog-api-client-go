// Update an execution policy returns "OK" response

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
	// there is a valid "execution_policy" in the system
	ExecutionPolicyDataID := os.Getenv("EXECUTION_POLICY_DATA_ID")

	body := datadogV2.ExecutionPolicyUpdateRequest{
		Data: datadogV2.ExecutionPolicyUpdateRequestData{
			Id:   ExecutionPolicyDataID,
			Type: datadogV2.EXECUTIONPOLICYTYPE_EXECUTION_POLICY,
			Attributes: datadogV2.ExecutionPolicyWriteAttributes{
				Name:   "Cassette Execution Policy Updated",
				Effect: datadogV2.EXECUTIONPOLICYEFFECT_ALLOW,
				ActionPattern: datadogV2.ExecutionPolicyActionPattern{
					Integration: datadogV2.EXECUTIONPOLICYINTEGRATION_INTEGRATION_SCRIPT,
					ActionFqns: []string{
						"com.datadoghq.script.*",
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateExecutionPolicy", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExecutionPolicyApi(apiClient)
	resp, r, err := api.UpdateExecutionPolicy(ctx, ExecutionPolicyDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionPolicyApi.UpdateExecutionPolicy`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExecutionPolicyApi.UpdateExecutionPolicy`:\n%s\n", responseContent)
}
