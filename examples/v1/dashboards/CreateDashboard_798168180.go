// Create a new dashboard with apm dependency stats widget

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Dashboard{
		Title: "Example-Create_a_new_dashboard_with_apm_dependency_stats_widget",
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
										FormulaAndFunctionApmDependencyStatsQueryDefinition: &datadog.FormulaAndFunctionApmDependencyStatsQueryDefinition{
											PrimaryTagValue: datadog.PtrString("edge-eu1.prod.dog"),
											Stat:            datadog.FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_AVG_DURATION,
											ResourceName:    "DELETE FROM monitor_history.monitor_state_change_history WHERE org_id = ? AND monitor_id IN ? AND group = ?",
											Name:            "query1",
											Service:         "cassandra",
											DataSource:      datadog.FORMULAANDFUNCTIONAPMDEPENDENCYSTATSDATASOURCE_APM_DEPENDENCY_STATS,
											Env:             "ci",
											PrimaryTagName:  datadog.PtrString("datacenter"),
											OperationName:   "cassandra.query",
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
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
