// Create a span-based metric returns "OK" response

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
	body := datadogV2.SpansMetricCreateRequest{
		Data: datadogV2.SpansMetricCreateData{
			Attributes: datadogV2.SpansMetricCreateAttributes{
				Compute: datadogV2.SpansMetricCompute{
					AggregationType:    datadogV2.SPANSMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
					IncludePercentiles: datadog.PtrBool(false),
					Path:               datadog.PtrString("@duration"),
				},
				Filter: &datadogV2.SpansMetricFilter{
					Query: datadog.PtrString("@http.status_code:200 service:my-service"),
				},
				GroupBy: []datadogV2.SpansMetricGroupBy{
					{
						Path:    "resource_name",
						TagName: datadog.PtrString("resource_name"),
					},
				},
			},
			Id:   "ExampleSpansMetric",
			Type: datadogV2.SPANSMETRICTYPE_SPANS_METRICS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSpansMetricsApi(apiClient)
	resp, r, err := api.CreateSpansMetric(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpansMetricsApi.CreateSpansMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SpansMetricsApi.CreateSpansMetric`:\n%s\n", responseContent)
}
