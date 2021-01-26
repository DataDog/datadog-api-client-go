# APIKeyResponseIncludedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**UserAttributes**](UserAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the user. | [optional] 
**Relationships** | Pointer to [**UserResponseRelationships**](UserResponseRelationships.md) |  | [optional] 
**Type** | Pointer to [**UsersType**](UsersType.md) |  | [optional] [default to "users"]

## Methods

### NewAPIKeyResponseIncludedItem

`func NewAPIKeyResponseIncludedItem() *APIKeyResponseIncludedItem`

NewAPIKeyResponseIncludedItem instantiates a new APIKeyResponseIncludedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyResponseIncludedItemWithDefaults

`func NewAPIKeyResponseIncludedItemWithDefaults() *APIKeyResponseIncludedItem`

NewAPIKeyResponseIncludedItemWithDefaults instantiates a new APIKeyResponseIncludedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *APIKeyResponseIncludedItem) GetAttributes() UserAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *APIKeyResponseIncludedItem) GetAttributesOk() (*UserAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *APIKeyResponseIncludedItem) SetAttributes(v UserAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *APIKeyResponseIncludedItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *APIKeyResponseIncludedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIKeyResponseIncludedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIKeyResponseIncludedItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIKeyResponseIncludedItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *APIKeyResponseIncludedItem) GetRelationships() UserResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *APIKeyResponseIncludedItem) GetRelationshipsOk() (*UserResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *APIKeyResponseIncludedItem) SetRelationships(v UserResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *APIKeyResponseIncludedItem) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *APIKeyResponseIncludedItem) GetType() UsersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *APIKeyResponseIncludedItem) GetTypeOk() (*UsersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *APIKeyResponseIncludedItem) SetType(v UsersType)`

SetType sets Type field to given value.

### HasType

`func (o *APIKeyResponseIncludedItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


