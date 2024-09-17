// Create a new dashboard with apm resource stats widget

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
				Definition: datadogV1.WidgetDefinition{
					TableWidgetDefinition: &datadogV1.TableWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE,
						Requests: []datadogV1.TableWidgetRequest{
							{
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionApmResourceStatsQueryDefinition: &datadogV1.FormulaAndFunctionApmResourceStatsQueryDefinition{
											PrimaryTagValue: datadog.PtrString("edge-eu1.prod.dog"),
											Stat:            datadogV1.FORMULAANDFUNCTIONAPMRESOURCESTATNAME_HITS,
											Name:            "query1",
											Service:         "cassandra",
											DataSource:      datadogV1.FORMULAANDFUNCTIONAPMRESOURCESTATSDATASOURCE_APM_RESOURCE_STATS,
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
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  4,
					Height: 4,
				},
			},
		},
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
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
