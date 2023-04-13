// Create a new dashboard returns "OK" response

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
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:      "Example-Dashboard with Profile Metrics Query",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					TimeseriesWidgetDefinition: &datadogV1.TimeseriesWidgetDefinition{
						Type: datadogV1.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
						Requests: []datadogV1.TimeseriesWidgetRequest{
							{
								ProfileMetricsQuery: &datadogV1.LogQueryDefinition{
									Compute: &datadogV1.LogsQueryCompute{
										Aggregation: "sum",
										Facet:       datadog.PtrString("@prof_core_cpu_cores"),
									},
									Search: &datadogV1.LogQueryDefinitionSearch{
										Query: "runtime:jvm",
									},
									GroupBy: []datadogV1.LogQueryDefinitionGroupBy{
										{
											Facet: "service",
											Limit: datadog.PtrInt64(10),
											Sort: &datadogV1.LogQueryDefinitionGroupBySort{
												Aggregation: "sum",
												Order:       datadogV1.WIDGETSORT_DESCENDING,
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
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
