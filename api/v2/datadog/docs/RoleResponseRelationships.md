# RoleResponseRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Org** | Pointer to [**RelationshipToOrganization**](RelationshipToOrganization.md) |  | [optional] 
**OtherOrgs** | Pointer to [**RelationshipToOrganizations**](RelationshipToOrganizations.md) |  | [optional] 
**OtherRoles** | Pointer to [**RelationshipToRoles**](RelationshipToRoles.md) |  | [optional] 
**Roles** | Pointer to [**RelationshipToRoles**](RelationshipToRoles.md) |  | [optional] 

## Methods

### NewRoleResponseRelationships

`func NewRoleResponseRelationships() *RoleResponseRelationships`

NewRoleResponseRelationships instantiates a new RoleResponseRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleResponseRelationshipsWithDefaults

`func NewRoleResponseRelationshipsWithDefaults() *RoleResponseRelationships`

NewRoleResponseRelationshipsWithDefaults instantiates a new RoleResponseRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrg

`func (o *RoleResponseRelationships) GetOrg() RelationshipToOrganization`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *RoleResponseRelationships) GetOrgOk() (*RelationshipToOrganization, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *RoleResponseRelationships) SetOrg(v RelationshipToOrganization)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *RoleResponseRelationships) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOtherOrgs

`func (o *RoleResponseRelationships) GetOtherOrgs() RelationshipToOrganizations`

GetOtherOrgs returns the OtherOrgs field if non-nil, zero value otherwise.

### GetOtherOrgsOk

`func (o *RoleResponseRelationships) GetOtherOrgsOk() (*RelationshipToOrganizations, bool)`

GetOtherOrgsOk returns a tuple with the OtherOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOrgs

`func (o *RoleResponseRelationships) SetOtherOrgs(v RelationshipToOrganizations)`

SetOtherOrgs sets OtherOrgs field to given value.

### HasOtherOrgs

`func (o *RoleResponseRelationships) HasOtherOrgs() bool`

HasOtherOrgs returns a boolean if a field has been set.

### GetOtherRoles

`func (o *RoleResponseRelationships) GetOtherRoles() RelationshipToRoles`

GetOtherRoles returns the OtherRoles field if non-nil, zero value otherwise.

### GetOtherRolesOk

`func (o *RoleResponseRelationships) GetOtherRolesOk() (*RelationshipToRoles, bool)`

GetOtherRolesOk returns a tuple with the OtherRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherRoles

`func (o *RoleResponseRelationships) SetOtherRoles(v RelationshipToRoles)`

SetOtherRoles sets OtherRoles field to given value.

### HasOtherRoles

`func (o *RoleResponseRelationships) HasOtherRoles() bool`

HasOtherRoles returns a boolean if a field has been set.

### GetRoles

`func (o *RoleResponseRelationships) GetRoles() RelationshipToRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleResponseRelationships) GetRolesOk() (*RelationshipToRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleResponseRelationships) SetRoles(v RelationshipToRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RoleResponseRelationships) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


