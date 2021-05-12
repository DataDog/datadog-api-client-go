# IncidentTeamCreateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**IncidentTeamCreateAttributes**](IncidentTeamCreateAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**IncidentTeamRelationships**](IncidentTeamRelationships.md) |  | [optional] 
**Type** | [**IncidentTeamType**](IncidentTeamType.md) |  | [default to INCIDENTTEAMTYPE_TEAMS]

## Methods

### NewIncidentTeamCreateData

`func NewIncidentTeamCreateData(type_ IncidentTeamType, ) *IncidentTeamCreateData`

NewIncidentTeamCreateData instantiates a new IncidentTeamCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTeamCreateDataWithDefaults

`func NewIncidentTeamCreateDataWithDefaults() *IncidentTeamCreateData`

NewIncidentTeamCreateDataWithDefaults instantiates a new IncidentTeamCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentTeamCreateData) GetAttributes() IncidentTeamCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentTeamCreateData) GetAttributesOk() (*IncidentTeamCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentTeamCreateData) SetAttributes(v IncidentTeamCreateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentTeamCreateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentTeamCreateData) GetRelationships() IncidentTeamRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentTeamCreateData) GetRelationshipsOk() (*IncidentTeamRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentTeamCreateData) SetRelationships(v IncidentTeamRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentTeamCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentTeamCreateData) GetType() IncidentTeamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentTeamCreateData) GetTypeOk() (*IncidentTeamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentTeamCreateData) SetType(v IncidentTeamType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


