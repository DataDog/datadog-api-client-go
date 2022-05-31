// Create a new dashboard with timeseries widget containing style attributes

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Dashboard{
		LayoutType: datadog.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:      "Example-Create_a_new_dashboard_with_timeseries_widget_containing_style_attributes with timeseries widget",
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					TimeseriesWidgetDefinition: &datadog.TimeseriesWidgetDefinition{
						Type: datadog.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
						Requests: []datadog.TimeseriesWidgetRequest{
							{
								Q:            datadog.PtrString("sum:trace.test.errors{env:prod,service:datadog-api-spec} by {resource_name}.as_count()"),
								OnRightYaxis: datadog.PtrBool(false),
								Style: &datadog.WidgetRequestStyle{
									Palette:   datadog.PtrString("warm"),
									LineType:  datadog.WIDGETLINETYPE_SOLID.Ptr(),
									LineWidth: datadog.WIDGETLINEWIDTH_NORMAL.Ptr(),
								},
								DisplayType: datadog.WIDGETDISPLAYTYPE_BARS.Ptr(),
							},
						},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsApi.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
