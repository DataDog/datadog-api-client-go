// Create a new dashboard with funnel widget

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
		Title:      "Example-Dashboard with funnel widget",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					FunnelWidgetDefinition: &datadogV1.FunnelWidgetDefinition{
						Type: datadogV1.FUNNELWIDGETDEFINITIONTYPE_FUNNEL,
						Requests: []datadogV1.FunnelWidgetRequest{
							{
								Query: datadogV1.FunnelQuery{
									DataSource:  datadogV1.FUNNELSOURCE_RUM,
									QueryString: "",
									Steps:       []datadogV1.FunnelStep{},
								},
								RequestType: datadogV1.FUNNELREQUESTTYPE_FUNNEL,
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
