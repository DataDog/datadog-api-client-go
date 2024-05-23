// Create a new dashboard with geomap widget

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
		Description: *datadog.NewNullableString(nil),
		Widgets: []datadogV1.Widget{
			{
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 30,
				},
				Definition: datadogV1.WidgetDefinition{
					GeomapWidgetDefinition: &datadogV1.GeomapWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Time:       &datadogV1.WidgetTime{},
						Type:       datadogV1.GEOMAPWIDGETDEFINITIONTYPE_GEOMAP,
						Requests: []datadogV1.GeomapWidgetRequest{
							{
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula: "query1",
									},
								},
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
											Name:       "query1",
											DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
											Search: &datadogV1.FormulaAndFunctionEventQueryDefinitionSearch{
												Query: "",
											},
											Indexes: []string{
												"*",
											},
											Compute: datadogV1.FormulaAndFunctionEventQueryDefinitionCompute{
												Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
											},
											GroupBy: []datadogV1.FormulaAndFunctionEventQueryGroupBy{
												{
													Facet: "@geo.country_iso_code",
													Limit: datadog.PtrInt64(250),
													Sort: &datadogV1.FormulaAndFunctionEventQueryGroupBySort{
														Order:       datadogV1.QUERYSORTORDER_DESC.Ptr(),
														Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
													},
												},
											},
										}},
								},
								Sort: &datadogV1.WidgetSortBy{
									Count: datadog.PtrInt64(250),
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
						Style: datadogV1.GeomapWidgetDefinitionStyle{
							Palette:     "hostmap_blues",
							PaletteFlip: false,
						},
						View: datadogV1.GeomapWidgetDefinitionView{
							Focus: "WORLD",
						},
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
