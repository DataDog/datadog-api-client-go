// Update a Product Analytics monitor returns "OK" response

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
	body := datadogV2.ProductAnalyticsMonitorUpdateRequest{
		Data: datadogV2.ProductAnalyticsMonitorUpdateData{
			Attributes: datadogV2.ProductAnalyticsMonitorUpdateAttributes{
				Name: datadog.PtrString("Funnel conversion drop"),
				Options: &datadogV2.ProductAnalyticsMonitorOptions{
					Thresholds: &datadogV2.ProductAnalyticsMonitorThresholds{
						Critical: datadog.PtrFloat64(1),
						Warning:  datadog.PtrFloat64(1.5),
					},
				},
			},
			Type: datadogV2.PRODUCTANALYTICSMONITORREQUESTTYPE_PA_MONITOR_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateProductAnalyticsMonitor", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	resp, r, err := api.UpdateProductAnalyticsMonitor(ctx, "12345", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.UpdateProductAnalyticsMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsApi.UpdateProductAnalyticsMonitor`:\n%s\n", responseContent)
}
