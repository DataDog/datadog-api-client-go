// Create a new dashboard with a query value widget using timeseries background

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	body := datadog.Dashboard{
		LayoutType: datadog.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:      "Example-Create_a_new_dashboard_with_a_query_value_widget_using_timeseries_background with QVW Timeseries Background",
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					QueryValueWidgetDefinition: &datadog.QueryValueWidgetDefinition{
						TitleSize:  common.PtrString("16"),
						Title:      common.PtrString(""),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Precision:  common.PtrInt64(2),
						Time:       &datadog.WidgetTime{},
						Autoscale:  common.PtrBool(true),
						Requests: []datadog.QueryValueWidgetRequest{
							{
								Formulas: []datadog.WidgetFormula{
									{
										Formula: "query1",
									},
								},
								ResponseFormat: datadog.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
								Queries: []datadog.FormulaAndFunctionQueryDefinition{
									datadog.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadog.FormulaAndFunctionMetricQueryDefinition{
											Query:      "sum:my.cool.count.metric{*}",
											DataSource: datadog.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query1",
											Aggregator: datadog.FORMULAANDFUNCTIONMETRICAGGREGATION_PERCENTILE.Ptr(),
										}},
								},
							},
						},
						Type: datadog.QUERYVALUEWIDGETDEFINITIONTYPE_QUERY_VALUE,
						TimeseriesBackground: &datadog.TimeseriesBackground{
							Type: datadog.TIMESERIESBACKGROUNDTYPE_AREA,
							Yaxis: &datadog.WidgetAxis{
								IncludeZero: common.PtrBool(true),
							},
						},
					}},
				Layout: &datadog.WidgetLayout{
					Y:      0,
					X:      0,
					Height: 2,
					Width:  2,
				},
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
