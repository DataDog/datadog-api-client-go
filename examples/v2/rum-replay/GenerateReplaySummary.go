// Generate replay summary returns "OK" response

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
	body := datadogV2.ReplaySummaryRequest{
		Data: datadogV2.ReplaySummaryDataRequest{
			Type: datadogV2.REPLAYSUMMARYREQUESTTYPE_REPLAY_SUMMARY_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.GenerateReplaySummary", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRumReplayApi(apiClient)
	resp, r, err := api.GenerateReplaySummary(ctx, "00000000-0000-0000-0000-000000000001", "rum", body, *datadogV2.NewGenerateReplaySummaryOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RumReplayApi.GenerateReplaySummary`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RumReplayApi.GenerateReplaySummary`:\n%s\n", responseContent)
}
