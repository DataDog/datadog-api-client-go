// Create a distribution widget using a histogram request containing a formulas and functions events query

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
		Title:       "Example-Create_a_distribution_widget_using_a_histogram_request_containing_a_formulas_and_functions_events_qu",
		Description: *common.NewNullableString(datadog.PtrString("Example-Create_a_distribution_widget_using_a_histogram_request_containing_a_formulas_and_functions_events_qu")),
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					DistributionWidgetDefinition: &datadog.DistributionWidgetDefinition{
						Title:      datadog.PtrString("Events Platform - Request latency HOP"),
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
									FormulaAndFunctionEventQueryDefinition: &datadog.FormulaAndFunctionEventQueryDefinition{
										Search: &datadog.FormulaAndFunctionEventQueryDefinitionSearch{
											Query: "",
										},
										DataSource: datadog.FORMULAANDFUNCTIONEVENTSDATASOURCE_EVENTS,
										Compute: datadog.FormulaAndFunctionEventQueryDefinitionCompute{
											Metric:      datadog.PtrString("@duration"),
											Aggregation: datadog.FORMULAANDFUNCTIONEVENTAGGREGATION_MIN,
										},
										Name: "query1",
										Indexes: []string{
											"*",
										},
										GroupBy: []datadog.FormulaAndFunctionEventQueryGroupBy{},
									}},
								RequestType: datadog.DISTRIBUTIONWIDGETHISTOGRAMREQUESTTYPE_HISTOGRAM.Ptr(),
							},
						},
					}},
				Layout: &datadog.WidgetLayout{
					X:      0,
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
