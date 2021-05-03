# UserInvitationDataAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**CreatedAt** | Pointer to **time.Time** | Creation time of the user invitation. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Time of invitation expiration. | [optional] 
**InviteType** | Pointer to **string** | Type of invitation. | [optional] 
**Uuid** | Pointer to **string** | UUID of the user invitation. | [optional] 

## Methods

### NewUserInvitationDataAttributes

`func NewUserInvitationDataAttributes() *UserInvitationDataAttributes`

NewUserInvitationDataAttributes instantiates a new UserInvitationDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationDataAttributesWithDefaults

`func NewUserInvitationDataAttributesWithDefaults() *UserInvitationDataAttributes`

NewUserInvitationDataAttributesWithDefaults instantiates a new UserInvitationDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UserInvitationDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInvitationDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInvitationDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserInvitationDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *UserInvitationDataAttributes) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserInvitationDataAttributes) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserInvitationDataAttributes) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UserInvitationDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetInviteType

`func (o *UserInvitationDataAttributes) GetInviteType() string`

GetInviteType returns the InviteType field if non-nil, zero value otherwise.

### GetInviteTypeOk

`func (o *UserInvitationDataAttributes) GetInviteTypeOk() (*string, bool)`

GetInviteTypeOk returns a tuple with the InviteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteType

`func (o *UserInvitationDataAttributes) SetInviteType(v string)`

SetInviteType sets InviteType field to given value.

### HasInviteType

`func (o *UserInvitationDataAttributes) HasInviteType() bool`

HasInviteType returns a boolean if a field has been set.

### GetUuid

`func (o *UserInvitationDataAttributes) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserInvitationDataAttributes) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserInvitationDataAttributes) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserInvitationDataAttributes) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


