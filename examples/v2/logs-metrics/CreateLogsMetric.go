// Create a log-based metric returns "OK" response

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
	body := datadog.LogsMetricCreateRequest{
		Data: datadog.LogsMetricCreateData{
			Id:   "Example-Create_a_log_based_metric_returns_OK_response",
			Type: datadog.LOGSMETRICTYPE_LOGS_METRICS,
			Attributes: datadog.LogsMetricCreateAttributes{
				Compute: datadog.LogsMetricCompute{
					AggregationType: datadog.LOGSMETRICCOMPUTEAGGREGATIONTYPE_COUNT,
				},
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.LogsMetricsApi(apiClient)
	resp, r, err := api.CreateLogsMetric(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.CreateLogsMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsMetricsApi.CreateLogsMetric`:\n%s\n", responseContent)
}
