/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestDashboardLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// create SLO for referencing in SLO widget (we're borrowing these from api_slo_test.go)
	testEventSLO := getTestEventSLO(ctx, t)
	sloResp, httpresp, err := Client(ctx).ServiceLevelObjectivesApi.CreateSLO(ctx).Body(testEventSLO).Execute()
	if err != nil {
		t.Fatalf("Error creating SLO %v for testing Dashboard SLO widget: Response %s: %v", testEventSLO, err.Error(), err)
	}
	slo := sloResp.GetData()[0]
	defer deleteSLOIfExists(ctx, slo.GetId())
	assert.Equal(httpresp.StatusCode, 200)

	widgetTime := datadog.NewWidgetTimeWithDefaults()
	widgetTime.SetLiveSpan(datadog.WIDGETLIVESPAN_PAST_FIFTEEN_MINUTES)

	widgetTimePastOneHour := datadog.NewWidgetTimeWithDefaults()
	widgetTimePastOneHour.SetLiveSpan(datadog.WIDGETLIVESPAN_PAST_ONE_HOUR)

	widgetTimePastOneDay := datadog.NewWidgetTimeWithDefaults()
	widgetTimePastOneDay.SetLiveSpan(datadog.WIDGETLIVESPAN_PAST_ONE_DAY)

	widgetTimePastOneMonth := datadog.NewWidgetTimeWithDefaults()
	widgetTimePastOneMonth.SetLiveSpan(datadog.WIDGETLIVESPAN_PAST_ONE_MONTH)

	widgetLayout := datadog.NewWidgetLayoutWithDefaults()
	widgetLayout.SetHeight(10)
	widgetLayout.SetWidth(10)
	widgetLayout.SetX(0)
	widgetLayout.SetY(0)

	// Alert Graph Widget
	alertGraphDefinition := datadog.NewAlertGraphWidgetDefinitionWithDefaults()
	alertGraphDefinition.SetAlertId("1234")
	alertGraphDefinition.SetVizType(datadog.WIDGETVIZTYPE_TIMESERIES)
	alertGraphDefinition.SetTitle("Test Alert Graph Widget")
	alertGraphDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	alertGraphDefinition.SetTitleSize("16")
	alertGraphDefinition.SetTime(*widgetTime)

	alertGraphWidget := datadog.NewWidgetWithDefaults()
	alertGraphWidget.SetDefinition(datadog.AlertGraphWidgetDefinitionAsWidgetDefinition(alertGraphDefinition))

	// Alert Value Widget
	alertValueDefinition := datadog.NewAlertValueWidgetDefinitionWithDefaults()
	alertValueDefinition.SetAlertId("1234")
	alertValueDefinition.SetPrecision(2)
	alertValueDefinition.SetUnit("ms")
	alertValueDefinition.SetTitleSize("12")
	alertValueDefinition.SetTextAlign(datadog.WIDGETTEXTALIGN_CENTER)
	alertValueDefinition.SetTitle("Test Alert Value Widget")
	alertValueDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_RIGHT)

	alertValueWidget := datadog.NewWidgetWithDefaults()
	alertValueWidget.SetDefinition(datadog.AlertValueWidgetDefinitionAsWidgetDefinition(alertValueDefinition))

	// Change Widget
	changeWidgetRequest := datadog.NewChangeWidgetRequest()
	changeWidgetRequest.SetQ("avg:system.load.1{*}")
	changeWidgetRequest.SetChangeType(datadog.WIDGETCHANGETYPE_ABSOLUTE)
	changeWidgetRequest.SetCompareTo(datadog.WIDGETCOMPARETO_HOUR_BEFORE)
	changeWidgetRequest.SetIncreaseGood(true)
	changeWidgetRequest.SetOrderBy(datadog.WIDGETORDERBY_CHANGE)
	changeWidgetRequest.SetOrderDir(datadog.WIDGETSORT_ASCENDING)
	changeWidgetRequest.SetShowPresent(true)

	changeWidgetDefinition := datadog.NewChangeWidgetDefinitionWithDefaults()
	changeWidgetDefinition.SetTitle("Test Change Widget")
	changeWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	changeWidgetDefinition.SetTitleSize("16")
	changeWidgetDefinition.SetTime(*widgetTime)
	changeWidgetDefinition.SetRequests([]datadog.ChangeWidgetRequest{
		*changeWidgetRequest,
	})

	changeWidget := datadog.NewWidgetWithDefaults()
	changeWidget.SetDefinition(datadog.ChangeWidgetDefinitionAsWidgetDefinition(changeWidgetDefinition))

	// Check Status Widget
	checkStatusWidgetDefinition := datadog.NewCheckStatusWidgetDefinitionWithDefaults()
	checkStatusWidgetDefinition.SetCheck("service_check.up")
	checkStatusWidgetDefinition.SetGrouping(datadog.WIDGETGROUPING_CHECK)
	checkStatusWidgetDefinition.SetGroup("*")
	checkStatusWidgetDefinition.SetTags([]string{"foo:bar"})
	checkStatusWidgetDefinition.SetGroupBy([]string{"bar"})
	checkStatusWidgetDefinition.SetTitle("Test Check Status Widget")
	checkStatusWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	checkStatusWidgetDefinition.SetTitleSize("16")
	checkStatusWidgetDefinition.SetTime(*widgetTime)

	checkStatusWidget := datadog.NewWidgetWithDefaults()
	checkStatusWidget.SetDefinition(datadog.CheckStatusWidgetDefinitionAsWidgetDefinition(checkStatusWidgetDefinition))

	// Distribution Widget
	distributionWidgetDefinition := datadog.NewDistributionWidgetDefinitionWithDefaults()

	distributionWidgetRequest := datadog.NewDistributionWidgetRequestWithDefaults()
	distributionWidgetRequest.SetQ("avg:system.load.1{*}")

	dogClassic := datadog.NewWidgetStyleWithDefaults()
	dogClassic.SetPalette("dog_classic")

	distributionWidgetRequest.SetStyle(*dogClassic)

	distributionWidgetDefinition.SetRequests([]datadog.DistributionWidgetRequest{
		*distributionWidgetRequest,
	})
	distributionWidgetDefinition.SetShowLegend(true)
	distributionWidgetDefinition.SetTitle("Test Distribution Widget")
	distributionWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	distributionWidgetDefinition.SetTitleSize("16")
	distributionWidgetDefinition.SetTime(*widgetTime)

	distributionWidget := datadog.NewWidget(datadog.DistributionWidgetDefinitionAsWidgetDefinition(distributionWidgetDefinition))

	// Event Stream Widget ONLY AVAILABLE ON FREE LAYOUTS
	eventStreamWidgetDefinition := datadog.NewEventStreamWidgetDefinitionWithDefaults()
	eventStreamWidgetDefinition.SetQuery("Build successful")
	eventStreamWidgetDefinition.SetEventSize(datadog.WIDGETEVENTSIZE_LARGE)
	eventStreamWidgetDefinition.SetTitle("Test Event Stream Widget")
	eventStreamWidgetDefinition.SetTitleSize("16")
	eventStreamWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	eventStreamWidgetDefinition.SetTime(*widgetTimePastOneDay)

	eventStreamWidget := datadog.NewWidgetWithDefaults()
	eventStreamWidget.SetDefinition(datadog.EventStreamWidgetDefinitionAsWidgetDefinition(eventStreamWidgetDefinition))
	eventStreamWidget.SetLayout(*widgetLayout)

	// Event Timeline Widget ONLY AVAILABLE ON FREE LAYOUTS
	eventTimelineWidgetDefinition := datadog.NewEventTimelineWidgetDefinitionWithDefaults()
	eventTimelineWidgetDefinition.SetQuery("Build Failed")
	eventTimelineWidgetDefinition.SetTitle("Test Event Timeline Widget")
	eventTimelineWidgetDefinition.SetTitleSize("16")
	eventTimelineWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_LEFT)
	eventTimelineWidgetDefinition.SetTime(*widgetTimePastOneMonth)

	eventTimelineWidget := datadog.NewWidgetWithDefaults()
	eventTimelineWidget.SetDefinition(datadog.EventTimelineWidgetDefinitionAsWidgetDefinition(eventTimelineWidgetDefinition))
	eventTimelineWidget.SetLayout(*widgetLayout)

	// Free Text Widget ONLY AVAILABLE ON FREE LAYOUTS
	freeTextWidgetDefinition := datadog.NewFreeTextWidgetDefinitionWithDefaults()
	freeTextWidgetDefinition.SetText("Test me text")
	freeTextWidgetDefinition.SetColor("blue")
	freeTextWidgetDefinition.SetFontSize("16")
	freeTextWidgetDefinition.SetTextAlign(datadog.WIDGETTEXTALIGN_CENTER)

	freeTextWidget := datadog.NewWidgetWithDefaults()
	freeTextWidget.SetDefinition(datadog.FreeTextWidgetDefinitionAsWidgetDefinition(freeTextWidgetDefinition))
	freeTextWidget.SetLayout(*widgetLayout)

	// Group Widget
	groupNoteWidgetDefinition := datadog.NewNoteWidgetDefinitionWithDefaults()
	groupNoteWidgetDefinition.SetContent("Test Note Inside Group")

	groupWidgetDefinition := datadog.NewGroupWidgetDefinitionWithDefaults()
	groupWidgetDefinition.SetLayoutType(datadog.WIDGETLAYOUTTYPE_ORDERED)
	groupWidgetDefinition.SetTitle("Test Group Widget")
	groupWidgetDefinition.SetWidgets([]datadog.Widget{
		*datadog.NewWidget(datadog.NoteWidgetDefinitionAsWidgetDefinition(groupNoteWidgetDefinition)),
	})

	groupWidget := datadog.NewWidget(datadog.GroupWidgetDefinitionAsWidgetDefinition(groupWidgetDefinition))

	// HeatMap Widget
	heatMapWidgetDefinition := datadog.NewHeatMapWidgetDefinitionWithDefaults()

	heatMapWidgetRequest := datadog.NewHeatMapWidgetRequest()
	heatMapWidgetRequest.SetStyle(datadog.WidgetStyle{Palette: datadog.PtrString("dog_classic")})
	heatMapWidgetRequest.SetQ("avg:system.load.1{*}")

	widgetAxis := datadog.NewWidgetAxis()
	widgetAxis.SetIncludeZero(true)
	widgetAxis.SetMin("0")
	widgetAxis.SetMax("100")
	widgetAxis.SetScale("linear")

	heatMapWidgetDefinition.SetRequests([]datadog.HeatMapWidgetRequest{*heatMapWidgetRequest})
	heatMapWidgetDefinition.SetYaxis(*widgetAxis)
	heatMapWidgetDefinition.SetEvents([]datadog.WidgetEvent{
		{
			Q:             "Build succeeded",
			TagsExecution: datadog.PtrString("tags"),
		},
	})

	heatMapWidgetDefinition.SetTitle("Test Headmap Widget")
	heatMapWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	heatMapWidgetDefinition.SetTitleSize("16")
	heatMapWidgetDefinition.SetTime(*widgetTime)
	heatMapWidgetDefinition.SetShowLegend(true)
	heatMapWidgetDefinition.SetLegendSize(datadog.WIDGETLEGENDSIZE_FOUR)

	heatMapWidget := datadog.NewWidget(datadog.HeatMapWidgetDefinitionAsWidgetDefinition(heatMapWidgetDefinition))

	// HostMap Widget
	hostMapWidgetDefinition := datadog.NewHostMapWidgetDefinitionWithDefaults()
	hostMapWidgetDefinition.SetNodeType(datadog.WIDGETNODETYPE_CONTAINER)
	hostMapWidgetDefinition.SetRequests(datadog.HostMapWidgetDefinitionRequests{
		Fill: &datadog.HostMapRequest{Q: datadog.PtrString("avg:system.load.1{*}")},
		Size: &datadog.HostMapRequest{Q: datadog.PtrString("avg:system.load.1{*}")},
	})
	hostMapWidgetDefinition.SetNoMetricHosts(true)
	hostMapWidgetDefinition.SetNoGroupHosts(true)
	hostMapWidgetDefinition.SetGroup([]string{"env:prod"})
	hostMapWidgetDefinition.SetScope([]string{"foo"})
	hostMapWidgetDefinition.SetStyle(datadog.HostMapWidgetDefinitionStyle{
		Palette:     datadog.PtrString("dog_classic"),
		PaletteFlip: datadog.PtrBool(true),
		FillMin:     datadog.PtrString("0"),
		FillMax:     datadog.PtrString("100"),
	})
	hostMapWidgetDefinition.SetTitle("Test HostMap Widget")
	hostMapWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	hostMapWidgetDefinition.SetTitleSize("16")

	hostMapWidget := datadog.NewWidget(datadog.HostMapWidgetDefinitionAsWidgetDefinition(hostMapWidgetDefinition))

	// Iframe Widget ONLY AVAILABLE ON FREE LAYOUTS
	iFrameWidgetDefinition := datadog.NewIFrameWidgetDefinitionWithDefaults()
	iFrameWidgetDefinition.SetUrl("https://datadoghq.com")

	iFrameWidget := datadog.NewWidget(datadog.IFrameWidgetDefinitionAsWidgetDefinition(iFrameWidgetDefinition))
	iFrameWidget.SetLayout(*widgetLayout)

	// Image Widget ONLY AVAILABLE ON FREE LAYOUTS
	imageWidgetDefinition := datadog.NewImageWidgetDefinitionWithDefaults()
	imageWidgetDefinition.SetUrl("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4")
	imageWidgetDefinition.SetSizing(datadog.WIDGETIMAGESIZING_CENTER)
	imageWidgetDefinition.SetMargin(datadog.WIDGETMARGIN_LARGE)

	imageWidget := datadog.NewWidget(datadog.ImageWidgetDefinitionAsWidgetDefinition(imageWidgetDefinition))
	imageWidget.SetLayout(*widgetLayout)

	// LogStream ONLY AVAILABLE ON FREE LAYOUTS
	logStreamWidgetDefinition := datadog.NewLogStreamWidgetDefinitionWithDefaults()
	logStreamWidgetDefinition.SetIndexes([]string{"main"})
	logStreamWidgetDefinition.SetLogset("106")
	logStreamWidgetDefinition.SetQuery("Route XYZ failed")
	logStreamWidgetDefinition.SetColumns([]string{"Route"})
	logStreamWidgetDefinition.SetTitle("Test Logstream Widget")
	logStreamWidgetDefinition.SetTitleSize("16")
	logStreamWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_RIGHT)
	logStreamWidgetDefinition.SetTime(datadog.WidgetTime{LiveSpan: datadog.WIDGETLIVESPAN_PAST_TWO_DAYS.Ptr()})
	logStreamWidgetDefinition.SetMessageDisplay(datadog.WIDGETMESSAGEDISPLAY_EXPANDED_LARGE)
	logStreamWidgetDefinition.SetShowDateColumn(true)
	logStreamWidgetDefinition.SetShowMessageColumn(true)
	logStreamWidgetDefinition.SetSort(datadog.WidgetFieldSort{Column: "Route", Order: datadog.WIDGETSORT_ASCENDING})

	logStreamWidget := datadog.NewWidget(datadog.LogStreamWidgetDefinitionAsWidgetDefinition(logStreamWidgetDefinition))
	logStreamWidget.SetLayout(*widgetLayout)

	// Monitor Summary ONLY AVAILABLE ON FREE LAYOUTS
	monitorSummaryWidgetDefiniton := datadog.NewMonitorSummaryWidgetDefinitionWithDefaults()
	monitorSummaryWidgetDefiniton.SetQuery("Errors are increasing")
	monitorSummaryWidgetDefiniton.SetSummaryType(datadog.WIDGETSUMMARYTYPE_COMBINED)
	monitorSummaryWidgetDefiniton.SetSort(datadog.WIDGETSORT_ASCENDING)
	monitorSummaryWidgetDefiniton.SetDisplayFormat(datadog.WIDGETMONITORSUMMARYDISPLAYFORMAT_COUNTS)
	monitorSummaryWidgetDefiniton.SetColorPreference(datadog.WIDGETCOLORPREFERENCE_BACKGROUND)
	monitorSummaryWidgetDefiniton.SetHideZeroCounts(false)
	monitorSummaryWidgetDefiniton.SetShowLastTriggered(true)
	monitorSummaryWidgetDefiniton.SetTitle("Test Monitor Summary Widget")
	monitorSummaryWidgetDefiniton.SetTitleSize("16")
	monitorSummaryWidgetDefiniton.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	monitorSummaryWidgetDefiniton.SetStart(0)
	monitorSummaryWidgetDefiniton.SetCount(5)


	monitorSummaryWidget := datadog.NewWidget(datadog.MonitorSummaryWidgetDefinitionAsWidgetDefinition(monitorSummaryWidgetDefiniton))
	monitorSummaryWidget.SetLayout(*widgetLayout)

	// Note Widget
	noteWidgetDefinition := datadog.NewNoteWidgetDefinitionWithDefaults()
	noteWidgetDefinition.SetFontSize("13")
	noteWidgetDefinition.SetContent("Test Note Widget Example")
	noteWidgetDefinition.SetBackgroundColor("blue")
	noteWidgetDefinition.SetTextAlign(datadog.WIDGETTEXTALIGN_CENTER)
	noteWidgetDefinition.SetShowTick(true)
	noteWidgetDefinition.SetTickPos("4")
	noteWidgetDefinition.SetTickEdge(datadog.WIDGETTICKEDGE_BOTTOM)

	noteWidget := datadog.NewWidget(datadog.NoteWidgetDefinitionAsWidgetDefinition(noteWidgetDefinition))

	// Query Value Widget
	queryValueWidgetDefinition := datadog.NewQueryValueWidgetDefinitionWithDefaults()
	queryValueWidgetDefinition.SetRequests([]datadog.QueryValueWidgetRequest{{
		Q:          datadog.PtrString("avg:system.load.1{*}"),
		Aggregator: datadog.WIDGETAGGREGATOR_AVERAGE.Ptr(),
		ConditionalFormats: &[]datadog.WidgetConditionalFormat{{
			Comparator:    datadog.WIDGETCOMPARATOR_GREATER_THAN,
			Value:         7.,
			Palette:       datadog.WIDGETPALETTE_RED_ON_WHITE,
			CustomBgColor: datadog.PtrString("blue"),
			CustomFgColor: datadog.PtrString("black"),
			ImageUrl:      datadog.PtrString("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4"),
		}},
	}})
	queryValueWidgetDefinition.SetAutoscale(true)
	queryValueWidgetDefinition.SetCustomUnit("ns")
	queryValueWidgetDefinition.SetPrecision(2)
	queryValueWidgetDefinition.SetTextAlign(datadog.WIDGETTEXTALIGN_CENTER)
	queryValueWidgetDefinition.SetTitle("Test Query Value Widget")
	queryValueWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	queryValueWidgetDefinition.SetTitleSize("16")
	queryValueWidgetDefinition.SetTime(*widgetTime)

	queryValueWidget := datadog.NewWidget(datadog.QueryValueWidgetDefinitionAsWidgetDefinition(queryValueWidgetDefinition))

	// Scatter Plot Widget
	scatterPlotWidgetDefinition := datadog.NewScatterPlotWidgetDefinitionWithDefaults()
	scatterPlotWidgetDefinition.SetRequests(datadog.ScatterPlotWidgetDefinitionRequests{
		X: datadog.ScatterPlotRequest{
			Q:          datadog.PtrString("avg:system.load.1{*}"),
			Aggregator: datadog.WIDGETAGGREGATOR_AVERAGE.Ptr()},
		Y: datadog.ScatterPlotRequest{
			Q:          datadog.PtrString("avg:system.load.1{*}"),
			Aggregator: datadog.WIDGETAGGREGATOR_AVERAGE.Ptr(),
		},
	})
	scatterPlotWidgetDefinition.SetYaxis(*widgetAxis)
	scatterPlotWidgetDefinition.SetYaxis(*widgetAxis)
	scatterPlotWidgetDefinition.SetColorByGroups([]string{"env"})
	scatterPlotWidgetDefinition.SetTitle("Test ScatterPlot Widget")
	scatterPlotWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	scatterPlotWidgetDefinition.SetTitleSize("16")
	scatterPlotWidgetDefinition.SetTime(*widgetTime)

	scatterPlotWidget := datadog.NewWidget(datadog.ScatterPlotWidgetDefinitionAsWidgetDefinition(scatterPlotWidgetDefinition))

	// SLO Widget
	sloWidgetDefinition := datadog.NewSLOWidgetDefinitionWithDefaults()
	sloWidgetDefinition.SetViewType("detail")
	sloWidgetDefinition.SetTitle("Test SLO Widget")
	sloWidgetDefinition.SetTitleSize("16")
	sloWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	sloWidgetDefinition.SetSloId(slo.GetId())
	sloWidgetDefinition.SetShowErrorBudget(true)
	sloWidgetDefinition.SetViewMode(datadog.WIDGETVIEWMODE_BOTH)
	sloWidgetDefinition.SetTimeWindows([]datadog.WidgetTimeWindows{
		datadog.WIDGETTIMEWINDOWS_SEVEN_DAYS,
	})

	sloWidget := datadog.NewWidget(datadog.SLOWidgetDefinitionAsWidgetDefinition(sloWidgetDefinition))

	// Service Map Widget
	serviceMapWidgetDefinition := datadog.NewServiceMapWidgetDefinitionWithDefaults()
	serviceMapWidgetDefinition.SetFilters([]string{"*"})
	serviceMapWidgetDefinition.SetService("1234")
	serviceMapWidgetDefinition.SetTitle("Test Service Map Widget")
	serviceMapWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	serviceMapWidgetDefinition.SetTitleSize("16")

	serviceMapWidget := datadog.NewWidget(datadog.ServiceMapWidgetDefinitionAsWidgetDefinition(serviceMapWidgetDefinition))

	// Service Summary Widget
	serviceSummaryWidgetDefinition := datadog.NewServiceSummaryWidgetDefinitionWithDefaults()
	serviceSummaryWidgetDefinition.SetEnv("prod")
	serviceSummaryWidgetDefinition.SetService("web")
	serviceSummaryWidgetDefinition.SetSpanName("endpoint")
	serviceSummaryWidgetDefinition.SetShowHits(true)
	serviceSummaryWidgetDefinition.SetShowErrors(true)
	serviceSummaryWidgetDefinition.SetShowLatency(true)
	serviceSummaryWidgetDefinition.SetShowBreakdown(true)
	serviceSummaryWidgetDefinition.SetShowDistribution(true)
	serviceSummaryWidgetDefinition.SetShowResourceList(true)
	serviceSummaryWidgetDefinition.SetSizeFormat(datadog.WIDGETSIZEFORMAT_LARGE)
	serviceSummaryWidgetDefinition.SetDisplayFormat(datadog.WIDGETSERVICESUMMARYDISPLAYFORMAT_TWO_COLUMN)
	serviceSummaryWidgetDefinition.SetTitle("Test Service Summary Widget")
	serviceSummaryWidgetDefinition.SetTitleSize("16")
	serviceSummaryWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	serviceSummaryWidgetDefinition.SetTime(*widgetTimePastOneHour)

	serviceSummaryWidget := datadog.NewWidget(datadog.ServiceSummaryWidgetDefinitionAsWidgetDefinition(serviceSummaryWidgetDefinition))
	serviceSummaryWidget.SetLayout(*widgetLayout)

	// Table Widget
	tableWidgetDefinition := datadog.NewTableWidgetDefinitionWithDefaults()
	tableWidgetDefinition.SetRequests([]datadog.TableWidgetRequest{{
		Q:          datadog.PtrString("avg:system.load.1{*}"),
		Alias:      datadog.PtrString("System Load"),
		Aggregator: datadog.WIDGETAGGREGATOR_AVERAGE.Ptr(),
		Limit:      datadog.PtrInt64(50),
		Order:      datadog.WIDGETSORT_ASCENDING.Ptr(),
		ConditionalFormats: &[]datadog.WidgetConditionalFormat{{
			Comparator:    datadog.WIDGETCOMPARATOR_GREATER_THAN,
			Value:         7.,
			Palette:       datadog.WIDGETPALETTE_RED_ON_WHITE,
			CustomBgColor: datadog.PtrString("blue"),
			CustomFgColor: datadog.PtrString("black"),
			ImageUrl:      datadog.PtrString("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4"),
		}},
	}})
	tableWidgetDefinition.SetTitle("Test Table Widget")
	tableWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	tableWidgetDefinition.SetTitleSize("16")
	tableWidgetDefinition.SetTime(*widgetTime)

	tableWidget := datadog.NewWidget(datadog.TableWidgetDefinitionAsWidgetDefinition(tableWidgetDefinition))

	// Timeseries Widget
	timeseriesWidgetDefinition := datadog.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinition.SetRequests([]datadog.TimeseriesWidgetRequest{{
		Q: datadog.PtrString("avg:system.load.1{*}"),
		Style: &datadog.WidgetRequestStyle{
			Palette:   datadog.PtrString("dog_classic"),
			LineType:  datadog.WIDGETLINETYPE_DASHED.Ptr(),
			LineWidth: datadog.WIDGETLINEWIDTH_THICK.Ptr(),
		},
		Metadata: &[]datadog.TimeseriesWidgetRequestMetadata{{
			Expression: "avg:system.load.1{*}",
			AliasName:  datadog.PtrString("Aliased metric"),
		}},
		DisplayType: datadog.WIDGETDISPLAYTYPE_LINE.Ptr(),
	}})
	timeseriesWidgetDefinition.SetYaxis(datadog.WidgetAxis{
		IncludeZero: datadog.PtrBool(true),
		Min:         datadog.PtrString("0"),
		Max:         datadog.PtrString("100"),
		Scale:       datadog.PtrString("linear"),
		Label:       datadog.PtrString("Widget Scale")})
	timeseriesWidgetDefinition.SetEvents([]datadog.WidgetEvent{{
		Q: "Build succeeded",
	}})
	timeseriesWidgetDefinition.SetMarkers([]datadog.WidgetMarker{{
		Value:       "y=15",
		DisplayType: datadog.PtrString("error dashed"),
		Label:       datadog.PtrString("error threshold"),
		Time:        datadog.PtrString("4h"),
	}})
	timeseriesWidgetDefinition.SetTitle("Test Timeseries Widget")
	timeseriesWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinition.SetTitleSize("16")
	timeseriesWidgetDefinition.SetTime(*widgetTime)
	timeseriesWidgetDefinition.SetShowLegend(true)
	timeseriesWidgetDefinition.SetLegendSize(datadog.WIDGETLEGENDSIZE_SIXTEEN)

	timeseriesWidget := datadog.NewWidget(datadog.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinition))

	// Timeseries Widget with Process query
	timeseriesWidgetDefinitionProcessQuery := datadog.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinitionProcessQuery.SetRequests([]datadog.TimeseriesWidgetRequest{{
		ProcessQuery: &datadog.ProcessQueryDefinition{
			Metric:   "process.stat.cpu.total_pct",
			FilterBy: &[]string{"account:test"},
			Limit:    datadog.PtrInt64(10),
			SearchBy: datadog.PtrString("editor"),
		},
		Style: &datadog.WidgetRequestStyle{
			Palette:   datadog.PtrString("dog_classic"),
			LineType:  datadog.WIDGETLINETYPE_DASHED.Ptr(),
			LineWidth: datadog.WIDGETLINEWIDTH_THICK.Ptr(),
		},
		Metadata: &[]datadog.TimeseriesWidgetRequestMetadata{{
			Expression: "avg:system.load.1{*}",
			AliasName:  datadog.PtrString("Aliased metric"),
		}},
		DisplayType: datadog.WIDGETDISPLAYTYPE_LINE.Ptr(),
	}})
	timeseriesWidgetDefinitionProcessQuery.SetYaxis(datadog.WidgetAxis{
		IncludeZero: datadog.PtrBool(true),
		Min:         datadog.PtrString("0"),
		Max:         datadog.PtrString("100"),
		Scale:       datadog.PtrString("linear")})
	timeseriesWidgetDefinitionProcessQuery.SetEvents([]datadog.WidgetEvent{{
		Q: "Build succeeded",
	}})
	timeseriesWidgetDefinitionProcessQuery.SetMarkers([]datadog.WidgetMarker{{
		Value:       "y=15",
		DisplayType: datadog.PtrString("error dashed"),
		Label:       datadog.PtrString("error threshold"),
		Time:        datadog.PtrString("4h"),
	}})
	timeseriesWidgetDefinitionProcessQuery.SetTitle("Test Timeseries Widget with Process Query")
	timeseriesWidgetDefinitionProcessQuery.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinitionProcessQuery.SetTitleSize("16")
	timeseriesWidgetDefinitionProcessQuery.SetTime(*widgetTime)
	timeseriesWidgetDefinitionProcessQuery.SetShowLegend(true)
	timeseriesWidgetDefinitionProcessQuery.SetLegendSize(datadog.WIDGETLEGENDSIZE_SIXTEEN)

	timeseriesWidgetProcessQuery := datadog.NewWidget(datadog.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinitionProcessQuery))

	// Timeseries Widget with Log query (APM/Log/Network/Rum share schemas, so only test one)
	timeseriesWidgetDefinitionLogQuery := datadog.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinitionLogQuery.SetRequests([]datadog.TimeseriesWidgetRequest{{
		LogQuery: &datadog.LogQueryDefinition{
			Index: datadog.PtrString("main"),
			Compute: &datadog.LogsQueryCompute{
				Aggregation: "count",
				Facet:       datadog.PtrString("host"),
				Interval:    datadog.PtrInt64(10),
			},
			Search: &datadog.LogQueryDefinitionSearch{Query: "Error parsing"},
			GroupBy: &[]datadog.LogQueryDefinitionGroupBy{{
				Facet: "host",
				Limit: datadog.PtrInt64(5),
				Sort: &datadog.LogQueryDefinitionSort{
					Aggregation: "count",
					Order:       datadog.WIDGETSORT_ASCENDING,
				},
			}},
		},
		Style: &datadog.WidgetRequestStyle{
			Palette:   datadog.PtrString("dog_classic"),
			LineType:  datadog.WIDGETLINETYPE_DASHED.Ptr(),
			LineWidth: datadog.WIDGETLINEWIDTH_THICK.Ptr(),
		},
		Metadata: &[]datadog.TimeseriesWidgetRequestMetadata{{
			Expression: "avg:system.load.1{*}",
			AliasName:  datadog.PtrString("Aliased metric"),
		}},
		DisplayType: datadog.WIDGETDISPLAYTYPE_LINE.Ptr(),
	}})
	timeseriesWidgetDefinitionLogQuery.SetYaxis(*widgetAxis)
	timeseriesWidgetDefinitionLogQuery.SetEvents([]datadog.WidgetEvent{{
		Q: "Build succeeded",
	}})
	timeseriesWidgetDefinitionLogQuery.SetMarkers([]datadog.WidgetMarker{{
		Value:       "y=15",
		DisplayType: datadog.PtrString("error dashed"),
		Label:       datadog.PtrString("error threshold"),
		Time:        datadog.PtrString("4h"),
	}})
	timeseriesWidgetDefinitionLogQuery.SetTitle("Test Timeseries Widget with Log Query")
	timeseriesWidgetDefinitionLogQuery.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinitionLogQuery.SetTitleSize("16")
	timeseriesWidgetDefinitionLogQuery.SetTime(*widgetTime)
	timeseriesWidgetDefinitionLogQuery.SetShowLegend(true)
	timeseriesWidgetDefinitionLogQuery.SetLegendSize(datadog.WIDGETLEGENDSIZE_SIXTEEN)

	timeseriesWidgetLogQuery := datadog.NewWidget(datadog.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinitionLogQuery))

	// Timeseries Widget with Event query
	timeseriesWidgetDefinitionEventQuery := datadog.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinitionEventQuery.SetRequests([]datadog.TimeseriesWidgetRequest{{
		EventQuery: &datadog.EventQueryDefinition{
			Search:        "Build failure",
			TagsExecution: "build",
		},
		Style: &datadog.WidgetRequestStyle{
			Palette:   datadog.PtrString("dog_classic"),
			LineType:  datadog.WIDGETLINETYPE_DASHED.Ptr(),
			LineWidth: datadog.WIDGETLINEWIDTH_THICK.Ptr()},
		Metadata: &[]datadog.TimeseriesWidgetRequestMetadata{{
			Expression: "avg:system.load.1{*}",
			AliasName:  datadog.PtrString("Aliased metric"),
		}},
		DisplayType: datadog.WIDGETDISPLAYTYPE_LINE.Ptr(),
	}})
	timeseriesWidgetDefinitionEventQuery.SetYaxis(*widgetAxis)
	timeseriesWidgetDefinitionEventQuery.SetEvents([]datadog.WidgetEvent{{
		Q: "Build succeeded",
	}})
	timeseriesWidgetDefinitionEventQuery.SetMarkers([]datadog.WidgetMarker{{
		Value:       "y=15",
		DisplayType: datadog.PtrString("error dashed"),
		Label:       datadog.PtrString("error threshold"),
		Time:        datadog.PtrString("4h"),
	}})
	timeseriesWidgetDefinitionEventQuery.SetTitle("Test Timeseries Widget with Event Query")
	timeseriesWidgetDefinitionEventQuery.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	timeseriesWidgetDefinitionEventQuery.SetTitleSize("16")
	timeseriesWidgetDefinitionEventQuery.SetTime(*widgetTime)
	timeseriesWidgetDefinitionEventQuery.SetShowLegend(true)
	timeseriesWidgetDefinitionEventQuery.SetLegendSize(datadog.WIDGETLEGENDSIZE_SIXTEEN)

	timeseriesWidgetEventQuery := datadog.NewWidget(datadog.TimeseriesWidgetDefinitionAsWidgetDefinition(timeseriesWidgetDefinitionEventQuery))

	// Toplist Widget
	toplistWidgetDefinition := datadog.NewToplistWidgetDefinitionWithDefaults()
	toplistWidgetDefinition.SetRequests([]datadog.ToplistWidgetRequest{{
		Q: datadog.PtrString("avg:system.load.1{*}"),
		ConditionalFormats: &[]datadog.WidgetConditionalFormat{{
			Comparator:    datadog.WIDGETCOMPARATOR_GREATER_THAN,
			Value:         7.,
			Palette:       datadog.WIDGETPALETTE_RED_ON_WHITE,
			CustomBgColor: datadog.PtrString("blue"),
			CustomFgColor: datadog.PtrString("black"),
			ImageUrl:      datadog.PtrString("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4"),
		}},
	}})
	toplistWidgetDefinition.SetTitle("Test Toplist Widget")
	toplistWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	toplistWidgetDefinition.SetTitleSize("16")
	toplistWidgetDefinition.SetTime(*widgetTime)

	toplistWidget := datadog.NewWidget(datadog.ToplistWidgetDefinitionAsWidgetDefinition(toplistWidgetDefinition))

	// Template Variables
	templateVariable := datadog.NewDashboardTemplateVariablesWithDefaults()
	templateVariable.SetName("test template var")
	templateVariable.SetPrefix("test-go")
	templateVariable.SetDefault("*")

	// Template Variable Presets
	dashboardTemplateVariablePreset := datadog.NewDashboardTemplateVariablePreset()
	dashboardTemplateVariablePreset.SetName("Test Preset")
	dashboardTemplateVariablePresetValue := datadog.NewDashboardTemplateVariablePresetValueWithDefaults()
	dashboardTemplateVariablePresetValue.SetName("test preset")
	dashboardTemplateVariablePresetValue.SetValue("*")
	dashboardTemplateVariablePreset.SetTemplateVariables([]datadog.DashboardTemplateVariablePresetValue{*dashboardTemplateVariablePresetValue})

	orderedWidgetList := []datadog.Widget{
		*alertGraphWidget,
		*alertValueWidget,
		*changeWidget,
		*checkStatusWidget,
		*distributionWidget,
		*groupWidget,
		*heatMapWidget,
		*hostMapWidget,
		*noteWidget,
		*queryValueWidget,
		*scatterPlotWidget,
		*sloWidget,
		*serviceMapWidget,
		*tableWidget,
		*timeseriesWidget,
		*timeseriesWidgetProcessQuery,
		*timeseriesWidgetLogQuery,
		*timeseriesWidgetEventQuery,
		*toplistWidget,
	}

	dashboard := datadog.NewDashboardWithDefaults()
	dashboard.SetLayoutType(datadog.DASHBOARDLAYOUTTYPE_ORDERED)
	dashboard.SetWidgets(orderedWidgetList)
	dashboard.SetTitle(fmt.Sprintf("%s-ordered", *tests.UniqueEntityName(ctx, t)))
	dashboard.SetDescription("Test dashboard for Go client")
	dashboard.SetIsReadOnly(false)
	dashboard.SetTemplateVariables([]datadog.DashboardTemplateVariables{*templateVariable})
	dashboard.SetTemplateVariablePresets([]datadog.DashboardTemplateVariablePreset{*dashboardTemplateVariablePreset})
	dashboard.SetNotifyList([]string{"test@datadoghq.com"})

	createdDashboard, httpresp, err := Client(ctx).DashboardsApi.CreateDashboard(ctx).Body(*dashboard).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboard(ctx, createdDashboard.GetId())
	assert.Equal(200, httpresp.StatusCode)

	getDashboard, httpresp, err := Client(ctx).DashboardsApi.GetDashboard(ctx, createdDashboard.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	assert.Equal(getDashboard, createdDashboard)

	assertDashboardDefinition := func(dashboard *datadog.Dashboard, response datadog.Dashboard) {
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
		assert.Equal(len(dashboard.GetNotifyList()), len(response.GetNotifyList()))

		// Template Variables
		assert.Equal(dashboard.GetTemplateVariables()[0], response.GetTemplateVariables()[0])

		for index, checkWidget := range response.GetWidgets() {
			assert.True(checkWidget.GetId() > 0)
			checkWidget.Id = nil
			checkWidgetDefinition := checkWidget.GetDefinition()
			checkWidgetInstance := checkWidgetDefinition.GetActualInstance()
			if def, ok := checkWidgetInstance.(*datadog.GroupWidgetDefinition); ok {
				for index := range def.GetWidgets() {
					def.Widgets[index].Id = nil
				}
			}
			assert.Equal(dashboard.GetWidgets()[index], checkWidget)
		}

	}

	assertDashboardDefinition(dashboard, createdDashboard)

	freeWidgetList := []datadog.Widget{
		*eventStreamWidget,
		*eventTimelineWidget,
		*freeTextWidget,
		*iFrameWidget,
		*imageWidget,
		*logStreamWidget,
		*monitorSummaryWidget,
		*serviceSummaryWidget,
	}

	freeDashboard := datadog.NewDashboardWithDefaults()
	freeDashboard.SetLayoutType(datadog.DASHBOARDLAYOUTTYPE_FREE)
	freeDashboard.SetWidgets(freeWidgetList)
	freeDashboard.SetTitle(fmt.Sprintf("%s-free", *tests.UniqueEntityName(ctx, t)))
	freeDashboard.SetDescription("Test Free layout dashboard for Go client")
	freeDashboard.SetIsReadOnly(false)
	freeDashboard.SetTemplateVariables([]datadog.DashboardTemplateVariables{*templateVariable})

	createdFreeDashboard, httpresp, err := Client(ctx).DashboardsApi.CreateDashboard(ctx).Body(*freeDashboard).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboard(ctx, createdFreeDashboard.GetId())
	assert.Equal(200, httpresp.StatusCode)

	getFreeDashboard, httpresp, err := Client(ctx).DashboardsApi.GetDashboard(ctx, createdFreeDashboard.GetId()).Execute()
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
	noteWidget.SetDefinition(datadog.NoteWidgetDefinitionAsWidgetDefinition(noteWidgetDefinition))

	noteWidgetIndex := 8
	dashboardWidgets[noteWidgetIndex] = *noteWidget
	dashboard.SetWidgets(dashboardWidgets)

	updateResponse, httpresp, err := Client(ctx).DashboardsApi.UpdateDashboard(ctx, createdDashboard.GetId()).Body(*dashboard).Execute()
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
		if def, ok := noteWidgetResponseInstance.(*datadog.NoteWidgetDefinition); ok {
			foundWidget = true
			assert.Equal("Updated content", def.GetContent())
			assert.Equal("30", def.GetFontSize())
			assert.Equal(noteWidgetIndex, index)
			break
		}
	}
	assert.True(foundWidget)
	assert.True(len(updateResponse.GetWidgets()) > 1)

	deleteResponse, httpresp, err := Client(ctx).DashboardsApi.DeleteDashboard(ctx, createdFreeDashboard.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error deleting dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Equal(createdFreeDashboard.GetId(), deleteResponse.GetDeletedDashboardId())

}

func TestDashboardGetAll(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	getAllResponse, httpresp, err := Client(ctx).DashboardsApi.ListDashboards(ctx).Execute()
	if err != nil {
		t.Fatalf("Error getting all dashboards: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(len(getAllResponse.GetDashboards()) >= 1)
}

func TestDashboardCreateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.Dashboard
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.Dashboard{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.Dashboard{}, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).DashboardsApi.CreateDashboard(ctx).Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardListErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).DashboardsApi.ListDashboards(ctx).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardDeleteErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).DashboardsApi.DeleteDashboard(ctx, "random").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardUpdateErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

	dashboardOK := *datadog.NewDashboardWithDefaults()
	dashboardOK.SetWidgets([]datadog.Widget{})
	dashboardOK.SetLayoutType(datadog.DASHBOARDLAYOUTTYPE_FREE)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		Body               datadog.Dashboard
		ExpectedStatusCode int
	}{
		"400 Bad Request": {WithTestAuth, datadog.Dashboard{}, 400},
		"403 Forbidden":   {WithFakeAuth, datadog.Dashboard{}, 403},
		"404 Not Found":   {WithTestAuth, dashboardOK, 404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).DashboardsApi.UpdateDashboard(ctx, "random").Body(tc.Body).Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDashboardGetErrors(t *testing.T) {
	ctx, close := tests.WithTestSpan(context.Background(), t)
	defer close()

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

			_, httpresp, err := Client(ctx).DashboardsApi.GetDashboard(ctx, "random").Execute()
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func deleteDashboard(ctx context.Context, dashboardID string) {
	_, httpresp, err := Client(ctx).DashboardsApi.DeleteDashboard(ctx, dashboardID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Error deleting Dashboard: %v, Another test may have already deleted this dashboard.", dashboardID)
	}
}
