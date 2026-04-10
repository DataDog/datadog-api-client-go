// Update Flaky Tests Management policies returns "OK" response

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
	body := datadogV2.TestOptimizationFlakyTestsManagementPoliciesUpdateRequest{
		Data: datadogV2.TestOptimizationFlakyTestsManagementPoliciesUpdateRequestData{
			Attributes: datadogV2.TestOptimizationFlakyTestsManagementPoliciesUpdateRequestAttributes{
				AttemptToFix: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesAttemptToFix{
					Retries: datadog.PtrInt64(3),
				},
				Disabled: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesDisabled{
					AutoDisableRule: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesAutoDisableRule{
						Enabled:       datadog.PtrBool(false),
						Status:        datadogV2.TESTOPTIMIZATIONFLAKYTESTSMANAGEMENTPOLICIESDISABLEDSTATUS_ACTIVE.Ptr(),
						WindowSeconds: datadog.PtrInt64(3600),
					},
					BranchRule: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesBranchRule{
						Branches: []string{
							"main",
						},
						Enabled:              datadog.PtrBool(true),
						ExcludedBranches:     []string{},
						ExcludedTestServices: []string{},
					},
					Enabled: datadog.PtrBool(false),
					FailureRateRule: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesDisabledFailureRateRule{
						Branches:  []string{},
						Enabled:   datadog.PtrBool(false),
						MinRuns:   datadog.PtrInt64(10),
						Status:    datadogV2.TESTOPTIMIZATIONFLAKYTESTSMANAGEMENTPOLICIESDISABLEDSTATUS_ACTIVE.Ptr(),
						Threshold: datadog.PtrFloat64(0.5),
					},
				},
				Quarantined: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesQuarantined{
					AutoQuarantineRule: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesAutoQuarantineRule{
						Enabled:       datadog.PtrBool(true),
						WindowSeconds: datadog.PtrInt64(3600),
					},
					BranchRule: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesBranchRule{
						Branches: []string{
							"main",
						},
						Enabled:              datadog.PtrBool(true),
						ExcludedBranches:     []string{},
						ExcludedTestServices: []string{},
					},
					Enabled: datadog.PtrBool(true),
					FailureRateRule: &datadogV2.TestOptimizationFlakyTestsManagementPoliciesQuarantinedFailureRateRule{
						Branches: []string{
							"main",
						},
						Enabled:   datadog.PtrBool(true),
						MinRuns:   datadog.PtrInt64(10),
						Threshold: datadog.PtrFloat64(0.5),
					},
				},
				RepositoryId: "github.com/datadog/shopist",
			},
			Type: datadogV2.TESTOPTIMIZATIONUPDATEFLAKYTESTSMANAGEMENTPOLICIESREQUESTDATATYPE_TEST_OPTIMIZATION_UPDATE_FLAKY_TESTS_MANAGEMENT_POLICIES_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateFlakyTestsManagementPolicies", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewTestOptimizationApi(apiClient)
	resp, r, err := api.UpdateFlakyTestsManagementPolicies(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestOptimizationApi.UpdateFlakyTestsManagementPolicies`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `TestOptimizationApi.UpdateFlakyTestsManagementPolicies`:\n%s\n", responseContent)
}
