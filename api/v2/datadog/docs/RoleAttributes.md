# RoleAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CreatedAt** | Pointer to **time.Time** | Creation time of the role. | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | Time of last role modification. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the role. | [optional] 
**UserCount** | Pointer to **int64** | Number of users with that role. | [optional] [readonly] 

## Methods

### NewRoleAttributes

`func NewRoleAttributes() *RoleAttributes`

NewRoleAttributes instantiates a new RoleAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRoleAttributesWithDefaults

`func NewRoleAttributesWithDefaults() *RoleAttributes`

NewRoleAttributesWithDefaults instantiates a new RoleAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *RoleAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoleAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *RoleAttributes) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RoleAttributes) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RoleAttributes) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RoleAttributes) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *RoleAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserCount

`func (o *RoleAttributes) GetUserCount() int64`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *RoleAttributes) GetUserCountOk() (*int64, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *RoleAttributes) SetUserCount(v int64)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *RoleAttributes) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


