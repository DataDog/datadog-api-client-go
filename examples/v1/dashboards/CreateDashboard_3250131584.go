// Create a new dashboard with event_timeline widget

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
		Title:       "Example-Create_a_new_dashboard_with_event_timeline_widget",
		Description: *datadog.NewNullableString(nil),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 9,
				},
				Definition: datadog.WidgetDefinition{
					EventTimelineWidgetDefinition: &datadog.EventTimelineWidgetDefinition{
						Title:         datadog.PtrString(""),
						TitleSize:     datadog.PtrString("16"),
						TitleAlign:    datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:          datadog.EVENTTIMELINEWIDGETDEFINITIONTYPE_EVENT_TIMELINE,
						Query:         "status:error priority:all",
						TagsExecution: datadog.PtrString("and"),
					}},
			},
		},
		TemplateVariables: []datadog.DashboardTemplateVariable{},
		LayoutType:        datadog.DASHBOARDLAYOUTTYPE_FREE,
		IsReadOnly:        datadog.PtrBool(false),
		NotifyList:        []string{},
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
