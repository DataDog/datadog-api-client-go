# RoleUpdateResponseData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**RoleUpdateAttributes**](RoleUpdateAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the role. | [optional] 
**Relationships** | Pointer to [**RoleResponseRelationships**](RoleResponseRelationships.md) |  | [optional] 
**Type** | [**RolesType**](RolesType.md) |  | [default to ROLESTYPE_ROLES]

## Methods

### NewRoleUpdateResponseData

`func NewRoleUpdateResponseData(type_ RolesType, ) *RoleUpdateResponseData`

NewRoleUpdateResponseData instantiates a new RoleUpdateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleUpdateResponseDataWithDefaults

`func NewRoleUpdateResponseDataWithDefaults() *RoleUpdateResponseData`

NewRoleUpdateResponseDataWithDefaults instantiates a new RoleUpdateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *RoleUpdateResponseData) GetAttributes() RoleUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RoleUpdateResponseData) GetAttributesOk() (*RoleUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RoleUpdateResponseData) SetAttributes(v RoleUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RoleUpdateResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *RoleUpdateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleUpdateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleUpdateResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleUpdateResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *RoleUpdateResponseData) GetRelationships() RoleResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RoleUpdateResponseData) GetRelationshipsOk() (*RoleResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RoleUpdateResponseData) SetRelationships(v RoleResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *RoleUpdateResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *RoleUpdateResponseData) GetType() RolesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleUpdateResponseData) GetTypeOk() (*RolesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleUpdateResponseData) SetType(v RolesType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


