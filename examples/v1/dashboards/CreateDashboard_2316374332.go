// Create a new dashboard with alert_value widget

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
	// there is a valid "monitor" in the system

	body := datadog.Dashboard{
		Title:       "Example-Create_a_new_dashboard_with_alert_value_widget",
		Description: *common.NewNullableString(common.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  15,
					Height: 8,
				},
				Definition: datadog.WidgetDefinition{
					AlertValueWidgetDefinition: &datadog.AlertValueWidgetDefinition{
						Title:      common.PtrString(""),
						TitleSize:  common.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadog.ALERTVALUEWIDGETDEFINITIONTYPE_ALERT_VALUE,
						AlertId:    "7",
						Unit:       common.PtrString("auto"),
						TextAlign:  datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Precision:  common.PtrInt64(2),
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
