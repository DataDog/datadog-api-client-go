// Create a new dashboard with hostmap widget

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Dashboard{
		Title:       "Example-Create_a_new_dashboard_with_hostmap_widget",
		Description: *common.NewNullableString(nil),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 22,
				},
				Definition: datadog.WidgetDefinition{
					HostMapWidgetDefinition: &datadog.HostMapWidgetDefinition{
						Title:      common.PtrString(""),
						TitleSize:  common.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadog.HOSTMAPWIDGETDEFINITIONTYPE_HOSTMAP,
						Requests: datadog.HostMapWidgetDefinitionRequests{
							Fill: &datadog.HostMapRequest{
								Q: common.PtrString("avg:system.cpu.user{*} by {host}"),
							},
						},
						NodeType:      datadog.WIDGETNODETYPE_HOST.Ptr(),
						NoMetricHosts: common.PtrBool(true),
						NoGroupHosts:  common.PtrBool(true),
						Style: &datadog.HostMapWidgetDefinitionStyle{
							Palette:     common.PtrString("green_to_orange"),
							PaletteFlip: common.PtrBool(false),
						},
					}},
			},
		},
		TemplateVariables: []datadog.DashboardTemplateVariable{},
		LayoutType:        datadog.DASHBOARDLAYOUTTYPE_FREE,
		IsReadOnly:        common.PtrBool(false),
		NotifyList:        []string{},
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
