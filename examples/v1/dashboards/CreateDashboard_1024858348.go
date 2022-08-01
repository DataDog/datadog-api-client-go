// Create a new dashboard with a formulas and functions treemap widget

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
		Title: "Example-Create_a_new_dashboard_with_a_formulas_and_functions_treemap_widget",
		Widgets: []datadog.Widget{
			{
				Definition: datadog.WidgetDefinition{
					TreeMapWidgetDefinition: &datadog.TreeMapWidgetDefinition{
						Title: common.PtrString(""),
						Type:  datadog.TREEMAPWIDGETDEFINITIONTYPE_TREEMAP,
						Requests: []datadog.TreeMapWidgetRequest{
							{
								Formulas: []datadog.WidgetFormula{
									{
										Formula: "hour_before(query1)",
									},
									{
										Formula: "query1",
									},
								},
								Queries: []datadog.FormulaAndFunctionQueryDefinition{
									datadog.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionEventQueryDefinition: &datadog.FormulaAndFunctionEventQueryDefinition{
											DataSource: datadog.FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
											Name:       "query1",
											Search: &datadog.FormulaAndFunctionEventQueryDefinitionSearch{
												Query: "",
											},
											Indexes: []string{
												"*",
											},
											Compute: datadog.FormulaAndFunctionEventQueryDefinitionCompute{
												Aggregation: datadog.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
											},
											GroupBy: []datadog.FormulaAndFunctionEventQueryGroupBy{},
										}},
								},
								ResponseFormat: datadog.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
							},
						},
					}},
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  4,
					Height: 4,
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
