# User

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AccessRole** | Pointer to [**AccessRole**](AccessRole.md) |  | [optional] [default to ACCESSROLE_STANDARD]
**Disabled** | Pointer to **bool** | The new disabled status of the user. | [optional] 
**Email** | Pointer to **string** | The new email of the user. | [optional] 
**Handle** | Pointer to **string** | The user handle, must be a valid email. | [optional] 
**Icon** | Pointer to **string** | Gravatar icon associated to the user. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the user. | [optional] 
**Verified** | Pointer to **bool** | Whether or not the user logged in Datadog at least once. | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAccessRole

`func (o *User) GetAccessRole() AccessRole`

GetAccessRole returns the AccessRole field if non-nil, zero value otherwise.

### GetAccessRoleOk

`func (o *User) GetAccessRoleOk() (*AccessRole, bool)`

GetAccessRoleOk returns a tuple with the AccessRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRole

`func (o *User) SetAccessRole(v AccessRole)`

SetAccessRole sets AccessRole field to given value.

### HasAccessRole

`func (o *User) HasAccessRole() bool`

HasAccessRole returns a boolean if a field has been set.

### GetDisabled

`func (o *User) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *User) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *User) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *User) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHandle

`func (o *User) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *User) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *User) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *User) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetIcon

`func (o *User) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *User) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *User) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *User) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVerified

`func (o *User) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *User) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *User) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *User) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


