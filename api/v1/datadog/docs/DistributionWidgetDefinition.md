# DistributionWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegendSize** | Pointer to [**WidgetLegendSize**](WidgetLegendSize.md) |  | [optional] 
**Requests** | Pointer to [**[]DistributionWidgetRequest**](DistributionWidgetRequest.md) |  | 
**ShowLegend** | Pointer to **bool** | Whether or not to display the legend on this widget | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "distribution"]

## Methods

### NewDistributionWidgetDefinition

`func NewDistributionWidgetDefinition(requests []DistributionWidgetRequest, type_ string, ) *DistributionWidgetDefinition`

NewDistributionWidgetDefinition instantiates a new DistributionWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionWidgetDefinitionWithDefaults

`func NewDistributionWidgetDefinitionWithDefaults() *DistributionWidgetDefinition`

NewDistributionWidgetDefinitionWithDefaults instantiates a new DistributionWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegendSize

`func (o *DistributionWidgetDefinition) GetLegendSize() WidgetLegendSize`

GetLegendSize returns the LegendSize field if non-nil, zero value otherwise.

### GetLegendSizeOk

`func (o *DistributionWidgetDefinition) GetLegendSizeOk() (WidgetLegendSize, bool)`

GetLegendSizeOk returns a tuple with the LegendSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegendSize

`func (o *DistributionWidgetDefinition) HasLegendSize() bool`

HasLegendSize returns a boolean if a field has been set.

### SetLegendSize

`func (o *DistributionWidgetDefinition) SetLegendSize(v WidgetLegendSize)`

SetLegendSize gets a reference to the given WidgetLegendSize and assigns it to the LegendSize field.

### GetRequests

`func (o *DistributionWidgetDefinition) GetRequests() []DistributionWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *DistributionWidgetDefinition) GetRequestsOk() ([]DistributionWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequests

`func (o *DistributionWidgetDefinition) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### SetRequests

`func (o *DistributionWidgetDefinition) SetRequests(v []DistributionWidgetRequest)`

SetRequests gets a reference to the given []DistributionWidgetRequest and assigns it to the Requests field.

### GetShowLegend

`func (o *DistributionWidgetDefinition) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *DistributionWidgetDefinition) GetShowLegendOk() (bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowLegend

`func (o *DistributionWidgetDefinition) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### SetShowLegend

`func (o *DistributionWidgetDefinition) SetShowLegend(v bool)`

SetShowLegend gets a reference to the given bool and assigns it to the ShowLegend field.

### GetTime

`func (o *DistributionWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DistributionWidgetDefinition) GetTimeOk() (WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *DistributionWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *DistributionWidgetDefinition) SetTime(v WidgetTime)`

SetTime gets a reference to the given WidgetTime and assigns it to the Time field.

### GetTitle

`func (o *DistributionWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DistributionWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *DistributionWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *DistributionWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *DistributionWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *DistributionWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *DistributionWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *DistributionWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *DistributionWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *DistributionWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *DistributionWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *DistributionWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *DistributionWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DistributionWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *DistributionWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *DistributionWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


### AsWidgetDefinition

`func (s *DistributionWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of DistributionWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


