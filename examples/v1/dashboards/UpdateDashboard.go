// Update a dashboard returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	// there is a valid "dashboard" in the system
	DashboardID := os.Getenv("DASHBOARD_ID")

	body := datadog.Dashboard{
		LayoutType:  datadog.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:       "Example-Update_a_dashboard_returns_OK_response with list_stream widget",
		Description: *datadog.NewNullableString(datadog.PtrString("Updated description")),
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
	resp, r, err := apiClient.DashboardsApi.UpdateDashboard(ctx, DashboardID, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.UpdateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.UpdateDashboard`:\n%s\n", responseContent)
}
