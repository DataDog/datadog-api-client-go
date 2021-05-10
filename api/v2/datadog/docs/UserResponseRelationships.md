# UserResponseRelationships

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Org** | Pointer to [**RelationshipToOrganization**](RelationshipToOrganization.md) |  | [optional] 
**OtherOrgs** | Pointer to [**RelationshipToOrganizations**](RelationshipToOrganizations.md) |  | [optional] 
**OtherUsers** | Pointer to [**RelationshipToUsers**](RelationshipToUsers.md) |  | [optional] 
**Roles** | Pointer to [**RelationshipToRoles**](RelationshipToRoles.md) |  | [optional] 

## Methods

### NewUserResponseRelationships

`func NewUserResponseRelationships() *UserResponseRelationships`

NewUserResponseRelationships instantiates a new UserResponseRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseRelationshipsWithDefaults

`func NewUserResponseRelationshipsWithDefaults() *UserResponseRelationships`

NewUserResponseRelationshipsWithDefaults instantiates a new UserResponseRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *UserResponseRelationships) GetOrg() RelationshipToOrganization`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *UserResponseRelationships) GetOrgOk() (*RelationshipToOrganization, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *UserResponseRelationships) SetOrg(v RelationshipToOrganization)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *UserResponseRelationships) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOtherOrgs

`func (o *UserResponseRelationships) GetOtherOrgs() RelationshipToOrganizations`

GetOtherOrgs returns the OtherOrgs field if non-nil, zero value otherwise.

### GetOtherOrgsOk

`func (o *UserResponseRelationships) GetOtherOrgsOk() (*RelationshipToOrganizations, bool)`

GetOtherOrgsOk returns a tuple with the OtherOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOrgs

`func (o *UserResponseRelationships) SetOtherOrgs(v RelationshipToOrganizations)`

SetOtherOrgs sets OtherOrgs field to given value.

### HasOtherOrgs

`func (o *UserResponseRelationships) HasOtherOrgs() bool`

HasOtherOrgs returns a boolean if a field has been set.

### GetOtherUsers

`func (o *UserResponseRelationships) GetOtherUsers() RelationshipToUsers`

GetOtherUsers returns the OtherUsers field if non-nil, zero value otherwise.

### GetOtherUsersOk

`func (o *UserResponseRelationships) GetOtherUsersOk() (*RelationshipToUsers, bool)`

GetOtherUsersOk returns a tuple with the OtherUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherUsers

`func (o *UserResponseRelationships) SetOtherUsers(v RelationshipToUsers)`

SetOtherUsers sets OtherUsers field to given value.

### HasOtherUsers

`func (o *UserResponseRelationships) HasOtherUsers() bool`

HasOtherUsers returns a boolean if a field has been set.

### GetRoles

`func (o *UserResponseRelationships) GetRoles() RelationshipToRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserResponseRelationships) GetRolesOk() (*RelationshipToRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserResponseRelationships) SetRoles(v RelationshipToRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserResponseRelationships) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


