# UserResponseAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Creation time of the user. | [optional] 
**Disabled** | Pointer to **bool** | Whether the user is disabled. | [optional] 
**Email** | Pointer to **string** | Email of the user. | [optional] 
**Handle** | Pointer to **string** | Handle of the user. | [optional] 
**Icon** | Pointer to **string** | URL of the user&#39;s icon. | [optional] 
**Name** | Pointer to **string** | Name of the user. | [optional] 
**Status** | Pointer to **string** | Status of the user. | [optional] 
**Title** | Pointer to **string** | Title of the user. | [optional] 
**Verified** | Pointer to **bool** | Whether the user is verified. | [optional] 

## Methods

### NewUserResponseAttributes

`func NewUserResponseAttributes() *UserResponseAttributes`

NewUserResponseAttributes instantiates a new UserResponseAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseAttributesWithDefaults

`func NewUserResponseAttributesWithDefaults() *UserResponseAttributes`

NewUserResponseAttributesWithDefaults instantiates a new UserResponseAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UserResponseAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserResponseAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserResponseAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserResponseAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *UserResponseAttributes) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UserResponseAttributes) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UserResponseAttributes) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UserResponseAttributes) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmail

`func (o *UserResponseAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResponseAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResponseAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserResponseAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHandle

`func (o *UserResponseAttributes) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UserResponseAttributes) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UserResponseAttributes) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UserResponseAttributes) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetIcon

`func (o *UserResponseAttributes) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UserResponseAttributes) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UserResponseAttributes) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UserResponseAttributes) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetName

`func (o *UserResponseAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserResponseAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserResponseAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserResponseAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *UserResponseAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserResponseAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserResponseAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserResponseAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *UserResponseAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserResponseAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserResponseAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserResponseAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVerified

`func (o *UserResponseAttributes) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserResponseAttributes) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserResponseAttributes) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserResponseAttributes) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


