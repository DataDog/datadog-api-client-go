// Create a new dashboard with geomap widget

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
		Title:       "Example-Create_a_new_dashboard_with_geomap_widget",
		Description: *common.NewNullableString(nil),
		Widgets: []datadog.Widget{
			{
				Layout: &datadog.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  47,
					Height: 30,
				},
				Definition: datadog.WidgetDefinition{
					GeomapWidgetDefinition: &datadog.GeomapWidgetDefinition{
						Title:      common.PtrString(""),
						TitleSize:  common.PtrString("16"),
						TitleAlign: datadog.WIDGETTEXTALIGN_LEFT.Ptr(),
						Time:       &datadog.WidgetTime{},
						Type:       datadog.GEOMAPWIDGETDEFINITIONTYPE_GEOMAP,
						Requests: []datadog.GeomapWidgetRequest{
							{
								Formulas: []datadog.WidgetFormula{
									{
										Formula: "query1",
										Limit: &datadog.WidgetFormulaLimit{
											Count: common.PtrInt64(250),
											Order: datadog.QUERYSORTORDER_DESC.Ptr(),
										},
									},
								},
								Queries: []datadog.FormulaAndFunctionQueryDefinition{
									datadog.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionEventQueryDefinition: &datadog.FormulaAndFunctionEventQueryDefinition{
											Name:       "query1",
											DataSource: datadog.FORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
											Search: &datadog.FormulaAndFunctionEventQueryDefinitionSearch{
												Query: "",
											},
											Indexes: []string{
												"*",
											},
											Compute: datadog.FormulaAndFunctionEventQueryDefinitionCompute{
												Aggregation: datadog.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
											},
											GroupBy: []datadog.FormulaAndFunctionEventQueryGroupBy{
												{
													Facet: "@geo.country_iso_code",
													Limit: common.PtrInt64(250),
													Sort: &datadog.FormulaAndFunctionEventQueryGroupBySort{
														Order:       datadog.QUERYSORTORDER_DESC.Ptr(),
														Aggregation: datadog.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
													},
												},
											},
										}},
								},
								ResponseFormat: datadog.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
							},
						},
						Style: datadog.GeomapWidgetDefinitionStyle{
							Palette:     "hostmap_blues",
							PaletteFlip: false,
						},
						View: datadog.GeomapWidgetDefinitionView{
							Focus: "WORLD",
						},
					}},
			},
		},
		TemplateVariables: []datadog.DashboardTemplateVariable{},
		LayoutType:        datadog.DASHBOARDLAYOUTTYPE_FREE,
		IsReadOnly:        common.PtrBool(false),
		NotifyList:        []string{},
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
