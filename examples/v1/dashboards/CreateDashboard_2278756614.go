// Create a new dashboard with split graph widget

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
		Title:       "Example-Dashboard",
		Description: *datadog.NewNullableString(datadog.PtrString("")),
		Widgets: []datadogV1.Widget{
			{
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  12,
					Height: 8,
				},
				Definition: datadogV1.WidgetDefinition{
					SplitGraphWidgetDefinition: &datadogV1.SplitGraphWidgetDefinition{
						Title: datadog.PtrString(""),
						Type:  datadogV1.SPLITGRAPHWIDGETDEFINITIONTYPE_SPLIT_GROUP,
						SourceWidgetDefinition: datadogV1.SplitGraphSourceWidgetDefinition{
							TimeseriesWidgetDefinition: &datadogV1.TimeseriesWidgetDefinition{
								Title:      datadog.PtrString(""),
								TitleSize:  datadog.PtrString("16"),
								TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
								Type:       datadogV1.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
								Requests: []datadogV1.TimeseriesWidgetRequest{
									{
										ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
										Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
											datadogV1.FormulaAndFunctionQueryDefinition{
												FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
													Name:       "query1",
													DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
													Query:      "avg:system.cpu.user{*}",
												}},
										},
										Style: &datadogV1.WidgetRequestStyle{
											Palette:   datadog.PtrString("dog_classic"),
											LineType:  datadogV1.WIDGETLINETYPE_SOLID.Ptr(),
											LineWidth: datadogV1.WIDGETLINEWIDTH_NORMAL.Ptr(),
										},
										DisplayType: datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
									},
								},
							}},
						SplitConfig: datadogV1.SplitConfig{
							SplitDimensions: []datadogV1.SplitDimension{
								{
									OneGraphPer: "service",
								},
							},
							Limit: 24,
							Sort: datadogV1.SplitSort{
								Compute: &datadogV1.SplitConfigSortCompute{
									Aggregation: "sum",
									Metric:      "system.cpu.user",
								},
								Order: datadogV1.WIDGETSORT_DESCENDING,
							},
							StaticSplits: [][]datadogV1.SplitVectorEntryItem{
								{
									{
										TagKey: "service",
										TagValues: []string{
											"cassandra",
										},
									},
									{
										TagKey:    "datacenter",
										TagValues: []string{},
									},
								},
								{
									{
										TagKey: "demo",
										TagValues: []string{
											"env",
										},
									},
								},
							},
						},
						Size:            datadogV1.SPLITGRAPHVIZSIZE_MD,
						HasUniformYAxes: datadog.PtrBool(true),
					}},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		IsReadOnly:        datadog.PtrBool(false),
		NotifyList:        *datadog.NewNullableList(&[]string{}),
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
