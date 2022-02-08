# AuthNMappingRelationships

## Properties

| Name                       | Type                                                                                           | Description | Notes      |
| -------------------------- | ---------------------------------------------------------------------------------------------- | ----------- | ---------- |
| **Role**                   | Pointer to [**RelationshipToRole**](RelationshipToRole.md)                                     |             | [optional] |
| **SamlAssertionAttribute** | Pointer to [**RelationshipToSAMLAssertionAttribute**](RelationshipToSAMLAssertionAttribute.md) |             | [optional] |

## Methods

### NewAuthNMappingRelationships

`func NewAuthNMappingRelationships() *AuthNMappingRelationships`

NewAuthNMappingRelationships instantiates a new AuthNMappingRelationships object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAuthNMappingRelationshipsWithDefaults

`func NewAuthNMappingRelationshipsWithDefaults() *AuthNMappingRelationships`

NewAuthNMappingRelationshipsWithDefaults instantiates a new AuthNMappingRelationships object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRole

`func (o *AuthNMappingRelationships) GetRole() RelationshipToRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AuthNMappingRelationships) GetRoleOk() (*RelationshipToRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AuthNMappingRelationships) SetRole(v RelationshipToRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *AuthNMappingRelationships) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSamlAssertionAttribute

`func (o *AuthNMappingRelationships) GetSamlAssertionAttribute() RelationshipToSAMLAssertionAttribute`

GetSamlAssertionAttribute returns the SamlAssertionAttribute field if non-nil, zero value otherwise.

### GetSamlAssertionAttributeOk

`func (o *AuthNMappingRelationships) GetSamlAssertionAttributeOk() (*RelationshipToSAMLAssertionAttribute, bool)`

GetSamlAssertionAttributeOk returns a tuple with the SamlAssertionAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAssertionAttribute

`func (o *AuthNMappingRelationships) SetSamlAssertionAttribute(v RelationshipToSAMLAssertionAttribute)`

SetSamlAssertionAttribute sets SamlAssertionAttribute field to given value.

### HasSamlAssertionAttribute

`func (o *AuthNMappingRelationships) HasSamlAssertionAttribute() bool`

HasSamlAssertionAttribute returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
