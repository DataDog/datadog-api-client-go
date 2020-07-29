# IFrameWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**IFrameWidgetDefinitionType**](IFrameWidgetDefinitionType.md) |  | [default to "iframe"]
**Url** | **string** | URL of the iframe. | 

## Methods

### NewIFrameWidgetDefinition

`func NewIFrameWidgetDefinition(type_ IFrameWidgetDefinitionType, url string, ) *IFrameWidgetDefinition`

NewIFrameWidgetDefinition instantiates a new IFrameWidgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIFrameWidgetDefinitionWithDefaults

`func NewIFrameWidgetDefinitionWithDefaults() *IFrameWidgetDefinition`

NewIFrameWidgetDefinitionWithDefaults instantiates a new IFrameWidgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IFrameWidgetDefinition) GetType() IFrameWidgetDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IFrameWidgetDefinition) GetTypeOk() (*IFrameWidgetDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IFrameWidgetDefinition) SetType(v IFrameWidgetDefinitionType)`

SetType sets Type field to given value.


### GetUrl

`func (o *IFrameWidgetDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IFrameWidgetDefinition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IFrameWidgetDefinition) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


