// Create a new dashboard returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/v2/api/v1/datadog"
)

func main() {
	body := datadog.Dashboard{
		LayoutType: datadog.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:      "Example-Create_a_new_dashboard_returns_OK_response with Profile Metrics Query",
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					TimeseriesWidgetDefinition: &datadog.TimeseriesWidgetDefinition{
						Type: datadog.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
						Requests: []datadog.TimeseriesWidgetRequest{
							{
								ProfileMetricsQuery: &datadog.LogQueryDefinition{
									Compute: &datadog.LogsQueryCompute{
										Aggregation: "sum",
										Facet:       common.PtrString("@prof_core_cpu_cores"),
									},
									Search: &datadog.LogQueryDefinitionSearch{
										Query: "runtime:jvm",
									},
									GroupBy: []datadog.LogQueryDefinitionGroupBy{
										{
											Facet: "service",
											Limit: common.PtrInt64(10),
											Sort: &datadog.LogQueryDefinitionGroupBySort{
												Aggregation: "sum",
												Order:       datadog.WIDGETSORT_DESCENDING,
												Facet:       common.PtrString("@prof_core_cpu_cores"),
											},
										},
									},
								},
							},
						},
					}},
			},
		},
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
