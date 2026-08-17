// Get per-file code coverage data returns "OK" response

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
	body := datadogV2.FilesCoverageRequest{
		Data: datadogV2.FilesCoverageRequestData{
			Attributes: datadogV2.FilesCoverageRequestAttributes{
				ChangedOnly:   datadog.PtrBool(true),
				CommitSha:     datadog.PtrString("66adc9350f2cc9b250b69abddab733dd55e1a588"),
				RepositoryUrl: datadog.PtrString("https://github.com/datadog/shopist"),
			},
			Type: datadogV2.FILESCOVERAGEREQUESTTYPE_CI_APP_COVERAGE_FILES_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCodeCoverageApi(apiClient)
	resp, r, err := api.GetCodeCoverageFiles(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodeCoverageApi.GetCodeCoverageFiles`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CodeCoverageApi.GetCodeCoverageFiles`:\n%s\n", responseContent)
}
