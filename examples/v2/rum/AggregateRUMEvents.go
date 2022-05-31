// Aggregate RUM events returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.RUMAggregateRequest{
		Compute: []datadog.RUMCompute{
			{
				Aggregation: datadog.RUMAGGREGATIONFUNCTION_PERCENTILE_90,
				Metric:      datadog.PtrString("@view.time_spent"),
				Type:        datadog.RUMCOMPUTETYPE_TOTAL.Ptr(),
			},
		},
		Filter: &datadog.RUMQueryFilter{
			From:  datadog.PtrString("now-15m"),
			Query: datadog.PtrString("@type:view AND @session.type:user"),
			To:    datadog.PtrString("now"),
		},
		GroupBy: []datadog.RUMGroupBy{
			{
				Facet: "@view.time_spent",
				Limit: datadog.PtrInt64(10),
				Total: &datadog.RUMGroupByTotal{
					RUMGroupByTotalBoolean: datadog.PtrBool(false)},
			},
		},
		Options: &datadog.RUMQueryOptions{
			Timezone: datadog.PtrString("GMT"),
		},
		Page: &datadog.RUMQueryPageOptions{
			Limit: datadog.PtrInt32(25),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.RUMApi.AggregateRUMEvents(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.AggregateRUMEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.AggregateRUMEvents`:\n%s\n", responseContent)
}
