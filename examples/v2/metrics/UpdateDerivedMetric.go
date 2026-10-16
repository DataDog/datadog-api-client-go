// Update a derived metric returns "OK" response

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
	// there is a valid "derived_metric" in the system
	DerivedMetricDataID := os.Getenv("DERIVED_METRIC_DATA_ID")

	body := datadogV2.DerivedMetricCreateRequest{
		Data: datadogV2.DerivedMetricCreateData{
			Attributes: datadogV2.DerivedMetricAttributes{
				Definition: datadogV2.DerivedMetricDefinition{
					Formula: "a * 2",
					Queries: []datadogV2.DerivedMetricQuery{
						{
							DataSource: "metrics",
							Name:       "a",
							Query:      "avg:system.cpu.user{*}",
						},
					},
				},
			},
			Id:   DerivedMetricDataID,
			Type: datadogV2.DERIVEDMETRICSTYPE_DERIVED_METRICS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.UpdateDerivedMetric(ctx, DerivedMetricDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateDerivedMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.UpdateDerivedMetric`:\n%s\n", responseContent)
}
