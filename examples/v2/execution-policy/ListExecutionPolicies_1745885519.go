// List execution policies with query parameters returns "OK" response

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
	ExecutionPolicyDataAttributesCreatedBy := os.Getenv("EXECUTION_POLICY_DATA_ATTRIBUTES_CREATED_BY")
	ExecutionPolicyDataAttributesName := os.Getenv("EXECUTION_POLICY_DATA_ATTRIBUTES_NAME")
	ExecutionPolicyDataID := os.Getenv("EXECUTION_POLICY_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.ListExecutionPolicies", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewExecutionPolicyApi(apiClient)
	resp, r, err := api.ListExecutionPolicies(ctx, *datadogV2.NewListExecutionPoliciesOptionalParameters().WithPageSize(10).WithPageNumber(0).WithFilterName(ExecutionPolicyDataAttributesName).WithFilterIds([]string{
		ExecutionPolicyDataID,
	}).WithFilterIntegration([]datadogV2.ExecutionPolicyIntegration{
		datadogV2.EXECUTIONPOLICYINTEGRATION_INTEGRATION_SCRIPT,
	}).WithFilterEffects([]datadogV2.ExecutionPolicyEffect{
		datadogV2.EXECUTIONPOLICYEFFECT_ALLOW,
	}).WithFilterCreatorIds([]string{
		ExecutionPolicyDataAttributesCreatedBy,
	}).WithSort([]string{
		"-created_at",
	}))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionPolicyApi.ListExecutionPolicies`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ExecutionPolicyApi.ListExecutionPolicies`:\n%s\n", responseContent)
}
