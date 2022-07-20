// Create a new dashboard with trace_service widget

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
		Title:       "Example-Create_a_new_dashboard_with_trace_service_widget",
		Description: *common.NewNullableString(common.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  72,
					Height: 72,
				},
				Definition: datadog.WidgetDefinition{
					ServiceSummaryWidgetDefinition: &datadog.ServiceSummaryWidgetDefinition{
						Title:            common.PtrString("Service Summary"),
						Time:             &datadog.WidgetTime{},
						Type:             datadog.SERVICESUMMARYWIDGETDEFINITIONTYPE_TRACE_SERVICE,
						Env:              "none",
						Service:          "",
						SpanName:         "",
						ShowHits:         common.PtrBool(true),
						ShowErrors:       common.PtrBool(true),
						ShowLatency:      common.PtrBool(true),
						ShowBreakdown:    common.PtrBool(true),
						ShowDistribution: common.PtrBool(true),
						ShowResourceList: common.PtrBool(false),
						SizeFormat:       datadog.WIDGETSIZEFORMAT_MEDIUM.Ptr(),
						DisplayFormat:    datadog.WIDGETSERVICESUMMARYDISPLAYFORMAT_TWO_COLUMN.Ptr(),
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
