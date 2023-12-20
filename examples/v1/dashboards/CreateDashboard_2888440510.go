// Create a new dashboard with split graph widget from distribution widget

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
							DistributionWidgetDefinition: &datadogV1.DistributionWidgetDefinition{
								Title:      datadog.PtrString(""),
								TitleSize:  datadog.PtrString("16"),
								TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
								Type:       datadogV1.DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION,
								Requests: []datadogV1.DistributionWidgetRequest{
									{
										ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
										Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
											datadogV1.FormulaAndFunctionQueryDefinition{
												FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
													DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
													Name:       "query1",
													Query:      "avg:system.cpu.user{*} by {service}",
													Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
												}},
										},
									},
								},
								Xaxis: &datadogV1.DistributionWidgetXAxis{
									Scale:       datadog.PtrString("linear"),
									IncludeZero: datadog.PtrBool(true),
									Min:         datadog.PtrString("auto"),
									Max:         datadog.PtrString("auto"),
								},
								Yaxis: &datadogV1.DistributionWidgetYAxis{
									Scale:       datadog.PtrString("linear"),
									IncludeZero: datadog.PtrBool(true),
									Min:         datadog.PtrString("auto"),
									Max:         datadog.PtrString("auto"),
								},
							}},
						SplitConfig: datadogV1.SplitConfig{
							SplitDimensions: []datadogV1.SplitDimension{
								{
									OneGraphPer: "service",
								},
							},
							Limit: 6,
							Sort: datadogV1.SplitSort{
								Compute: &datadogV1.SplitConfigSortCompute{
									Aggregation: "sum",
									Metric:      "system.cpu.user",
								},
								Order: datadogV1.WIDGETSORT_DESCENDING,
							},
						},
						Size: datadogV1.SPLITGRAPHVIZSIZE_MD,
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
