// Aggregate tests events returns "OK" response

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
	body := datadogV2.CIAppTestsAggregateRequest{
		Compute: []datadogV2.CIAppCompute{
			{
				Aggregation: datadogV2.CIAPPAGGREGATIONFUNCTION_COUNT,
				Metric:      datadog.PtrString("@test.is_flaky"),
				Type:        datadogV2.CIAPPCOMPUTETYPE_TOTAL.Ptr(),
			},
		},
		Filter: &datadogV2.CIAppTestsQueryFilter{
			From:  datadog.PtrString("now-15m"),
			Query: datadog.PtrString("@language:(python OR go)"),
			To:    datadog.PtrString("now"),
		},
		GroupBy: []datadogV2.CIAppTestsGroupBy{
			{
				Facet: "@git.branch",
				Limit: datadog.PtrInt64(10),
				Sort: &datadogV2.CIAppAggregateSort{
					Order: datadogV2.CIAPPSORTORDER_ASCENDING.Ptr(),
				},
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
	api := datadogV2.NewCIVisibilityTestsApi(apiClient)
	resp, r, err := api.AggregateCIAppTestEvents(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIVisibilityTestsApi.AggregateCIAppTestEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CIVisibilityTestsApi.AggregateCIAppTestEvents`:\n%s\n", responseContent)
}
