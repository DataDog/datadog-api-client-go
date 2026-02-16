// Get code coverage summary for a commit returns "OK" response

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
	body := datadogV2.CommitCoverageSummaryRequest{
		Data: datadogV2.CommitCoverageSummaryRequestData{
			Attributes: datadogV2.CommitCoverageSummaryRequestAttributes{
				CommitSha:    "66adc9350f2cc9b250b69abddab733dd55e1a588",
				RepositoryId: "github.com/datadog/shopist",
			},
			Type: datadogV2.COMMITCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_COMMIT_SUMMARY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GetCodeCoverageCommitSummary", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCodeCoverageApi(apiClient)
	resp, r, err := api.GetCodeCoverageCommitSummary(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeCoverageApi.GetCodeCoverageCommitSummary`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CodeCoverageApi.GetCodeCoverageCommitSummary`:\n%s\n", responseContent)
}
