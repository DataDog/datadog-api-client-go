# HeatMapWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]WidgetEvent**](WidgetEvent.md) | List of widget events. | [optional] 
**LegendSize** | Pointer to [**WidgetLegendSize**](WidgetLegendSize.md) |  | [optional] 
**Requests** | Pointer to [**[]HeatMapWidgetRequest**](HeatMapWidgetRequest.md) | List of widget types. | 
**ShowLegend** | Pointer to **bool** | Whether or not to display the legend on this widget. | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | Pointer to **string** | Type of the widget. | [readonly] [default to "heatmap"]
**Yaxis** | Pointer to [**WidgetAxis**](WidgetAxis.md) |  | [optional] 

## Methods

### NewHeatMapWidgetDefinition

`func NewHeatMapWidgetDefinition(requests []HeatMapWidgetRequest, type_ string, ) *HeatMapWidgetDefinition`

NewHeatMapWidgetDefinition instantiates a new HeatMapWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeatMapWidgetDefinitionWithDefaults

`func NewHeatMapWidgetDefinitionWithDefaults() *HeatMapWidgetDefinition`

NewHeatMapWidgetDefinitionWithDefaults instantiates a new HeatMapWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *HeatMapWidgetDefinition) GetEvents() []WidgetEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HeatMapWidgetDefinition) GetEventsOk() (*[]WidgetEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HeatMapWidgetDefinition) SetEvents(v []WidgetEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HeatMapWidgetDefinition) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLegendSize

`func (o *HeatMapWidgetDefinition) GetLegendSize() WidgetLegendSize`

GetLegendSize returns the LegendSize field if non-nil, zero value otherwise.

### GetLegendSizeOk

`func (o *HeatMapWidgetDefinition) GetLegendSizeOk() (*WidgetLegendSize, bool)`

GetLegendSizeOk returns a tuple with the LegendSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendSize

`func (o *HeatMapWidgetDefinition) SetLegendSize(v WidgetLegendSize)`

SetLegendSize sets LegendSize field to given value.

### HasLegendSize

`func (o *HeatMapWidgetDefinition) HasLegendSize() bool`

HasLegendSize returns a boolean if a field has been set.

### GetRequests

`func (o *HeatMapWidgetDefinition) GetRequests() []HeatMapWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *HeatMapWidgetDefinition) GetRequestsOk() (*[]HeatMapWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *HeatMapWidgetDefinition) SetRequests(v []HeatMapWidgetRequest)`

SetRequests sets Requests field to given value.


### GetShowLegend

`func (o *HeatMapWidgetDefinition) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *HeatMapWidgetDefinition) GetShowLegendOk() (*bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegend

`func (o *HeatMapWidgetDefinition) SetShowLegend(v bool)`

SetShowLegend sets ShowLegend field to given value.

### HasShowLegend

`func (o *HeatMapWidgetDefinition) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### GetTime

`func (o *HeatMapWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *HeatMapWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *HeatMapWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *HeatMapWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *HeatMapWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HeatMapWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HeatMapWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HeatMapWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *HeatMapWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *HeatMapWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *HeatMapWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *HeatMapWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *HeatMapWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *HeatMapWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *HeatMapWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *HeatMapWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *HeatMapWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeatMapWidgetDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeatMapWidgetDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetYaxis

`func (o *HeatMapWidgetDefinition) GetYaxis() WidgetAxis`

GetYaxis returns the Yaxis field if non-nil, zero value otherwise.

### GetYaxisOk

`func (o *HeatMapWidgetDefinition) GetYaxisOk() (*WidgetAxis, bool)`

GetYaxisOk returns a tuple with the Yaxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaxis

`func (o *HeatMapWidgetDefinition) SetYaxis(v WidgetAxis)`

SetYaxis sets Yaxis field to given value.

### HasYaxis

`func (o *HeatMapWidgetDefinition) HasYaxis() bool`

HasYaxis returns a boolean if a field has been set.


### AsWidgetDefinition

`func (s *HeatMapWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of HeatMapWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


