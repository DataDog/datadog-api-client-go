// Get a pruned trace by ID returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetPrunedTraceByID", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPMTraceApi(apiClient)
	resp, r, err := api.GetPrunedTraceByID(ctx, "trace_id", *datadogV2.NewGetPrunedTraceByIDOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APMTraceApi.GetPrunedTraceByID`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `APMTraceApi.GetPrunedTraceByID`:\n%s\n", responseContent)
}
