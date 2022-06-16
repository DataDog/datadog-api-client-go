// Create a distribution widget using a histogram request containing a formulas and functions APM Stats query

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	body := datadog.Dashboard{
		Title:       "Example-Create_a_distribution_widget_using_a_histogram_request_containing_a_formulas_and_functions_APM_Stats",
		Description: *datadog.NewNullableString(datadog.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					DistributionWidgetDefinition: &datadog.DistributionWidgetDefinition{
						Title:      datadog.PtrString("APM Stats - Request latency HOP"),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						ShowLegend: datadog.PtrBool(false),
						Type:       datadog.DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION,
						Xaxis: &datadog.DistributionWidgetXAxis{
							Max:         datadog.PtrString("auto"),
							IncludeZero: datadog.PtrBool(true),
							Scale:       datadog.PtrString("linear"),
							Min:         datadog.PtrString("auto"),
						},
						Yaxis: &datadog.DistributionWidgetYAxis{
							Max:         datadog.PtrString("auto"),
							IncludeZero: datadog.PtrBool(true),
							Scale:       datadog.PtrString("linear"),
							Min:         datadog.PtrString("auto"),
						},
						Requests: []datadog.DistributionWidgetRequest{
							{
								Query: &datadog.DistributionWidgetHistogramRequestQuery{
									FormulaAndFunctionApmResourceStatsQueryDefinition: &datadog.FormulaAndFunctionApmResourceStatsQueryDefinition{
										PrimaryTagValue: datadog.PtrString("*"),
										Stat:            datadog.FORMULAANDFUNCTIONAPMRESOURCESTATNAME_LATENCY_DISTRIBUTION,
										DataSource:      datadog.FORMULAANDFUNCTIONAPMRESOURCESTATSDATASOURCE_APM_RESOURCE_STATS,
										Name:            "query1",
										Service:         "azure-bill-import",
										GroupBy: []string{
											"resource_name",
										},
										Env:            "staging",
										PrimaryTagName: datadog.PtrString("datacenter"),
										OperationName:  datadog.PtrString("universal.http.client"),
									}},
								RequestType: datadog.DISTRIBUTIONWIDGETHISTOGRAMREQUESTTYPE_HISTOGRAM.Ptr(),
								Style: &datadog.WidgetStyle{
									Palette: datadog.PtrString("dog_classic"),
								},
							},
						},
					}},
				Layout: &datadog.WidgetLayout{
					X:      8,
					Y:      0,
					Width:  4,
					Height: 2,
				},
			},
		},
		LayoutType: datadog.DASHBOARDLAYOUTTYPE_ORDERED,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsApi.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
