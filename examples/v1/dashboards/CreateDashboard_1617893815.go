// Create a new dashboard with formula and function distribution widget

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
		Title: "Example-Dashboard",
		Widgets: []datadogV1.Widget{
			{
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 15,
				},
				Definition: datadogV1.WidgetDefinition{
					DistributionWidgetDefinition: &datadogV1.DistributionWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Time: &datadogV1.WidgetTime{
							WidgetLegacyLiveSpan: &datadogV1.WidgetLegacyLiveSpan{}},
						Type: datadogV1.DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION,
						Requests: []datadogV1.DistributionWidgetRequest{
							{
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
											Name:       "query1",
											Search: &datadogV1.FormulaAndFunctionEventQueryDefinitionSearch{
												Query: "",
											},
											Indexes: []string{
												"*",
											},
											Compute: datadogV1.FormulaAndFunctionEventQueryDefinitionCompute{
												Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_AVG,
												Metric:      datadog.PtrString("@duration"),
											},
											GroupBy: []datadogV1.FormulaAndFunctionEventQueryGroupBy{
												{
													Facet: "service",
													Limit: datadog.PtrInt64(1000),
													Sort: &datadogV1.FormulaAndFunctionEventQueryGroupBySort{
														Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
														Order:       datadogV1.QUERYSORTORDER_DESC.Ptr(),
													},
												},
											},
											Storage: datadog.PtrString("hot"),
										}},
								},
							},
						},
					}},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_FREE,
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
