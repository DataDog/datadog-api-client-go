// Aggregate RUM events returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v2/datadog"
)

func main() {
	body := datadog.RUMAggregateRequest{
		Compute: []datadog.RUMCompute{
			{
				Aggregation: datadog.RUMAGGREGATIONFUNCTION_PERCENTILE_90,
				Metric:      common.PtrString("@view.time_spent"),
				Type:        datadog.RUMCOMPUTETYPE_TOTAL.Ptr(),
			},
		},
		Filter: &datadog.RUMQueryFilter{
			From:  common.PtrString("now-15m"),
			Query: common.PtrString("@type:view AND @session.type:user"),
			To:    common.PtrString("now"),
		},
		GroupBy: []datadog.RUMGroupBy{
			{
				Facet: "@view.time_spent",
				Limit: common.PtrInt64(10),
				Total: &datadog.RUMGroupByTotal{
					RUMGroupByTotalBoolean: common.PtrBool(false)},
			},
		},
		Options: &datadog.RUMQueryOptions{
			Timezone: common.PtrString("GMT"),
		},
		Page: &datadog.RUMQueryPageOptions{
			Limit: common.PtrInt32(25),
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewRUMApi(apiClient)
	resp, r, err := api.AggregateRUMEvents(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMApi.AggregateRUMEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMApi.AggregateRUMEvents`:\n%s\n", responseContent)
}
