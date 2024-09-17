// Create a new dashboard with a formulas and functions treemap widget

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
					TreeMapWidgetDefinition: &datadogV1.TreeMapWidgetDefinition{
						Title: datadog.PtrString(""),
						Type:  datadogV1.TREEMAPWIDGETDEFINITIONTYPE_TREEMAP,
						Requests: []datadogV1.TreeMapWidgetRequest{
							{
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula: "hour_before(query1)",
									},
									{
										Formula: "query1",
									},
								},
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
											Name:       "query1",
											Search: &datadogV1.FormulaAndFunctionEventQueryDefinitionSearch{
												Query: "",
											},
											Indexes: []string{
												"*",
											},
											Compute: datadogV1.FormulaAndFunctionEventQueryDefinitionCompute{
												Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
											},
											GroupBy: []datadogV1.FormulaAndFunctionEventQueryGroupBy{},
										}},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
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
