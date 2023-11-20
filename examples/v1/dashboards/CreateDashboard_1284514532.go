// Create a new dashboard with a timeseries widget using formulas and functions cloud cost query

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
					TimeseriesWidgetDefinition: &datadogV1.TimeseriesWidgetDefinition{
						Title:      datadog.PtrString("Example Cloud Cost Query"),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES,
						Requests: []datadogV1.TimeseriesWidgetRequest{
							{
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula: "query1",
									},
								},
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionCloudCostQueryDefinition: &datadogV1.FormulaAndFunctionCloudCostQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONCLOUDCOSTDATASOURCE_CLOUD_COST,
											Name:       "query1",
											Query:      "sum:aws.cost.amortized{*} by {aws_product}.rollup(sum, monthly)",
										}},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
								Style: &datadogV1.WidgetRequestStyle{
									Palette:   datadog.PtrString("dog_classic"),
									LineType:  datadogV1.WIDGETLINETYPE_SOLID.Ptr(),
									LineWidth: datadogV1.WIDGETLINEWIDTH_NORMAL.Ptr(),
								},
								DisplayType: datadogV1.WIDGETDISPLAYTYPE_BARS.Ptr(),
							},
						},
						Time: &datadogV1.WidgetTime{
							LiveSpan: datadogV1.WIDGETLIVESPAN_WEEK_TO_DATE.Ptr(),
						},
					}},
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
