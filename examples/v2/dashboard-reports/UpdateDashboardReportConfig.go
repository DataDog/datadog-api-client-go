// Update Dashboard Report returns "OK" response

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
	// there is a valid "dashboard" in the system
	DashboardID := os.Getenv("DASHBOARD_ID")

	// there is a valid "dashboard_report" in the system
	DashboardReportDataID := os.Getenv("DASHBOARD_REPORT_DATA_ID")

	body := datadogV2.DashboardReportUpdateRequest{
		Data: datadogV2.DashboardReportUpdateRequestData{
			Attributes: &datadogV2.DashboardReportUpdateAttributes{
				Schedule: &datadogV2.DashboardReportSchedule{
					Active: datadog.PtrBool(false),
				},
				Title: datadog.PtrString("Summary of performance for Web Application Backend"),
			},
			Id:   DashboardReportDataID,
			Type: datadogV2.DASHBOARDREPORTTYPE_DASHBOARD_REPORTS_TYPE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardReportsApi(apiClient)
	resp, r, err := api.UpdateDashboardReportConfig(ctx, DashboardID, DashboardReportDataID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardReportsApi.UpdateDashboardReportConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardReportsApi.UpdateDashboardReportConfig`:\n%s\n", responseContent)
}
