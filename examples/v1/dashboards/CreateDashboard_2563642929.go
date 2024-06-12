// Create a new dashboard with a toplist widget sorted by group

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
					Width:  47,
					Height: 15,
				},
				Definition: datadogV1.WidgetDefinition{
					ToplistWidgetDefinition: &datadogV1.ToplistWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Time:       &datadogV1.WidgetTime{},
						Style: &datadogV1.ToplistWidgetStyle{
							Display: &datadogV1.ToplistWidgetDisplay{
								ToplistWidgetStacked: &datadogV1.ToplistWidgetStacked{
									Type:   datadogV1.TOPLISTWIDGETSTACKEDTYPE_STACKED,
									Legend: datadogV1.TOPLISTWIDGETLEGEND_INLINE,
								}},
							Scaling: datadogV1.TOPLISTWIDGETSCALING_RELATIVE.Ptr(),
						},
						Type: datadogV1.TOPLISTWIDGETDEFINITIONTYPE_TOPLIST,
						Requests: []datadogV1.ToplistWidgetRequest{
							{
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query1",
											Query:      "avg:system.cpu.user{*} by {service}",
											Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
										}},
								},
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula: "query1",
									},
								},
								Sort: &datadogV1.WidgetSortBy{
									Count: datadog.PtrInt64(10),
									OrderBy: []datadogV1.WidgetSortOrderBy{
										datadogV1.WidgetSortOrderBy{
											WidgetGroupSort: &datadogV1.WidgetGroupSort{
												Type:  datadogV1.GROUPTYPE_GROUP,
												Name:  "service",
												Order: datadogV1.WIDGETSORT_ASCENDING,
											}},
									},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
							},
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
