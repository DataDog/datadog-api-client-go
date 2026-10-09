// Create a heatgrid widget with a preset gradient

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
		Title:      "Example-Dashboard",
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					HeatgridWidgetDefinition: &datadogV1.HeatgridWidgetDefinition{
						Type: datadogV1.HEATGRIDWIDGETDEFINITIONTYPE_HEATGRID,
						Requests: []datadogV1.HeatgridWidgetRequest{
							{
								ResponseFormat: datadogV1.HEATGRIDWIDGETRESPONSEFORMAT_TIMESERIES,
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query1",
											Query:      "avg:system.cpu.user{*} by {host}",
										}},
								},
								Formulas: []datadogV1.HeatgridWidgetFormula{
									{
										Formula: "query1",
									},
								},
							},
						},
						Sort: datadogV1.HeatgridSort{
							NestingDisplay: datadogV1.HEATGRIDNESTINGDISPLAY_FLAT,
							SortBy: datadogV1.HeatgridSortBy{
								HeatgridSortByValue: &datadogV1.HeatgridSortByValue{
									Property:    datadogV1.HEATGRIDSORTBYVALUEPROPERTY_VALUE,
									Order:       datadogV1.HEATGRIDSORTORDER_DESC,
									Aggregation: datadogV1.HEATGRIDSORTAGGREGATION_AVG,
								}},
						},
						Color: &datadogV1.HeatgridColorConfig{
							HeatgridGradientPresetColor: &datadogV1.HeatgridGradientPresetColor{
								Mode:       datadogV1.HEATGRIDGRADIENTMODE_GRADIENT,
								Source:     datadogV1.HEATGRIDPRESETCOLORSOURCE_PRESET,
								PresetName: "classic",
							}},
						Legend: &datadogV1.HeatgridLegend{
							ShowCaption: datadog.PtrBool(true),
						},
						LabelColumn: &datadogV1.HeatgridLabelColumn{
							Width: datadogV1.HEATGRIDLABELCOLUMNWIDTH_M,
						},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
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
