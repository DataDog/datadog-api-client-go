// Create a new dashboard with servicemap widget

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
		Title:       "Example-Create_a_new_dashboard_with_servicemap_widget",
		Description: *common.NewNullableString(common.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 15,
				},
				Definition: datadog.WidgetDefinition{
					ServiceMapWidgetDefinition: &datadog.ServiceMapWidgetDefinition{
						Title:      common.PtrString(""),
						TitleSize:  common.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadog.SERVICEMAPWIDGETDEFINITIONTYPE_SERVICEMAP,
						Service:    "",
						Filters: []string{
							"env:none",
							"environment:*",
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
