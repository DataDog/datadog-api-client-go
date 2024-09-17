// Aggregate spans returns "OK" response

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
	body := datadogV2.SpansAggregateRequest{
		Data: &datadogV2.SpansAggregateData{
			Attributes: &datadogV2.SpansAggregateRequestAttributes{
				Compute: []datadogV2.SpansCompute{
					{
						Aggregation: datadogV2.SPANSAGGREGATIONFUNCTION_COUNT,
						Interval:    datadog.PtrString("5m"),
						Type:        datadogV2.SPANSCOMPUTETYPE_TIMESERIES.Ptr(),
					},
				},
				Filter: &datadogV2.SpansQueryFilter{
					From:  datadog.PtrString("now-15m"),
					Query: datadog.PtrString("*"),
					To:    datadog.PtrString("now"),
				},
			},
			Type: datadogV2.SPANSAGGREGATEREQUESTTYPE_AGGREGATE_REQUEST.Ptr(),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSpansApi(apiClient)
	resp, r, err := api.AggregateSpans(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpansApi.AggregateSpans`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SpansApi.AggregateSpans`:\n%s\n", responseContent)
}
