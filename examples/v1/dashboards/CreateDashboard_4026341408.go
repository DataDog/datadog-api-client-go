// Create a new dashboard with apm_issue_stream list_stream widget

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
		Title:      "Example-Create_a_new_dashboard_with_apm_issue_stream_list_stream_widget with list_stream widget",
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					ListStreamWidgetDefinition: &datadog.ListStreamWidgetDefinition{
						Type: datadog.LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM,
						Requests: []datadog.ListStreamWidgetRequest{
							{
								Columns: []datadog.ListStreamColumn{
									{
										Width: datadog.LISTSTREAMCOLUMNWIDTH_AUTO,
										Field: "timestamp",
									},
								},
								Query: datadog.ListStreamQuery{
									DataSource:  datadog.LISTSTREAMSOURCE_APM_ISSUE_STREAM,
									QueryString: "",
								},
								ResponseFormat: datadog.LISTSTREAMRESPONSEFORMAT_EVENT_LIST,
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
