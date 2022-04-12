// Create a new dashboard returns "OK" response

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
										Facet:       datadog.PtrString("@prof_core_cpu_cores"),
									},
									Search: &datadog.LogQueryDefinitionSearch{
										Query: "runtime:jvm",
									},
									GroupBy: []datadog.LogQueryDefinitionGroupBy{
										{
											Facet: "service",
											Limit: datadog.PtrInt64(10),
											Sort: &datadog.LogQueryDefinitionGroupBySort{
												Aggregation: "sum",
												Order:       datadog.WIDGETSORT_DESCENDING,
												Facet:       datadog.PtrString("@prof_core_cpu_cores"),
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
