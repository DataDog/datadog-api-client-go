// Get a trace by ID returns "OK" response

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
	configuration.SetUnstableOperationEnabled("v2.GetTraceByID", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAPMTraceApi(apiClient)
	resp, r, err := api.GetTraceByID(ctx, "trace_id", *datadogV2.NewGetTraceByIDOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APMTraceApi.GetTraceByID`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `APMTraceApi.GetTraceByID`:\n%s\n", responseContent)
}
