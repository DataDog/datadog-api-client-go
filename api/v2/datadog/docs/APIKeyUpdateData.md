# APIKeyUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**APIKeyUpdateAttributes**](APIKeyUpdateAttributes.md) |  | 
**Id** | **string** | ID of the API key. | 
**Type** | [**APIKeysType**](APIKeysType.md) |  | [default to APIKEYSTYPE_API_KEYS]

## Methods

### NewAPIKeyUpdateData

`func NewAPIKeyUpdateData(attributes APIKeyUpdateAttributes, id string, type_ APIKeysType, ) *APIKeyUpdateData`

NewAPIKeyUpdateData instantiates a new APIKeyUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyUpdateDataWithDefaults

`func NewAPIKeyUpdateDataWithDefaults() *APIKeyUpdateData`

NewAPIKeyUpdateDataWithDefaults instantiates a new APIKeyUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *APIKeyUpdateData) GetAttributes() APIKeyUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *APIKeyUpdateData) GetAttributesOk() (*APIKeyUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *APIKeyUpdateData) SetAttributes(v APIKeyUpdateAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *APIKeyUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIKeyUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIKeyUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *APIKeyUpdateData) GetType() APIKeysType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *APIKeyUpdateData) GetTypeOk() (*APIKeysType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *APIKeyUpdateData) SetType(v APIKeysType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


