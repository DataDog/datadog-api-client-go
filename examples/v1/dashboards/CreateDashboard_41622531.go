// Create a new dashboard with timeseries widget and formula style attributes

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
		Title: "Example-Dashboard with formula style",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					TimeseriesWidgetDefinition: &datadogV1.TimeseriesWidgetDefinition{
						Title:        datadog.PtrString("styled timeseries"),
						ShowLegend:   datadog.PtrBool(true),
						LegendLayout: datadogV1.TIMESERIESWIDGETLEGENDLAYOUT_AUTO.Ptr(),
						LegendColumns: []datadogV1.TimeseriesWidgetLegendColumn{
							datadogV1.TIMESERIESWIDGETLEGENDCOLUMN_AVG,
							datadogV1.TIMESERIESWIDGETLEGENDCOLUMN_MIN,
							datadogV1.TIMESERIESWIDGETLEGENDCOLUMN_MAX,
							datadogV1.TIMESERIESWIDGETLEGENDCOLUMN_VALUE,
							datadogV1.TIMESERIESWIDGETLEGENDCOLUMN_SUM,
						},
						Time: &datadogV1.WidgetTime{},
						Type: datadogV1.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
						Requests: []datadogV1.TimeseriesWidgetRequest{
							{
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula: "query1",
										Style: &datadogV1.WidgetFormulaStyle{
											PaletteIndex: datadog.PtrInt64(4),
											Palette:      datadog.PtrString("classic"),
										},
									},
								},
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											Query:      "avg:system.cpu.user{*}",
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query1",
										}},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
								Style: &datadogV1.WidgetRequestStyle{
									Palette:   datadog.PtrString("dog_classic"),
									LineType:  datadogV1.WIDGETLINETYPE_SOLID.Ptr(),
									LineWidth: datadogV1.WIDGETLINEWIDTH_NORMAL.Ptr(),
								},
								DisplayType: datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
							},
						},
					}},
			},
		},
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		ReflowType: datadogV1.DASHBOARDREFLOWTYPE_AUTO.Ptr(),
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
