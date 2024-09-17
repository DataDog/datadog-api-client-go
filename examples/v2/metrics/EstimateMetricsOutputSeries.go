// Tag Configuration Cardinality Estimator returns "Success" response

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
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.EstimateMetricsOutputSeries(ctx, "system.cpu.idle", *datadogV2.NewEstimateMetricsOutputSeriesOptionalParameters().WithFilterGroups("app,host").WithFilterNumAggregations(4))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.EstimateMetricsOutputSeries`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.EstimateMetricsOutputSeries`:\n%s\n", responseContent)
}
