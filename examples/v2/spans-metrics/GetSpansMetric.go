// Get a span-based metric returns "OK" response

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
	// there is a valid "spans_metric" in the system
	SpansMetricDataID := os.Getenv("SPANS_METRIC_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSpansMetricsApi(apiClient)
	resp, r, err := api.GetSpansMetric(ctx, SpansMetricDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpansMetricsApi.GetSpansMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SpansMetricsApi.GetSpansMetric`:\n%s\n", responseContent)
}
