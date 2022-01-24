# IncidentCreateRelationships

## Properties

| Name              | Type                                                            | Description | Notes |
| ----------------- | --------------------------------------------------------------- | ----------- | ----- |
| **CommanderUser** | [**NullableRelationshipToUser**](NullableRelationshipToUser.md) |             |

## Methods

### NewIncidentCreateRelationships

`func NewIncidentCreateRelationships(commanderUser NullableRelationshipToUser) *IncidentCreateRelationships`

NewIncidentCreateRelationships instantiates a new IncidentCreateRelationships object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentCreateRelationshipsWithDefaults

`func NewIncidentCreateRelationshipsWithDefaults() *IncidentCreateRelationships`

NewIncidentCreateRelationshipsWithDefaults instantiates a new IncidentCreateRelationships object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCommanderUser

`func (o *IncidentCreateRelationships) GetCommanderUser() NullableRelationshipToUser`

GetCommanderUser returns the CommanderUser field if non-nil, zero value otherwise.

### GetCommanderUserOk

`func (o *IncidentCreateRelationships) GetCommanderUserOk() (*NullableRelationshipToUser, bool)`

GetCommanderUserOk returns a tuple with the CommanderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommanderUser

`func (o *IncidentCreateRelationships) SetCommanderUser(v NullableRelationshipToUser)`

SetCommanderUser sets CommanderUser field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
