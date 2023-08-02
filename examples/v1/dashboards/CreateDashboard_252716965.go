// Create a distribution widget using a histogram request containing a formulas and functions metrics query

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
					DistributionWidgetDefinition: &datadogV1.DistributionWidgetDefinition{
						Title:      datadog.PtrString("Metrics HOP"),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						ShowLegend: datadog.PtrBool(false),
						Type:       datadogV1.DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION,
						CustomLinks: []datadogV1.WidgetCustomLink{
							{
								Label: datadog.PtrString("Example"),
								Link:  datadog.PtrString("https://example.org/"),
							},
						},
						Xaxis: &datadogV1.DistributionWidgetXAxis{
							Max:         datadog.PtrString("auto"),
							IncludeZero: datadog.PtrBool(true),
							Scale:       datadog.PtrString("linear"),
							Min:         datadog.PtrString("auto"),
						},
						Yaxis: &datadogV1.DistributionWidgetYAxis{
							Max:         datadog.PtrString("auto"),
							IncludeZero: datadog.PtrBool(true),
							Scale:       datadog.PtrString("linear"),
							Min:         datadog.PtrString("auto"),
						},
						Requests: []datadogV1.DistributionWidgetRequest{
							{
								Query: &datadogV1.DistributionWidgetHistogramRequestQuery{
									FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
										Query:      "histogram:trace.Load{*}",
										DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
										Name:       "query1",
									}},
								RequestType: datadogV1.DISTRIBUTIONWIDGETHISTOGRAMREQUESTTYPE_HISTOGRAM.Ptr(),
								Style: &datadogV1.WidgetStyle{
									Palette: datadog.PtrString("dog_classic"),
								},
							},
						},
					}},
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  4,
					Height: 2,
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
