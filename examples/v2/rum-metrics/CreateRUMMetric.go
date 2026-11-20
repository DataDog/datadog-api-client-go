// Create a RUM-based metric returns "Created" response

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
	body := datadogV2.RUMMetricCreateRequest{
		Data: datadogV2.RUMMetricCreateData{
			Attributes: datadogV2.RUMMetricCreateAttributes{
				Compute: datadogV2.RUMMetricCompute{
					AggregationType:    datadogV2.RUMMETRICCOMPUTEAGGREGATIONTYPE_DISTRIBUTION,
					IncludePercentiles: datadog.PtrBool(true),
					Path:               datadog.PtrString("@duration"),
				},
				EventType: datadogV2.RUMMETRICEVENTTYPE_SESSION,
				Filter: &datadogV2.RUMMetricFilter{
					Query: "@service:web-ui",
				},
				GroupBy: []datadogV2.RUMMetricGroupBy{
					{
						Path:    "@browser.name",
						TagName: datadog.PtrString("browser_name"),
					},
				},
				Uniqueness: &datadogV2.RUMMetricUniqueness{
					When: datadogV2.RUMMETRICUNIQUENESSWHEN_WHEN_MATCH,
				},
			},
			Id:   "examplerummetric",
			Type: datadogV2.RUMMETRICTYPE_RUM_METRICS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMMetricsApi(apiClient)
	resp, r, err := api.CreateRUMMetric(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMMetricsApi.CreateRUMMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMMetricsApi.CreateRUMMetric`:\n%s\n", responseContent)
}
