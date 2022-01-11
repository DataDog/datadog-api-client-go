# IncidentCreateData

## Properties

| Name              | Type                                                                         | Description | Notes                               |
| ----------------- | ---------------------------------------------------------------------------- | ----------- | ----------------------------------- |
| **Attributes**    | [**IncidentCreateAttributes**](IncidentCreateAttributes.md)                  |             |
| **Relationships** | Pointer to [**IncidentCreateRelationships**](IncidentCreateRelationships.md) |             | [optional]                          |
| **Type**          | [**IncidentType**](IncidentType.md)                                          |             | [default to INCIDENTTYPE_INCIDENTS] |

## Methods

### NewIncidentCreateData

`func NewIncidentCreateData(attributes IncidentCreateAttributes, type_ IncidentType) *IncidentCreateData`

NewIncidentCreateData instantiates a new IncidentCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentCreateDataWithDefaults

`func NewIncidentCreateDataWithDefaults() *IncidentCreateData`

NewIncidentCreateDataWithDefaults instantiates a new IncidentCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *IncidentCreateData) GetAttributes() IncidentCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentCreateData) GetAttributesOk() (*IncidentCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentCreateData) SetAttributes(v IncidentCreateAttributes)`

SetAttributes sets Attributes field to given value.

### GetRelationships

`func (o *IncidentCreateData) GetRelationships() IncidentCreateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentCreateData) GetRelationshipsOk() (*IncidentCreateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentCreateData) SetRelationships(v IncidentCreateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentCreateData) GetType() IncidentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentCreateData) GetTypeOk() (*IncidentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentCreateData) SetType(v IncidentType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
