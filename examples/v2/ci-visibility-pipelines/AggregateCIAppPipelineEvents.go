// Aggregate pipelines events returns "OK" response

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
	body := datadogV2.CIAppPipelinesAggregateRequest{
		Compute: []datadogV2.CIAppCompute{
			{
				Aggregation: datadogV2.CIAPPAGGREGATIONFUNCTION_PERCENTILE_90,
				Metric:      datadog.PtrString("@duration"),
				Type:        datadogV2.CIAPPCOMPUTETYPE_TOTAL.Ptr(),
			},
		},
		Filter: &datadogV2.CIAppPipelinesQueryFilter{
			From:  datadog.PtrString("now-15m"),
			Query: datadog.PtrString("@ci.provider.name:(gitlab OR github)"),
			To:    datadog.PtrString("now"),
		},
		GroupBy: []datadogV2.CIAppPipelinesGroupBy{
			{
				Facet: "@ci.status",
				Limit: datadog.PtrInt64(10),
				Total: &datadogV2.CIAppGroupByTotal{
					CIAppGroupByTotalBoolean: datadog.PtrBool(false)},
			},
		},
		Options: &datadogV2.CIAppQueryOptions{
			Timezone: datadog.PtrString("GMT"),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewCIVisibilityPipelinesApi(apiClient)
	resp, r, err := api.AggregateCIAppPipelineEvents(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityPipelinesApi.AggregateCIAppPipelineEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityPipelinesApi.AggregateCIAppPipelineEvents`:\n%s\n", responseContent)
}
