// Create a new dashboard with manage_status widget and show_priority parameter

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
					Width:  50,
					Height: 25,
				},
				Definition: datadogV1.WidgetDefinition{
					MonitorSummaryWidgetDefinition: &datadogV1.MonitorSummaryWidgetDefinition{
						Type:              datadogV1.MONITORSUMMARYWIDGETDEFINITIONTYPE_MANAGE_STATUS,
						SummaryType:       datadogV1.WIDGETSUMMARYTYPE_MONITORS.Ptr(),
						DisplayFormat:     datadogV1.WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS_AND_LIST.Ptr(),
						ColorPreference:   datadogV1.WIDGETCOLORPREFERENCE_TEXT.Ptr(),
						HideZeroCounts:    datadog.PtrBool(true),
						ShowLastTriggered: datadog.PtrBool(false),
						Query:             "",
						Sort:              datadogV1.WIDGETMONITORSUMMARYSORT_PRIORITY_ASCENDING.Ptr(),
						Count:             datadog.PtrInt64(50),
						Start:             datadog.PtrInt64(0),
						ShowPriority:      datadog.PtrBool(false),
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
