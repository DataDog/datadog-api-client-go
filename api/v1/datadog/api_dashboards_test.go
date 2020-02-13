package datadog_test

import (
	"log"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestDashboardLifecycle(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

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
	alertGraphWidget.SetDefinition(alertGraphDefinition.AsWidgetDefinition())

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
	alertValueWidget.SetDefinition(alertValueDefinition.AsWidgetDefinition())

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
	changeWidget.SetDefinition(changeWidgetDefinition.AsWidgetDefinition())

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
	checkStatusWidget.SetDefinition(checkStatusWidgetDefinition.AsWidgetDefinition())

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

	distributionWidget := datadog.NewWidget(distributionWidgetDefinition.AsWidgetDefinition())

	// Event Stream Widget ONLY AVAILABLE ON FREE LAYOUTS
	eventStreamWidgetDefinition := datadog.NewEventStreamWidgetDefinitionWithDefaults()
	eventStreamWidgetDefinition.SetQuery("Build successful")
	eventStreamWidgetDefinition.SetEventSize(datadog.WIDGETEVENTSIZE_LARGE)
	eventStreamWidgetDefinition.SetTitle("Test Event Stream Widget")
	eventStreamWidgetDefinition.SetTitleSize("16")
	eventStreamWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	eventStreamWidgetDefinition.SetTime(*widgetTimePastOneDay)

	eventStreamWidget := datadog.NewWidgetWithDefaults()
	eventStreamWidget.SetDefinition(eventStreamWidgetDefinition.AsWidgetDefinition())
	eventStreamWidget.SetLayout(*widgetLayout)

	// Event Timeline Widget ONLY AVAILABLE ON FREE LAYOUTS
	eventTimelineWidgetDefinition := datadog.NewEventTimelineWidgetDefinitionWithDefaults()
	eventTimelineWidgetDefinition.SetQuery("Build Failed")
	eventTimelineWidgetDefinition.SetTitle("Test Event Timeline Widget")
	eventTimelineWidgetDefinition.SetTitleSize("16")
	eventTimelineWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_LEFT)
	eventTimelineWidgetDefinition.SetTime(*widgetTimePastOneMonth)

	eventTimelineWidget := datadog.NewWidgetWithDefaults()
	eventTimelineWidget.SetDefinition(eventTimelineWidgetDefinition.AsWidgetDefinition())
	eventTimelineWidget.SetLayout(*widgetLayout)

	// Free Text Widget ONLY AVAILABLE ON FREE LAYOUTS
	freeTextWidgetDefinition := datadog.NewFreeTextWidgetDefinitionWithDefaults()
	freeTextWidgetDefinition.SetText("Test me text")
	freeTextWidgetDefinition.SetColor("blue")
	freeTextWidgetDefinition.SetFontSize("16")
	freeTextWidgetDefinition.SetTextAlign(datadog.WIDGETTEXTALIGN_CENTER)

	freeTextWidget := datadog.NewWidgetWithDefaults()
	freeTextWidget.SetDefinition(freeTextWidgetDefinition.AsWidgetDefinition())
	freeTextWidget.SetLayout(*widgetLayout)

	// Group Widget
	groupNoteWidgetDefinition := datadog.NewNoteWidgetDefinitionWithDefaults()
	groupNoteWidgetDefinition.SetContent("Test Note Inside Group")

	groupWidgetDefinition := datadog.NewGroupWidgetDefinitionWithDefaults()
	groupWidgetDefinition.SetLayoutType(datadog.WIDGETLAYOUTTYPE_ORDERED)
	groupWidgetDefinition.SetTitle("Test Group Widget")
	groupWidgetDefinition.SetWidgets([]datadog.Widget{
		*datadog.NewWidget(groupNoteWidgetDefinition.AsWidgetDefinition()),
	})

	groupWidget := datadog.NewWidget(groupWidgetDefinition.AsWidgetDefinition())

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

	heatMapWidget := datadog.NewWidget(heatMapWidgetDefinition.AsWidgetDefinition())

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

	hostMapWidget := datadog.NewWidget(hostMapWidgetDefinition.AsWidgetDefinition())

	// Iframe Widget ONLY AVAILABLE ON FREE LAYOUTS
	iFrameWidgetDefinition := datadog.NewIFrameWidgetDefinitionWithDefaults()
	iFrameWidgetDefinition.SetUrl("https://datadoghq.com")

	iFrameWidget := datadog.NewWidget(iFrameWidgetDefinition.AsWidgetDefinition())
	iFrameWidget.SetLayout(*widgetLayout)

	// Image Widget ONLY AVAILABLE ON FREE LAYOUTS
	imageWidgetDefinition := datadog.NewImageWidgetDefinitionWithDefaults()
	imageWidgetDefinition.SetUrl("https://docs.datadoghq.com/images/dashboards/widgets/image/image.mp4")
	imageWidgetDefinition.SetSizing(datadog.WIDGETIMAGESIZING_CENTER)
	imageWidgetDefinition.SetMargin(datadog.WIDGETMARGIN_LARGE)

	imageWidget := datadog.NewWidget(imageWidgetDefinition.AsWidgetDefinition())
	imageWidget.SetLayout(*widgetLayout)

	// LogStream ONLY AVAILABLE ON FREE LAYOUTS
	logStreamWidgetDefinition := datadog.NewLogStreamWidgetDefinitionWithDefaults()
	logStreamWidgetDefinition.SetIndexes([]string{"main"})
	logStreamWidgetDefinition.SetQuery("Route XYZ failed")
	logStreamWidgetDefinition.SetColumns([]string{"Route"})
	logStreamWidgetDefinition.SetTitle("Test Logstream Widget")
	logStreamWidgetDefinition.SetTitleSize("16")
	logStreamWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_RIGHT)
	logStreamWidgetDefinition.SetTime(datadog.WidgetTime{LiveSpan: datadog.WIDGETLIVESPAN_PAST_TWO_DAYS.Ptr()})

	logStreamWidget := datadog.NewWidget(logStreamWidgetDefinition.AsWidgetDefinition())
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

	monitorSummaryWidget := datadog.NewWidget(monitorSummaryWidgetDefiniton.AsWidgetDefinition())
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

	noteWidget := datadog.NewWidget(noteWidgetDefinition.AsWidgetDefinition())

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

	queryValueWidget := datadog.NewWidget(queryValueWidgetDefinition.AsWidgetDefinition())

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

	scatterPlotWidget := datadog.NewWidget(scatterPlotWidgetDefinition.AsWidgetDefinition())

	// SLO Widget
	sloWidgetDefinition := datadog.NewSloWidgetDefinitionWithDefaults()
	sloWidgetDefinition.SetViewType("detail")
	sloWidgetDefinition.SetTitle("Test SLO Widget")
	sloWidgetDefinition.SetTitleSize("16")
	sloWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	sloWidgetDefinition.SetSloId("1234L")
	sloWidgetDefinition.SetShowErrorBudget(true)
	sloWidgetDefinition.SetViewMode(datadog.WIDGETVIEWMODE_BOTH)
	sloWidgetDefinition.SetTimeWindows([]datadog.WidgetTimeWindows{
		datadog.WIDGETTIMEWINDOWS_SEVEN_DAYS,
	})

	sloWidget := datadog.NewWidget(sloWidgetDefinition.AsWidgetDefinition())

	// Service Map Widget
	serviceMapWidgetDefinition := datadog.NewServiceMapWidgetDefinitionWithDefaults()
	serviceMapWidgetDefinition.SetFilters([]string{"*"})
	serviceMapWidgetDefinition.SetService("1234")
	serviceMapWidgetDefinition.SetTitle("Test Service Map Widget")
	serviceMapWidgetDefinition.SetTitleAlign(datadog.WIDGETTEXTALIGN_CENTER)
	serviceMapWidgetDefinition.SetTitleSize("16")

	serviceMapWidget := datadog.NewWidget(serviceMapWidgetDefinition.AsWidgetDefinition())

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

	serviceSummaryWidget := datadog.NewWidget(serviceSummaryWidgetDefinition.AsWidgetDefinition())
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

	tableWidget := datadog.NewWidget(tableWidgetDefinition.AsWidgetDefinition())

	// Timeseries Widget
	timeseriesWidgetDefinition := datadog.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinition.SetRequests([]datadog.TimeseriesWidgetRequest{{
		Q: datadog.PtrString("avg:system.load.1{*}"),
		Style: &datadog.TimeseriesWidgetRequestStyle{
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

	timeseriesWidget := datadog.NewWidget(timeseriesWidgetDefinition.AsWidgetDefinition())

	// Timeseries Widget with Process query
	timeseriesWidgetDefinitionProcessQuery := datadog.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinitionProcessQuery.SetRequests([]datadog.TimeseriesWidgetRequest{{
		ProcessQuery: &datadog.ProcessQueryDefinition{
			Metric:   "process.stat.cpu.total_pct",
			FilterBy: &[]string{"account:test"},
			Limit:    datadog.PtrInt64(10),
			SearchBy: datadog.PtrString("editor"),
		},
		Style: &datadog.TimeseriesWidgetRequestStyle{
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

	timeseriesWidgetProcessQuery := datadog.NewWidget(timeseriesWidgetDefinitionProcessQuery.AsWidgetDefinition())

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
		Style: &datadog.TimeseriesWidgetRequestStyle{
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

	timeseriesWidgetLogQuery := datadog.NewWidget(timeseriesWidgetDefinitionLogQuery.AsWidgetDefinition())

	// Timeseries Widget with Event query
	timeseriesWidgetDefinitionEventQuery := datadog.NewTimeseriesWidgetDefinitionWithDefaults()
	timeseriesWidgetDefinitionEventQuery.SetRequests([]datadog.TimeseriesWidgetRequest{{
		EventQuery: &datadog.EventQueryDefinition{
			Search:        "Build failure",
			TagsExecution: "build",
		},
		Style: &datadog.TimeseriesWidgetRequestStyle{
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

	timeseriesWidgetEventQuery := datadog.NewWidget(timeseriesWidgetDefinitionEventQuery.AsWidgetDefinition())

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

	toplistWidget := datadog.NewWidget(toplistWidgetDefinition.AsWidgetDefinition())

	// Template Variables
	templateVariable := datadog.NewDashboardTemplateVariablesWithDefaults()
	templateVariable.SetName("test template var")
	templateVariable.SetPrefix("test-java")
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
	dashboard.SetTitle("Go Client Test ORDERED Dashboard")
	dashboard.SetDescription("Test dashboard for Go client")
	dashboard.SetIsReadOnly(false)
	dashboard.SetTemplateVariables([]datadog.DashboardTemplateVariables{*templateVariable})
	dashboard.SetTemplateVariablePresets([]datadog.DashboardTemplateVariablePreset{*dashboardTemplateVariablePreset})

	createdDashboard, httpresp, err := TESTAPICLIENT.DashboardsApi.CreateDashboard(TESTAUTH).Body(*dashboard).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboard(createdDashboard.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	getDashboard, httpresp, err := TESTAPICLIENT.DashboardsApi.GetDashboard(TESTAUTH, createdDashboard.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.DeepEqual(t, getDashboard, createdDashboard)

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
	freeDashboard.SetTitle("Go Client Test Free Dashboard")
	freeDashboard.SetDescription("Test Free layout dashboard for Go client")
	freeDashboard.SetIsReadOnly(false)
	freeDashboard.SetTemplateVariables([]datadog.DashboardTemplateVariables{*templateVariable})

	createdFreeDashboard, httpresp, err := TESTAPICLIENT.DashboardsApi.CreateDashboard(TESTAUTH).Body(*freeDashboard).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	defer deleteDashboard(createdFreeDashboard.GetId())
	assert.Equal(t, 200, httpresp.StatusCode)

	getFreeDashboard, httpresp, err := TESTAPICLIENT.DashboardsApi.GetDashboard(TESTAUTH, createdFreeDashboard.GetId()).Execute()
	if err != nil {
		t.Fatalf("Error creating dashboard: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(t, 200, httpresp.StatusCode)
	assert.DeepEqual(t, getFreeDashboard, createdFreeDashboard)
	/*
		assertTrue(response.getUrl(),contains(fmt.Sprintf("/dashboard/%s", response.getId())))
		assertNotNull(response.getCreatedAt())
		assertNotNull(response.getModifiedAt())
		assertNotNull(response.getAuthorHandle())
		assert.Equal(t, dashboard.getLayoutType(), response.getLayoutType())
		assert.Equal(t, dashboard.getNotifyList(), response.getNotifyList())
		assert.Equal(t, templateVariables, response.getTemplateVariables())
		for _, checkWidget := range response.getWidgets() {
			assertNotNull(checkWidget.getId())
			checkWidget.setId(nil)
			if checkWidget.getDefinition(),getType(),equals("group") {
				for _, subWidget := range checkWidget.getDefinition(),(*GroupWidgetDefinition),getWidgets() {
					subWidget.id(nil)
				}
			}
		}
		for _, checkWidget := range getFreeResponse.getWidgets() {
			assertNotNull(checkWidget.getId())
			checkWidget.id(nil)
		}
		dashboard.description("Updated dashboard description"),addWidgetsItem(noteWidget.definition(noteDefinition.SetContent("Updated content"),fontSize("30")))
		updateResponse := api.updateDashboard(response.getId()),body(dashboard),execute()
		assert.Equal(t, dashboard.getDescription(), updateResponse.getDescription())
		assert.Equal(t, dashboard.getTitle(), updateResponse.getTitle())
		assert.Equal(t, dashboard.getWidgets(),get(0), updateResponse.getWidgets(),get(0),id(nil))
		foundWidget := false
		for _, noteWidgetResponse := range updateResponse.getWidgets() {
			if noteWidgetResponse.getDefinition(),getType(),equals("note") {
				def, ok := noteWidgetResponse.getDefinition(),(*NoteWidgetDefinition)
				if !ok {
					panic("XXX Cast fail for *parser.GoCastType")
				}
				foundWidget = true
				assert.Equal(t, "Updated content", def.getContent())
				assert.Equal(t, "30", def.getFontSize())
				break
			}
		}
		assertTrue(foundWidget)
		assertTrue(updateResponse.getWidgets(),
			Size: datadog.PtrString() > 1)
		deleteResponse := api.deleteDashboard(response.getId()),execute()
		assert.Equal(t, deleteResponse.getDeletedDashboardId(), response.getId())
	*/
}

func TestGetAllDashboardTest(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	/*
		getAllResponse, httpresp, err := TESTAPICLIENT.DashboardsApi.GetAllDashboards(TESTAUTH),Execute()

		assertNotNull(getAllResponse.getDashboards(),get(0),getAuthorHandle())
		assertNotNull(getAllResponse.getDashboards(),get(0),getCreatedAt())
		assertNotNull(getAllResponse.getDashboards(),get(0),getModifiedAt())
		assertNotNull(getAllResponse.getDashboards(),get(0),getId())
		assertNotNull(getAllResponse.getDashboards(),get(0),getIsReadOnly())
		assertNotNull(getAllResponse.getDashboards(),get(0),getLayoutType())
		assertNotNull(getAllResponse.getDashboards(),get(0),getTitle())
		assertNotNull(getAllResponse.getDashboards(),get(0),getUrl())
	*/
}

func deleteDashboard(dashboardID string) {
	_, httpresp, err := TESTAPICLIENT.DashboardsApi.DeleteDashboard(TESTAUTH, dashboardID).Execute()
	if err != nil && httpresp.StatusCode != 404 {
		log.Printf("Error deleting Dashboard: %v, Another test may have already deleted this dashboard.", dashboardID)
	}
}
