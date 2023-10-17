/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestDashboardLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	sloApi := datadogV1.NewServiceLevelObjectivesApi(Client(ctx))
	api := datadogV1.NewDashboardsApi(Client(ctx))

	// create SLO for referencing in SLO widget (we're borrowing these from api_slo_test.go)
	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err := sloApi.CreateSLO(ctx, testEventSLO)
	if err != nil {
		t.Fatalf("Error creating SLO %v for testing Dashboard SLO widget: Response %s: %v", testEventSLO, err.Error(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, t, slo.GetId())
	assert.Equal(httpresp.StatusCode, 200)

	widgetTime := datadogV1.NewWidgetTimeWithDefaults()
	widgetTime.SetLiveSpan(datadogV1.WIDGETLIVESPAN_PAST_FIFTEEN_MINUTES)

	widgetTimePastOneHour := datadogV1.NewWidgetTimeWithDefaults()
	widgetTimePastOneHour.SetLiveSpan(datadogV1.WIDGETLIVESPAN_PAST_ONE_HOUR)

	widgetTimePastOneDay := datadogV1.NewWidgetTimeWithDefaults()
	widgetTimePastOneDay.SetLiveSpan(datadogV1.WIDGETLIVESPAN_PAST_ONE_DAY)

	widgetTimePastOneMonth := datadogV1.NewWidgetTimeWithDefaults()
	widgetTimePastOneMonth.SetLiveSpan(datadogV1.WIDGETLIVESPAN_PAST_ONE_MONTH)

	widgetLayout := datadogV1.NewWidgetLayoutWithDefaults()
	widgetLayout.SetHeight(10)
	widgetLayout.SetWidth(10)
	widgetLayout.SetX(0)
	widgetLayout.SetY(0)

	// Custom Links
	customLink := datadogV1.WidgetCustomLink{}
	customLink.SetLabel("Test Custom Link label")
	customLink.SetLink("https://app.datadoghq.com/dashboard/lists")

	// Alert Graph Widget
	alertGraphDefinition := datadogV1.NewAlertGraphWidgetDefinitionWithDefaults()
	alertGraphDefinition.SetAlertId("1234")
	alertGraphDefinition.SetVizType(datadogV1.WIDGETVIZTYPE_TIMESERIES)
	alertGraphDefinition.SetTitle("Test Alert Graph Widget")
	alertGraphDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	alertGraphDefinition.SetTitleSize("16")
	alertGraphDefinition.SetTime(*widgetTime)

	alertGraphWidget := datadogV1.NewWidgetWithDefaults()
	alertGraphWidget.SetDefinition(datadogV1.AlertGraphWidgetDefinitionAsWidgetDefinition(alertGraphDefinition))

	// Alert Value Widget
	alertValueDefinition := datadogV1.NewAlertValueWidgetDefinitionWithDefaults()
	alertValueDefinition.SetAlertId("1234")
	alertValueDefinition.SetPrecision(2)
	alertValueDefinition.SetUnit("ms")
	alertValueDefinition.SetTitleSize("12")
	alertValueDefinition.SetTextAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	alertValueDefinition.SetTitle("Test Alert Value Widget")
	alertValueDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_RIGHT)

	alertValueWidget := datadogV1.NewWidgetWithDefaults()
	alertValueWidget.SetDefinition(datadogV1.AlertValueWidgetDefinitionAsWidgetDefinition(alertValueDefinition))

	// Change Widget
	changeWidgetRequest := datadogV1.NewChangeWidgetRequest()
	changeWidgetRequest.SetQ("avg:system.load.1{*}")
	changeWidgetRequest.SetChangeType(datadogV1.WIDGETCHANGETYPE_ABSOLUTE)
	changeWidgetRequest.SetCompareTo(datadogV1.WIDGETCOMPARETO_HOUR_BEFORE)
	changeWidgetRequest.SetIncreaseGood(true)
	changeWidgetRequest.SetOrderBy(datadogV1.WIDGETORDERBY_CHANGE)
	changeWidgetRequest.SetOrderDir(datadogV1.WIDGETSORT_ASCENDING)
	changeWidgetRequest.SetShowPresent(true)

	changeWidgetDefinition := datadogV1.NewChangeWidgetDefinitionWithDefaults()
	changeWidgetDefinition.SetTitle("Test Change Widget")
	changeWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	changeWidgetDefinition.SetTitleSize("16")
	changeWidgetDefinition.SetTime(*widgetTime)
	changeWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})
	changeWidgetDefinition.SetRequests([]datadogV1.ChangeWidgetRequest{
		*changeWidgetRequest,
	})

	changeWidget := datadogV1.NewWidgetWithDefaults()
	changeWidget.SetDefinition(datadogV1.ChangeWidgetDefinitionAsWidgetDefinition(changeWidgetDefinition))

	// Check Status Widget
	checkStatusWidgetDefinition := datadogV1.NewCheckStatusWidgetDefinitionWithDefaults()
	checkStatusWidgetDefinition.SetCheck("service_check.up")
	checkStatusWidgetDefinition.SetGrouping(datadogV1.WIDGETGROUPING_CHECK)
	checkStatusWidgetDefinition.SetGroup("*")
	checkStatusWidgetDefinition.SetTags([]string{"foo:bar"})
	checkStatusWidgetDefinition.SetGroupBy([]string{"bar"})
	checkStatusWidgetDefinition.SetTitle("Test Check Status Widget")
	checkStatusWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	checkStatusWidgetDefinition.SetTitleSize("16")
	checkStatusWidgetDefinition.SetTime(*widgetTime)

	checkStatusWidget := datadogV1.NewWidgetWithDefaults()
	checkStatusWidget.SetDefinition(datadogV1.CheckStatusWidgetDefinitionAsWidgetDefinition(checkStatusWidgetDefinition))

	// Distribution Widget
	distributionWidgetDefinition := datadogV1.NewDistributionWidgetDefinitionWithDefaults()

	distributionWidgetRequest := datadogV1.NewDistributionWidgetRequestWithDefaults()
	distributionWidgetRequest.SetQ("avg:system.load.1{*}")

	dogClassic := datadogV1.NewWidgetStyleWithDefaults()
	dogClassic.SetPalette("dog_classic")

	distributionWidgetRequest.SetStyle(*dogClassic)

	distributionWidgetDefinition.SetRequests([]datadogV1.DistributionWidgetRequest{
		*distributionWidgetRequest,
	})
	distributionWidgetDefinition.SetTitle("Test Distribution Widget")
	distributionWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	distributionWidgetDefinition.SetTitleSize("16")
	distributionWidgetDefinition.SetTime(*widgetTime)

	distributionWidget := datadogV1.NewWidget(datadogV1.DistributionWidgetDefinitionAsWidgetDefinition(distributionWidgetDefinition))

	// Event Stream Widget ONLY AVAILABLE ON FREE LAYOUTS
	eventStreamWidgetDefinition := datadogV1.NewEventStreamWidgetDefinitionWithDefaults()
	eventStreamWidgetDefinition.SetQuery("Build successful")
	eventStreamWidgetDefinition.SetEventSize(datadogV1.WIDGETEVENTSIZE_LARGE)
	eventStreamWidgetDefinition.SetTitle("Test Event Stream Widget")
	eventStreamWidgetDefinition.SetTitleSize("16")
	eventStreamWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	eventStreamWidgetDefinition.SetTime(*widgetTimePastOneDay)

	eventStreamWidget := datadogV1.NewWidgetWithDefaults()
	eventStreamWidget.SetDefinition(datadogV1.EventStreamWidgetDefinitionAsWidgetDefinition(eventStreamWidgetDefinition))
	eventStreamWidget.SetLayout(*widgetLayout)

	// Event Timeline Widget ONLY AVAILABLE ON FREE LAYOUTS
	eventTimelineWidgetDefinition := datadogV1.NewEventTimelineWidgetDefinitionWithDefaults()
	eventTimelineWidgetDefinition.SetQuery("Build Failed")
	eventTimelineWidgetDefinition.SetTitle("Test Event Timeline Widget")
	eventTimelineWidgetDefinition.SetTitleSize("16")
	eventTimelineWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_LEFT)
	eventTimelineWidgetDefinition.SetTime(*widgetTimePastOneMonth)

	eventTimelineWidget := datadogV1.NewWidgetWithDefaults()
	eventTimelineWidget.SetDefinition(datadogV1.EventTimelineWidgetDefinitionAsWidgetDefinition(eventTimelineWidgetDefinition))
	eventTimelineWidget.SetLayout(*widgetLayout)

	// Free Text Widget ONLY AVAILABLE ON FREE LAYOUTS
	freeTextWidgetDefinition := datadogV1.NewFreeTextWidgetDefinitionWithDefaults()
	freeTextWidgetDefinition.SetText("Test me text")
	freeTextWidgetDefinition.SetColor("blue")
	freeTextWidgetDefinition.SetFontSize("16")
	freeTextWidgetDefinition.SetTextAlign(datadogV1.WIDGETTEXTALIGN_CENTER)

	freeTextWidget := datadogV1.NewWidgetWithDefaults()
	freeTextWidget.SetDefinition(datadogV1.FreeTextWidgetDefinitionAsWidgetDefinition(freeTextWidgetDefinition))
	freeTextWidget.SetLayout(*widgetLayout)

	// Geomap with Formulas and Functions Query
	geoMapWidgetDefinitionFormulaFunctionsQuery := datadogV1.NewGeomapWidgetDefinitionWithDefaults()
	geoMapWidgetDefinitionFormulaFunctionsQuery.SetRequests([]datadogV1.GeomapWidgetRequest{{
		Formulas: []datadogV1.WidgetFormula{{
			Formula: "query1",
		}},
		ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
		Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
			{
				FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
					DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
					Compute: datadogV1.FormulaAndFunctionEventQueryDefinitionCompute{
						Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
					},
					GroupBy: []datadogV1.FormulaAndFunctionEventQueryGroupBy{{
						Facet: "@geo.country_iso_code",
						Limit: datadog.PtrInt64(250),
						Sort: &datadogV1.FormulaAndFunctionEventQueryGroupBySort{
							Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
						}}},
					Indexes: []string{"*"},
					Name:    "query1",
				},
			},
		}}})
	geoMapWidgetDefinitionFormulaFunctionsQuery.SetTitle("Test Formulas and Functions Metric + Event query")
	geoMapWidgetDefinitionFormulaFunctionsQuery.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	geoMapWidgetDefinitionFormulaFunctionsQuery.SetTitleSize("16")
	geoMapWidgetDefinitionFormulaFunctionsQuery.SetTime(*widgetTime)
	geoMapWidgetDefinitionFormulaFunctionsQuery.SetStyle(datadogV1.GeomapWidgetDefinitionStyle{
		Palette:     *datadog.PtrString("dog_classic"),
		PaletteFlip: *datadog.PtrBool(true),
	})
	geoMapWidgetDefinitionFormulaFunctionsQuery.SetView(datadogV1.GeomapWidgetDefinitionView{
		Focus: *datadog.PtrString("WORLD"),
	})

	geoMapWidgetFormulaFunctionsQuery := datadogV1.NewWidget(datadogV1.GeomapWidgetDefinitionAsWidgetDefinition(geoMapWidgetDefinitionFormulaFunctionsQuery))

	// Geomap Widget
	geoMapWidgetDefinition := datadogV1.NewGeomapWidgetDefinitionWithDefaults()
	geoMapWidgetDefinition.SetRequests([]datadogV1.GeomapWidgetRequest{{
		LogQuery: &datadogV1.LogQueryDefinition{
			Index: datadog.PtrString("*"),
			Compute: &datadogV1.LogsQueryCompute{
				Aggregation: "count",
			},
			GroupBy: []datadogV1.LogQueryDefinitionGroupBy{{
				Facet: "@geo.country_iso_code",
				Limit: datadog.PtrInt64(250),
				Sort: &datadogV1.LogQueryDefinitionGroupBySort{
					Aggregation: "count",
					Order:       datadogV1.WIDGETSORT_DESCENDING,
				},
			}},
		},
	}})
	geoMapWidgetDefinition.SetStyle(datadogV1.GeomapWidgetDefinitionStyle{
		Palette:     *datadog.PtrString("dog_classic"),
		PaletteFlip: *datadog.PtrBool(true),
	})
	geoMapWidgetDefinition.SetView(datadogV1.GeomapWidgetDefinitionView{
		Focus: *datadog.PtrString("WORLD"),
	})
	geoMapWidget := datadogV1.NewWidget(datadogV1.GeomapWidgetDefinitionAsWidgetDefinition(geoMapWidgetDefinition))

	// Group Widget
	groupNoteWidgetDefinition := datadogV1.NewNoteWidgetDefinitionWithDefaults()
	groupNoteWidgetDefinition.SetContent("Test Note Inside Group")

	groupWidgetDefinition := datadogV1.NewGroupWidgetDefinitionWithDefaults()
	groupWidgetDefinition.SetLayoutType(datadogV1.WIDGETLAYOUTTYPE_ORDERED)
	groupWidgetDefinition.SetTitle("Test Group Widget")
	groupWidgetDefinition.SetWidgets([]datadogV1.GroupWidgetItem{
		*datadogV1.NewGroupWidgetItem(datadogV1.NoteWidgetDefinitionAsGroupWidgetItemDefinition(groupNoteWidgetDefinition)),
	})

	groupWidget := datadogV1.NewWidget(datadogV1.GroupWidgetDefinitionAsWidgetDefinition(groupWidgetDefinition))

	// HeatMap Widget
	heatMapWidgetDefinition := datadogV1.NewHeatMapWidgetDefinitionWithDefaults()

	heatMapWidgetRequest := datadogV1.NewHeatMapWidgetRequest()
	heatMapWidgetRequest.SetStyle(datadogV1.WidgetStyle{Palette: datadog.PtrString("dog_classic")})
	heatMapWidgetRequest.SetQ("avg:system.load.1{*}")

	widgetAxis := datadogV1.NewWidgetAxis()
	widgetAxis.SetIncludeZero(true)
	widgetAxis.SetMin("0")
	widgetAxis.SetMax("100")
	widgetAxis.SetScale("linear")

	heatMapWidgetDefinition.SetRequests([]datadogV1.HeatMapWidgetRequest{*heatMapWidgetRequest})
	heatMapWidgetDefinition.SetYaxis(*widgetAxis)
	heatMapWidgetDefinition.SetEvents([]datadogV1.WidgetEvent{
		{
			Q:             "Build succeeded",
			TagsExecution: datadog.PtrString("tags"),
		},
	})

	heatMapWidgetDefinition.SetTitle("Test Headmap Widget")
	heatMapWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	heatMapWidgetDefinition.SetTitleSize("16")
	heatMapWidgetDefinition.SetTime(*widgetTime)
	heatMapWidgetDefinition.SetShowLegend(true)
	heatMapWidgetDefinition.SetLegendSize("4")
	heatMapWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	heatMapWidget := datadogV1.NewWidget(datadogV1.HeatMapWidgetDefinitionAsWidgetDefinition(heatMapWidgetDefinition))

	// HostMap Widget
	hostMapWidgetDefinition := datadogV1.NewHostMapWidgetDefinitionWithDefaults()
	hostMapWidgetDefinition.SetNodeType(datadogV1.WIDGETNODETYPE_CONTAINER)
	hostMapWidgetDefinition.SetRequests(datadogV1.HostMapWidgetDefinitionRequests{
		Fill: &datadogV1.HostMapRequest{Q: datadog.PtrString("avg:system.load.1{*}")},
		Size: &datadogV1.HostMapRequest{Q: datadog.PtrString("avg:system.load.1{*}")},
	})
	hostMapWidgetDefinition.SetNoMetricHosts(true)
	hostMapWidgetDefinition.SetNoGroupHosts(true)
	hostMapWidgetDefinition.SetGroup([]string{"env:prod"})
	hostMapWidgetDefinition.SetScope([]string{"foo"})
	hostMapWidgetDefinition.SetStyle(datadogV1.HostMapWidgetDefinitionStyle{
		Palette:     datadog.PtrString("dog_classic"),
		PaletteFlip: datadog.PtrBool(true),
		FillMin:     datadog.PtrString("0"),
		FillMax:     datadog.PtrString("100"),
	})
	hostMapWidgetDefinition.SetTitle("Test HostMap Widget")
	hostMapWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	hostMapWidgetDefinition.SetTitleSize("16")
	hostMapWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	hostMapWidget := datadogV1.NewWidget(datadogV1.HostMapWidgetDefinitionAsWidgetDefinition(hostMapWidgetDefinition))

	// Iframe Widget ONLY AVAILABLE ON FREE LAYOUTS
	iFrameWidgetDefinition := datadogV1.NewIFrameWidgetDefinitionWithDefaults()
	iFrameWidgetDefinition.SetUrl("https://datadoghq.com")

	iFrameWidget := datadogV1.NewWidget(datadogV1.IFrameWidgetDefinitionAsWidgetDefinition(iFrameWidgetDefinition))
	iFrameWidget.SetLayout(*widgetLayout)

	// Image Widget ONLY AVAILABLE ON FREE LAYOUTS
	imageWidgetDefinition := datadogV1.NewImageWidgetDefinitionWithDefaults()
	imageWidgetDefinition.SetUrl("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4")
	imageWidgetDefinition.SetSizing(datadogV1.WIDGETIMAGESIZING_CENTER)
	imageWidgetDefinition.SetMargin(datadogV1.WIDGETMARGIN_LARGE)

	imageWidget := datadogV1.NewWidget(datadogV1.ImageWidgetDefinitionAsWidgetDefinition(imageWidgetDefinition))
	imageWidget.SetLayout(*widgetLayout)

	// LogStream ONLY AVAILABLE ON FREE LAYOUTS
	logStreamWidgetDefinition := datadogV1.NewLogStreamWidgetDefinitionWithDefaults()
	logStreamWidgetDefinition.SetIndexes([]string{"main"})
	logStreamWidgetDefinition.SetLogset("106")
	logStreamWidgetDefinition.SetQuery("Route XYZ failed")
	logStreamWidgetDefinition.SetColumns([]string{"Route"})
	logStreamWidgetDefinition.SetTitle("Test Logstream Widget")
	logStreamWidgetDefinition.SetTitleSize("16")
	logStreamWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_RIGHT)
	logStreamWidgetDefinition.SetTime(datadogV1.WidgetTime{LiveSpan: datadogV1.WIDGETLIVESPAN_PAST_TWO_DAYS.Ptr()})
	logStreamWidgetDefinition.SetMessageDisplay(datadogV1.WIDGETMESSAGEDISPLAY_EXPANDED_LARGE)
	logStreamWidgetDefinition.SetShowDateColumn(true)
	logStreamWidgetDefinition.SetShowMessageColumn(true)
	logStreamWidgetDefinition.SetSort(datadogV1.WidgetFieldSort{Column: "Route", Order: datadogV1.WIDGETSORT_ASCENDING})

	logStreamWidget := datadogV1.NewWidget(datadogV1.LogStreamWidgetDefinitionAsWidgetDefinition(logStreamWidgetDefinition))
	logStreamWidget.SetLayout(*widgetLayout)

	// Monitor Summary ONLY AVAILABLE ON FREE LAYOUTS
	monitorSummaryWidgetDefinition := datadogV1.NewMonitorSummaryWidgetDefinitionWithDefaults()
	monitorSummaryWidgetDefinition.SetQuery("Errors are increasing")
	monitorSummaryWidgetDefinition.SetSummaryType(datadogV1.WIDGETSUMMARYTYPE_COMBINED)
	monitorSummaryWidgetDefinition.SetSort(datadogV1.WIDGETMONITORSUMMARYSORT_NAME_ASCENDING)
	monitorSummaryWidgetDefinition.SetDisplayFormat(datadogV1.WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS)
	monitorSummaryWidgetDefinition.SetColorPreference(datadogV1.WIDGETCOLORPREFERENCE_BACKGROUND)
	monitorSummaryWidgetDefinition.SetHideZeroCounts(false)
	monitorSummaryWidgetDefinition.SetShowLastTriggered(true)
	monitorSummaryWidgetDefinition.SetTitle("Test Monitor Summary Widget")
	monitorSummaryWidgetDefinition.SetTitleSize("16")
	monitorSummaryWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	monitorSummaryWidgetDefinition.SetStart(0)
	monitorSummaryWidgetDefinition.SetCount(5)
	monitorSummaryWidgetDefinition.SetShowPriority(false)

	monitorSummaryWidget := datadogV1.NewWidget(datadogV1.MonitorSummaryWidgetDefinitionAsWidgetDefinition(monitorSummaryWidgetDefinition))
	monitorSummaryWidget.SetLayout(*widgetLayout)

	// Note Widget
	noteWidgetDefinition := datadogV1.NewNoteWidgetDefinitionWithDefaults()
	noteWidgetDefinition.SetFontSize("13")
	noteWidgetDefinition.SetContent("Test Note Widget Example")
	noteWidgetDefinition.SetBackgroundColor("blue")
	noteWidgetDefinition.SetTextAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	noteWidgetDefinition.SetShowTick(true)
	noteWidgetDefinition.SetTickPos("4")
	noteWidgetDefinition.SetTickEdge(datadogV1.WIDGETTICKEDGE_BOTTOM)

	noteWidget := datadogV1.NewWidget(datadogV1.NoteWidgetDefinitionAsWidgetDefinition(noteWidgetDefinition))

	// Query Value Widget
	queryValueWidgetDefinition := datadogV1.NewQueryValueWidgetDefinitionWithDefaults()
	queryValueWidgetDefinition.SetRequests([]datadogV1.QueryValueWidgetRequest{{
		Q:          datadog.PtrString("avg:system.load.1{*}"),
		Aggregator: datadogV1.WIDGETAGGREGATOR_AVERAGE.Ptr(),
		ConditionalFormats: []datadogV1.WidgetConditionalFormat{{
			Comparator:    datadogV1.WIDGETCOMPARATOR_GREATER_THAN,
			Value:         7.,
			Palette:       datadogV1.WIDGETPALETTE_RED_ON_WHITE,
			CustomBgColor: datadog.PtrString("blue"),
			CustomFgColor: datadog.PtrString("black"),
			ImageUrl:      datadog.PtrString("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4"),
		}},
	}})
	queryValueWidgetDefinition.SetAutoscale(true)
	queryValueWidgetDefinition.SetCustomUnit("ns")
	queryValueWidgetDefinition.SetPrecision(2)
	queryValueWidgetDefinition.SetTextAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	queryValueWidgetDefinition.SetTitle("Test Query Value Widget")
	queryValueWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	queryValueWidgetDefinition.SetTitleSize("16")
	queryValueWidgetDefinition.SetTime(*widgetTime)
	queryValueWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	queryValueWidget := datadogV1.NewWidget(datadogV1.QueryValueWidgetDefinitionAsWidgetDefinition(queryValueWidgetDefinition))

	// Query Widget with Formulas and Functions Query
	queryValueWidgetDefinitionFormulaFunctionsQuery := datadogV1.NewQueryValueWidgetDefinitionWithDefaults()

	queryValueWidgetDefinitionFormulaFunctionsQuery.SetRequests([]datadogV1.QueryValueWidgetRequest{{
		Formulas: []datadogV1.WidgetFormula{{
			Formula: "(((errors * 0.2)) / (query * 0.3))",
			Alias:   datadog.PtrString("sample_performance_calculator"),
		}},
		ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
		Queries: []datadogV1.FormulaAndFunctionQueryDefinition{{
			FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
				DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
				Query:      "avg:dd.metrics.query.sq.by_source{service:query}.as_count()",
				Name:       "query",
			},
		},
			{
				FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
					DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
					Compute: datadogV1.FormulaAndFunctionEventQueryDefinitionCompute{
						Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
					},
					Search: &datadogV1.FormulaAndFunctionEventQueryDefinitionSearch{
						Query: "service:query Errors",
					},
					GroupBy: []datadogV1.FormulaAndFunctionEventQueryGroupBy{{
						Facet: "host",
					}},
					Indexes: []string{"*"},
					Name:    "errors",
				},
			},
			{
				FormulaAndFunctionProcessQueryDefinition: &datadogV1.FormulaAndFunctionProcessQueryDefinition{
					DataSource: datadogV1.FORMULAANDFUNCTIONPROCESSQUERYDATASOURCE_PROCESS,
					TextFilter: datadog.PtrString(""),
					Metric:     "process.stat.cpu.total_pct",
					Limit:      datadog.PtrInt64(10),
					Name:       "process_query",
				},
			},
		}}})
	queryValueWidgetDefinitionFormulaFunctionsQuery.SetTitle("Test Formulas and Functions Metric + Event query")
	queryValueWidgetDefinitionFormulaFunctionsQuery.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	queryValueWidgetDefinitionFormulaFunctionsQuery.SetTitleSize("16")
	queryValueWidgetDefinitionFormulaFunctionsQuery.SetTime(*widgetTime)

	queryValueWidgetFormulaFunctionsQuery := datadogV1.NewWidget(datadogV1.QueryValueWidgetDefinitionAsWidgetDefinition(queryValueWidgetDefinitionFormulaFunctionsQuery))

	// Scatter Plot Widget
	scatterPlotWidgetDefinition := datadogV1.NewScatterPlotWidgetDefinitionWithDefaults()
	scatterPlotWidgetDefinition.SetRequests(datadogV1.ScatterPlotWidgetDefinitionRequests{
		X: &datadogV1.ScatterPlotRequest{
			Q:          datadog.PtrString("avg:system.load.1{*}"),
			Aggregator: datadogV1.SCATTERPLOTWIDGETAGGREGATOR_AVERAGE.Ptr()},
		Y: &datadogV1.ScatterPlotRequest{
			Q:          datadog.PtrString("avg:system.load.1{*}"),
			Aggregator: datadogV1.SCATTERPLOTWIDGETAGGREGATOR_AVERAGE.Ptr(),
		},
	})
	scatterPlotWidgetDefinition.SetYaxis(*widgetAxis)
	scatterPlotWidgetDefinition.SetYaxis(*widgetAxis)
	scatterPlotWidgetDefinition.SetColorByGroups([]string{"env"})
	scatterPlotWidgetDefinition.SetTitle("Test ScatterPlot Widget")
	scatterPlotWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	scatterPlotWidgetDefinition.SetTitleSize("16")
	scatterPlotWidgetDefinition.SetTime(*widgetTime)
	scatterPlotWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	scatterPlotWidget := datadogV1.NewWidget(datadogV1.ScatterPlotWidgetDefinitionAsWidgetDefinition(scatterPlotWidgetDefinition))

	// SLO Widget
	sloWidgetDefinition := datadogV1.NewSLOWidgetDefinitionWithDefaults()
	sloWidgetDefinition.SetViewType("detail")
	sloWidgetDefinition.SetTitle("Test SLO Widget")
	sloWidgetDefinition.SetTitleSize("16")
	sloWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	sloWidgetDefinition.SetSloId(slo.GetId())
	sloWidgetDefinition.SetShowErrorBudget(true)
	sloWidgetDefinition.SetViewMode(datadogV1.WIDGETVIEWMODE_BOTH)
	sloWidgetDefinition.SetTimeWindows([]datadogV1.WidgetTimeWindows{
		datadogV1.WIDGETTIMEWINDOWS_SEVEN_DAYS,
	})

	sloWidget := datadogV1.NewWidget(datadogV1.SLOWidgetDefinitionAsWidgetDefinition(sloWidgetDefinition))

	// Service Map Widget
	serviceMapWidgetDefinition := datadogV1.NewServiceMapWidgetDefinitionWithDefaults()
	serviceMapWidgetDefinition.SetFilters([]string{"*"})
	serviceMapWidgetDefinition.SetService("1234")
	serviceMapWidgetDefinition.SetTitle("Test Service Map Widget")
	serviceMapWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	serviceMapWidgetDefinition.SetTitleSize("16")
	serviceMapWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	serviceMapWidget := datadogV1.NewWidget(datadogV1.ServiceMapWidgetDefinitionAsWidgetDefinition(serviceMapWidgetDefinition))

	// Service Summary Widget
	serviceSummaryWidgetDefinition := datadogV1.NewServiceSummaryWidgetDefinitionWithDefaults()
	serviceSummaryWidgetDefinition.SetEnv("prod")
	serviceSummaryWidgetDefinition.SetService("web")
	serviceSummaryWidgetDefinition.SetSpanName("endpoint")
	serviceSummaryWidgetDefinition.SetShowHits(true)
	serviceSummaryWidgetDefinition.SetShowErrors(true)
	serviceSummaryWidgetDefinition.SetShowLatency(true)
	serviceSummaryWidgetDefinition.SetShowBreakdown(true)
	serviceSummaryWidgetDefinition.SetShowDistribution(true)
	serviceSummaryWidgetDefinition.SetShowResourceList(true)
	serviceSummaryWidgetDefinition.SetSizeFormat(datadogV1.WIDGETSIZEFORMAT_LARGE)
	serviceSummaryWidgetDefinition.SetDisplayFormat(datadogV1.WIDGETSERVICESUMMARYDISPLAYFORMAT_TWO_COLUMN)
	serviceSummaryWidgetDefinition.SetTitle("Test Service Summary Widget")
	serviceSummaryWidgetDefinition.SetTitleSize("16")
	serviceSummaryWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	serviceSummaryWidgetDefinition.SetTime(*widgetTimePastOneHour)

	serviceSummaryWidget := datadogV1.NewWidget(datadogV1.ServiceSummaryWidgetDefinitionAsWidgetDefinition(serviceSummaryWidgetDefinition))
	serviceSummaryWidget.SetLayout(*widgetLayout)

	// Table Widget
	tableWidgetDefinition := datadogV1.NewTableWidgetDefinitionWithDefaults()
	tableWidgetDefinition.SetRequests([]datadogV1.TableWidgetRequest{{
		Q:          datadog.PtrString("avg:system.load.1{*}"),
		Alias:      datadog.PtrString("System Load"),
		Aggregator: datadogV1.WIDGETAGGREGATOR_AVERAGE.Ptr(),
		Limit:      datadog.PtrInt64(50),
		Order:      datadogV1.WIDGETSORT_ASCENDING.Ptr(),
		ConditionalFormats: []datadogV1.WidgetConditionalFormat{{
			Comparator:    datadogV1.WIDGETCOMPARATOR_GREATER_THAN,
			Value:         7.,
			Palette:       datadogV1.WIDGETPALETTE_RED_ON_WHITE,
			CustomBgColor: datadog.PtrString("blue"),
			CustomFgColor: datadog.PtrString("black"),
			ImageUrl:      datadog.PtrString("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4"),
		}},
		CellDisplayMode: []datadogV1.TableWidgetCellDisplayMode{datadogV1.TABLEWIDGETCELLDISPLAYMODE_NUMBER},
	}})
	tableWidgetDefinition.SetTitle("Test Table Widget")
	tableWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	tableWidgetDefinition.SetTitleSize("16")
	tableWidgetDefinition.SetTime(*widgetTime)
	tableWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})
	tableWidgetDefinition.SetHasSearchBar(datadogV1.TABLEWIDGETHASSEARCHBAR_AUTO)

	tableWidget := datadogV1.NewWidget(datadogV1.TableWidgetDefinitionAsWidgetDefinition(tableWidgetDefinition))

	// Table Widget with Formulas and Functions Query
	tableWidgetDefinitionFormulaFunctionsQuery := datadogV1.NewTableWidgetDefinitionWithDefaults()

	tableWidgetDefinitionFormulaFunctionsQuery.SetRequests([]datadogV1.TableWidgetRequest{{
		Formulas: []datadogV1.WidgetFormula{{
			Formula: "(((errors * 0.2)) / (query * 0.3))",
			Alias:   datadog.PtrString("sample_performance_calculator"),
			ConditionalFormats: []datadogV1.WidgetConditionalFormat{{
				Comparator:    datadogV1.WIDGETCOMPARATOR_GREATER_THAN,
				Value:         7.,
				Palette:       datadogV1.WIDGETPALETTE_RED_ON_WHITE,
				CustomBgColor: datadog.PtrString("blue"),
				CustomFgColor: datadog.PtrString("black"),
				ImageUrl:      datadog.PtrString("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4"),
			}},
			CellDisplayMode: datadogV1.TABLEWIDGETCELLDISPLAYMODE_NUMBER.Ptr(),
		}},
		ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
		Queries: []datadogV1.FormulaAndFunctionQueryDefinition{{
			FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
				DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
				Query:      "avg:dd.metrics.query.sq.by_source{service:query}.as_count()",
				Name:       "query",
			},
		},
		}}})
	tableWidgetDefinitionFormulaFunctionsQuery.SetTitle("Test Formulas and Functions Metric + Event query")
	tableWidgetDefinitionFormulaFunctionsQuery.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	tableWidgetDefinitionFormulaFunctionsQuery.SetTitleSize("16")
	tableWidgetDefinitionFormulaFunctionsQuery.SetTime(*widgetTime)

	tableWidgetFormulaFunctionsQuery := datadogV1.NewWidget(datadogV1.TableWidgetDefinitionAsWidgetDefinition(tableWidgetDefinitionFormulaFunctionsQuery))

	// Table Widget with APM Stats data
	tableWidgetApmStatsDefinition := datadogV1.NewTableWidgetDefinitionWithDefaults()
	tableWidgetApmStatsDefinition.SetRequests([]datadogV1.TableWidgetRequest{{
		ApmStatsQuery: &datadogV1.ApmStatsQueryDefinition{
			Env:        "prod",
			Name:       "web",
			PrimaryTag: "foo:*",
			Resource:   datadog.PtrString("endpoint"),
			RowType:    datadogV1.APMSTATSQUERYROWTYPE_SPAN,
			Columns: []datadogV1.ApmStatsQueryColumnType{{
				Name: "baz",
			}},
		},
	}})
	tableWidgetApmStatsDefinition.SetTitle("Test Table Widget with APM Stats Data")
	tableWidgetApmStatsDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	tableWidgetApmStatsDefinition.SetTitleSize("16")
	tableWidgetApmStatsDefinition.SetTime(*widgetTime)
	tableWidgetApmStatsDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	tableWidgetApmStats := datadogV1.NewWidget(datadogV1.TableWidgetDefinitionAsWidgetDefinition(tableWidgetApmStatsDefinition))

	// Timeseries Widget
	timeseriesWidgetDefinition := datadogV1.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinition.SetRequests([]datadogV1.TimeseriesWidgetRequest{{
		Q: datadog.PtrString("avg:system.load.1{*}"),
		Style: &datadogV1.WidgetRequestStyle{
			Palette:   datadog.PtrString("dog_classic"),
			LineType:  datadogV1.WIDGETLINETYPE_DASHED.Ptr(),
			LineWidth: datadogV1.WIDGETLINEWIDTH_THICK.Ptr(),
		},
		Metadata: []datadogV1.TimeseriesWidgetExpressionAlias{{
			Expression: "avg:system.load.1{*}",
			AliasName:  datadog.PtrString("Aliased metric"),
		}},
		DisplayType:  datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
		OnRightYaxis: datadog.PtrBool(true),
	}})
	timeseriesWidgetDefinition.SetYaxis(datadogV1.WidgetAxis{
		IncludeZero: datadog.PtrBool(true),
		Min:         datadog.PtrString("0"),
		Max:         datadog.PtrString("100"),
		Scale:       datadog.PtrString("linear"),
		Label:       datadog.PtrString("Widget Scale")})
	timeseriesWidgetDefinition.SetRightYaxis(datadogV1.WidgetAxis{
		IncludeZero: datadog.PtrBool(true),
		Min:         datadog.PtrString("0"),
		Max:         datadog.PtrString("100"),
		Scale:       datadog.PtrString("linear"),
		Label:       datadog.PtrString("Widget Scale")})
	timeseriesWidgetDefinition.SetEvents([]datadogV1.WidgetEvent{{
		Q: "Build succeeded",
	}})
	timeseriesWidgetDefinition.SetMarkers([]datadogV1.WidgetMarker{{
		Value:       "y=15",
		DisplayType: datadog.PtrString("error dashed"),
		Label:       datadog.PtrString("error threshold"),
		Time:        datadog.PtrString("4h"),
	}})
	timeseriesWidgetDefinition.SetTitle("Test Timeseries Widget")
	timeseriesWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinition.SetTitleSize("16")
	timeseriesWidgetDefinition.SetTime(*widgetTime)
	timeseriesWidgetDefinition.SetShowLegend(true)
	timeseriesWidgetDefinition.SetLegendSize("16")
	timeseriesWidgetDefinition.SetLegendLayout("horizontal")
	timeseriesWidgetDefinition.SetLegendColumns([]datadogV1.TimeseriesWidgetLegendColumn{"value", "min", "max"})
	timeseriesWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	timeseriesWidget := datadogV1.NewWidget(datadogV1.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinition))

	// Timeseries Widget with Process query
	timeseriesWidgetDefinitionProcessQuery := datadogV1.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinitionProcessQuery.SetRequests([]datadogV1.TimeseriesWidgetRequest{{
		ProcessQuery: &datadogV1.ProcessQueryDefinition{
			Metric:   "process.stat.cpu.total_pct",
			FilterBy: []string{"account:test"},
			Limit:    datadog.PtrInt64(10),
			SearchBy: datadog.PtrString("editor"),
		},
		Style: &datadogV1.WidgetRequestStyle{
			Palette:   datadog.PtrString("dog_classic"),
			LineType:  datadogV1.WIDGETLINETYPE_DASHED.Ptr(),
			LineWidth: datadogV1.WIDGETLINEWIDTH_THICK.Ptr(),
		},
		Metadata: []datadogV1.TimeseriesWidgetExpressionAlias{{
			Expression: "avg:system.load.1{*}",
			AliasName:  datadog.PtrString("Aliased metric"),
		}},
		DisplayType:  datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
		OnRightYaxis: datadog.PtrBool(true),
	}})
	timeseriesWidgetDefinitionProcessQuery.SetYaxis(datadogV1.WidgetAxis{
		IncludeZero: datadog.PtrBool(true),
		Min:         datadog.PtrString("0"),
		Max:         datadog.PtrString("100"),
		Scale:       datadog.PtrString("linear")})
	timeseriesWidgetDefinitionProcessQuery.SetRightYaxis(datadogV1.WidgetAxis{
		IncludeZero: datadog.PtrBool(true),
		Min:         datadog.PtrString("0"),
		Max:         datadog.PtrString("100"),
		Scale:       datadog.PtrString("linear")})
	timeseriesWidgetDefinitionProcessQuery.SetEvents([]datadogV1.WidgetEvent{{
		Q: "Build succeeded",
	}})
	timeseriesWidgetDefinitionProcessQuery.SetMarkers([]datadogV1.WidgetMarker{{
		Value:       "y=15",
		DisplayType: datadog.PtrString("error dashed"),
		Label:       datadog.PtrString("error threshold"),
		Time:        datadog.PtrString("4h"),
	}})
	timeseriesWidgetDefinitionProcessQuery.SetTitle("Test Timeseries Widget with Process Query")
	timeseriesWidgetDefinitionProcessQuery.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinitionProcessQuery.SetTitleSize("16")
	timeseriesWidgetDefinitionProcessQuery.SetTime(*widgetTime)
	timeseriesWidgetDefinitionProcessQuery.SetShowLegend(true)
	timeseriesWidgetDefinitionProcessQuery.SetLegendSize("16")
	timeseriesWidgetDefinitionProcessQuery.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	timeseriesWidgetProcessQuery := datadogV1.NewWidget(datadogV1.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinitionProcessQuery))

	// Timeseries Widget with Log query (APM/Log/Network/Rum/Event share schemas, so only test one)
	timeseriesWidgetDefinitionLogQuery := datadogV1.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinitionLogQuery.SetRequests([]datadogV1.TimeseriesWidgetRequest{{
		LogQuery: &datadogV1.LogQueryDefinition{
			Index: datadog.PtrString("main"),
			Compute: &datadogV1.LogsQueryCompute{
				Aggregation: "count",
				Facet:       datadog.PtrString("host"),
				Interval:    datadog.PtrInt64(10),
			},
			Search: &datadogV1.LogQueryDefinitionSearch{Query: "Error parsing"},
			GroupBy: []datadogV1.LogQueryDefinitionGroupBy{{
				Facet: "host",
				Limit: datadog.PtrInt64(5),
				Sort: &datadogV1.LogQueryDefinitionGroupBySort{
					Aggregation: "count",
					Order:       datadogV1.WIDGETSORT_ASCENDING,
				},
			}},
		},
		Style: &datadogV1.WidgetRequestStyle{
			Palette:   datadog.PtrString("dog_classic"),
			LineType:  datadogV1.WIDGETLINETYPE_DASHED.Ptr(),
			LineWidth: datadogV1.WIDGETLINEWIDTH_THICK.Ptr(),
		},
		Metadata: []datadogV1.TimeseriesWidgetExpressionAlias{{
			Expression: "avg:system.load.1{*}",
			AliasName:  datadog.PtrString("Aliased metric"),
		}},
		DisplayType:  datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
		OnRightYaxis: datadog.PtrBool(true),
	}})
	timeseriesWidgetDefinitionLogQuery.SetYaxis(*widgetAxis)
	timeseriesWidgetDefinitionLogQuery.SetRightYaxis(*widgetAxis)
	timeseriesWidgetDefinitionLogQuery.SetEvents([]datadogV1.WidgetEvent{{
		Q: "Build succeeded",
	}})
	timeseriesWidgetDefinitionLogQuery.SetMarkers([]datadogV1.WidgetMarker{{
		Value:       "y=15",
		DisplayType: datadog.PtrString("error dashed"),
		Label:       datadog.PtrString("error threshold"),
		Time:        datadog.PtrString("4h"),
	}})
	timeseriesWidgetDefinitionLogQuery.SetTitle("Test Timeseries Widget with Log Query")
	timeseriesWidgetDefinitionLogQuery.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinitionLogQuery.SetTitleSize("16")
	timeseriesWidgetDefinitionLogQuery.SetTime(*widgetTime)
	timeseriesWidgetDefinitionLogQuery.SetShowLegend(true)
	timeseriesWidgetDefinitionLogQuery.SetLegendSize("16")
	timeseriesWidgetDefinitionLogQuery.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	timeseriesWidgetLogQuery := datadogV1.NewWidget(datadogV1.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinitionLogQuery))

	// Timeseries Widget with Event query
	timeseriesWidgetDefinitionEventQuery := datadogV1.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinitionEventQuery.SetRequests([]datadogV1.TimeseriesWidgetRequest{{
		EventQuery: &datadogV1.LogQueryDefinition{
			Index: datadog.PtrString("*"),
			Compute: &datadogV1.LogsQueryCompute{
				Aggregation: "count",
				Facet:       datadog.PtrString("host"),
				Interval:    datadog.PtrInt64(10),
			},
			Search: &datadogV1.LogQueryDefinitionSearch{Query: "source:kubernetes"},
			GroupBy: []datadogV1.LogQueryDefinitionGroupBy{{
				Facet: "host",
				Limit: datadog.PtrInt64(5),
				Sort: &datadogV1.LogQueryDefinitionGroupBySort{
					Aggregation: "count",
					Order:       datadogV1.WIDGETSORT_ASCENDING,
				},
			}},
		},
		Style: &datadogV1.WidgetRequestStyle{
			Palette:   datadog.PtrString("dog_classic"),
			LineType:  datadogV1.WIDGETLINETYPE_DASHED.Ptr(),
			LineWidth: datadogV1.WIDGETLINEWIDTH_THICK.Ptr()},
		Metadata: []datadogV1.TimeseriesWidgetExpressionAlias{{
			Expression: "avg:system.load.1{*}",
			AliasName:  datadog.PtrString("Aliased metric"),
		}},
		DisplayType:  datadogV1.WIDGETDISPLAYTYPE_LINE.Ptr(),
		OnRightYaxis: datadog.PtrBool(true),
	}})
	timeseriesWidgetDefinitionEventQuery.SetYaxis(*widgetAxis)
	timeseriesWidgetDefinitionEventQuery.SetRightYaxis(*widgetAxis)
	timeseriesWidgetDefinitionEventQuery.SetEvents([]datadogV1.WidgetEvent{{
		Q: "Build succeeded",
	}})
	timeseriesWidgetDefinitionEventQuery.SetMarkers([]datadogV1.WidgetMarker{{
		Value:       "y=15",
		DisplayType: datadog.PtrString("error dashed"),
		Label:       datadog.PtrString("error threshold"),
		Time:        datadog.PtrString("4h"),
	}})
	timeseriesWidgetDefinitionEventQuery.SetTitle("Test Timeseries Widget with Event Query")
	timeseriesWidgetDefinitionEventQuery.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinitionEventQuery.SetTitleSize("16")
	timeseriesWidgetDefinitionEventQuery.SetTime(*widgetTime)
	timeseriesWidgetDefinitionEventQuery.SetShowLegend(true)
	timeseriesWidgetDefinitionEventQuery.SetLegendSize("16")
	timeseriesWidgetDefinitionEventQuery.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	timeseriesWidgetEventQuery := datadogV1.NewWidget(datadogV1.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinitionEventQuery))

	// Timeseries Widget with Formulas and Functions Query
	timeseriesWidgetDefinitionFormulaFunctionsQuery := datadogV1.NewTimeseriesWidgetDefinitionWithDefaults()

	timeseriesWidgetDefinitionFormulaFunctionsQuery.SetRequests([]datadogV1.TimeseriesWidgetRequest{{
		Formulas: []datadogV1.WidgetFormula{{
			Formula: "(((errors * 0.2)) / (query * 0.3))",
			Alias:   datadog.PtrString("sample_performance_calculator"),
		}},
		ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
		Queries: []datadogV1.FormulaAndFunctionQueryDefinition{{
			FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
				DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
				Query:      "avg:dd.metrics.query.sq.by_source{service:query}.as_count()",
				Name:       "query",
			},
		},
			{
				FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
					DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
					Compute: datadogV1.FormulaAndFunctionEventQueryDefinitionCompute{
						Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
					},
					Search: &datadogV1.FormulaAndFunctionEventQueryDefinitionSearch{
						Query: "service:query Errors",
					},
					GroupBy: []datadogV1.FormulaAndFunctionEventQueryGroupBy{{
						Facet: "host",
					}},
					Indexes: []string{"*"},
					Name:    "errors",
				},
			},
			{
				FormulaAndFunctionProcessQueryDefinition: &datadogV1.FormulaAndFunctionProcessQueryDefinition{
					DataSource: datadogV1.FORMULAANDFUNCTIONPROCESSQUERYDATASOURCE_PROCESS,
					TextFilter: datadog.PtrString(""),
					Metric:     "process.stat.cpu.total_pct",
					Limit:      datadog.PtrInt64(10),
					Name:       "process_query",
				},
			},
		}}})
	timeseriesWidgetDefinitionFormulaFunctionsQuery.SetTitle("Test Formulas and Functions Metric + Event query")
	timeseriesWidgetDefinitionFormulaFunctionsQuery.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinitionFormulaFunctionsQuery.SetTitleSize("16")
	timeseriesWidgetDefinitionFormulaFunctionsQuery.SetTime(*widgetTime)

	timeseriesWidgetFormulaFunctionsQuery := datadogV1.NewWidget(datadogV1.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinitionFormulaFunctionsQuery))

	// Toplist Widget
	toplistWidgetDefinition := datadogV1.NewToplistWidgetDefinitionWithDefaults()
	toplistWidgetDefinition.SetRequests([]datadogV1.ToplistWidgetRequest{{
		Q: datadog.PtrString("avg:system.load.1{*}"),
		ConditionalFormats: []datadogV1.WidgetConditionalFormat{{
			Comparator:    datadogV1.WIDGETCOMPARATOR_GREATER_THAN,
			Value:         7.,
			Palette:       datadogV1.WIDGETPALETTE_RED_ON_WHITE,
			CustomBgColor: datadog.PtrString("blue"),
			CustomFgColor: datadog.PtrString("black"),
			ImageUrl:      datadog.PtrString("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4"),
		}},
	}})
	toplistWidgetDefinition.SetTitle("Test Toplist Widget")
	toplistWidgetDefinition.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	toplistWidgetDefinition.SetTitleSize("16")
	toplistWidgetDefinition.SetTime(*widgetTime)
	toplistWidgetDefinition.SetCustomLinks([]datadogV1.WidgetCustomLink{
		customLink,
	})

	toplistWidget := datadogV1.NewWidget(datadogV1.ToplistWidgetDefinitionAsWidgetDefinition(toplistWidgetDefinition))

	// Toplist Widget with Formulas and Functions Query
	toplistWidgetDefinitionFormulaFunctionsQuery := datadogV1.NewToplistWidgetDefinitionWithDefaults()

	toplistWidgetDefinitionFormulaFunctionsQuery.SetRequests([]datadogV1.ToplistWidgetRequest{{
		Formulas: []datadogV1.WidgetFormula{{
			Formula: "(((errors * 0.2)) / (query * 0.3))",
			Alias:   datadog.PtrString("sample_performance_calculator"),
		}},
		ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_TIMESERIES.Ptr(),
		Queries: []datadogV1.FormulaAndFunctionQueryDefinition{{
			FormulaAndFunctionMetricQueryDefinition: &datadogV1.FormulaAndFunctionMetricQueryDefinition{
				DataSource: datadogV1.FORMULAANDFUNCTIONMETRICDATASOURCE_METRICS,
				Query:      "avg:dd.metrics.query.sq.by_source{service:query}.as_count()",
				Name:       "query",
			},
		},
			{
				FormulaAndFunctionEventQueryDefinition: &datadogV1.FormulaAndFunctionEventQueryDefinition{
					DataSource: datadogV1.FORMULAANDFUNCTIONEVENTSDATASOURCE_LOGS,
					Compute: datadogV1.FormulaAndFunctionEventQueryDefinitionCompute{
						Aggregation: datadogV1.FORMULAANDFUNCTIONEVENTAGGREGATION_COUNT,
					},
					Search: &datadogV1.FormulaAndFunctionEventQueryDefinitionSearch{
						Query: "service:query Errors",
					},
					GroupBy: []datadogV1.FormulaAndFunctionEventQueryGroupBy{{
						Facet: "host",
					}},
					Indexes: []string{"*"},
					Name:    "errors",
				},
			},
			{
				FormulaAndFunctionProcessQueryDefinition: &datadogV1.FormulaAndFunctionProcessQueryDefinition{
					DataSource: datadogV1.FORMULAANDFUNCTIONPROCESSQUERYDATASOURCE_PROCESS,
					TextFilter: datadog.PtrString(""),
					Metric:     "process.stat.cpu.total_pct",
					Limit:      datadog.PtrInt64(10),
					Name:       "process_query",
				},
			},
		}}})
	toplistWidgetDefinitionFormulaFunctionsQuery.SetTitle("Test Formulas and Functions Metric + Event query")
	toplistWidgetDefinitionFormulaFunctionsQuery.SetTitleAlign(datadogV1.WIDGETTEXTALIGN_CENTER)
	toplistWidgetDefinitionFormulaFunctionsQuery.SetTitleSize("16")
	toplistWidgetDefinitionFormulaFunctionsQuery.SetTime(*widgetTime)

	toplistWidgetFormulaFunctionsQuery := datadogV1.NewWidget(datadogV1.ToplistWidgetDefinitionAsWidgetDefinition(toplistWidgetDefinitionFormulaFunctionsQuery))

	// Template Variables
	templateVariable := datadogV1.NewDashboardTemplateVariableWithDefaults()
	templateVariable.SetName("test template var")
	templateVariable.SetPrefix("test-go")
	templateVariable.SetDefault("*")
	templateVariable.SetAvailableValues([]string{"available-value-1, available-value-2"})

	// Template Variable Presets
	dashboardTemplateVariablePreset := datadogV1.NewDashboardTemplateVariablePreset()
	dashboardTemplateVariablePreset.SetName("Test Preset")
	dashboardTemplateVariablePresetValue := datadogV1.NewDashboardTemplateVariablePresetValueWithDefaults()
	dashboardTemplateVariablePresetValue.SetName("test preset")
	dashboardTemplateVariablePresetValue.SetValue("*")
	dashboardTemplateVariablePreset.SetTemplateVariables([]datadogV1.DashboardTemplateVariablePresetValue{*dashboardTemplateVariablePresetValue})

	orderedWidgetList := []datadogV1.Widget{
		*alertGraphWidget,
		*alertValueWidget,
		*changeWidget,
		*checkStatusWidget,
		*distributionWidget,
		*geoMapWidget,
		*geoMapWidgetFormulaFunctionsQuery,
		*groupWidget,
		*heatMapWidget,
		*hostMapWidget,
		*noteWidget,
		*queryValueWidget,
		*queryValueWidgetFormulaFunctionsQuery,
		*scatterPlotWidget,
		*sloWidget,
		*serviceMapWidget,
		*tableWidget,
		*tableWidgetFormulaFunctionsQuery,
		*tableWidgetApmStats,
		*timeseriesWidget,
		*timeseriesWidgetProcessQuery,
		*timeseriesWidgetLogQuery,
		*timeseriesWidgetEventQuery,
		*timeseriesWidgetFormulaFunctionsQuery,
		*toplistWidgetFormulaFunctionsQuery,
		*toplistWidget,
	}

	dashboard := datadogV1.NewDashboardWithDefaults()
	dashboard.SetLayoutType(datadogV1.DASHBOARDLAYOUTTYPE_ORDERED)
	dashboard.SetWidgets(orderedWidgetList)
	dashboard.SetTitle(fmt.Sprintf("%s-ordered", *tests.UniqueEntityName(ctx, t)))
	dashboard.SetDescription("Test dashboard for Go client")
	dashboard.SetIsReadOnly(false)
	dashboard.SetTemplateVariables([]datadogV1.DashboardTemplateVariable{*templateVariable})
	dashboard.SetTemplateVariablePresets([]datadogV1.DashboardTemplateVariablePreset{*dashboardTemplateVariablePreset})
	// FIXME dashboard.SetNotifyList([]string{"test@datadoghq.com"})

	createdDashboard, httpresp, err := api.CreateDashboard(ctx, *dashboard)
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboard(ctx, t, createdDashboard.GetId())
	assert.Equal(200, httpresp.StatusCode)

	getDashboard, httpresp, err := api.GetDashboard(ctx, createdDashboard.GetId())
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(getDashboard, createdDashboard)

	assertDashboardDefinition := func(dashboard *datadogV1.Dashboard, response datadogV1.Dashboard) {
		// Assert root dashboard items on the create response
		assert.Equal(dashboard.GetTitle(), response.GetTitle())
		assert.Equal(dashboard.GetDescription(), response.GetDescription())
		assert.Equal(dashboard.GetIsReadOnly(), response.GetIsReadOnly())
		// The end of the url is a normalized version of the title, so lets just check the beginning of the URL
		assert.True(strings.HasPrefix(response.GetUrl(), fmt.Sprintf("/dashboard/%s", response.GetId())))
		assert.Greater(response.GetCreatedAt().Unix(), int64(0))
		assert.Greater(response.GetModifiedAt().Unix(), int64(0))
		assert.Greater(len(response.GetAuthorHandle()), 0)
		assert.Equal(dashboard.GetLayoutType(), response.GetLayoutType())
		// FIXME assert.Equal(len(dashboard.GetNotifyList()), len(response.GetNotifyList()))

		// Template Variables
		assert.Equal(dashboard.GetTemplateVariables()[0], response.GetTemplateVariables()[0])

		for index, checkWidget := range response.GetWidgets() {
			assert.True(checkWidget.GetId() > 0)
			checkWidget.Id = nil
			checkWidgetDefinition := checkWidget.GetDefinition()
			checkWidgetInstance := checkWidgetDefinition.GetActualInstance()
			if def, ok := checkWidgetInstance.(*datadogV1.GroupWidgetDefinition); ok {
				for index := range def.GetWidgets() {
					def.Widgets[index].Id = nil
				}
			}
			assert.Equal(dashboard.GetWidgets()[index], checkWidget)
		}

	}

	assertDashboardDefinition(dashboard, createdDashboard)

	freeWidgetList := []datadogV1.Widget{
		*eventStreamWidget,
		*eventTimelineWidget,
		*freeTextWidget,
		*iFrameWidget,
		*imageWidget,
		*logStreamWidget,
		*monitorSummaryWidget,
		*serviceSummaryWidget,
	}

	freeDashboard := datadogV1.NewDashboardWithDefaults()
	freeDashboard.SetLayoutType(datadogV1.DASHBOARDLAYOUTTYPE_FREE)
	freeDashboard.SetWidgets(freeWidgetList)
	freeDashboard.SetTitle(fmt.Sprintf("%s-free", *tests.UniqueEntityName(ctx, t)))
	freeDashboard.SetDescription("Test Free layout dashboard for Go client")
	freeDashboard.SetIsReadOnly(false)
	freeDashboard.SetTemplateVariables([]datadogV1.DashboardTemplateVariable{*templateVariable})

	createdFreeDashboard, httpresp, err := api.CreateDashboard(ctx, *freeDashboard)
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboard(ctx, t, createdFreeDashboard.GetId())
	assert.Equal(200, httpresp.StatusCode)

	getFreeDashboard, httpresp, err := api.GetDashboard(ctx, createdFreeDashboard.GetId())
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(getFreeDashboard, createdFreeDashboard)

	assertDashboardDefinition(freeDashboard, createdFreeDashboard)

	// Update the dashboard and set all nullable fields to null
	dashboard.SetDescriptionNil()
	dashboard.SetTemplateVariables(nil)
	dashboard.SetTemplateVariablePresets(nil)
	dashboard.SetNotifyList(nil)

	// Update dashboard widgets
	dashboardWidgets := dashboard.GetWidgets()

	noteWidgetDefinition.SetContent("Updated content")
	noteWidgetDefinition.SetFontSize("30")
	noteWidget.SetDefinition(datadogV1.NoteWidgetDefinitionAsWidgetDefinition(noteWidgetDefinition))

	noteWidgetIndex := 8
	dashboardWidgets[noteWidgetIndex] = *noteWidget
	dashboard.SetWidgets(dashboardWidgets)

	updateResponse, httpresp, err := api.UpdateDashboard(ctx, createdDashboard.GetId(), *dashboard)
	if err != nil {
		t.Fatalf("Error updating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(dashboard.GetTitle(), updateResponse.GetTitle())
	v, ok := updateResponse.GetDescriptionOk()
	assert.Nil(v)
	assert.True(ok)
	assert.Nil(updateResponse.GetTemplateVariables())
	assert.Nil(updateResponse.GetTemplateVariablePresets())
	assert.Nil(updateResponse.GetNotifyList())

	updateResponse.Id = nil
	updateResponse.GetWidgets()[0].Id = nil
	assert.Equal(dashboard.GetWidgets()[0], updateResponse.GetWidgets()[0])

	foundWidget := false
	for index, noteWidgetResponse := range updateResponse.GetWidgets() {
		noteWidgetResponseDefinition := noteWidgetResponse.GetDefinition()
		noteWidgetResponseInstance := noteWidgetResponseDefinition.GetActualInstance()
		if def, ok := noteWidgetResponseInstance.(*datadogV1.NoteWidgetDefinition); ok {
			foundWidget = true
			assert.Equal("Updated content", def.GetContent())
			assert.Equal("30", def.GetFontSize())
			assert.Equal(noteWidgetIndex, index)
			break
		}
	}
	assert.True(foundWidget)
	assert.True(len(updateResponse.GetWidgets()) > 1)

	deleteResponse, httpresp, err := api.DeleteDashboard(ctx, createdFreeDashboard.GetId())
	if err != nil {
		t.Fatalf("Error deleting dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(createdFreeDashboard.GetId(), deleteResponse.GetDeletedDashboardId())

}

func TestDashboardGetAll(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	api := datadogV1.NewDashboardsApi(Client(ctx))

	getAllResponse, httpresp, err := api.ListDashboards(ctx)
	if err != nil {
		t.Fatalf("Error getting all dashboards: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(len(getAllResponse.GetDashboards()) >= 1)
}

func TestDashboardCreateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.Dashboard
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.Dashboard{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.Dashboard{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDashboardsApi(Client(ctx))

			_, httpresp, err := api.CreateDashboard(ctx, tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDashboardsApi(Client(ctx))

			_, httpresp, err := api.ListDashboards(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardDeleteErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDashboardsApi(Client(ctx))

			_, httpresp, err := api.DeleteDashboard(ctx, "123-abc-xyz")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardUpdateErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	dashboardOK := *datadogV1.NewDashboardWithDefaults()
	dashboardOK.SetWidgets([]datadogV1.Widget{})
	dashboardOK.SetLayoutType(datadogV1.DASHBOARDLAYOUTTYPE_FREE)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadogV1.Dashboard
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadogV1.Dashboard{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadogV1.Dashboard{}, 403},
		"404 Not Found":   {WithTestAuth, dashboardOK, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDashboardsApi(Client(ctx))

			_, httpresp, err := api.UpdateDashboard(ctx, "123-abc-xyz", tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardGetErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
		"404 Not Found": {WithTestAuth, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)
			api := datadogV1.NewDashboardsApi(Client(ctx))

			_, httpresp, err := api.GetDashboard(ctx, "123-abc-xyz")
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadogV1.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteDashboard(ctx context.Context, t *testing.T, dashboardID string) {
	api := datadogV1.NewDashboardsApi(Client(ctx))

	_, httpresp, err := api.DeleteDashboard(ctx, dashboardID)
	if err != nil && httpresp.StatusCode != 404 {
		t.Logf("Error deleting Dashboard: %v, Another test may have already deleted this dashboard.", dashboardID)
	}
}
