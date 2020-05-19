# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**PermissionAttributes**](PermissionAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the permission. | [optional] 
**Type** | Pointer to [**PermissionsType**](PermissionsType.md) |  | [default to "permissions"]

## Methods

### NewPermission

`func NewPermission(type_ PermissionsType, ) *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Permission) GetAttributes() PermissionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Permission) GetAttributesOk() (*PermissionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Permission) SetAttributes(v PermissionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Permission) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *Permission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Permission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Permission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Permission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Permission) GetType() PermissionsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Permission) GetTypeOk() (*PermissionsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Permission) SetType(v PermissionsType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


