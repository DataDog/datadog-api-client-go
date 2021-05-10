# APIKeyCreateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**APIKeyCreateAttributes**](APIKeyCreateAttributes.md) |  | 
**Type** | [**APIKeysType**](APIKeysType.md) |  | [default to APIKEYSTYPE_API_KEYS]

## Methods

### NewAPIKeyCreateData

`func NewAPIKeyCreateData(attributes APIKeyCreateAttributes, type_ APIKeysType, ) *APIKeyCreateData`

NewAPIKeyCreateData instantiates a new APIKeyCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyCreateDataWithDefaults

`func NewAPIKeyCreateDataWithDefaults() *APIKeyCreateData`

NewAPIKeyCreateDataWithDefaults instantiates a new APIKeyCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *APIKeyCreateData) GetAttributes() APIKeyCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *APIKeyCreateData) GetAttributesOk() (*APIKeyCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *APIKeyCreateData) SetAttributes(v APIKeyCreateAttributes)`

SetAttributes sets Attributes field to given value.


### GetType

`func (o *APIKeyCreateData) GetType() APIKeysType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *APIKeyCreateData) GetTypeOk() (*APIKeysType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *APIKeyCreateData) SetType(v APIKeysType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


