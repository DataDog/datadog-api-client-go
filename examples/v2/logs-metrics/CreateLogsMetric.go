// Create a log-based metric returns "OK" response

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
	body := datadogV2.LogsMetricCreateRequest{
		Data: datadogV2.LogsMetricCreateData{
			Id:   "ExampleLogsMetric",
			Type: datadogV2.LOGSMETRICTYPE_LOGS_METRICS,
			Attributes: datadogV2.LogsMetricCreateAttributes{
				Compute: datadogV2.LogsMetricCompute{
					AggregationType:    datadogV2.LOGSMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
					IncludePercentiles: datadog.PtrBool(true),
					Path:               datadog.PtrString("@duration"),
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsMetricsApi(apiClient)
	resp, r, err := api.CreateLogsMetric(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsMetricsApi.CreateLogsMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsMetricsApi.CreateLogsMetric`:\n%s\n", responseContent)
}
