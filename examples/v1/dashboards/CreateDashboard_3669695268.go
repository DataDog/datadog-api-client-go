// Create a new dashboard with logs query table widget and storage parameter

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
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:      "Example-Dashboard with query table widget and storage parameter",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					TableWidgetDefinition: &datadogV1.TableWidgetDefinition{
						Type: datadogV1.TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE,
						Requests: []datadogV1.TableWidgetRequest{
							{
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
											Storage: datadog.PtrString("online_archives"),
										}},
								},
								Formulas: []datadogV1.WidgetFormula{
									{
										ConditionalFormats: []datadogV1.WidgetConditionalFormat{},
										CellDisplayMode:    datadogV1.TABLEWIDGETCELLDISPLAYMODE_BAR.Ptr(),
										Formula:            "query1",
									},
								},
								Sort: &datadogV1.WidgetSortBy{
									Count: datadog.PtrInt64(50),
									OrderBy: []datadogV1.WidgetSortOrderBy{
										datadogV1.WidgetSortOrderBy{
											WidgetFormulaSort: &datadogV1.WidgetFormulaSort{
												Type:  datadogV1.FORMULATYPE_FORMULA,
												Index: 0,
												Order: datadogV1.WIDGETSORT_DESCENDING,
											}},
									},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
							},
						},
					}},
			},
		},
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
