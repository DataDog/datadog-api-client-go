// Create a new dashboard with funnel widget

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
		Title:      "Example-Create_a_new_dashboard_with_funnel_widget with funnel widget",
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					FunnelWidgetDefinition: &datadog.FunnelWidgetDefinition{
						Type: datadog.FUNNELWIDGETDEFINITIONTYPE_FUNNEL,
						Requests: []datadog.FunnelWidgetRequest{
							{
								Query: datadog.FunnelQuery{
									DataSource:  datadog.FUNNELSOURCE_RUM,
									QueryString: "",
									Steps:       []datadog.FunnelStep{},
								},
								RequestType: datadog.FUNNELREQUESTTYPE_FUNNEL,
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
