# TimeseriesWidgetDefinition

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CustomLinks** | Pointer to [**[]WidgetCustomLink**](WidgetCustomLink.md) | List of custom links. | [optional] 
**Events** | Pointer to [**[]WidgetEvent**](WidgetEvent.md) | List of widget events. | [optional] 
**LegendColumns** | Pointer to [**[]TimeseriesWidgetLegendColumn**](TimeseriesWidgetLegendColumn.md) | Columns displayed in the legend. | [optional] 
**LegendLayout** | Pointer to [**TimeseriesWidgetLegendLayout**](TimeseriesWidgetLegendLayout.md) |  | [optional] 
**LegendSize** | Pointer to **string** | Available legend sizes for a widget. Should be one of \&quot;0\&quot;, \&quot;2\&quot;, \&quot;4\&quot;, \&quot;8\&quot;, \&quot;16\&quot;, or \&quot;auto\&quot;. | [optional] 
**Markers** | Pointer to [**[]WidgetMarker**](WidgetMarker.md) | List of markers. | [optional] 
**Requests** | [**[]TimeseriesWidgetRequest**](TimeseriesWidgetRequest.md) | List of timeseries widget requests. | 
**RightYaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 
**ShowLegend** | Pointer to **bool** | (screenboard only) Show the legend for this widget. | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of your widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | [**TimeseriesWidgetDefinitionType**](TimeseriesWidgetDefinitionType.md) |  | [default to TIMESERIESWIDGETDEFINITIONTYPE_TIMESERIES]
**Yaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 

## Methods

### NewTimeseriesWidgetDefinition

`func NewTimeseriesWidgetDefinition(requests []TimeseriesWidgetRequest, type_ TimeseriesWidgetDefinitionType, ) *TimeseriesWidgetDefinition`

NewTimeseriesWidgetDefinition instantiates a new TimeseriesWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesWidgetDefinitionWithDefaults

`func NewTimeseriesWidgetDefinitionWithDefaults() *TimeseriesWidgetDefinition`

NewTimeseriesWidgetDefinitionWithDefaults instantiates a new TimeseriesWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomLinks

`func (o *TimeseriesWidgetDefinition) GetCustomLinks() []WidgetCustomLink`

GetCustomLinks returns the CustomLinks field if non-nil, zero value otherwise.

### GetCustomLinksOk

`func (o *TimeseriesWidgetDefinition) GetCustomLinksOk() (*[]WidgetCustomLink, bool)`

GetCustomLinksOk returns a tuple with the CustomLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLinks

`func (o *TimeseriesWidgetDefinition) SetCustomLinks(v []WidgetCustomLink)`

SetCustomLinks sets CustomLinks field to given value.

### HasCustomLinks

`func (o *TimeseriesWidgetDefinition) HasCustomLinks() bool`

HasCustomLinks returns a boolean if a field has been set.

### GetEvents

`func (o *TimeseriesWidgetDefinition) GetEvents() []WidgetEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TimeseriesWidgetDefinition) GetEventsOk() (*[]WidgetEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TimeseriesWidgetDefinition) SetEvents(v []WidgetEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TimeseriesWidgetDefinition) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLegendColumns

`func (o *TimeseriesWidgetDefinition) GetLegendColumns() []TimeseriesWidgetLegendColumn`

GetLegendColumns returns the LegendColumns field if non-nil, zero value otherwise.

### GetLegendColumnsOk

`func (o *TimeseriesWidgetDefinition) GetLegendColumnsOk() (*[]TimeseriesWidgetLegendColumn, bool)`

GetLegendColumnsOk returns a tuple with the LegendColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendColumns

`func (o *TimeseriesWidgetDefinition) SetLegendColumns(v []TimeseriesWidgetLegendColumn)`

SetLegendColumns sets LegendColumns field to given value.

### HasLegendColumns

`func (o *TimeseriesWidgetDefinition) HasLegendColumns() bool`

HasLegendColumns returns a boolean if a field has been set.

### GetLegendLayout

`func (o *TimeseriesWidgetDefinition) GetLegendLayout() TimeseriesWidgetLegendLayout`

GetLegendLayout returns the LegendLayout field if non-nil, zero value otherwise.

### GetLegendLayoutOk

`func (o *TimeseriesWidgetDefinition) GetLegendLayoutOk() (*TimeseriesWidgetLegendLayout, bool)`

GetLegendLayoutOk returns a tuple with the LegendLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendLayout

`func (o *TimeseriesWidgetDefinition) SetLegendLayout(v TimeseriesWidgetLegendLayout)`

SetLegendLayout sets LegendLayout field to given value.

### HasLegendLayout

`func (o *TimeseriesWidgetDefinition) HasLegendLayout() bool`

HasLegendLayout returns a boolean if a field has been set.

### GetLegendSize

`func (o *TimeseriesWidgetDefinition) GetLegendSize() string`

GetLegendSize returns the LegendSize field if non-nil, zero value otherwise.

### GetLegendSizeOk

`func (o *TimeseriesWidgetDefinition) GetLegendSizeOk() (*string, bool)`

GetLegendSizeOk returns a tuple with the LegendSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendSize

`func (o *TimeseriesWidgetDefinition) SetLegendSize(v string)`

SetLegendSize sets LegendSize field to given value.

### HasLegendSize

`func (o *TimeseriesWidgetDefinition) HasLegendSize() bool`

HasLegendSize returns a boolean if a field has been set.

### GetMarkers

`func (o *TimeseriesWidgetDefinition) GetMarkers() []WidgetMarker`

GetMarkers returns the Markers field if non-nil, zero value otherwise.

### GetMarkersOk

`func (o *TimeseriesWidgetDefinition) GetMarkersOk() (*[]WidgetMarker, bool)`

GetMarkersOk returns a tuple with the Markers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkers

`func (o *TimeseriesWidgetDefinition) SetMarkers(v []WidgetMarker)`

SetMarkers sets Markers field to given value.

### HasMarkers

`func (o *TimeseriesWidgetDefinition) HasMarkers() bool`

HasMarkers returns a boolean if a field has been set.

### GetRequests

`func (o *TimeseriesWidgetDefinition) GetRequests() []TimeseriesWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TimeseriesWidgetDefinition) GetRequestsOk() (*[]TimeseriesWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TimeseriesWidgetDefinition) SetRequests(v []TimeseriesWidgetRequest)`

SetRequests sets Requests field to given value.


### GetRightYaxis

`func (o *TimeseriesWidgetDefinition) GetRightYaxis() WidgetAxis`

GetRightYaxis returns the RightYaxis field if non-nil, zero value otherwise.

### GetRightYaxisOk

`func (o *TimeseriesWidgetDefinition) GetRightYaxisOk() (*WidgetAxis, bool)`

GetRightYaxisOk returns a tuple with the RightYaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightYaxis

`func (o *TimeseriesWidgetDefinition) SetRightYaxis(v WidgetAxis)`

SetRightYaxis sets RightYaxis field to given value.

### HasRightYaxis

`func (o *TimeseriesWidgetDefinition) HasRightYaxis() bool`

HasRightYaxis returns a boolean if a field has been set.

### GetShowLegend

`func (o *TimeseriesWidgetDefinition) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *TimeseriesWidgetDefinition) GetShowLegendOk() (*bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegend

`func (o *TimeseriesWidgetDefinition) SetShowLegend(v bool)`

SetShowLegend sets ShowLegend field to given value.

### HasShowLegend

`func (o *TimeseriesWidgetDefinition) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### GetTime

`func (o *TimeseriesWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TimeseriesWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TimeseriesWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *TimeseriesWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *TimeseriesWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TimeseriesWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TimeseriesWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TimeseriesWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *TimeseriesWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *TimeseriesWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *TimeseriesWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *TimeseriesWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *TimeseriesWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *TimeseriesWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *TimeseriesWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *TimeseriesWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *TimeseriesWidgetDefinition) GetType() TimeseriesWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeseriesWidgetDefinition) GetTypeOk() (*TimeseriesWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeseriesWidgetDefinition) SetType(v TimeseriesWidgetDefinitionType)`

SetType sets Type field to given value.


### GetYaxis

`func (o *TimeseriesWidgetDefinition) GetYaxis() WidgetAxis`

GetYaxis returns the Yaxis field if non-nil, zero value otherwise.

### GetYaxisOk

`func (o *TimeseriesWidgetDefinition) GetYaxisOk() (*WidgetAxis, bool)`

GetYaxisOk returns a tuple with the Yaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaxis

`func (o *TimeseriesWidgetDefinition) SetYaxis(v WidgetAxis)`

SetYaxis sets Yaxis field to given value.

### HasYaxis

`func (o *TimeseriesWidgetDefinition) HasYaxis() bool`

HasYaxis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


