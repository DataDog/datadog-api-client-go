# ImageWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Margin** | Pointer to [**WidgetMargin**](WidgetMargin.md) |  | [optional] 
**Sizing** | Pointer to [**WidgetImageSizing**](WidgetImageSizing.md) |  | [optional] 
**Type** | [**ImageWidgetDefinitionType**](ImageWidgetDefinitionType.md) |  | [default to "image"]
**Url** | **string** | URL of the image. | 

## Methods

### NewImageWidgetDefinition

`func NewImageWidgetDefinition(type_ ImageWidgetDefinitionType, url string, ) *ImageWidgetDefinition`

NewImageWidgetDefinition instantiates a new ImageWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWidgetDefinitionWithDefaults

`func NewImageWidgetDefinitionWithDefaults() *ImageWidgetDefinition`

NewImageWidgetDefinitionWithDefaults instantiates a new ImageWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMargin

`func (o *ImageWidgetDefinition) GetMargin() WidgetMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *ImageWidgetDefinition) GetMarginOk() (*WidgetMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *ImageWidgetDefinition) SetMargin(v WidgetMargin)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *ImageWidgetDefinition) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetSizing

`func (o *ImageWidgetDefinition) GetSizing() WidgetImageSizing`

GetSizing returns the Sizing field if non-nil, zero value otherwise.

### GetSizingOk

`func (o *ImageWidgetDefinition) GetSizingOk() (*WidgetImageSizing, bool)`

GetSizingOk returns a tuple with the Sizing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizing

`func (o *ImageWidgetDefinition) SetSizing(v WidgetImageSizing)`

SetSizing sets Sizing field to given value.

### HasSizing

`func (o *ImageWidgetDefinition) HasSizing() bool`

HasSizing returns a boolean if a field has been set.

### GetType

`func (o *ImageWidgetDefinition) GetType() ImageWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageWidgetDefinition) GetTypeOk() (*ImageWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageWidgetDefinition) SetType(v ImageWidgetDefinitionType)`

SetType sets Type field to given value.


### GetUrl

`func (o *ImageWidgetDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageWidgetDefinition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageWidgetDefinition) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


