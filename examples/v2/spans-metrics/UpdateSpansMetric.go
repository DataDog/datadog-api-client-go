// Update a span-based metric returns "OK" response

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
	// there is a valid "spans_metric" in the system
	SpansMetricDataID := os.Getenv("SPANS_METRIC_DATA_ID")

	body := datadogV2.SpansMetricUpdateRequest{
		Data: datadogV2.SpansMetricUpdateData{
			Attributes: datadogV2.SpansMetricUpdateAttributes{
				Compute: &datadogV2.SpansMetricUpdateCompute{
					IncludePercentiles: datadog.PtrBool(false),
				},
				Filter: &datadogV2.SpansMetricFilter{
					Query: datadog.PtrString("@http.status_code:200 service:my-service-updated"),
				},
				GroupBy: []datadogV2.SpansMetricGroupBy{
					{
						Path:    "resource_name",
						TagName: datadog.PtrString("resource_name"),
					},
				},
			},
			Type: datadogV2.SPANSMETRICTYPE_SPANS_METRICS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewSpansMetricsApi(apiClient)
	resp, r, err := api.UpdateSpansMetric(ctx, SpansMetricDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpansMetricsApi.UpdateSpansMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `SpansMetricsApi.UpdateSpansMetric`:\n%s\n", responseContent)
}
