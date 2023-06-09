// Create a new dashboard with trace_service widget

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
					Width:  72,
					Height: 72,
				},
				Definition: datadogV1.WidgetDefinition{
					ServiceSummaryWidgetDefinition: &datadogV1.ServiceSummaryWidgetDefinition{
						Title:            datadog.PtrString("Service Summary"),
						Time:             &datadogV1.WidgetTime{},
						Type:             datadogV1.SERVICESUMMARYWIDGETDEFINITIONTYPE_TRACE_SERVICE,
						Env:              "none",
						Service:          "",
						SpanName:         "",
						ShowHits:         datadog.PtrBool(true),
						ShowErrors:       datadog.PtrBool(true),
						ShowLatency:      datadog.PtrBool(true),
						ShowBreakdown:    datadog.PtrBool(true),
						ShowDistribution: datadog.PtrBool(true),
						ShowResourceList: datadog.PtrBool(false),
						SizeFormat:       datadogV1.WIDGETSIZEFORMAT_MEDIUM.Ptr(),
						DisplayFormat:    datadogV1.WIDGETSERVICESUMMARYDISPLAYFORMAT_TWO_COLUMN.Ptr(),
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
