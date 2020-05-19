# AlertGraphWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertId** | Pointer to **string** | ID of the alert to use in the widget. | 
**Time** | Pointer to [**WidgetTime**](WidgetTime.md) |  | [optional] 
**Title** | Pointer to **string** | The title of the widget. | [optional] 
**TitleAlign** | Pointer to [**WidgetTextAlign**](WidgetTextAlign.md) |  | [optional] 
**TitleSize** | Pointer to **string** | Size of the title. | [optional] 
**Type** | Pointer to [**AlertGraphWidgetDefinitionType**](AlertGraphWidgetDefinitionType.md) |  | [default to "alert_graph"]
**VizType** | Pointer to [**WidgetVizType**](WidgetVizType.md) |  | 

## Methods

### NewAlertGraphWidgetDefinition

`func NewAlertGraphWidgetDefinition(alertId string, type_ AlertGraphWidgetDefinitionType, vizType WidgetVizType, ) *AlertGraphWidgetDefinition`

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

`func (o *AlertGraphWidgetDefinition) GetAlertIdOk() (*string, bool)`

GetAlertIdOk returns a tuple with the AlertId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertId

`func (o *AlertGraphWidgetDefinition) SetAlertId(v string)`

SetAlertId sets AlertId field to given value.


### GetTime

`func (o *AlertGraphWidgetDefinition) GetTime() WidgetTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AlertGraphWidgetDefinition) GetTimeOk() (*WidgetTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AlertGraphWidgetDefinition) SetTime(v WidgetTime)`

SetTime sets Time field to given value.

### HasTime

`func (o *AlertGraphWidgetDefinition) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitle

`func (o *AlertGraphWidgetDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertGraphWidgetDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertGraphWidgetDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlertGraphWidgetDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleAlign

`func (o *AlertGraphWidgetDefinition) GetTitleAlign() WidgetTextAlign`

GetTitleAlign returns the TitleAlign field if non-nil, zero value otherwise.

### GetTitleAlignOk

`func (o *AlertGraphWidgetDefinition) GetTitleAlignOk() (*WidgetTextAlign, bool)`

GetTitleAlignOk returns a tuple with the TitleAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleAlign

`func (o *AlertGraphWidgetDefinition) SetTitleAlign(v WidgetTextAlign)`

SetTitleAlign sets TitleAlign field to given value.

### HasTitleAlign

`func (o *AlertGraphWidgetDefinition) HasTitleAlign() bool`

HasTitleAlign returns a boolean if a field has been set.

### GetTitleSize

`func (o *AlertGraphWidgetDefinition) GetTitleSize() string`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *AlertGraphWidgetDefinition) GetTitleSizeOk() (*string, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *AlertGraphWidgetDefinition) SetTitleSize(v string)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *AlertGraphWidgetDefinition) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetType

`func (o *AlertGraphWidgetDefinition) GetType() AlertGraphWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertGraphWidgetDefinition) GetTypeOk() (*AlertGraphWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertGraphWidgetDefinition) SetType(v AlertGraphWidgetDefinitionType)`

SetType sets Type field to given value.


### GetVizType

`func (o *AlertGraphWidgetDefinition) GetVizType() WidgetVizType`

GetVizType returns the VizType field if non-nil, zero value otherwise.

### GetVizTypeOk

`func (o *AlertGraphWidgetDefinition) GetVizTypeOk() (*WidgetVizType, bool)`

GetVizTypeOk returns a tuple with the VizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizType

`func (o *AlertGraphWidgetDefinition) SetVizType(v WidgetVizType)`

SetVizType sets VizType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


