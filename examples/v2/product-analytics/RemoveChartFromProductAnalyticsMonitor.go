// Unlink a Product Analytics monitor from its chart returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	body := datadogV2.ProductAnalyticsMonitorUnlinkRequest{
		Data: datadogV2.ProductAnalyticsMonitorUnlinkData{
			Attributes: datadogV2.ProductAnalyticsMonitorUnlinkAttributes{
				ChartId:     "chart-1",
				DashboardId: datadog.PtrString("dash-1"),
			},
			Type: datadogV2.PRODUCTANALYTICSMONITORUNLINKREQUESTTYPE_PA_MONITOR_UNLINK_REQUEST,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.RemoveChartFromProductAnalyticsMonitor", true)
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewProductAnalyticsApi(apiClient)
	r, err := api.RemoveChartFromProductAnalyticsMonitor(ctx, "12345", body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsApi.RemoveChartFromProductAnalyticsMonitor`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
