// Create a new Dashboard Report returns "Created" response

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

	body := datadogV2.DashboardReportCreateRequest{
		Data: datadogV2.DashboardReportCreate{
			Attributes: &datadogV2.DashboardReportCreateAttributes{
				Description: *datadog.NewNullableString(datadog.PtrString("This report summarizes the recent errors, latency, and uptime of our Web Application Backend.")),
				Destinations: datadogV2.DashboardReportDestination{
					Email: &datadogV2.DashboardReportDestinationEmail{
						RecipientAddresses: []string{
							"jane.doe@example.com",
						},
					},
				},
				Schedule: datadogV2.DashboardReportSchedule{
					Active:             datadog.PtrBool(true),
					Frequency:          datadogV2.DASHBOARDREPORTSCHEDULEFREQUENCY_DASHBOARD_REPORT_SCHEDULE_FREQUENCY_1D.Ptr(),
					RepeatAt:           datadogV2.DASHBOARDREPORTSCHEDULEREPEATAT_DASHBOARD_REPORT_SCHEDULE_REPEAT_AT_13_30.Ptr(),
					RepeatOnDayOfMonth: *datadogV2.NewNullableDashboardReportScheduleRepeatOnDayOfMonth(nil),
					RepeatOnDayOfWeek:  *datadogV2.NewNullableDashboardReportScheduleRepeatOnDayOfWeek(nil),
					Timezone:           datadog.PtrString("America/New_York"),
				},
				Timeframe: datadogV2.DASHBOARDREPORTTIMEFRAME_DASHBOARD_REPORT_TIMEFRAME_1W,
				Title:     "Summary of recent stability and performance for Web Application Backend",
			},
			Type: datadogV2.DASHBOARDREPORTTYPE_DASHBOARD_REPORTS_TYPE,
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewDashboardReportsApi(apiClient)
	resp, r, err := api.CreateDashboardReportConfig(ctx, DashboardID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardReportsApi.CreateDashboardReportConfig`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardReportsApi.CreateDashboardReportConfig`:\n%s\n", responseContent)
}
