# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**RoleAttributes**](RoleAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | Role ID. | [optional] 
**Relationships** | Pointer to [**RoleRelationships**](RoleRelationships.md) |  | [optional] 
**Type** | Pointer to **string** | TODO | [optional] [default to "roles"]

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Role) GetAttributes() RoleAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Role) GetAttributesOk() (*RoleAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Role) SetAttributes(v RoleAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Role) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Role) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *Role) GetRelationships() RoleRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Role) GetRelationshipsOk() (*RoleRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Role) SetRelationships(v RoleRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Role) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *Role) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Role) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Role) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Role) HasType() bool`

HasType returns a boolean if a field has been set.


### AsUserResponseIncludedItem

`func (s *Role) AsUserResponseIncludedItem() UserResponseIncludedItem`

Convenience method to wrap this instance of Role in UserResponseIncludedItem

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


