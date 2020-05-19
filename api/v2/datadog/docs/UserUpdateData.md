# UserUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**UserUpdateAttributes**](UserUpdateAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the user. | [optional] 
**Type** | Pointer to [**UsersType**](UsersType.md) |  | [optional] [default to "users"]

## Methods

### NewUserUpdateData

`func NewUserUpdateData() *UserUpdateData`

NewUserUpdateData instantiates a new UserUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateDataWithDefaults

`func NewUserUpdateDataWithDefaults() *UserUpdateData`

NewUserUpdateDataWithDefaults instantiates a new UserUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UserUpdateData) GetAttributes() UserUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserUpdateData) GetAttributesOk() (*UserUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserUpdateData) SetAttributes(v UserUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserUpdateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *UserUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserUpdateData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserUpdateData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserUpdateData) GetType() UsersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserUpdateData) GetTypeOk() (*UsersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserUpdateData) SetType(v UsersType)`

SetType sets Type field to given value.

### HasType

`func (o *UserUpdateData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


