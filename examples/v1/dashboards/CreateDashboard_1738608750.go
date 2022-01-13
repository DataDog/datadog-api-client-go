// Create a new dashboard with free_text widget

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
		Title:       "Example-Create_a_new_dashboard_with_free_text_widget",
		Description: *datadog.NewNullableString(nil),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  24,
					Height: 6,
				},
				Definition: datadog.WidgetDefinition{
					FreeTextWidgetDefinition: &datadog.FreeTextWidgetDefinition{
						Type:      datadog.FREETEXTWIDGETDEFINITIONTYPE_FREE_TEXT,
						Text:      "Example free text",
						Color:     datadog.PtrString("#4d4d4d"),
						FontSize:  datadog.PtrString("auto"),
						TextAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
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
