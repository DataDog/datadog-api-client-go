// Create a new dashboard with event_stream widget

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
		Title:       "Example-Create_a_new_dashboard_with_event_stream_widget",
		Description: *datadog.NewNullableString(datadog.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 38,
				},
				Definition: datadog.WidgetDefinition{
					EventStreamWidgetDefinition: &datadog.EventStreamWidgetDefinition{
						Title:         datadog.PtrString(""),
						TitleSize:     datadog.PtrString("16"),
						TitleAlign:    datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:          datadog.EVENTSTREAMWIDGETDEFINITIONTYPE_EVENT_STREAM,
						Query:         "example-query",
						TagsExecution: datadog.PtrString("and"),
						EventSize:     datadog.WIDGETEVENTSIZE_SMALL.Ptr(),
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
