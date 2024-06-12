// Create a new dashboard with query_table widget

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
					Width:  54,
					Height: 32,
				},
				Definition: datadogV1.WidgetDefinition{
					TableWidgetDefinition: &datadogV1.TableWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Time:       &datadogV1.WidgetTime{},
						Type:       datadogV1.TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE,
						Requests: []datadogV1.TableWidgetRequest{
							{
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query1",
											Query:      "avg:system.cpu.user{*} by {host}",
											Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
										}},
								},
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula:            "query1",
										ConditionalFormats: []datadogV1.WidgetConditionalFormat{},
										CellDisplayMode:    datadogV1.TABLEWIDGETCELLDISPLAYMODE_BAR.Ptr(),
									},
								},
								Sort: &datadogV1.WidgetSortBy{
									Count: datadog.PtrInt64(500),
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
						HasSearchBar: datadogV1.TABLEWIDGETHASSEARCHBAR_AUTO.Ptr(),
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
