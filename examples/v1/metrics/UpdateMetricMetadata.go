// Edit metric metadata returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.MetricMetadata{
		PerUnit: datadog.PtrString("second"),
		Type:    datadog.PtrString("count"),
		Unit:    datadog.PtrString("byte"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewMetricsApi(apiClient)
	resp, r, err := api.UpdateMetricMetadata(ctx, "metric_name", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.UpdateMetricMetadata`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.UpdateMetricMetadata`:\n%s\n", responseContent)
}
