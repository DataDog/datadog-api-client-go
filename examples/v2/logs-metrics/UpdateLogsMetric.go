// Update a log-based metric returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "logs_metric" in the system
	LogsMetricDataID := os.Getenv("LOGS_METRIC_DATA_ID")

	body := datadog.LogsMetricUpdateRequest{
		Data: datadog.LogsMetricUpdateData{
			Type: datadog.LOGSMETRICTYPE_LOGS_METRICS,
			Attributes: datadog.LogsMetricUpdateAttributes{
				Filter: &datadog.LogsMetricFilter{
					Query: datadog.PtrString("service:web* AND @http.status_code:[200 TO 299]-updated"),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsMetricsApi.UpdateLogsMetric(ctx, LogsMetricDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.UpdateLogsMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsMetricsApi.UpdateLogsMetric`:\n%s\n", responseContent)
}
