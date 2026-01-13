// Create a new dashboard with heatmap widget with markers and num_buckets

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
		Title: "Example-Dashboard",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					HeatMapWidgetDefinition: &datadogV1.HeatMapWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.HEATMAPWIDGETDEFINITIONTYPE_HEATMAP,
						Xaxis: &datadogV1.HeatMapWidgetXAxis{
							NumBuckets: datadog.PtrInt64(75),
						},
						Yaxis: &datadogV1.WidgetAxis{
							Scale:       datadog.PtrString("linear"),
							Min:         datadog.PtrString("auto"),
							Max:         datadog.PtrString("auto"),
							IncludeZero: datadog.PtrBool(true),
						},
						Markers: []datadogV1.WidgetMarker{
							{
								DisplayType: datadog.PtrString("percentile"),
								Value:       "50",
							},
							{
								DisplayType: datadog.PtrString("percentile"),
								Value:       "99",
							},
						},
						Requests: []datadogV1.HeatMapWidgetRequest{
							{
								RequestType: datadogV1.WIDGETHISTOGRAMREQUESTTYPE_HISTOGRAM.Ptr(),
								Query: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
									DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
									Name:       "query1",
									Query:      "histogram:trace.servlet.request{*}",
								},
							},
						},
					}},
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  4,
					Height: 4,
				},
			},
		},
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
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
