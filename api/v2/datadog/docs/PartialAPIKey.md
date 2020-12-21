# PartialAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**PartialAPIKeyAttributes**](PartialAPIKeyAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the API key. | [optional] 
**Relationships** | Pointer to [**APIKeyRelationships**](APIKeyRelationships.md) |  | [optional] 
**Type** | Pointer to [**APIKeysType**](APIKeysType.md) |  | [optional] [default to "api_keys"]

## Methods

### NewPartialAPIKey

`func NewPartialAPIKey() *PartialAPIKey`

NewPartialAPIKey instantiates a new PartialAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialAPIKeyWithDefaults

`func NewPartialAPIKeyWithDefaults() *PartialAPIKey`

NewPartialAPIKeyWithDefaults instantiates a new PartialAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PartialAPIKey) GetAttributes() PartialAPIKeyAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PartialAPIKey) GetAttributesOk() (*PartialAPIKeyAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PartialAPIKey) SetAttributes(v PartialAPIKeyAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PartialAPIKey) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *PartialAPIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartialAPIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartialAPIKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartialAPIKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *PartialAPIKey) GetRelationships() APIKeyRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PartialAPIKey) GetRelationshipsOk() (*APIKeyRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PartialAPIKey) SetRelationships(v APIKeyRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PartialAPIKey) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *PartialAPIKey) GetType() APIKeysType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartialAPIKey) GetTypeOk() (*APIKeysType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartialAPIKey) SetType(v APIKeysType)`

SetType sets Type field to given value.

### HasType

`func (o *PartialAPIKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


