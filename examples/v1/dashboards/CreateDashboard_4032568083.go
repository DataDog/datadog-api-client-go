// Create a new dashboard with a live default_timeframe returns "OK" response

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
		Title:      "Example-Dashboard",
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					NoteWidgetDefinition: &datadogV1.NoteWidgetDefinition{
						Type:            datadogV1.NOTEWIDGETDEFINITIONTYPE_NOTE,
						Content:         "test",
						BackgroundColor: datadog.PtrString("white"),
						FontSize:        datadog.PtrString("14"),
						TextAlign:       datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						ShowTick:        datadog.PtrBool(false),
						TickPos:         datadog.PtrString("50%"),
						TickEdge:        datadogV1.WIDGETTICKEDGE_LEFT.Ptr(),
					}},
			},
		},
		DefaultTimeframe: &datadogV1.DashboardDefaultTimeframeSetting{
			DashboardLiveTimeframe: &datadogV1.DashboardLiveTimeframe{
				Type:  datadogV1.DASHBOARDLIVETIMEFRAMETYPE_LIVE,
				Unit:  datadogV1.WIDGETLIVESPANUNIT_HOUR,
				Value: 4,
			}},
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
