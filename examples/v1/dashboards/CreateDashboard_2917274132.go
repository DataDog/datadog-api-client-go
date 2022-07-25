// Create a new dashboard with manage_status widget

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
		Title:       "Example-Create_a_new_dashboard_with_manage_status_widget",
		Description: *common.NewNullableString(common.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  50,
					Height: 25,
				},
				Definition: datadog.WidgetDefinition{
					MonitorSummaryWidgetDefinition: &datadog.MonitorSummaryWidgetDefinition{
						Type:              datadog.MONITORSUMMARYWIDGETDEFINITIONTYPE_MANAGE_STATUS,
						SummaryType:       datadog.WIDGETSUMMARYTYPE_MONITORS.Ptr(),
						DisplayFormat:     datadog.WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS_AND_LIST.Ptr(),
						ColorPreference:   datadog.WIDGETCOLORPREFERENCE_TEXT.Ptr(),
						HideZeroCounts:    common.PtrBool(true),
						ShowLastTriggered: common.PtrBool(false),
						Query:             "",
						Sort:              datadog.WIDGETMONITORSUMMARYSORT_STATUS_ASCENDING.Ptr(),
						Count:             common.PtrInt64(50),
						Start:             common.PtrInt64(0),
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
