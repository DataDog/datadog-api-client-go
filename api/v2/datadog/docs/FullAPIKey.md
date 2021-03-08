# FullAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**FullAPIKeyAttributes**](FullAPIKeyAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the API key. | [optional] 
**Relationships** | Pointer to [**APIKeyRelationships**](APIKeyRelationships.md) |  | [optional] 
**Type** | Pointer to [**APIKeysType**](APIKeysType.md) |  | [optional] [default to APIKEYSTYPE_API_KEYS]

## Methods

### NewFullAPIKey

`func NewFullAPIKey() *FullAPIKey`

NewFullAPIKey instantiates a new FullAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullAPIKeyWithDefaults

`func NewFullAPIKeyWithDefaults() *FullAPIKey`

NewFullAPIKeyWithDefaults instantiates a new FullAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *FullAPIKey) GetAttributes() FullAPIKeyAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FullAPIKey) GetAttributesOk() (*FullAPIKeyAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FullAPIKey) SetAttributes(v FullAPIKeyAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FullAPIKey) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *FullAPIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullAPIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullAPIKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FullAPIKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *FullAPIKey) GetRelationships() APIKeyRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FullAPIKey) GetRelationshipsOk() (*APIKeyRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FullAPIKey) SetRelationships(v APIKeyRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FullAPIKey) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *FullAPIKey) GetType() APIKeysType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FullAPIKey) GetTypeOk() (*APIKeysType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FullAPIKey) SetType(v APIKeysType)`

SetType sets Type field to given value.

### HasType

`func (o *FullAPIKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


