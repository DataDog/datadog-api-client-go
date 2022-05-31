// Create a new dashboard with an audit logs query

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
		LayoutType: datadog.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:      "Example-Create_a_new_dashboard_with_an_audit_logs_query with Audit Logs Query",
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					TimeseriesWidgetDefinition: &datadog.TimeseriesWidgetDefinition{
						Type: datadog.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
						Requests: []datadog.TimeseriesWidgetRequest{
							{
								ResponseFormat: datadog.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
								Queries: []datadog.FormulaAndFunctionQueryDefinition{
									datadog.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionEventQueryDefinition: &datadog.FormulaAndFunctionEventQueryDefinition{
											Search: &datadog.FormulaAndFunctionEventQueryDefinitionSearch{
												Query: "",
											},
											DataSource: datadog.FORMULAANDFUNCTIONEVENTSDATASOURCE_AUDIT,
											Compute: datadog.FormulaAndFunctionEventQueryDefinitionCompute{
												Aggregation: datadog.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
											},
											Name: "query1",
											Indexes: []string{
												"*",
											},
											GroupBy: []datadog.FormulaAndFunctionEventQueryGroupBy{},
										}},
								},
							},
						},
					}},
				Layout: &datadog.WidgetLayout{
					X:      2,
					Y:      0,
					Width:  4,
					Height: 2,
				},
			},
		},
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
