// Get a RUM-based metric returns "OK" response

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
	// there is a valid "rum_metric" in the system
	RumMetricDataID := os.Getenv("RUM_METRIC_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMMetricsApi(apiClient)
	resp, r, err := api.GetRUMMetric(ctx, RumMetricDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMMetricsApi.GetRUMMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMMetricsApi.GetRUMMetric`:\n%s\n", responseContent)
}
