# AlertValueWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | Pointer to **string** | ID of the alert to use in the widget | 
**Precision** | Pointer to **int64** | Number of decimal to show. If not defined, will use the raw value | [optional] 
**TextAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**Title** | Pointer to **string** | Title of the widget | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of value in the widget | [optional] 
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "alert_value"]
**Unit** | Pointer to **string** | Unit to display with the value | [optional] 

## Methods

### NewAlertValueWidgetDefinition

`func NewAlertValueWidgetDefinition(alertId string, type_ string, ) *AlertValueWidgetDefinition`

NewAlertValueWidgetDefinition instantiates a new AlertValueWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertValueWidgetDefinitionWithDefaults

`func NewAlertValueWidgetDefinitionWithDefaults() *AlertValueWidgetDefinition`

NewAlertValueWidgetDefinitionWithDefaults instantiates a new AlertValueWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertId

`func (o *AlertValueWidgetDefinition) GetAlertId() string`

GetAlertId returns the AlertId field if non-nil, zero value otherwise.

### GetAlertIdOk

`func (o *AlertValueWidgetDefinition) GetAlertIdOk() (string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlertId

`func (o *AlertValueWidgetDefinition) HasAlertId() bool`

HasAlertId returns a boolean if a field has been set.

### SetAlertId

`func (o *AlertValueWidgetDefinition) SetAlertId(v string)`

SetAlertId gets a reference to the given string and assigns it to the AlertId field.

### GetPrecision

`func (o *AlertValueWidgetDefinition) GetPrecision() int64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *AlertValueWidgetDefinition) GetPrecisionOk() (int64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrecision

`func (o *AlertValueWidgetDefinition) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### SetPrecision

`func (o *AlertValueWidgetDefinition) SetPrecision(v int64)`

SetPrecision gets a reference to the given int64 and assigns it to the Precision field.

### GetTextAlign

`func (o *AlertValueWidgetDefinition) GetTextAlign() WidgetTextAlign`

GetTextAlign returns the TextAlign field if non-nil, zero value otherwise.

### GetTextAlignOk

`func (o *AlertValueWidgetDefinition) GetTextAlignOk() (WidgetTextAlign, bool)`

GetTextAlignOk returns a tuple with the TextAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTextAlign

`func (o *AlertValueWidgetDefinition) HasTextAlign() bool`

HasTextAlign returns a boolean if a field has been set.

### SetTextAlign

`func (o *AlertValueWidgetDefinition) SetTextAlign(v WidgetTextAlign)`

SetTextAlign gets a reference to the given WidgetTextAlign and assigns it to the TextAlign field.

### GetTitle

`func (o *AlertValueWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertValueWidgetDefinition) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *AlertValueWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *AlertValueWidgetDefinition) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetTitleAlign

`func (o *AlertValueWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *AlertValueWidgetDefinition) GetTitleAlignOk() (WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleAlign

`func (o *AlertValueWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### SetTitleAlign

`func (o *AlertValueWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign gets a reference to the given WidgetTextAlign and assigns it to the TitleAlign field.

### GetTitleSize

`func (o *AlertValueWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *AlertValueWidgetDefinition) GetTitleSizeOk() (string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitleSize

`func (o *AlertValueWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### SetTitleSize

`func (o *AlertValueWidgetDefinition) SetTitleSize(v string)`

SetTitleSize gets a reference to the given string and assigns it to the TitleSize field.

### GetType

`func (o *AlertValueWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertValueWidgetDefinition) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *AlertValueWidgetDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *AlertValueWidgetDefinition) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetUnit

`func (o *AlertValueWidgetDefinition) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AlertValueWidgetDefinition) GetUnitOk() (string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnit

`func (o *AlertValueWidgetDefinition) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnit

`func (o *AlertValueWidgetDefinition) SetUnit(v string)`

SetUnit gets a reference to the given string and assigns it to the Unit field.


### AsWidgetDefinition

`func (s *AlertValueWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of AlertValueWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


