# UserResponseIncludedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**RoleAttributes**](RoleAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the role. | [optional] 
**Type** | [**RolesType**](RolesType.md) |  | [default to ROLESTYPE_ROLES]
**Relationships** | Pointer to [**RoleResponseRelationships**](RoleResponseRelationships.md) |  | [optional] 

## Methods

### NewUserResponseIncludedItem

`func NewUserResponseIncludedItem(type_ RolesType, ) *UserResponseIncludedItem`

NewUserResponseIncludedItem instantiates a new UserResponseIncludedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseIncludedItemWithDefaults

`func NewUserResponseIncludedItemWithDefaults() *UserResponseIncludedItem`

NewUserResponseIncludedItemWithDefaults instantiates a new UserResponseIncludedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UserResponseIncludedItem) GetAttributes() RoleAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserResponseIncludedItem) GetAttributesOk() (*RoleAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserResponseIncludedItem) SetAttributes(v RoleAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserResponseIncludedItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *UserResponseIncludedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponseIncludedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponseIncludedItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserResponseIncludedItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserResponseIncludedItem) GetType() RolesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserResponseIncludedItem) GetTypeOk() (*RolesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserResponseIncludedItem) SetType(v RolesType)`

SetType sets Type field to given value.


### GetRelationships

`func (o *UserResponseIncludedItem) GetRelationships() RoleResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UserResponseIncludedItem) GetRelationshipsOk() (*RoleResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UserResponseIncludedItem) SetRelationships(v RoleResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UserResponseIncludedItem) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


