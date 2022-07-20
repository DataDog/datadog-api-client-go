// Aggregate compute events with group by returns "OK" response

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
	body := datadog.LogsAggregateRequest{
		Compute: []datadog.LogsCompute{
			{
				Aggregation: datadog.LOGSAGGREGATIONFUNCTION_COUNT,
				Interval:    common.PtrString("5m"),
				Type:        datadog.LOGSCOMPUTETYPE_TIMESERIES.Ptr(),
			},
		},
		Filter: &datadog.LogsQueryFilter{
			From: common.PtrString("now-15m"),
			Indexes: []string{
				"main",
			},
			Query: common.PtrString("*"),
			To:    common.PtrString("now"),
		},
		GroupBy: []datadog.LogsGroupBy{
			{
				Facet: "host",
				Missing: &datadog.LogsGroupByMissing{
					LogsGroupByMissingString: common.PtrString("miss")},
				Sort: &datadog.LogsAggregateSort{
					Type:        datadog.LOGSAGGREGATESORTTYPE_MEASURE.Ptr(),
					Order:       datadog.LOGSSORTORDER_ASCENDING.Ptr(),
					Aggregation: datadog.LOGSAGGREGATIONFUNCTION_PERCENTILE_90.Ptr(),
					Metric:      common.PtrString("@duration"),
				},
				Total: &datadog.LogsGroupByTotal{
					LogsGroupByTotalString: common.PtrString("recall")},
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewLogsApi(apiClient)
	resp, r, err := api.AggregateLogs(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.AggregateLogs`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `LogsApi.AggregateLogs`:\n%s\n", responseContent)
}
