# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRole** | Pointer to [**AccessRole**](AccessRole.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRole

`func (o *User) GetAccessRole() AccessRole`

GetAccessRole returns the AccessRole field if non-nil, zero value otherwise.

### GetAccessRoleOk

`func (o *User) GetAccessRoleOk() (AccessRole, bool)`

GetAccessRoleOk returns a tuple with the AccessRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccessRole

`func (o *User) HasAccessRole() bool`

HasAccessRole returns a boolean if a field has been set.

### SetAccessRole

`func (o *User) SetAccessRole(v AccessRole)`

SetAccessRole gets a reference to the given AccessRole and assigns it to the AccessRole field.

### GetDisabled

`func (o *User) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *User) GetDisabledOk() (bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisabled

`func (o *User) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabled

`func (o *User) SetDisabled(v bool)`

SetDisabled gets a reference to the given bool and assigns it to the Disabled field.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail gets a reference to the given string and assigns it to the Email field.

### GetHandle

`func (o *User) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *User) GetHandleOk() (string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHandle

`func (o *User) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### SetHandle

`func (o *User) SetHandle(v string)`

SetHandle gets a reference to the given string and assigns it to the Handle field.

### GetIcon

`func (o *User) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *User) GetIconOk() (string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIcon

`func (o *User) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIcon

`func (o *User) SetIcon(v string)`

SetIcon gets a reference to the given string and assigns it to the Icon field.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *User) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetVerified

`func (o *User) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *User) GetVerifiedOk() (bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVerified

`func (o *User) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### SetVerified

`func (o *User) SetVerified(v bool)`

SetVerified gets a reference to the given bool and assigns it to the Verified field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


