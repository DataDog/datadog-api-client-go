// Delete a Dashboard Report returns "No Content" response

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	// there is a valid "dashboard" in the system
	DashboardID := os.Getenv("DASHBOARD_ID")

	// there is a valid "dashboard_report" in the system
	DashboardReportDataID := os.Getenv("DASHBOARD_REPORT_DATA_ID")

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardReportsApi(apiClient)
	r, err := api.DeleteDashboardReportConfig(ctx, DashboardID, DashboardReportDataID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardReportsApi.DeleteDashboardReportConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
