// Create a new dashboard with apm resource stats widget

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
		Title: "Example-Create_a_new_dashboard_with_apm_resource_stats_widget",
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					TableWidgetDefinition: &datadog.TableWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadog.TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE,
						Requests: []datadog.TableWidgetRequest{
							{
								ResponseFormat: datadog.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
								Queries: []datadog.FormulaAndFunctionQueryDefinition{
									datadog.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionApmResourceStatsQueryDefinition: &datadog.FormulaAndFunctionApmResourceStatsQueryDefinition{
											PrimaryTagValue: datadog.PtrString("edge-eu1.prod.dog"),
											Stat:            datadog.FORMULAANDFUNCTIONAPMRESOURCESTATNAME_HITS,
											Name:            "query1",
											Service:         "cassandra",
											DataSource:      datadog.FORMULAANDFUNCTIONAPMRESOURCESTATSDATASOURCE_APM_RESOURCE_STATS,
											Env:             "ci",
											PrimaryTagName:  datadog.PtrString("datacenter"),
											OperationName:   datadog.PtrString("cassandra.query"),
											GroupBy: []string{
												"resource_name",
											},
										}},
								},
							},
						},
					}},
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  4,
					Height: 4,
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
