// Create a new dashboard with check_status widget

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
		Title:       "Example-Dashboard",
		Description: *datadog.NewNullableString(datadog.PtrString("")),
		Widgets: []datadogV1.Widget{
			{
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  15,
					Height: 8,
				},
				Definition: datadogV1.WidgetDefinition{
					CheckStatusWidgetDefinition: &datadogV1.CheckStatusWidgetDefinition{
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.CHECKSTATUSWIDGETDEFINITIONTYPE_CHECK_STATUS,
						Check:      "datadog.agent.up",
						Grouping:   datadogV1.WIDGETGROUPING_CHECK,
						Tags: []string{
							"*",
						},
					}},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_FREE,
		IsReadOnly:        datadog.PtrBool(false),
		NotifyList:        *datadog.NewNullableList(&[]string{}),
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
