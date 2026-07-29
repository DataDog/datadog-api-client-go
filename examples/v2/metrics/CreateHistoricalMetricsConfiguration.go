// Enable historical metrics ingestion returns "OK" response

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
	body := datadogV2.HistoricalMetricsConfigurationCreateRequest{
		Data: datadogV2.HistoricalMetricsConfigurationCreateData{
			Id:   "dd.test.metric",
			Type: datadogV2.HISTORICALMETRICSCONFIGURATIONTYPE_HISTORICAL_METRICS_CONFIGURATIONS,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.CreateHistoricalMetricsConfiguration", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewMetricsApi(apiClient)
	resp, r, err := api.CreateHistoricalMetricsConfiguration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateHistoricalMetricsConfiguration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.CreateHistoricalMetricsConfiguration`:\n%s\n", responseContent)
}
