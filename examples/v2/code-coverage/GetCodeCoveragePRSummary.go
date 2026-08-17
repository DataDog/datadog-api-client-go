// Get code coverage summary for a pull request returns "OK" response

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
	body := datadogV2.PRCoverageSummaryRequest{
		Data: datadogV2.PRCoverageSummaryRequestData{
			Attributes: datadogV2.PRCoverageSummaryRequestAttributes{
				PrNumber:      42,
				RepositoryUrl: "https://github.com/datadog/shopist",
			},
			Type: datadogV2.PRCOVERAGESUMMARYREQUESTTYPE_CI_APP_COVERAGE_PR_SUMMARY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCodeCoverageApi(apiClient)
	resp, r, err := api.GetCodeCoveragePRSummary(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeCoverageApi.GetCodeCoveragePRSummary`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CodeCoverageApi.GetCodeCoveragePRSummary`:\n%s\n", responseContent)
}
