// Create a new dashboard with scatterplot widget

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
					ScatterPlotWidgetDefinition: &datadogV1.ScatterPlotWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Time: &datadogV1.WidgetTime{
							WidgetLegacyLiveSpan: &datadogV1.WidgetLegacyLiveSpan{}},
						Type: datadogV1.SCATTERPLOTWIDGETDEFINITIONTYPE_SCATTERPLOT,
						Requests: datadogV1.ScatterPlotWidgetDefinitionRequests{
							Table: &datadogV1.ScatterplotTableRequest{
								Formulas: []datadogV1.ScatterplotWidgetFormula{
									{
										Formula:   "query1",
										Dimension: datadogV1.SCATTERPLOTDIMENSION_X,
										Alias:     datadog.PtrString(""),
									},
									{
										Formula:   "query2",
										Dimension: datadogV1.SCATTERPLOTDIMENSION_Y,
										Alias:     datadog.PtrString(""),
									},
								},
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query1",
											Query:      "avg:system.cpu.user{*} by {service}",
											Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
										}},
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
											DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
											Name:       "query2",
											Query:      "avg:system.mem.used{*} by {service}",
											Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
										}},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
							},
						},
						Xaxis: &datadogV1.WidgetAxis{
							Scale:       datadog.PtrString("linear"),
							IncludeZero: datadog.PtrBool(true),
							Min:         datadog.PtrString("auto"),
							Max:         datadog.PtrString("auto"),
						},
						Yaxis: &datadogV1.WidgetAxis{
							Scale:       datadog.PtrString("linear"),
							IncludeZero: datadog.PtrBool(true),
							Min:         datadog.PtrString("auto"),
							Max:         datadog.PtrString("auto"),
						},
						ColorByGroups: []string{},
					}},
			},
			{
				Layout: &datadogV1.WidgetLayout{
					X:      48,
					Y:      0,
					Width:  47,
					Height: 15,
				},
				Definition: datadogV1.WidgetDefinition{
					ScatterPlotWidgetDefinition: &datadogV1.ScatterPlotWidgetDefinition{
						Title:      datadog.PtrString("Data Projection Scatterplot"),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.SCATTERPLOTWIDGETDEFINITIONTYPE_SCATTERPLOT,
						Requests: datadogV1.ScatterPlotWidgetDefinitionRequests{
							Table: &datadogV1.ScatterplotTableRequest{
								RequestType: datadogV1.SCATTERPLOTTABLEREQUESTTYPE_DATA_PROJECTION.Ptr(),
								Query: &datadogV1.ScatterplotDataProjectionQuery{
									QueryString: "@service:web-store",
									DataSource:  datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_SPANS,
								},
								Projection: &datadogV1.ScatterplotDataProjectionProjection{
									Type: datadogV1.SCATTERPLOTDATAPROJECTIONPROJECTIONTYPE_SCATTERPLOT,
									Dimensions: []datadogV1.ScatterplotDataProjectionDimension{
										{
											Column:    "duration",
											Dimension: datadogV1.SCATTERPLOTDIMENSION_X,
										},
										{
											Column:    "@resource_name",
											Dimension: datadogV1.SCATTERPLOTDIMENSION_Y,
										},
									},
								},
								Limit: datadog.PtrInt64(200),
							},
						},
						Xaxis: &datadogV1.WidgetAxis{
							Scale:       datadog.PtrString("linear"),
							IncludeZero: datadog.PtrBool(true),
							Min:         datadog.PtrString("auto"),
							Max:         datadog.PtrString("auto"),
						},
						Yaxis: &datadogV1.WidgetAxis{
							Scale:       datadog.PtrString("linear"),
							IncludeZero: datadog.PtrBool(true),
							Min:         datadog.PtrString("auto"),
							Max:         datadog.PtrString("auto"),
						},
					}},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_FREE,
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
