# RoleCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**RoleCreateAttributes**](RoleCreateAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the role. | [optional] 
**Relationships** | Pointer to [**RoleResponseRelationships**](RoleResponseRelationships.md) |  | [optional] 
**Type** | [**RolesType**](RolesType.md) |  | [default to ROLESTYPE_ROLES]

## Methods

### NewRoleCreateResponseData

`func NewRoleCreateResponseData(type_ RolesType, ) *RoleCreateResponseData`

NewRoleCreateResponseData instantiates a new RoleCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleCreateResponseDataWithDefaults

`func NewRoleCreateResponseDataWithDefaults() *RoleCreateResponseData`

NewRoleCreateResponseDataWithDefaults instantiates a new RoleCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *RoleCreateResponseData) GetAttributes() RoleCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RoleCreateResponseData) GetAttributesOk() (*RoleCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RoleCreateResponseData) SetAttributes(v RoleCreateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RoleCreateResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *RoleCreateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleCreateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleCreateResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleCreateResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *RoleCreateResponseData) GetRelationships() RoleResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RoleCreateResponseData) GetRelationshipsOk() (*RoleResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RoleCreateResponseData) SetRelationships(v RoleResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *RoleCreateResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *RoleCreateResponseData) GetType() RolesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleCreateResponseData) GetTypeOk() (*RolesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleCreateResponseData) SetType(v RolesType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


