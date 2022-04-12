// Create a new dashboard with formulas and functions scatterplot widget

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
		Title: "Example-Create_a_new_dashboard_with_formulas_and_functions_scatterplot_widget",
		Widgets: []datadog.Widget{
			{
				Id: datadog.PtrInt64(5346764334358972),
				Definition: datadog.WidgetDefinition{
					ScatterPlotWidgetDefinition: &datadog.ScatterPlotWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadog.SCATTERPLOTWIDGETDEFINITIONTYPE_SCATTERPLOT,
						Requests: datadog.ScatterPlotWidgetDefinitionRequests{
							Table: &datadog.ScatterplotTableRequest{
								Formulas: []datadog.ScatterplotWidgetFormula{
									{
										Formula:   "query1",
										Dimension: datadog.SCATTERPLOTDIMENSION_X,
										Alias:     datadog.PtrString("my-query1"),
									},
									{
										Formula:   "query2",
										Dimension: datadog.SCATTERPLOTDIMENSION_Y,
										Alias:     datadog.PtrString("my-query2"),
									},
								},
								Queries: []datadog.FormulaAndFunctionQueryDefinition{
									datadog.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadog.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadog.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query1",
											Query:      "avg:system.cpu.user{*} by {service}",
											Aggregator: datadog.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
										}},
									datadog.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadog.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadog.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query2",
											Query:      "avg:system.mem.used{*} by {service}",
											Aggregator: datadog.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
										}},
								},
								ResponseFormat: datadog.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
							},
						},
					}},
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  4,
					Height: 2,
				},
			},
		},
		LayoutType: datadog.DASHBOARDLAYOUTTYPE_ORDERED,
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
