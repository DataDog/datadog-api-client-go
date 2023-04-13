// Create a new timeseries widget with ci_tests data source

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
		Title: "Example-Dashboard with ci_tests datasource",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					TimeseriesWidgetDefinition: &datadogV1.TimeseriesWidgetDefinition{
						Title:        datadog.PtrString(""),
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
									},
								},
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_CI_TESTS,
											Name:       "query1",
											Search: &datadogV1.FormulaAndFunctionEventQueryDefinitionSearch{
												Query: "test_level:test",
											},
											Indexes: []string{
												"*",
											},
											Compute: datadogV1.FormulaAndFunctionEventQueryDefinitionCompute{
												Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
											},
											GroupBy: []datadogV1.FormulaAndFunctionEventQueryGroupBy{},
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
