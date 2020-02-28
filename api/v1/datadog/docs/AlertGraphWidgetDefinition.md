# AlertGraphWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | Pointer to **string** | ID of the alert to use in the widget | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | The title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "alert_graph"]
**VizType** | Pointer to [**WidgetVizType**](WidgetVizType.md) |  | 

## Methods

### NewAlertGraphWidgetDefinition

`func NewAlertGraphWidgetDefinition(alertId string, type_ string, vizType WidgetVizType, ) *AlertGraphWidgetDefinition`

NewAlertGraphWidgetDefinition instantiates a new AlertGraphWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGraphWidgetDefinitionWithDefaults

`func NewAlertGraphWidgetDefinitionWithDefaults() *AlertGraphWidgetDefinition`

NewAlertGraphWidgetDefinitionWithDefaults instantiates a new AlertGraphWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *AlertGraphWidgetDefinition) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *AlertGraphWidgetDefinition) GetAlertIdOk() (string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlertId

`func (o *AlertGraphWidgetDefinition) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### SetAlertId

`func (o *AlertGraphWidgetDefinition) SetAlertId(v string)`

SetAlertId gets a reference to the given string and assigns it to the AlertId field.

### GetTime

`func (o *AlertGraphWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AlertGraphWidgetDefinition) GetTimeOk() (WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *AlertGraphWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *AlertGraphWidgetDefinition) SetTime(v WidgetTime)`

SetTime gets a reference to the given WidgetTime and assigns it to the Time field.

### GetTitle

`func (o *AlertGraphWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertGraphWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *AlertGraphWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *AlertGraphWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *AlertGraphWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *AlertGraphWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *AlertGraphWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *AlertGraphWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *AlertGraphWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *AlertGraphWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *AlertGraphWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *AlertGraphWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *AlertGraphWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertGraphWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *AlertGraphWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *AlertGraphWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetVizType

`func (o *AlertGraphWidgetDefinition) GetVizType() WidgetVizType`

GetVizType returns the VizType field if non-nil, zero value otherwise.

### GetVizTypeOk

`func (o *AlertGraphWidgetDefinition) GetVizTypeOk() (WidgetVizType, bool)`

GetVizTypeOk returns a tuple with the VizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVizType

`func (o *AlertGraphWidgetDefinition) HasVizType() bool`

HasVizType returns a boolean if a field has been set.

### SetVizType

`func (o *AlertGraphWidgetDefinition) SetVizType(v WidgetVizType)`

SetVizType gets a reference to the given WidgetVizType and assigns it to the VizType field.


### AsWidgetDefinition

`func (s *AlertGraphWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of AlertGraphWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


