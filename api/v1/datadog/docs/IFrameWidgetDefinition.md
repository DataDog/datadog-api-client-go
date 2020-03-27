# IFrameWidgetDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the widget | [readonly] [default to "iframe"]
**Url** | Pointer to **string** | URL of the iframe | 

## Methods

### NewIFrameWidgetDefinition

`func NewIFrameWidgetDefinition(type_ string, url string, ) *IFrameWidgetDefinition`

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

`func (o *IFrameWidgetDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IFrameWidgetDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IFrameWidgetDefinition) SetType(v string)`

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



### AsWidgetDefinition

`func (s *IFrameWidgetDefinition) AsWidgetDefinition() WidgetDefinition`

Convenience method to wrap this instance of IFrameWidgetDefinition in WidgetDefinition

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


