// Get code coverage summary for a branch returns "OK" response

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
	body := datadogV2.BranchCoverageSummaryRequest{
		Data: datadogV2.BranchCoverageSummaryRequestData{
			Attributes: datadogV2.BranchCoverageSummaryRequestAttributes{
				Branch:       "prod",
				RepositoryId: "github.com/datadog/shopist",
			},
			Type: datadogV2.BRANCHCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_BRANCH_SUMMARY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetCodeCoverageBranchSummary", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCodeCoverageApi(apiClient)
	resp, r, err := api.GetCodeCoverageBranchSummary(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeCoverageApi.GetCodeCoverageBranchSummary`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CodeCoverageApi.GetCodeCoverageBranchSummary`:\n%s\n", responseContent)
}
