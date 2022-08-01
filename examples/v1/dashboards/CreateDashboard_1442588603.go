// Create a distribution widget using a histogram request containing a formulas and functions APM Stats query

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
		Title:       "Example-Create_a_distribution_widget_using_a_histogram_request_containing_a_formulas_and_functions_APM_Stats",
		Description: *common.NewNullableString(common.PtrString("")),
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					DistributionWidgetDefinition: &datadog.DistributionWidgetDefinition{
						Title:      common.PtrString("APM Stats - Request latency HOP"),
						TitleSize:  common.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						ShowLegend: common.PtrBool(false),
						Type:       datadog.DISTRIBUTIONWIDGETDEFINITIONTYPE_DISTRIBUTION,
						Xaxis: &datadog.DistributionWidgetXAxis{
							Max:         common.PtrString("auto"),
							IncludeZero: common.PtrBool(true),
							Scale:       common.PtrString("linear"),
							Min:         common.PtrString("auto"),
						},
						Yaxis: &datadog.DistributionWidgetYAxis{
							Max:         common.PtrString("auto"),
							IncludeZero: common.PtrBool(true),
							Scale:       common.PtrString("linear"),
							Min:         common.PtrString("auto"),
						},
						Requests: []datadog.DistributionWidgetRequest{
							{
								Query: &datadog.DistributionWidgetHistogramRequestQuery{
									FormulaAndFunctionApmResourceStatsQueryDefinition: &datadog.FormulaAndFunctionApmResourceStatsQueryDefinition{
										PrimaryTagValue: common.PtrString("*"),
										Stat:            datadog.FORMULAANDFUNCTIONAPMRESOURCESTATNAME_LATENCY_DISTRIBUTION,
										DataSource:      datadog.FORMULAANDFUNCTIONAPMRESOURCESTATSDATASOURCE_APM_RESOURCE_STATS,
										Name:            "query1",
										Service:         "azure-bill-import",
										GroupBy: []string{
											"resource_name",
										},
										Env:            "staging",
										PrimaryTagName: common.PtrString("datacenter"),
										OperationName:  common.PtrString("universal.http.client"),
									}},
								RequestType: datadog.DISTRIBUTIONWIDGETHISTOGRAMREQUESTTYPE_HISTOGRAM.Ptr(),
								Style: &datadog.WidgetStyle{
									Palette: common.PtrString("dog_classic"),
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
