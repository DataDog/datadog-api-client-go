# UserAttributes

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

### NewUserAttributes

`func NewUserAttributes() *UserAttributes`

NewUserAttributes instantiates a new UserAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesWithDefaults

`func NewUserAttributesWithDefaults() *UserAttributes`

NewUserAttributesWithDefaults instantiates a new UserAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UserAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *UserAttributes) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UserAttributes) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UserAttributes) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UserAttributes) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEmail

`func (o *UserAttributes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAttributes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAttributes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAttributes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHandle

`func (o *UserAttributes) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UserAttributes) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UserAttributes) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UserAttributes) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetIcon

`func (o *UserAttributes) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UserAttributes) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UserAttributes) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UserAttributes) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetName

`func (o *UserAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *UserAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *UserAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVerified

`func (o *UserAttributes) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserAttributes) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserAttributes) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserAttributes) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


