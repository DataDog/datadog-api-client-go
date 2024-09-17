// Update a log-based metric with include_percentiles field returns "OK" response

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
	// there is a valid "logs_metric_percentile" in the system
	LogsMetricPercentileDataID := os.Getenv("LOGS_METRIC_PERCENTILE_DATA_ID")

	body := datadogV2.LogsMetricUpdateRequest{
		Data: datadogV2.LogsMetricUpdateData{
			Type: datadogV2.LOGSMETRICTYPE_LOGS_METRICS,
			Attributes: datadogV2.LogsMetricUpdateAttributes{
				Compute: &datadogV2.LogsMetricUpdateCompute{
					IncludePercentiles: datadog.PtrBool(false),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsMetricsApi(apiClient)
	resp, r, err := api.UpdateLogsMetric(ctx, LogsMetricPercentileDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.UpdateLogsMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsMetricsApi.UpdateLogsMetric`:\n%s\n", responseContent)
}
