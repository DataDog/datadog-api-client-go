# QueryValueWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoscale** | Pointer to **bool** | Whether to use autoscaling or not. | [optional] 
**CustomUnit** | Pointer to **string** | Display a unit of your choice on the widget. | [optional] 
**Precision** | Pointer to **int64** | Number of decimals to show. If not defined, the widget uses the raw value. | [optional] 
**Requests** | Pointer to [**[]QueryValueWidgetRequest**](QueryValueWidgetRequest.md) |  | 
**TextAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | Title of your widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "query_value"]

## Methods

### NewQueryValueWidgetDefinition

`func NewQueryValueWidgetDefinition(requests []QueryValueWidgetRequest, type_ string, ) *QueryValueWidgetDefinition`

NewQueryValueWidgetDefinition instantiates a new QueryValueWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryValueWidgetDefinitionWithDefaults

`func NewQueryValueWidgetDefinitionWithDefaults() *QueryValueWidgetDefinition`

NewQueryValueWidgetDefinitionWithDefaults instantiates a new QueryValueWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscale

`func (o *QueryValueWidgetDefinition) GetAutoscale() bool`

GetAutoscale returns the Autoscale field if non-nil, zero value otherwise.

### GetAutoscaleOk

`func (o *QueryValueWidgetDefinition) GetAutoscaleOk() (*bool, bool)`

GetAutoscaleOk returns a tuple with the Autoscale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscale

`func (o *QueryValueWidgetDefinition) SetAutoscale(v bool)`

SetAutoscale sets Autoscale field to given value.

### HasAutoscale

`func (o *QueryValueWidgetDefinition) HasAutoscale() bool`

HasAutoscale returns a boolean if a field has been set.

### GetCustomUnit

`func (o *QueryValueWidgetDefinition) GetCustomUnit() string`

GetCustomUnit returns the CustomUnit field if non-nil, zero value otherwise.

### GetCustomUnitOk

`func (o *QueryValueWidgetDefinition) GetCustomUnitOk() (*string, bool)`

GetCustomUnitOk returns a tuple with the CustomUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUnit

`func (o *QueryValueWidgetDefinition) SetCustomUnit(v string)`

SetCustomUnit sets CustomUnit field to given value.

### HasCustomUnit

`func (o *QueryValueWidgetDefinition) HasCustomUnit() bool`

HasCustomUnit returns a boolean if a field has been set.

### GetPrecision

`func (o *QueryValueWidgetDefinition) GetPrecision() int64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *QueryValueWidgetDefinition) GetPrecisionOk() (*int64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *QueryValueWidgetDefinition) SetPrecision(v int64)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *QueryValueWidgetDefinition) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetRequests

`func (o *QueryValueWidgetDefinition) GetRequests() []QueryValueWidgetRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *QueryValueWidgetDefinition) GetRequestsOk() (*[]QueryValueWidgetRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *QueryValueWidgetDefinition) SetRequests(v []QueryValueWidgetRequest)`

SetRequests sets Requests field to given value.


### GetTextAlign

`func (o *QueryValueWidgetDefinition) GetTextAlign() WidgetTextAlign`

GetTextAlign returns the TextAlign field if non-nil, zero value otherwise.

### GetTextAlignOk

`func (o *QueryValueWidgetDefinition) GetTextAlignOk() (*WidgetTextAlign, bool)`

GetTextAlignOk returns a tuple with the TextAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAlign

`func (o *QueryValueWidgetDefinition) SetTextAlign(v WidgetTextAlign)`

SetTextAlign sets TextAlign field to given value.

### HasTextAlign

`func (o *QueryValueWidgetDefinition) HasTextAlign() bool`

HasTextAlign returns a boolean if a field has been set.

### GetTime

`func (o *QueryValueWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueryValueWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueryValueWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *QueryValueWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *QueryValueWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QueryValueWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QueryValueWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QueryValueWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *QueryValueWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *QueryValueWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *QueryValueWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *QueryValueWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *QueryValueWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *QueryValueWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *QueryValueWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *QueryValueWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *QueryValueWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryValueWidgetDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryValueWidgetDefinition) SetType(v string)`

SetType sets Type field to given value.



### AsWidgetDefinition

`func (s *QueryValueWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of QueryValueWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


