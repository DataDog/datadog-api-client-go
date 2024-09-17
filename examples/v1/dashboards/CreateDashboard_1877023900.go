// Create a new dashboard with list_stream widget with a valid sort parameter ASC

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
		Title:      "Example-Dashboard with list_stream widget",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					ListStreamWidgetDefinition: &datadogV1.ListStreamWidgetDefinition{
						Type: datadogV1.LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM,
						Requests: []datadogV1.ListStreamWidgetRequest{
							{
								Columns: []datadogV1.ListStreamColumn{
									{
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
										Field: "timestamp",
									},
								},
								Query: datadogV1.ListStreamQuery{
									DataSource:  datadogV1.LISTSTREAMSOURCE_EVENT_STREAM,
									QueryString: "",
									EventSize:   datadogV1.WIDGETEVENTSIZE_LARGE.Ptr(),
									Sort: &datadogV1.WidgetFieldSort{
										Column: "timestamp",
										Order:  datadogV1.WIDGETSORT_ASCENDING,
									},
								},
								ResponseFormat: datadogV1.LISTSTREAMRESPONSEFORMAT_EVENT_LIST,
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
