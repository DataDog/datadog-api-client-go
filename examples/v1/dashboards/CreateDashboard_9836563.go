// Create a geomap widget with conditional formats and text formats

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
		Description: *datadog.NewNullableString(datadog.PtrString("Example-Dashboard")),
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					GeomapWidgetDefinition: &datadogV1.GeomapWidgetDefinition{
						Title: datadog.PtrString("Log Count by Service and Source"),
						Type:  datadogV1.GEOMAPWIDGETDEFINITIONTYPE_GEOMAP,
						Requests: []datadogV1.GeomapWidgetRequest{
							{
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
											Name:       "query1",
											Search: &datadogV1.FormulaAndFunctionEventQueryDefinitionSearch{
												Query: "@type:session",
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
								ConditionalFormats: []datadogV1.WidgetConditionalFormat{
									{
										Comparator: datadogV1.WIDGETCOMPARATOR_GREATER_THAN,
										Value:      1000,
										Palette:    datadogV1.WIDGETPALETTE_WHITE_ON_GREEN,
									},
								},
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula: "query1",
									},
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
							},
							{
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_EVENT_LIST.Ptr(),
								Query: &datadogV1.ListStreamQuery{
									DataSource:  datadogV1.LISTSTREAMSOURCE_LOGS_STREAM,
									QueryString: "",
									Indexes:     []string{},
									Storage:     datadog.PtrString("hot"),
								},
								Columns: []datadogV1.ListStreamColumn{
									{
										Field: "@network.client.geoip.location.latitude",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
									{
										Field: "@network.client.geoip.location.longitude",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
									{
										Field: "@network.client.geoip.country.iso_code",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
									{
										Field: "@network.client.geoip.subdivision.name",
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
									},
								},
								Style: &datadogV1.GeomapWidgetRequestStyle{
									ColorBy: datadog.PtrString("status"),
								},
								TextFormats: []datadogV1.TableWidgetTextFormatRule{
									{
										Match: datadogV1.TableWidgetTextFormatMatch{
											Type:  datadogV1.TABLEWIDGETTEXTFORMATMATCHTYPE_IS,
											Value: "error",
										},
										Palette: datadogV1.TABLEWIDGETTEXTFORMATPALETTE_WHITE_ON_RED.Ptr(),
									},
								},
							},
						},
						Style: datadogV1.GeomapWidgetDefinitionStyle{
							Palette:     "hostmap_blues",
							PaletteFlip: false,
						},
						View: datadogV1.GeomapWidgetDefinitionView{
							Focus: "NORTH_AMERICA",
						},
					}},
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  12,
					Height: 6,
				},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		NotifyList:        *datadog.NewNullableList(&[]string{}),
		ReflowType:        datadogV1.DASHBOARDREFLOWTYPE_FIXED.Ptr(),
		Tags:              *datadog.NewNullableList(&[]string{}),
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
