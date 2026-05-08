// Get replay analysis issue returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetReplayAnalysisIssue", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumReplayAnalysisApi(apiClient)
	resp, r, err := api.GetReplayAnalysisIssue(ctx, "00000000-0000-0000-0000-000000000001")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumReplayAnalysisApi.GetReplayAnalysisIssue`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumReplayAnalysisApi.GetReplayAnalysisIssue`:\n%s\n", responseContent)
}
