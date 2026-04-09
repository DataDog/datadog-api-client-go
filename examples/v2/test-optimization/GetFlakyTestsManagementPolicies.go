// Get Flaky Tests Management policies returns "OK" response

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
	body := datadogV2.TestOptimizationFlakyTestsManagementPoliciesGetRequest{
		Data: datadogV2.TestOptimizationFlakyTestsManagementPoliciesGetRequestData{
			Attributes: datadogV2.TestOptimizationFlakyTestsManagementPoliciesGetRequestAttributes{
				RepositoryId: "github.com/datadog/shopist",
			},
			Type: datadogV2.TESTOPTIMIZATIONGETFLAKYTESTSMANAGEMENTPOLICIESREQUESTDATATYPE_TEST_OPTIMIZATION_GET_FLAKY_TESTS_MANAGEMENT_POLICIES_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetFlakyTestsManagementPolicies", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	resp, r, err := api.GetFlakyTestsManagementPolicies(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.GetFlakyTestsManagementPolicies`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TestOptimizationApi.GetFlakyTestsManagementPolicies`:\n%s\n", responseContent)
}
