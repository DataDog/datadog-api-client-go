// Aggregate compute events returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.LogsAggregateRequest{
		Compute: []datadog.LogsCompute{
			{
				Aggregation: datadog.LOGSAGGREGATIONFUNCTION_COUNT,
				Interval:    datadog.PtrString("5m"),
				Type:        datadog.LOGSCOMPUTETYPE_TIMESERIES.Ptr(),
			},
		},
		Filter: &datadog.LogsQueryFilter{
			From: datadog.PtrString("now-15m"),
			Indexes: []string{
				"main",
			},
			Query: datadog.PtrString("*"),
			To:    datadog.PtrString("now"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.LogsApi.AggregateLogs(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.AggregateLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.AggregateLogs`:\n%s\n", responseContent)
}
