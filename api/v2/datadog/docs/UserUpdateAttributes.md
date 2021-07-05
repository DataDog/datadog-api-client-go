# UserUpdateAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Disabled** | Pointer to **bool** | If the user is enabled or disabled. | [optional] 
**Email** | Pointer to **string** | The email of the user. | [optional] 
**Name** | Pointer to **string** | The name of the user. | [optional] 

## Methods

### NewUserUpdateAttributes

`func NewUserUpdateAttributes() *UserUpdateAttributes`

NewUserUpdateAttributes instantiates a new UserUpdateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUserUpdateAttributesWithDefaults

`func NewUserUpdateAttributesWithDefaults() *UserUpdateAttributes`

NewUserUpdateAttributesWithDefaults instantiates a new UserUpdateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDisabled

`func (o *UserUpdateAttributes) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UserUpdateAttributes) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UserUpdateAttributes) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UserUpdateAttributes) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmail

`func (o *UserUpdateAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUpdateAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUpdateAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserUpdateAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UserUpdateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserUpdateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserUpdateAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserUpdateAttributes) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


