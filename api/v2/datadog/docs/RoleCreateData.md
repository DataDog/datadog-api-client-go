# RoleCreateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | [**RoleCreateAttributes**](RoleCreateAttributes.md) |  | 
**Relationships** | Pointer to [**RoleRelationships**](RoleRelationships.md) |  | [optional] 
**Type** | Pointer to [**RolesType**](RolesType.md) |  | [optional] [default to ROLESTYPE_ROLES]

## Methods

### NewRoleCreateData

`func NewRoleCreateData(attributes RoleCreateAttributes) *RoleCreateData`

NewRoleCreateData instantiates a new RoleCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRoleCreateDataWithDefaults

`func NewRoleCreateDataWithDefaults() *RoleCreateData`

NewRoleCreateDataWithDefaults instantiates a new RoleCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *RoleCreateData) GetAttributes() RoleCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RoleCreateData) GetAttributesOk() (*RoleCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RoleCreateData) SetAttributes(v RoleCreateAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *RoleCreateData) GetRelationships() RoleRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RoleCreateData) GetRelationshipsOk() (*RoleRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RoleCreateData) SetRelationships(v RoleRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *RoleCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *RoleCreateData) GetType() RolesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleCreateData) GetTypeOk() (*RolesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleCreateData) SetType(v RolesType)`

SetType sets Type field to given value.

### HasType

`func (o *RoleCreateData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


