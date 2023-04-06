// Create a new dashboard with a query value widget using the percentile aggregator

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
		Title:      "Example-Dashboard with QVW Percentile Aggregator",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					QueryValueWidgetDefinition: &datadogV1.QueryValueWidgetDefinition{
						TitleSize:  datadog.PtrString("16"),
						Title:      datadog.PtrString(""),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Precision:  datadog.PtrInt64(2),
						Time:       &datadogV1.WidgetTime{},
						Autoscale:  datadog.PtrBool(true),
						Requests: []datadogV1.QueryValueWidgetRequest{
							{
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula: "query1",
									},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											Query:      "p90:dist.dd.dogweb.latency{*}",
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query1",
											Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_PERCENTILE.Ptr(),
										}},
								},
							},
						},
						Type: datadogV1.QUERYVALUEWIDGETDEFINITIONTYPE_QUERY_VALUE,
					}},
				Layout: &datadogV1.WidgetLayout{
					Y:      0,
					X:      0,
					Height: 2,
					Width:  2,
				},
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
