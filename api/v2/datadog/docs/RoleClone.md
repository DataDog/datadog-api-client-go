# RoleClone

## Properties

| Name           | Type                                              | Description | Notes                        |
| -------------- | ------------------------------------------------- | ----------- | ---------------------------- |
| **Attributes** | [**RoleCloneAttributes**](RoleCloneAttributes.md) |             |
| **Type**       | [**RolesType**](RolesType.md)                     |             | [default to ROLESTYPE_ROLES] |

## Methods

### NewRoleClone

`func NewRoleClone(attributes RoleCloneAttributes, type_ RolesType) *RoleClone`

NewRoleClone instantiates a new RoleClone object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRoleCloneWithDefaults

`func NewRoleCloneWithDefaults() *RoleClone`

NewRoleCloneWithDefaults instantiates a new RoleClone object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *RoleClone) GetAttributes() RoleCloneAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RoleClone) GetAttributesOk() (*RoleCloneAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RoleClone) SetAttributes(v RoleCloneAttributes)`

SetAttributes sets Attributes field to given value.

### GetType

`func (o *RoleClone) GetType() RolesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleClone) GetTypeOk() (*RolesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleClone) SetType(v RolesType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
