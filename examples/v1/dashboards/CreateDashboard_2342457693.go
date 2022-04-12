// Create a new dashboard with scatterplot widget

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
		Title:       "Example-Create_a_new_dashboard_with_scatterplot_widget",
		Description: *datadog.NewNullableString(datadog.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 15,
				},
				Definition: datadog.WidgetDefinition{
					ScatterPlotWidgetDefinition: &datadog.ScatterPlotWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Time:       &datadog.WidgetTime{},
						Type:       datadog.SCATTERPLOTWIDGETDEFINITIONTYPE_SCATTERPLOT,
						Requests: datadog.ScatterPlotWidgetDefinitionRequests{
							Table: &datadog.ScatterplotTableRequest{
								Formulas: []datadog.ScatterplotWidgetFormula{
									{
										Formula:   "query1",
										Dimension: datadog.SCATTERPLOTDIMENSION_X,
										Alias:     datadog.PtrString(""),
									},
									{
										Formula:   "query2",
										Dimension: datadog.SCATTERPLOTDIMENSION_Y,
										Alias:     datadog.PtrString(""),
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
						Xaxis: &datadog.WidgetAxis{
							Scale:       datadog.PtrString("linear"),
							IncludeZero: datadog.PtrBool(true),
							Min:         datadog.PtrString("auto"),
							Max:         datadog.PtrString("auto"),
						},
						Yaxis: &datadog.WidgetAxis{
							Scale:       datadog.PtrString("linear"),
							IncludeZero: datadog.PtrBool(true),
							Min:         datadog.PtrString("auto"),
							Max:         datadog.PtrString("auto"),
						},
						ColorByGroups: []string{},
					}},
			},
		},
		TemplateVariables: []datadog.DashboardTemplateVariable{},
		LayoutType:        datadog.DASHBOARDLAYOUTTYPE_FREE,
		IsReadOnly:        datadog.PtrBool(false),
		NotifyList:        []string{},
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
