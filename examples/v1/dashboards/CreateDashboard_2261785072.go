// Create a new dashboard with a timeseries widget and an overlay request

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
	body := datadogV1.Dashboard{
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:      "Example-Dashboard",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					TimeseriesWidgetDefinition: &datadogV1.TimeseriesWidgetDefinition{
						Type: datadogV1.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
						Requests: []datadogV1.TimeseriesWidgetRequest{
							{
								OnRightYaxis: datadog.PtrBool(false),
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "mymetric",
											Query:      "avg:system.cpu.user{*}",
										}},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
								DisplayType:    datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
							},
							{
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "mymetricoverlay",
											Query:      "avg:system.cpu.user{*}",
										}},
								},
								Style: &datadogV1.WidgetRequestStyle{
									Palette:   datadog.PtrString("purple"),
									LineType:  datadogV1.WIDGETLINETYPE_SOLID.Ptr(),
									LineWidth: datadogV1.WIDGETLINEWIDTH_NORMAL.Ptr(),
								},
								DisplayType: datadogV1.WIDGETDISPLAYTYPE_OVERLAY.Ptr(),
							},
						},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
