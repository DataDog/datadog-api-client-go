// Update a RUM-based metric returns "OK" response

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
	// there is a valid "rum_metric" in the system
	RumMetricDataID := os.Getenv("RUM_METRIC_DATA_ID")

	body := datadogV2.RUMMetricUpdateRequest{
		Data: datadogV2.RUMMetricUpdateData{
			Id:   datadog.PtrString(RumMetricDataID),
			Type: datadogV2.RUMMETRICTYPE_RUM_METRICS,
			Attributes: datadogV2.RUMMetricUpdateAttributes{
				Compute: &datadogV2.RUMMetricUpdateCompute{
					IncludePercentiles: datadog.PtrBool(false),
				},
				Filter: &datadogV2.RUMMetricFilter{
					Query: "@service:rum-config",
				},
				GroupBy: []datadogV2.RUMMetricGroupBy{
					{
						Path:    "@browser.version",
						TagName: datadog.PtrString("browser_version"),
					},
				},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewRUMMetricsApi(apiClient)
	resp, r, err := api.UpdateRUMMetric(ctx, RumMetricDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RUMMetricsApi.UpdateRUMMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `RUMMetricsApi.UpdateRUMMetric`:\n%s\n", responseContent)
}
