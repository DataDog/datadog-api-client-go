// Create a new dashboard with topology_map custom widget

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
					TopologyMapWidgetDefinition: &datadogV1.TopologyMapWidgetDefinition{
						TopologyMapWidgetDefinitionCustom: &datadogV1.TopologyMapWidgetDefinitionCustom{
							Title:      datadog.PtrString(""),
							TitleSize:  datadog.PtrString("16"),
							TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
							Type:       datadogV1.TOPOLOGYMAPWIDGETDEFINITIONTYPE_TOPOLOGY_MAP,
							Layout: datadogV1.TopologyMapLayout{
								Type:      datadogV1.TOPOLOGYMAPLAYOUTTYPE_HIERARCHICAL,
								Direction: datadogV1.TOPOLOGYMAPLAYOUTDIRECTION_LEFT_TO_RIGHT,
							},
							Data: &datadogV1.TopologyMapInlineData{
								Nodes: &datadogV1.TopologyMapInlineNodesData{
									Columns: []datadogV1.TopologyMapInlineNodeColumn{
										datadogV1.TopologyMapInlineNodeColumn{
											TopologyMapInlineNodeIdColumn: &datadogV1.TopologyMapInlineNodeIdColumn{
												Kind: datadogV1.TOPOLOGYMAPINLINECOLUMNKINDSTRING_STRING,
												Id:   datadogV1.TOPOLOGYMAPINLINENODEIDCOLUMNID_ID,
												Values: []string{
													"service:web",
													"service:db",
												},
											}},
										datadogV1.TopologyMapInlineNodeColumn{
											TopologyMapInlineNodeStringColumn: &datadogV1.TopologyMapInlineNodeStringColumn{
												Kind: datadogV1.TOPOLOGYMAPINLINECOLUMNKINDSTRING_STRING,
												Id:   datadogV1.TOPOLOGYMAPINLINENODESTRINGCOLUMNID_SUBTITLE,
												Values: []*string{
													datadog.PtrString("api-team"),
													nil,
												},
											}},
										datadogV1.TopologyMapInlineNodeColumn{
											TopologyMapInlineNodeValueColumn: &datadogV1.TopologyMapInlineNodeValueColumn{
												Kind: datadogV1.TOPOLOGYMAPINLINECOLUMNKINDSCALAR_SCALAR,
												Id:   "values[1]",
												Values: []*float64{
													datadog.PtrFloat64(1.5),
													nil,
												},
												Meta: &datadogV1.TopologyMapInlineNodeValueColumnMeta{
													Unit: &datadogV1.NumberFormatUnit{
														NumberFormatUnitCanonical: &datadogV1.NumberFormatUnitCanonical{
															Type:     datadogV1.NUMBERFORMATUNITSCALETYPE_CANONICAL_UNIT.Ptr(),
															UnitName: datadog.PtrString("second"),
														}},
												},
											}},
									},
								},
								Edges: &datadogV1.TopologyMapInlineEdgesData{
									Columns: []datadogV1.TopologyMapInlineEdgeColumn{
										{
											Kind: datadogV1.TOPOLOGYMAPINLINECOLUMNKINDSTRING_STRING,
											Id:   datadogV1.TOPOLOGYMAPINLINEEDGECOLUMNID_SOURCE,
											Values: []string{
												"service:web",
											},
										},
										{
											Kind: datadogV1.TOPOLOGYMAPINLINECOLUMNKINDSTRING_STRING,
											Id:   datadogV1.TOPOLOGYMAPINLINEEDGECOLUMNID_TARGET,
											Values: []string{
												"service:db",
											},
										},
									},
								},
							},
							Requests: []datadogV1.TopologyRequestCustom{
								{
									ResponseFormat: datadogV1.TOPOLOGYREQUESTCUSTOMRESPONSEFORMAT_SCALAR,
									Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
										datadogV1.FormulaAndFunctionQueryDefinition{
											FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
												DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
												Name:       "a",
												Query:      "avg:system.cpu.user{*} by {service}",
												Aggregator: datadogV1.FORMULAANDFUNCTIONMETRICAGGREGATION_AVG.Ptr(),
											}},
									},
									Formulas: []datadogV1.WidgetFormula{
										{
											Formula: "a",
										},
									},
									Projection: datadogV1.TopologyRequestCustomProjection{
										Type: datadogV1.TOPOLOGYREQUESTCUSTOMPROJECTIONTYPE_CUSTOM_TOPOLOGY_MAP_NODE,
										Dimensions: []datadogV1.TopologyRequestCustomProjectionDimension{
											{
												Column:    "a",
												Dimension: "values[0]",
											},
										},
									},
								},
							},
						}}},
			},
		},
		TemplateVariables: []datadogV1.DashboardTemplateVariable{},
		LayoutType:        datadogV1.DASHBOARDLAYOUTTYPE_FREE,
		NotifyList:        *datadog.NewNullableList(&[]string{}),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	ctx = context.WithValue(ctx, datadog.ContextAccessToken, os.Getenv("DD_BEARER_TOKEN"))
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
