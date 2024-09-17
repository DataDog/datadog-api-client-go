// Delete a log-based metric returns "OK" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "logs_metric" in the system
	LogsMetricDataID := os.Getenv("LOGS_METRIC_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsMetricsApi(apiClient)
	r, err := api.DeleteLogsMetric(ctx, LogsMetricDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.DeleteLogsMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
