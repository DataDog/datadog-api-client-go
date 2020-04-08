# TimeseriesWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]WidgetEvent**](WidgetEvent.md) | TODO. | [optional] 
**LegendSize** | Pointer to [**WidgetLegendSize**](WidgetLegendSize.md) |  | [optional] 
**Markers** | Pointer to [**[]WidgetMarker**](WidgetMarker.md) | TODO. | [optional] 
**Requests** | Pointer to [**[]TimeseriesWidgetRequest**](TimeseriesWidgetRequest.md) | TODO. | 
**ShowLegend** | Pointer to **bool** | (screenboard only) Show the legend for this widget | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of your widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "timeseries"]
**Yaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 

## Methods

### NewTimeseriesWidgetDefinition

`func NewTimeseriesWidgetDefinition(requests []TimeseriesWidgetRequest, type_ string, ) *TimeseriesWidgetDefinition`

NewTimeseriesWidgetDefinition instantiates a new TimeseriesWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesWidgetDefinitionWithDefaults

`func NewTimeseriesWidgetDefinitionWithDefaults() *TimeseriesWidgetDefinition`

NewTimeseriesWidgetDefinitionWithDefaults instantiates a new TimeseriesWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetLegendSize

`func (o *TimeseriesWidgetDefinition) GetLegendSize() WidgetLegendSize`

GetLegendSize returns the LegendSize field if non-nil, zero value otherwise.

### GetLegendSizeOk

`func (o *TimeseriesWidgetDefinition) GetLegendSizeOk() (*WidgetLegendSize, bool)`

GetLegendSizeOk returns a tuple with the LegendSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendSize

`func (o *TimeseriesWidgetDefinition) SetLegendSize(v WidgetLegendSize)`

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

`func (o *TimeseriesWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeseriesWidgetDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeseriesWidgetDefinition) SetType(v string)`

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


### AsWidgetDefinition

`func (s *TimeseriesWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of TimeseriesWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


