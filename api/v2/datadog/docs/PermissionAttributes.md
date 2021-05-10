# PermissionAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Created** | Pointer to **time.Time** | Creation time of the permission. | [optional] 
**Description** | Pointer to **string** | Description of the permission. | [optional] 
**DisplayName** | Pointer to **string** | Displayed name for the permission. | [optional] 
**DisplayType** | Pointer to **string** | Display type. | [optional] 
**GroupName** | Pointer to **string** | Name of the permission group. | [optional] 
**Name** | Pointer to **string** | Name of the permission. | [optional] 
**Restricted** | Pointer to **bool** | Whether or not the permission is restricted. | [optional] 

## Methods

### NewPermissionAttributes

`func NewPermissionAttributes() *PermissionAttributes`

NewPermissionAttributes instantiates a new PermissionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionAttributesWithDefaults

`func NewPermissionAttributesWithDefaults() *PermissionAttributes`

NewPermissionAttributesWithDefaults instantiates a new PermissionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PermissionAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PermissionAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PermissionAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PermissionAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *PermissionAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PermissionAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PermissionAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PermissionAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *PermissionAttributes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PermissionAttributes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PermissionAttributes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PermissionAttributes) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayType

`func (o *PermissionAttributes) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *PermissionAttributes) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *PermissionAttributes) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *PermissionAttributes) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetGroupName

`func (o *PermissionAttributes) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *PermissionAttributes) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *PermissionAttributes) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *PermissionAttributes) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetName

`func (o *PermissionAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRestricted

`func (o *PermissionAttributes) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *PermissionAttributes) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *PermissionAttributes) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *PermissionAttributes) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


