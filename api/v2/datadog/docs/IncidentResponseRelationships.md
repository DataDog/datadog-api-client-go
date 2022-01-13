# IncidentResponseRelationships

## Properties

| Name                   | Type                                                                                                       | Description | Notes      |
| ---------------------- | ---------------------------------------------------------------------------------------------------------- | ----------- | ---------- |
| **CommanderUser**      | Pointer to [**RelationshipToUser**](RelationshipToUser.md)                                                 |             | [optional] |
| **CreatedByUser**      | Pointer to [**RelationshipToUser**](RelationshipToUser.md)                                                 |             | [optional] |
| **Integrations**       | Pointer to [**RelationshipToIncidentIntegrationMetadatas**](RelationshipToIncidentIntegrationMetadatas.md) |             | [optional] |
| **LastModifiedByUser** | Pointer to [**RelationshipToUser**](RelationshipToUser.md)                                                 |             | [optional] |
| **Postmortem**         | Pointer to [**RelationshipToIncidentPostmortem**](RelationshipToIncidentPostmortem.md)                     |             | [optional] |

## Methods

### NewIncidentResponseRelationships

`func NewIncidentResponseRelationships() *IncidentResponseRelationships`

NewIncidentResponseRelationships instantiates a new IncidentResponseRelationships object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentResponseRelationshipsWithDefaults

`func NewIncidentResponseRelationshipsWithDefaults() *IncidentResponseRelationships`

NewIncidentResponseRelationshipsWithDefaults instantiates a new IncidentResponseRelationships object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCommanderUser

`func (o *IncidentResponseRelationships) GetCommanderUser() RelationshipToUser`

GetCommanderUser returns the CommanderUser field if non-nil, zero value otherwise.

### GetCommanderUserOk

`func (o *IncidentResponseRelationships) GetCommanderUserOk() (*RelationshipToUser, bool)`

GetCommanderUserOk returns a tuple with the CommanderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommanderUser

`func (o *IncidentResponseRelationships) SetCommanderUser(v RelationshipToUser)`

SetCommanderUser sets CommanderUser field to given value.

### HasCommanderUser

`func (o *IncidentResponseRelationships) HasCommanderUser() bool`

HasCommanderUser returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *IncidentResponseRelationships) GetCreatedByUser() RelationshipToUser`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *IncidentResponseRelationships) GetCreatedByUserOk() (*RelationshipToUser, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *IncidentResponseRelationships) SetCreatedByUser(v RelationshipToUser)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *IncidentResponseRelationships) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetIntegrations

`func (o *IncidentResponseRelationships) GetIntegrations() RelationshipToIncidentIntegrationMetadatas`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *IncidentResponseRelationships) GetIntegrationsOk() (*RelationshipToIncidentIntegrationMetadatas, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *IncidentResponseRelationships) SetIntegrations(v RelationshipToIncidentIntegrationMetadatas)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *IncidentResponseRelationships) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetLastModifiedByUser

`func (o *IncidentResponseRelationships) GetLastModifiedByUser() RelationshipToUser`

GetLastModifiedByUser returns the LastModifiedByUser field if non-nil, zero value otherwise.

### GetLastModifiedByUserOk

`func (o *IncidentResponseRelationships) GetLastModifiedByUserOk() (*RelationshipToUser, bool)`

GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUser

`func (o *IncidentResponseRelationships) SetLastModifiedByUser(v RelationshipToUser)`

SetLastModifiedByUser sets LastModifiedByUser field to given value.

### HasLastModifiedByUser

`func (o *IncidentResponseRelationships) HasLastModifiedByUser() bool`

HasLastModifiedByUser returns a boolean if a field has been set.

### GetPostmortem

`func (o *IncidentResponseRelationships) GetPostmortem() RelationshipToIncidentPostmortem`

GetPostmortem returns the Postmortem field if non-nil, zero value otherwise.

### GetPostmortemOk

`func (o *IncidentResponseRelationships) GetPostmortemOk() (*RelationshipToIncidentPostmortem, bool)`

GetPostmortemOk returns a tuple with the Postmortem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmortem

`func (o *IncidentResponseRelationships) SetPostmortem(v RelationshipToIncidentPostmortem)`

SetPostmortem sets Postmortem field to given value.

### HasPostmortem

`func (o *IncidentResponseRelationships) HasPostmortem() bool`

HasPostmortem returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
