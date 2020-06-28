# WidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | **string** | ID of the alert to use in the widget. | 
**Time** | [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | **string** | Title of your widget. | [optional] 
**TitleAlign** | [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | **string** | Size of the title. | [optional] 
**Type** | [**ToplistWidgetDefinitionType**](ToplistWidgetDefinitionType.md) |  | 
**VizType** | [**WidgetVizType**](WidgetVizType.md) |  | 
**Precision** | **int64** | Number of decimals to show. If not defined, the widget uses the raw value. | [optional] 
**TextAlign** | [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**Unit** | **string** | Unit to display with the value. | [optional] 
**Requests** | [**[]ToplistWidgetRequest**](ToplistWidgetRequest.md) | List of top list widget requests. | 
**Check** | **string** | Name of the check to use in the widget. | 
**Group** | **[]string** | List of tag prefixes to group by. | [optional] 
**GroupBy** | **[]string** | List of tag prefixes to group by in the case of a cluster check. | [optional] 
**Grouping** | [**WidgetGrouping**](WidgetGrouping.md) |  | 
**Tags** | **[]string** | List of tags used to filter the groups reporting a cluster check. | [optional] 
**LegendSize** | [**WidgetLegendSize**](WidgetLegendSize.md) |  | [optional] 
**ShowLegend** | **bool** | (screenboard only) Show the legend for this widget. | [optional] 
**EventSize** | [**WidgetEventSize**](WidgetEventSize.md) |  | [optional] 
**Query** | **string** | Query to filter the monitors with. | 
**TagsExecution** | **string** | The execution method for multi-value filters. Can be either and or or. | [optional] 
**Color** | **string** | Color of the text. | [optional] 
**FontSize** | **string** | Size of the text. | [optional] 
**Text** | **string** | Text to display. | 
**LayoutType** | [**WidgetLayoutType**](WidgetLayoutType.md) |  | 
**Widgets** | [**[]Widget**](Widget.md) | List of widget groups. | 
**Events** | [**[]WidgetEvent**](WidgetEvent.md) | List of widget events. | [optional] 
**Yaxis** | [**WidgetAxis**](WidgetAxis.md) |  | [optional] 
**NoGroupHosts** | **bool** | Whether to show the hosts that don’t fit in a group. | [optional] 
**NoMetricHosts** | **bool** | Whether to show the hosts with no metrics. | [optional] 
**NodeType** | [**WidgetNodeType**](WidgetNodeType.md) |  | [optional] 
**Notes** | **string** | Notes on the title. | [optional] 
**Scope** | **[]string** | List of tags used to filter the map. | [optional] 
**Style** | [**HostMapWidgetDefinitionStyle**](HostMapWidgetDefinition_style.md) |  | [optional] 
**Url** | **string** | URL of the image. | 
**Margin** | [**WidgetMargin**](WidgetMargin.md) |  | [optional] 
**Sizing** | [**WidgetImageSizing**](WidgetImageSizing.md) |  | [optional] 
**Columns** | **[]string** | Which columns to display on the widget. | [optional] 
**Indexes** | **[]string** | An array of index names to query in the stream. Use [] to query all indexes at once. | [optional] 
**Logset** | **string** | ID of the log set to use. | [optional] 
**MessageDisplay** | [**WidgetMessageDisplay**](WidgetMessageDisplay.md) |  | [optional] 
**ShowDateColumn** | **bool** | Whether to show the date column or not | [optional] 
**ShowMessageColumn** | **bool** | Whether to show the message column or not | [optional] 
**Sort** | [**WidgetMonitorSummarySort**](WidgetMonitorSummarySort.md) |  | [optional] 
**ColorPreference** | [**WidgetColorPreference**](WidgetColorPreference.md) |  | [optional] 
**Count** | **int64** | The number of monitors to display. | [optional] 
**DisplayFormat** | [**WidgetServiceSummaryDisplayFormat**](WidgetServiceSummaryDisplayFormat.md) |  | [optional] 
**HideZeroCounts** | **bool** | Whether to show counts of 0 or not. | [optional] 
**ShowLastTriggered** | **bool** | Whether to show the time that has elapsed since the monitor/group triggered. | [optional] 
**Start** | **int64** | The start of the list. Typically 0. | [optional] 
**SummaryType** | [**WidgetSummaryType**](WidgetSummaryType.md) |  | [optional] 
**BackgroundColor** | **string** | Background color of the note. | [optional] 
**Content** | **string** | Content of the note. | 
**ShowTick** | **bool** | Whether to show a tick or not. | [optional] 
**TickEdge** | [**WidgetTickEdge**](WidgetTickEdge.md) |  | [optional] 
**TickPos** | **string** | Where to position the tick on an edge. | [optional] 
**Autoscale** | **bool** | Whether to use auto-scaling or not. | [optional] 
**CustomUnit** | **string** | Display a unit of your choice on the widget. | [optional] 
**ColorByGroups** | **[]string** | List of groups used for colors. | [optional] 
**Xaxis** | [**WidgetAxis**](WidgetAxis.md) |  | [optional] 
**ShowErrorBudget** | **bool** | Defined error budget. | [optional] 
**SloId** | **string** | ID of the SLO displayed. | [optional] 
**TimeWindows** | [**[]WidgetTimeWindows**](WidgetTimeWindows.md) | Times being monitored. | [optional] 
**ViewMode** | [**WidgetViewMode**](WidgetViewMode.md) |  | [optional] 
**ViewType** | **string** | Type of view displayed by the widget. | [default to detail]
**Filters** | **[]string** | Your environment and primary tag (or * if enabled for your account). | 
**Service** | **string** | APM service. | 
**Env** | **string** | APM environment. | 
**ShowBreakdown** | **bool** | Whether to show the latency breakdown or not. | [optional] 
**ShowDistribution** | **bool** | Whether to show the latency distribution or not. | [optional] 
**ShowErrors** | **bool** | Whether to show the error metrics or not. | [optional] 
**ShowHits** | **bool** | Whether to show the hits metrics or not. | [optional] 
**ShowLatency** | **bool** | Whether to show the latency metrics or not. | [optional] 
**ShowResourceList** | **bool** | Whether to show the resource list or not. | [optional] 
**SizeFormat** | [**WidgetSizeFormat**](WidgetSizeFormat.md) |  | [optional] 
**SpanName** | **string** | APM span name. | 
**Markers** | [**[]WidgetMarker**](WidgetMarker.md) | List of markers. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


