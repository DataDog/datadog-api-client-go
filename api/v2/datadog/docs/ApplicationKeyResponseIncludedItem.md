# ApplicationKeyResponseIncludedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**UserAttributes**](UserAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the user. | [optional] 
**Relationships** | Pointer to [**UserResponseRelationships**](UserResponseRelationships.md) |  | [optional] 
**Type** | Pointer to [**UsersType**](UsersType.md) |  | [optional] [default to "users"]

## Methods

### NewApplicationKeyResponseIncludedItem

`func NewApplicationKeyResponseIncludedItem() *ApplicationKeyResponseIncludedItem`

NewApplicationKeyResponseIncludedItem instantiates a new ApplicationKeyResponseIncludedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationKeyResponseIncludedItemWithDefaults

`func NewApplicationKeyResponseIncludedItemWithDefaults() *ApplicationKeyResponseIncludedItem`

NewApplicationKeyResponseIncludedItemWithDefaults instantiates a new ApplicationKeyResponseIncludedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ApplicationKeyResponseIncludedItem) GetAttributes() UserAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationKeyResponseIncludedItem) GetAttributesOk() (*UserAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationKeyResponseIncludedItem) SetAttributes(v UserAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplicationKeyResponseIncludedItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *ApplicationKeyResponseIncludedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationKeyResponseIncludedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationKeyResponseIncludedItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationKeyResponseIncludedItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *ApplicationKeyResponseIncludedItem) GetRelationships() UserResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ApplicationKeyResponseIncludedItem) GetRelationshipsOk() (*UserResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ApplicationKeyResponseIncludedItem) SetRelationships(v UserResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ApplicationKeyResponseIncludedItem) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *ApplicationKeyResponseIncludedItem) GetType() UsersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationKeyResponseIncludedItem) GetTypeOk() (*UsersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationKeyResponseIncludedItem) SetType(v UsersType)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationKeyResponseIncludedItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


