// Get a log-based metric returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "logs_metric" in the system
	LogsMetricDataID := os.Getenv("LOGS_METRIC_DATA_ID")

	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.LogsMetricsApi(apiClient)
	resp, r, err := api.GetLogsMetric(ctx, LogsMetricDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.GetLogsMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsMetricsApi.GetLogsMetric`:\n%s\n", responseContent)
}
