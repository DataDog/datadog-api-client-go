# RoleUpdateAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CreatedAt** | Pointer to **time.Time** | Creation time of the role. | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | Time of last role modification. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the role. | [optional] 

## Methods

### NewRoleUpdateAttributes

`func NewRoleUpdateAttributes() *RoleUpdateAttributes`

NewRoleUpdateAttributes instantiates a new RoleUpdateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleUpdateAttributesWithDefaults

`func NewRoleUpdateAttributesWithDefaults() *RoleUpdateAttributes`

NewRoleUpdateAttributesWithDefaults instantiates a new RoleUpdateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RoleUpdateAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleUpdateAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleUpdateAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoleUpdateAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *RoleUpdateAttributes) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RoleUpdateAttributes) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RoleUpdateAttributes) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *RoleUpdateAttributes) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *RoleUpdateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleUpdateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleUpdateAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleUpdateAttributes) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


