// Create a heatgrid widget with custom discrete thresholds

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
								HeatgridSortByLabel: &datadogV1.HeatgridSortByLabel{
									Property: datadogV1.HEATGRIDSORTBYLABELPROPERTY_LABEL,
									Order:    datadogV1.HEATGRIDSORTORDER_ASC,
								}},
						},
						Color: &datadogV1.HeatgridColorConfig{
							HeatgridDiscreteCustomColor: &datadogV1.HeatgridDiscreteCustomColor{
								Mode:   datadogV1.HEATGRIDDISCRETEMODE_DISCRETE,
								Source: datadogV1.HEATGRIDCUSTOMCOLORSOURCE_CUSTOM,
								Bins: []datadogV1.HeatgridColorBin{
									{
										Color: datadogV1.HeatgridColor{
											String: datadog.PtrString("#00FF00")},
									},
									{
										Color: datadogV1.HeatgridColor{
											String: datadog.PtrString("#FF0000")},
										LowerBound: datadog.PtrFloat64(80),
									},
								},
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
