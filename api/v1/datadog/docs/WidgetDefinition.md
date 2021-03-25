# WidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | **string** | ID of the alert to use in the widget. | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of your widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | [**ToplistWidgetDefinitionType**](ToplistWidgetDefinitionType.md) |  | [default to TOPLISTWIDGETDEFINITIONTYPE_TOPLIST]
**VizType** | [**WidgetVizType**](WidgetVizType.md) |  | 
**Precision** | Pointer to **int64** | Number of decimals to show. If not defined, the widget uses the raw value. | [optional] 
**TextAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**Unit** | Pointer to **string** | Unit to display with the value. | [optional] 
**CustomLinks** | Pointer to [**[]WidgetCustomLink**](WidgetCustomLink.md) | List of custom links. | [optional] 
**Requests** | [**[]ToplistWidgetRequest**](ToplistWidgetRequest.md) | List of top list widget requests. | 
**Check** | **string** | Name of the check to use in the widget. | 
**Group** | Pointer to **[]string** | List of tag prefixes to group by. | [optional] 
**GroupBy** | Pointer to **[]string** | List of tag prefixes to group by in the case of a cluster check. | [optional] 
**Grouping** | [**WidgetGrouping**](WidgetGrouping.md) |  | 
**Tags** | Pointer to **[]string** | List of tags used to filter the groups reporting a cluster check. | [optional] 
**LegendSize** | Pointer to **string** | Available legend sizes for a widget. Should be one of \&quot;0\&quot;, \&quot;2\&quot;, \&quot;4\&quot;, \&quot;8\&quot;, \&quot;16\&quot;, or \&quot;auto\&quot;. | [optional] 
**ShowLegend** | Pointer to **bool** | (screenboard only) Show the legend for this widget. | [optional] 
**EventSize** | Pointer to [**WidgetEventSize**](WidgetEventSize.md) |  | [optional] 
**Query** | **string** | Query to filter the monitors with. | 
**TagsExecution** | Pointer to **string** | The execution method for multi-value filters. Can be either and or or. | [optional] 
**Color** | Pointer to **string** | Color of the text. | [optional] 
**FontSize** | Pointer to **string** | Size of the text. | [optional] 
**Text** | **string** | Text to display. | 
**Style** | [**HostMapWidgetDefinitionStyle**](HostMapWidgetDefinitionStyle.md) |  | 
**View** | [**GeomapWidgetDefinitionView**](GeomapWidgetDefinitionView.md) |  | 
**LayoutType** | [**WidgetLayoutType**](WidgetLayoutType.md) |  | 
**Widgets** | [**[]Widget**](Widget.md) | List of widget groups. | 
**Events** | Pointer to [**[]WidgetEvent**](WidgetEvent.md) | List of widget events. | [optional] 
**Yaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 
**NoGroupHosts** | Pointer to **bool** | Whether to show the hosts that don’t fit in a group. | [optional] 
**NoMetricHosts** | Pointer to **bool** | Whether to show the hosts with no metrics. | [optional] 
**NodeType** | Pointer to [**WidgetNodeType**](WidgetNodeType.md) |  | [optional] 
**Notes** | Pointer to **string** | Notes on the title. | [optional] 
**Scope** | Pointer to **[]string** | List of tags used to filter the map. | [optional] 
**Url** | **string** | URL of the image. | 
**Margin** | Pointer to [**WidgetMargin**](WidgetMargin.md) |  | [optional] 
**Sizing** | Pointer to [**WidgetImageSizing**](WidgetImageSizing.md) |  | [optional] 
**Columns** | Pointer to **[]string** | Which columns to display on the widget. | [optional] 
**Indexes** | Pointer to **[]string** | An array of index names to query in the stream. Use [] to query all indexes at once. | [optional] 
**Logset** | Pointer to **string** | ID of the log set to use. | [optional] 
**MessageDisplay** | Pointer to [**WidgetMessageDisplay**](WidgetMessageDisplay.md) |  | [optional] 
**ShowDateColumn** | Pointer to **bool** | Whether to show the date column or not | [optional] 
**ShowMessageColumn** | Pointer to **bool** | Whether to show the message column or not | [optional] 
**Sort** | Pointer to [**WidgetMonitorSummarySort**](WidgetMonitorSummarySort.md) |  | [optional] 
**ColorPreference** | Pointer to [**WidgetColorPreference**](WidgetColorPreference.md) |  | [optional] 
**Count** | Pointer to **int64** | The number of monitors to display. | [optional] 
**DisplayFormat** | Pointer to [**WidgetServiceSummaryDisplayFormat**](WidgetServiceSummaryDisplayFormat.md) |  | [optional] 
**HideZeroCounts** | Pointer to **bool** | Whether to show counts of 0 or not. | [optional] 
**ShowLastTriggered** | Pointer to **bool** | Whether to show the time that has elapsed since the monitor/group triggered. | [optional] 
**Start** | Pointer to **int64** | The start of the list. Typically 0. | [optional] 
**SummaryType** | Pointer to [**WidgetSummaryType**](WidgetSummaryType.md) |  | [optional] 
**BackgroundColor** | Pointer to **string** | Background color of the note. | [optional] 
**Content** | **string** | Content of the note. | 
**ShowTick** | Pointer to **bool** | Whether to show a tick or not. | [optional] 
**TickEdge** | Pointer to [**WidgetTickEdge**](WidgetTickEdge.md) |  | [optional] 
**TickPos** | Pointer to **string** | Where to position the tick on an edge. | [optional] 
**Autoscale** | Pointer to **bool** | Whether to use auto-scaling or not. | [optional] 
**CustomUnit** | Pointer to **string** | Display a unit of your choice on the widget. | [optional] 
**ColorByGroups** | Pointer to **[]string** | List of groups used for colors. | [optional] 
**Xaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 
**GlobalTimeTarget** | Pointer to **string** | Defined global time target. | [optional] 
**ShowErrorBudget** | Pointer to **bool** | Defined error budget. | [optional] 
**SloId** | Pointer to **string** | ID of the SLO displayed. | [optional] 
**TimeWindows** | Pointer to [**[]WidgetTimeWindows**](WidgetTimeWindows.md) | Times being monitored. | [optional] 
**ViewMode** | Pointer to [**WidgetViewMode**](WidgetViewMode.md) |  | [optional] 
**ViewType** | **string** | Type of view displayed by the widget. | [default to "detail"]
**Filters** | **[]string** | Your environment and primary tag (or * if enabled for your account). | 
**Service** | **string** | APM service. | 
**Env** | **string** | APM environment. | 
**ShowBreakdown** | Pointer to **bool** | Whether to show the latency breakdown or not. | [optional] 
**ShowDistribution** | Pointer to **bool** | Whether to show the latency distribution or not. | [optional] 
**ShowErrors** | Pointer to **bool** | Whether to show the error metrics or not. | [optional] 
**ShowHits** | Pointer to **bool** | Whether to show the hits metrics or not. | [optional] 
**ShowLatency** | Pointer to **bool** | Whether to show the latency metrics or not. | [optional] 
**ShowResourceList** | Pointer to **bool** | Whether to show the resource list or not. | [optional] 
**SizeFormat** | Pointer to [**WidgetSizeFormat**](WidgetSizeFormat.md) |  | [optional] 
**SpanName** | **string** | APM span name. | 
**HasSearchBar** | Pointer to [**TableWidgetHasSearchBar**](TableWidgetHasSearchBar.md) |  | [optional] 
**LegendColumns** | Pointer to [**[]TimeseriesWidgetLegendColumn**](TimeseriesWidgetLegendColumn.md) | Columns displayed in the legend. | [optional] 
**LegendLayout** | Pointer to [**TimeseriesWidgetLegendLayout**](TimeseriesWidgetLegendLayout.md) |  | [optional] 
**Markers** | Pointer to [**[]WidgetMarker**](WidgetMarker.md) | List of markers. | [optional] 
**RightYaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 

## Methods

### NewWidgetDefinition

`func NewWidgetDefinition(alertId string, type_ ToplistWidgetDefinitionType, vizType WidgetVizType, requests []ToplistWidgetRequest, check string, grouping WidgetGrouping, query string, text string, style HostMapWidgetDefinitionStyle, view GeomapWidgetDefinitionView, layoutType WidgetLayoutType, widgets []Widget, url string, content string, viewType string, filters []string, service string, env string, spanName string, ) *WidgetDefinition`

NewWidgetDefinition instantiates a new WidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetDefinitionWithDefaults

`func NewWidgetDefinitionWithDefaults() *WidgetDefinition`

NewWidgetDefinitionWithDefaults instantiates a new WidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *WidgetDefinition) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *WidgetDefinition) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *WidgetDefinition) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.


### GetTime

`func (o *WidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *WidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *WidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *WidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *WidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *WidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *WidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *WidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *WidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *WidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *WidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *WidgetDefinition) GetType() ToplistWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WidgetDefinition) GetTypeOk() (*ToplistWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WidgetDefinition) SetType(v ToplistWidgetDefinitionType)`

SetType sets Type field to given value.


### GetVizType

`func (o *WidgetDefinition) GetVizType() WidgetVizType`

GetVizType returns the VizType field if non-nil, zero value otherwise.

### GetVizTypeOk

`func (o *WidgetDefinition) GetVizTypeOk() (*WidgetVizType, bool)`

GetVizTypeOk returns a tuple with the VizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizType

`func (o *WidgetDefinition) SetVizType(v WidgetVizType)`

SetVizType sets VizType field to given value.


### GetPrecision

`func (o *WidgetDefinition) GetPrecision() int64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *WidgetDefinition) GetPrecisionOk() (*int64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *WidgetDefinition) SetPrecision(v int64)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *WidgetDefinition) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetTextAlign

`func (o *WidgetDefinition) GetTextAlign() WidgetTextAlign`

GetTextAlign returns the TextAlign field if non-nil, zero value otherwise.

### GetTextAlignOk

`func (o *WidgetDefinition) GetTextAlignOk() (*WidgetTextAlign, bool)`

GetTextAlignOk returns a tuple with the TextAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAlign

`func (o *WidgetDefinition) SetTextAlign(v WidgetTextAlign)`

SetTextAlign sets TextAlign field to given value.

### HasTextAlign

`func (o *WidgetDefinition) HasTextAlign() bool`

HasTextAlign returns a boolean if a field has been set.

### GetUnit

`func (o *WidgetDefinition) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *WidgetDefinition) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *WidgetDefinition) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *WidgetDefinition) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetCustomLinks

`func (o *WidgetDefinition) GetCustomLinks() []WidgetCustomLink`

GetCustomLinks returns the CustomLinks field if non-nil, zero value otherwise.

### GetCustomLinksOk

`func (o *WidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool)`

GetCustomLinksOk returns a tuple with the CustomLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLinks

`func (o *WidgetDefinition) SetCustomLinks(v []WidgetCustomLink)`

SetCustomLinks sets CustomLinks field to given value.

### HasCustomLinks

`func (o *WidgetDefinition) HasCustomLinks() bool`

HasCustomLinks returns a boolean if a field has been set.

### GetRequests

`func (o *WidgetDefinition) GetRequests() []ToplistWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *WidgetDefinition) GetRequestsOk() (*[]ToplistWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *WidgetDefinition) SetRequests(v []ToplistWidgetRequest)`

SetRequests sets Requests field to given value.


### GetCheck

`func (o *WidgetDefinition) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *WidgetDefinition) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *WidgetDefinition) SetCheck(v string)`

SetCheck sets Check field to given value.


### GetGroup

`func (o *WidgetDefinition) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WidgetDefinition) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WidgetDefinition) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WidgetDefinition) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupBy

`func (o *WidgetDefinition) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *WidgetDefinition) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *WidgetDefinition) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *WidgetDefinition) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGrouping

`func (o *WidgetDefinition) GetGrouping() WidgetGrouping`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *WidgetDefinition) GetGroupingOk() (*WidgetGrouping, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *WidgetDefinition) SetGrouping(v WidgetGrouping)`

SetGrouping sets Grouping field to given value.


### GetTags

`func (o *WidgetDefinition) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WidgetDefinition) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WidgetDefinition) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WidgetDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLegendSize

`func (o *WidgetDefinition) GetLegendSize() string`

GetLegendSize returns the LegendSize field if non-nil, zero value otherwise.

### GetLegendSizeOk

`func (o *WidgetDefinition) GetLegendSizeOk() (*string, bool)`

GetLegendSizeOk returns a tuple with the LegendSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendSize

`func (o *WidgetDefinition) SetLegendSize(v string)`

SetLegendSize sets LegendSize field to given value.

### HasLegendSize

`func (o *WidgetDefinition) HasLegendSize() bool`

HasLegendSize returns a boolean if a field has been set.

### GetShowLegend

`func (o *WidgetDefinition) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *WidgetDefinition) GetShowLegendOk() (*bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegend

`func (o *WidgetDefinition) SetShowLegend(v bool)`

SetShowLegend sets ShowLegend field to given value.

### HasShowLegend

`func (o *WidgetDefinition) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### GetEventSize

`func (o *WidgetDefinition) GetEventSize() WidgetEventSize`

GetEventSize returns the EventSize field if non-nil, zero value otherwise.

### GetEventSizeOk

`func (o *WidgetDefinition) GetEventSizeOk() (*WidgetEventSize, bool)`

GetEventSizeOk returns a tuple with the EventSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSize

`func (o *WidgetDefinition) SetEventSize(v WidgetEventSize)`

SetEventSize sets EventSize field to given value.

### HasEventSize

`func (o *WidgetDefinition) HasEventSize() bool`

HasEventSize returns a boolean if a field has been set.

### GetQuery

`func (o *WidgetDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *WidgetDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *WidgetDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTagsExecution

`func (o *WidgetDefinition) GetTagsExecution() string`

GetTagsExecution returns the TagsExecution field if non-nil, zero value otherwise.

### GetTagsExecutionOk

`func (o *WidgetDefinition) GetTagsExecutionOk() (*string, bool)`

GetTagsExecutionOk returns a tuple with the TagsExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsExecution

`func (o *WidgetDefinition) SetTagsExecution(v string)`

SetTagsExecution sets TagsExecution field to given value.

### HasTagsExecution

`func (o *WidgetDefinition) HasTagsExecution() bool`

HasTagsExecution returns a boolean if a field has been set.

### GetColor

`func (o *WidgetDefinition) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WidgetDefinition) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WidgetDefinition) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WidgetDefinition) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetFontSize

`func (o *WidgetDefinition) GetFontSize() string`

GetFontSize returns the FontSize field if non-nil, zero value otherwise.

### GetFontSizeOk

`func (o *WidgetDefinition) GetFontSizeOk() (*string, bool)`

GetFontSizeOk returns a tuple with the FontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSize

`func (o *WidgetDefinition) SetFontSize(v string)`

SetFontSize sets FontSize field to given value.

### HasFontSize

`func (o *WidgetDefinition) HasFontSize() bool`

HasFontSize returns a boolean if a field has been set.

### GetText

`func (o *WidgetDefinition) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WidgetDefinition) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WidgetDefinition) SetText(v string)`

SetText sets Text field to given value.


### GetStyle

`func (o *WidgetDefinition) GetStyle() HostMapWidgetDefinitionStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *WidgetDefinition) GetStyleOk() (*HostMapWidgetDefinitionStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *WidgetDefinition) SetStyle(v HostMapWidgetDefinitionStyle)`

SetStyle sets Style field to given value.


### GetView

`func (o *WidgetDefinition) GetView() GeomapWidgetDefinitionView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *WidgetDefinition) GetViewOk() (*GeomapWidgetDefinitionView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *WidgetDefinition) SetView(v GeomapWidgetDefinitionView)`

SetView sets View field to given value.


### GetLayoutType

`func (o *WidgetDefinition) GetLayoutType() WidgetLayoutType`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *WidgetDefinition) GetLayoutTypeOk() (*WidgetLayoutType, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *WidgetDefinition) SetLayoutType(v WidgetLayoutType)`

SetLayoutType sets LayoutType field to given value.


### GetWidgets

`func (o *WidgetDefinition) GetWidgets() []Widget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *WidgetDefinition) GetWidgetsOk() (*[]Widget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *WidgetDefinition) SetWidgets(v []Widget)`

SetWidgets sets Widgets field to given value.


### GetEvents

`func (o *WidgetDefinition) GetEvents() []WidgetEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WidgetDefinition) GetEventsOk() (*[]WidgetEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WidgetDefinition) SetEvents(v []WidgetEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WidgetDefinition) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetYaxis

`func (o *WidgetDefinition) GetYaxis() WidgetAxis`

GetYaxis returns the Yaxis field if non-nil, zero value otherwise.

### GetYaxisOk

`func (o *WidgetDefinition) GetYaxisOk() (*WidgetAxis, bool)`

GetYaxisOk returns a tuple with the Yaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaxis

`func (o *WidgetDefinition) SetYaxis(v WidgetAxis)`

SetYaxis sets Yaxis field to given value.

### HasYaxis

`func (o *WidgetDefinition) HasYaxis() bool`

HasYaxis returns a boolean if a field has been set.

### GetNoGroupHosts

`func (o *WidgetDefinition) GetNoGroupHosts() bool`

GetNoGroupHosts returns the NoGroupHosts field if non-nil, zero value otherwise.

### GetNoGroupHostsOk

`func (o *WidgetDefinition) GetNoGroupHostsOk() (*bool, bool)`

GetNoGroupHostsOk returns a tuple with the NoGroupHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoGroupHosts

`func (o *WidgetDefinition) SetNoGroupHosts(v bool)`

SetNoGroupHosts sets NoGroupHosts field to given value.

### HasNoGroupHosts

`func (o *WidgetDefinition) HasNoGroupHosts() bool`

HasNoGroupHosts returns a boolean if a field has been set.

### GetNoMetricHosts

`func (o *WidgetDefinition) GetNoMetricHosts() bool`

GetNoMetricHosts returns the NoMetricHosts field if non-nil, zero value otherwise.

### GetNoMetricHostsOk

`func (o *WidgetDefinition) GetNoMetricHostsOk() (*bool, bool)`

GetNoMetricHostsOk returns a tuple with the NoMetricHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMetricHosts

`func (o *WidgetDefinition) SetNoMetricHosts(v bool)`

SetNoMetricHosts sets NoMetricHosts field to given value.

### HasNoMetricHosts

`func (o *WidgetDefinition) HasNoMetricHosts() bool`

HasNoMetricHosts returns a boolean if a field has been set.

### GetNodeType

`func (o *WidgetDefinition) GetNodeType() WidgetNodeType`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *WidgetDefinition) GetNodeTypeOk() (*WidgetNodeType, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *WidgetDefinition) SetNodeType(v WidgetNodeType)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *WidgetDefinition) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetNotes

`func (o *WidgetDefinition) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *WidgetDefinition) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *WidgetDefinition) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *WidgetDefinition) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetScope

`func (o *WidgetDefinition) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WidgetDefinition) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *WidgetDefinition) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *WidgetDefinition) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUrl

`func (o *WidgetDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WidgetDefinition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WidgetDefinition) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMargin

`func (o *WidgetDefinition) GetMargin() WidgetMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *WidgetDefinition) GetMarginOk() (*WidgetMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *WidgetDefinition) SetMargin(v WidgetMargin)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *WidgetDefinition) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetSizing

`func (o *WidgetDefinition) GetSizing() WidgetImageSizing`

GetSizing returns the Sizing field if non-nil, zero value otherwise.

### GetSizingOk

`func (o *WidgetDefinition) GetSizingOk() (*WidgetImageSizing, bool)`

GetSizingOk returns a tuple with the Sizing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizing

`func (o *WidgetDefinition) SetSizing(v WidgetImageSizing)`

SetSizing sets Sizing field to given value.

### HasSizing

`func (o *WidgetDefinition) HasSizing() bool`

HasSizing returns a boolean if a field has been set.

### GetColumns

`func (o *WidgetDefinition) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *WidgetDefinition) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *WidgetDefinition) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *WidgetDefinition) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetIndexes

`func (o *WidgetDefinition) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *WidgetDefinition) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *WidgetDefinition) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *WidgetDefinition) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetLogset

`func (o *WidgetDefinition) GetLogset() string`

GetLogset returns the Logset field if non-nil, zero value otherwise.

### GetLogsetOk

`func (o *WidgetDefinition) GetLogsetOk() (*string, bool)`

GetLogsetOk returns a tuple with the Logset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogset

`func (o *WidgetDefinition) SetLogset(v string)`

SetLogset sets Logset field to given value.

### HasLogset

`func (o *WidgetDefinition) HasLogset() bool`

HasLogset returns a boolean if a field has been set.

### GetMessageDisplay

`func (o *WidgetDefinition) GetMessageDisplay() WidgetMessageDisplay`

GetMessageDisplay returns the MessageDisplay field if non-nil, zero value otherwise.

### GetMessageDisplayOk

`func (o *WidgetDefinition) GetMessageDisplayOk() (*WidgetMessageDisplay, bool)`

GetMessageDisplayOk returns a tuple with the MessageDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDisplay

`func (o *WidgetDefinition) SetMessageDisplay(v WidgetMessageDisplay)`

SetMessageDisplay sets MessageDisplay field to given value.

### HasMessageDisplay

`func (o *WidgetDefinition) HasMessageDisplay() bool`

HasMessageDisplay returns a boolean if a field has been set.

### GetShowDateColumn

`func (o *WidgetDefinition) GetShowDateColumn() bool`

GetShowDateColumn returns the ShowDateColumn field if non-nil, zero value otherwise.

### GetShowDateColumnOk

`func (o *WidgetDefinition) GetShowDateColumnOk() (*bool, bool)`

GetShowDateColumnOk returns a tuple with the ShowDateColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDateColumn

`func (o *WidgetDefinition) SetShowDateColumn(v bool)`

SetShowDateColumn sets ShowDateColumn field to given value.

### HasShowDateColumn

`func (o *WidgetDefinition) HasShowDateColumn() bool`

HasShowDateColumn returns a boolean if a field has been set.

### GetShowMessageColumn

`func (o *WidgetDefinition) GetShowMessageColumn() bool`

GetShowMessageColumn returns the ShowMessageColumn field if non-nil, zero value otherwise.

### GetShowMessageColumnOk

`func (o *WidgetDefinition) GetShowMessageColumnOk() (*bool, bool)`

GetShowMessageColumnOk returns a tuple with the ShowMessageColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMessageColumn

`func (o *WidgetDefinition) SetShowMessageColumn(v bool)`

SetShowMessageColumn sets ShowMessageColumn field to given value.

### HasShowMessageColumn

`func (o *WidgetDefinition) HasShowMessageColumn() bool`

HasShowMessageColumn returns a boolean if a field has been set.

### GetSort

`func (o *WidgetDefinition) GetSort() WidgetMonitorSummarySort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *WidgetDefinition) GetSortOk() (*WidgetMonitorSummarySort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *WidgetDefinition) SetSort(v WidgetMonitorSummarySort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *WidgetDefinition) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetColorPreference

`func (o *WidgetDefinition) GetColorPreference() WidgetColorPreference`

GetColorPreference returns the ColorPreference field if non-nil, zero value otherwise.

### GetColorPreferenceOk

`func (o *WidgetDefinition) GetColorPreferenceOk() (*WidgetColorPreference, bool)`

GetColorPreferenceOk returns a tuple with the ColorPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorPreference

`func (o *WidgetDefinition) SetColorPreference(v WidgetColorPreference)`

SetColorPreference sets ColorPreference field to given value.

### HasColorPreference

`func (o *WidgetDefinition) HasColorPreference() bool`

HasColorPreference returns a boolean if a field has been set.

### GetCount

`func (o *WidgetDefinition) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WidgetDefinition) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WidgetDefinition) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *WidgetDefinition) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDisplayFormat

`func (o *WidgetDefinition) GetDisplayFormat() WidgetServiceSummaryDisplayFormat`

GetDisplayFormat returns the DisplayFormat field if non-nil, zero value otherwise.

### GetDisplayFormatOk

`func (o *WidgetDefinition) GetDisplayFormatOk() (*WidgetServiceSummaryDisplayFormat, bool)`

GetDisplayFormatOk returns a tuple with the DisplayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFormat

`func (o *WidgetDefinition) SetDisplayFormat(v WidgetServiceSummaryDisplayFormat)`

SetDisplayFormat sets DisplayFormat field to given value.

### HasDisplayFormat

`func (o *WidgetDefinition) HasDisplayFormat() bool`

HasDisplayFormat returns a boolean if a field has been set.

### GetHideZeroCounts

`func (o *WidgetDefinition) GetHideZeroCounts() bool`

GetHideZeroCounts returns the HideZeroCounts field if non-nil, zero value otherwise.

### GetHideZeroCountsOk

`func (o *WidgetDefinition) GetHideZeroCountsOk() (*bool, bool)`

GetHideZeroCountsOk returns a tuple with the HideZeroCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideZeroCounts

`func (o *WidgetDefinition) SetHideZeroCounts(v bool)`

SetHideZeroCounts sets HideZeroCounts field to given value.

### HasHideZeroCounts

`func (o *WidgetDefinition) HasHideZeroCounts() bool`

HasHideZeroCounts returns a boolean if a field has been set.

### GetShowLastTriggered

`func (o *WidgetDefinition) GetShowLastTriggered() bool`

GetShowLastTriggered returns the ShowLastTriggered field if non-nil, zero value otherwise.

### GetShowLastTriggeredOk

`func (o *WidgetDefinition) GetShowLastTriggeredOk() (*bool, bool)`

GetShowLastTriggeredOk returns a tuple with the ShowLastTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastTriggered

`func (o *WidgetDefinition) SetShowLastTriggered(v bool)`

SetShowLastTriggered sets ShowLastTriggered field to given value.

### HasShowLastTriggered

`func (o *WidgetDefinition) HasShowLastTriggered() bool`

HasShowLastTriggered returns a boolean if a field has been set.

### GetStart

`func (o *WidgetDefinition) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *WidgetDefinition) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *WidgetDefinition) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *WidgetDefinition) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetSummaryType

`func (o *WidgetDefinition) GetSummaryType() WidgetSummaryType`

GetSummaryType returns the SummaryType field if non-nil, zero value otherwise.

### GetSummaryTypeOk

`func (o *WidgetDefinition) GetSummaryTypeOk() (*WidgetSummaryType, bool)`

GetSummaryTypeOk returns a tuple with the SummaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryType

`func (o *WidgetDefinition) SetSummaryType(v WidgetSummaryType)`

SetSummaryType sets SummaryType field to given value.

### HasSummaryType

`func (o *WidgetDefinition) HasSummaryType() bool`

HasSummaryType returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *WidgetDefinition) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *WidgetDefinition) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *WidgetDefinition) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *WidgetDefinition) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetContent

`func (o *WidgetDefinition) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WidgetDefinition) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WidgetDefinition) SetContent(v string)`

SetContent sets Content field to given value.


### GetShowTick

`func (o *WidgetDefinition) GetShowTick() bool`

GetShowTick returns the ShowTick field if non-nil, zero value otherwise.

### GetShowTickOk

`func (o *WidgetDefinition) GetShowTickOk() (*bool, bool)`

GetShowTickOk returns a tuple with the ShowTick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTick

`func (o *WidgetDefinition) SetShowTick(v bool)`

SetShowTick sets ShowTick field to given value.

### HasShowTick

`func (o *WidgetDefinition) HasShowTick() bool`

HasShowTick returns a boolean if a field has been set.

### GetTickEdge

`func (o *WidgetDefinition) GetTickEdge() WidgetTickEdge`

GetTickEdge returns the TickEdge field if non-nil, zero value otherwise.

### GetTickEdgeOk

`func (o *WidgetDefinition) GetTickEdgeOk() (*WidgetTickEdge, bool)`

GetTickEdgeOk returns a tuple with the TickEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickEdge

`func (o *WidgetDefinition) SetTickEdge(v WidgetTickEdge)`

SetTickEdge sets TickEdge field to given value.

### HasTickEdge

`func (o *WidgetDefinition) HasTickEdge() bool`

HasTickEdge returns a boolean if a field has been set.

### GetTickPos

`func (o *WidgetDefinition) GetTickPos() string`

GetTickPos returns the TickPos field if non-nil, zero value otherwise.

### GetTickPosOk

`func (o *WidgetDefinition) GetTickPosOk() (*string, bool)`

GetTickPosOk returns a tuple with the TickPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickPos

`func (o *WidgetDefinition) SetTickPos(v string)`

SetTickPos sets TickPos field to given value.

### HasTickPos

`func (o *WidgetDefinition) HasTickPos() bool`

HasTickPos returns a boolean if a field has been set.

### GetAutoscale

`func (o *WidgetDefinition) GetAutoscale() bool`

GetAutoscale returns the Autoscale field if non-nil, zero value otherwise.

### GetAutoscaleOk

`func (o *WidgetDefinition) GetAutoscaleOk() (*bool, bool)`

GetAutoscaleOk returns a tuple with the Autoscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscale

`func (o *WidgetDefinition) SetAutoscale(v bool)`

SetAutoscale sets Autoscale field to given value.

### HasAutoscale

`func (o *WidgetDefinition) HasAutoscale() bool`

HasAutoscale returns a boolean if a field has been set.

### GetCustomUnit

`func (o *WidgetDefinition) GetCustomUnit() string`

GetCustomUnit returns the CustomUnit field if non-nil, zero value otherwise.

### GetCustomUnitOk

`func (o *WidgetDefinition) GetCustomUnitOk() (*string, bool)`

GetCustomUnitOk returns a tuple with the CustomUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUnit

`func (o *WidgetDefinition) SetCustomUnit(v string)`

SetCustomUnit sets CustomUnit field to given value.

### HasCustomUnit

`func (o *WidgetDefinition) HasCustomUnit() bool`

HasCustomUnit returns a boolean if a field has been set.

### GetColorByGroups

`func (o *WidgetDefinition) GetColorByGroups() []string`

GetColorByGroups returns the ColorByGroups field if non-nil, zero value otherwise.

### GetColorByGroupsOk

`func (o *WidgetDefinition) GetColorByGroupsOk() (*[]string, bool)`

GetColorByGroupsOk returns a tuple with the ColorByGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorByGroups

`func (o *WidgetDefinition) SetColorByGroups(v []string)`

SetColorByGroups sets ColorByGroups field to given value.

### HasColorByGroups

`func (o *WidgetDefinition) HasColorByGroups() bool`

HasColorByGroups returns a boolean if a field has been set.

### GetXaxis

`func (o *WidgetDefinition) GetXaxis() WidgetAxis`

GetXaxis returns the Xaxis field if non-nil, zero value otherwise.

### GetXaxisOk

`func (o *WidgetDefinition) GetXaxisOk() (*WidgetAxis, bool)`

GetXaxisOk returns a tuple with the Xaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXaxis

`func (o *WidgetDefinition) SetXaxis(v WidgetAxis)`

SetXaxis sets Xaxis field to given value.

### HasXaxis

`func (o *WidgetDefinition) HasXaxis() bool`

HasXaxis returns a boolean if a field has been set.

### GetGlobalTimeTarget

`func (o *WidgetDefinition) GetGlobalTimeTarget() string`

GetGlobalTimeTarget returns the GlobalTimeTarget field if non-nil, zero value otherwise.

### GetGlobalTimeTargetOk

`func (o *WidgetDefinition) GetGlobalTimeTargetOk() (*string, bool)`

GetGlobalTimeTargetOk returns a tuple with the GlobalTimeTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTimeTarget

`func (o *WidgetDefinition) SetGlobalTimeTarget(v string)`

SetGlobalTimeTarget sets GlobalTimeTarget field to given value.

### HasGlobalTimeTarget

`func (o *WidgetDefinition) HasGlobalTimeTarget() bool`

HasGlobalTimeTarget returns a boolean if a field has been set.

### GetShowErrorBudget

`func (o *WidgetDefinition) GetShowErrorBudget() bool`

GetShowErrorBudget returns the ShowErrorBudget field if non-nil, zero value otherwise.

### GetShowErrorBudgetOk

`func (o *WidgetDefinition) GetShowErrorBudgetOk() (*bool, bool)`

GetShowErrorBudgetOk returns a tuple with the ShowErrorBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowErrorBudget

`func (o *WidgetDefinition) SetShowErrorBudget(v bool)`

SetShowErrorBudget sets ShowErrorBudget field to given value.

### HasShowErrorBudget

`func (o *WidgetDefinition) HasShowErrorBudget() bool`

HasShowErrorBudget returns a boolean if a field has been set.

### GetSloId

`func (o *WidgetDefinition) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *WidgetDefinition) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *WidgetDefinition) SetSloId(v string)`

SetSloId sets SloId field to given value.

### HasSloId

`func (o *WidgetDefinition) HasSloId() bool`

HasSloId returns a boolean if a field has been set.

### GetTimeWindows

`func (o *WidgetDefinition) GetTimeWindows() []WidgetTimeWindows`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *WidgetDefinition) GetTimeWindowsOk() (*[]WidgetTimeWindows, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *WidgetDefinition) SetTimeWindows(v []WidgetTimeWindows)`

SetTimeWindows sets TimeWindows field to given value.

### HasTimeWindows

`func (o *WidgetDefinition) HasTimeWindows() bool`

HasTimeWindows returns a boolean if a field has been set.

### GetViewMode

`func (o *WidgetDefinition) GetViewMode() WidgetViewMode`

GetViewMode returns the ViewMode field if non-nil, zero value otherwise.

### GetViewModeOk

`func (o *WidgetDefinition) GetViewModeOk() (*WidgetViewMode, bool)`

GetViewModeOk returns a tuple with the ViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMode

`func (o *WidgetDefinition) SetViewMode(v WidgetViewMode)`

SetViewMode sets ViewMode field to given value.

### HasViewMode

`func (o *WidgetDefinition) HasViewMode() bool`

HasViewMode returns a boolean if a field has been set.

### GetViewType

`func (o *WidgetDefinition) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *WidgetDefinition) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *WidgetDefinition) SetViewType(v string)`

SetViewType sets ViewType field to given value.


### GetFilters

`func (o *WidgetDefinition) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *WidgetDefinition) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *WidgetDefinition) SetFilters(v []string)`

SetFilters sets Filters field to given value.


### GetService

`func (o *WidgetDefinition) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *WidgetDefinition) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *WidgetDefinition) SetService(v string)`

SetService sets Service field to given value.


### GetEnv

`func (o *WidgetDefinition) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *WidgetDefinition) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *WidgetDefinition) SetEnv(v string)`

SetEnv sets Env field to given value.


### GetShowBreakdown

`func (o *WidgetDefinition) GetShowBreakdown() bool`

GetShowBreakdown returns the ShowBreakdown field if non-nil, zero value otherwise.

### GetShowBreakdownOk

`func (o *WidgetDefinition) GetShowBreakdownOk() (*bool, bool)`

GetShowBreakdownOk returns a tuple with the ShowBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBreakdown

`func (o *WidgetDefinition) SetShowBreakdown(v bool)`

SetShowBreakdown sets ShowBreakdown field to given value.

### HasShowBreakdown

`func (o *WidgetDefinition) HasShowBreakdown() bool`

HasShowBreakdown returns a boolean if a field has been set.

### GetShowDistribution

`func (o *WidgetDefinition) GetShowDistribution() bool`

GetShowDistribution returns the ShowDistribution field if non-nil, zero value otherwise.

### GetShowDistributionOk

`func (o *WidgetDefinition) GetShowDistributionOk() (*bool, bool)`

GetShowDistributionOk returns a tuple with the ShowDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDistribution

`func (o *WidgetDefinition) SetShowDistribution(v bool)`

SetShowDistribution sets ShowDistribution field to given value.

### HasShowDistribution

`func (o *WidgetDefinition) HasShowDistribution() bool`

HasShowDistribution returns a boolean if a field has been set.

### GetShowErrors

`func (o *WidgetDefinition) GetShowErrors() bool`

GetShowErrors returns the ShowErrors field if non-nil, zero value otherwise.

### GetShowErrorsOk

`func (o *WidgetDefinition) GetShowErrorsOk() (*bool, bool)`

GetShowErrorsOk returns a tuple with the ShowErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowErrors

`func (o *WidgetDefinition) SetShowErrors(v bool)`

SetShowErrors sets ShowErrors field to given value.

### HasShowErrors

`func (o *WidgetDefinition) HasShowErrors() bool`

HasShowErrors returns a boolean if a field has been set.

### GetShowHits

`func (o *WidgetDefinition) GetShowHits() bool`

GetShowHits returns the ShowHits field if non-nil, zero value otherwise.

### GetShowHitsOk

`func (o *WidgetDefinition) GetShowHitsOk() (*bool, bool)`

GetShowHitsOk returns a tuple with the ShowHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHits

`func (o *WidgetDefinition) SetShowHits(v bool)`

SetShowHits sets ShowHits field to given value.

### HasShowHits

`func (o *WidgetDefinition) HasShowHits() bool`

HasShowHits returns a boolean if a field has been set.

### GetShowLatency

`func (o *WidgetDefinition) GetShowLatency() bool`

GetShowLatency returns the ShowLatency field if non-nil, zero value otherwise.

### GetShowLatencyOk

`func (o *WidgetDefinition) GetShowLatencyOk() (*bool, bool)`

GetShowLatencyOk returns a tuple with the ShowLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLatency

`func (o *WidgetDefinition) SetShowLatency(v bool)`

SetShowLatency sets ShowLatency field to given value.

### HasShowLatency

`func (o *WidgetDefinition) HasShowLatency() bool`

HasShowLatency returns a boolean if a field has been set.

### GetShowResourceList

`func (o *WidgetDefinition) GetShowResourceList() bool`

GetShowResourceList returns the ShowResourceList field if non-nil, zero value otherwise.

### GetShowResourceListOk

`func (o *WidgetDefinition) GetShowResourceListOk() (*bool, bool)`

GetShowResourceListOk returns a tuple with the ShowResourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResourceList

`func (o *WidgetDefinition) SetShowResourceList(v bool)`

SetShowResourceList sets ShowResourceList field to given value.

### HasShowResourceList

`func (o *WidgetDefinition) HasShowResourceList() bool`

HasShowResourceList returns a boolean if a field has been set.

### GetSizeFormat

`func (o *WidgetDefinition) GetSizeFormat() WidgetSizeFormat`

GetSizeFormat returns the SizeFormat field if non-nil, zero value otherwise.

### GetSizeFormatOk

`func (o *WidgetDefinition) GetSizeFormatOk() (*WidgetSizeFormat, bool)`

GetSizeFormatOk returns a tuple with the SizeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeFormat

`func (o *WidgetDefinition) SetSizeFormat(v WidgetSizeFormat)`

SetSizeFormat sets SizeFormat field to given value.

### HasSizeFormat

`func (o *WidgetDefinition) HasSizeFormat() bool`

HasSizeFormat returns a boolean if a field has been set.

### GetSpanName

`func (o *WidgetDefinition) GetSpanName() string`

GetSpanName returns the SpanName field if non-nil, zero value otherwise.

### GetSpanNameOk

`func (o *WidgetDefinition) GetSpanNameOk() (*string, bool)`

GetSpanNameOk returns a tuple with the SpanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanName

`func (o *WidgetDefinition) SetSpanName(v string)`

SetSpanName sets SpanName field to given value.


### GetHasSearchBar

`func (o *WidgetDefinition) GetHasSearchBar() TableWidgetHasSearchBar`

GetHasSearchBar returns the HasSearchBar field if non-nil, zero value otherwise.

### GetHasSearchBarOk

`func (o *WidgetDefinition) GetHasSearchBarOk() (*TableWidgetHasSearchBar, bool)`

GetHasSearchBarOk returns a tuple with the HasSearchBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSearchBar

`func (o *WidgetDefinition) SetHasSearchBar(v TableWidgetHasSearchBar)`

SetHasSearchBar sets HasSearchBar field to given value.

### HasHasSearchBar

`func (o *WidgetDefinition) HasHasSearchBar() bool`

HasHasSearchBar returns a boolean if a field has been set.

### GetLegendColumns

`func (o *WidgetDefinition) GetLegendColumns() []TimeseriesWidgetLegendColumn`

GetLegendColumns returns the LegendColumns field if non-nil, zero value otherwise.

### GetLegendColumnsOk

`func (o *WidgetDefinition) GetLegendColumnsOk() (*[]TimeseriesWidgetLegendColumn, bool)`

GetLegendColumnsOk returns a tuple with the LegendColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendColumns

`func (o *WidgetDefinition) SetLegendColumns(v []TimeseriesWidgetLegendColumn)`

SetLegendColumns sets LegendColumns field to given value.

### HasLegendColumns

`func (o *WidgetDefinition) HasLegendColumns() bool`

HasLegendColumns returns a boolean if a field has been set.

### GetLegendLayout

`func (o *WidgetDefinition) GetLegendLayout() TimeseriesWidgetLegendLayout`

GetLegendLayout returns the LegendLayout field if non-nil, zero value otherwise.

### GetLegendLayoutOk

`func (o *WidgetDefinition) GetLegendLayoutOk() (*TimeseriesWidgetLegendLayout, bool)`

GetLegendLayoutOk returns a tuple with the LegendLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendLayout

`func (o *WidgetDefinition) SetLegendLayout(v TimeseriesWidgetLegendLayout)`

SetLegendLayout sets LegendLayout field to given value.

### HasLegendLayout

`func (o *WidgetDefinition) HasLegendLayout() bool`

HasLegendLayout returns a boolean if a field has been set.

### GetMarkers

`func (o *WidgetDefinition) GetMarkers() []WidgetMarker`

GetMarkers returns the Markers field if non-nil, zero value otherwise.

### GetMarkersOk

`func (o *WidgetDefinition) GetMarkersOk() (*[]WidgetMarker, bool)`

GetMarkersOk returns a tuple with the Markers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkers

`func (o *WidgetDefinition) SetMarkers(v []WidgetMarker)`

SetMarkers sets Markers field to given value.

### HasMarkers

`func (o *WidgetDefinition) HasMarkers() bool`

HasMarkers returns a boolean if a field has been set.

### GetRightYaxis

`func (o *WidgetDefinition) GetRightYaxis() WidgetAxis`

GetRightYaxis returns the RightYaxis field if non-nil, zero value otherwise.

### GetRightYaxisOk

`func (o *WidgetDefinition) GetRightYaxisOk() (*WidgetAxis, bool)`

GetRightYaxisOk returns a tuple with the RightYaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightYaxis

`func (o *WidgetDefinition) SetRightYaxis(v WidgetAxis)`

SetRightYaxis sets RightYaxis field to given value.

### HasRightYaxis

`func (o *WidgetDefinition) HasRightYaxis() bool`

HasRightYaxis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


