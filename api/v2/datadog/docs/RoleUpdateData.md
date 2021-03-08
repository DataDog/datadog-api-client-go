# RoleUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**RoleUpdateAttributes**](RoleUpdateAttributes.md) |  | 
**Id** | **string** | ID of the role. | 
**Type** | [**RolesType**](RolesType.md) |  | [default to ROLESTYPE_ROLES]

## Methods

### NewRoleUpdateData

`func NewRoleUpdateData(attributes RoleUpdateAttributes, id string, type_ RolesType, ) *RoleUpdateData`

NewRoleUpdateData instantiates a new RoleUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleUpdateDataWithDefaults

`func NewRoleUpdateDataWithDefaults() *RoleUpdateData`

NewRoleUpdateDataWithDefaults instantiates a new RoleUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *RoleUpdateData) GetAttributes() RoleUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RoleUpdateData) GetAttributesOk() (*RoleUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RoleUpdateData) SetAttributes(v RoleUpdateAttributes)`

SetAttributes sets Attributes field to given value.


### GetId

`func (o *RoleUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RoleUpdateData) GetType() RolesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleUpdateData) GetTypeOk() (*RolesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleUpdateData) SetType(v RolesType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


