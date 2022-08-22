// Aggregate RUM events returns "OK" response

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
	body := datadogV2.RUMAggregateRequest{
		Compute: []datadogV2.RUMCompute{
			{
				Aggregation: datadogV2.RUMAGGREGATIONFUNCTION_PERCENTILE_90,
				Metric:      datadog.PtrString("@view.time_spent"),
				Type:        datadogV2.RUMCOMPUTETYPE_TOTAL.Ptr(),
			},
		},
		Filter: &datadogV2.RUMQueryFilter{
			From:  datadog.PtrString("now-15m"),
			Query: datadog.PtrString("@type:view AND @session.type:user"),
			To:    datadog.PtrString("now"),
		},
		GroupBy: []datadogV2.RUMGroupBy{
			{
				Facet: "@view.time_spent",
				Limit: datadog.PtrInt64(10),
				Total: &datadogV2.RUMGroupByTotal{
					RUMGroupByTotalBoolean: datadog.PtrBool(false)},
			},
		},
		Options: &datadogV2.RUMQueryOptions{
			Timezone: datadog.PtrString("GMT"),
		},
		Page: &datadogV2.RUMQueryPageOptions{
			Limit: datadog.PtrInt32(25),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMApi(apiClient)
	resp, r, err := api.AggregateRUMEvents(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.AggregateRUMEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.AggregateRUMEvents`:\n%s\n", responseContent)
}
