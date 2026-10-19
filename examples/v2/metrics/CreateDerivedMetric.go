// Create a derived metric returns "Created" response

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
	body := datadogV2.DerivedMetricCreateRequest{
		Data: datadogV2.DerivedMetricCreateData{
			Attributes: datadogV2.DerivedMetricAttributes{
				Definition: datadogV2.DerivedMetricDefinition{
					Formula: "a",
					Queries: []datadogV2.DerivedMetricQuery{
						{
							DataSource: "metrics",
							Name:       "a",
							Query:      "avg:system.cpu.user{*}",
						},
					},
				},
			},
			Id:   "test.derived.examplemetric",
			Type: datadogV2.DERIVEDMETRICSTYPE_DERIVED_METRICS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.CreateDerivedMetric(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateDerivedMetric`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateDerivedMetric`:\n%s\n", responseContent)
}
