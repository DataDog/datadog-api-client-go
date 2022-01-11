# IncidentServiceCreateData

## Properties

| Name              | Type                                                                                 | Description | Notes                                     |
| ----------------- | ------------------------------------------------------------------------------------ | ----------- | ----------------------------------------- |
| **Attributes**    | Pointer to [**IncidentServiceCreateAttributes**](IncidentServiceCreateAttributes.md) |             | [optional]                                |
| **Relationships** | Pointer to [**IncidentServiceRelationships**](IncidentServiceRelationships.md)       |             | [optional]                                |
| **Type**          | [**IncidentServiceType**](IncidentServiceType.md)                                    |             | [default to INCIDENTSERVICETYPE_SERVICES] |

## Methods

### NewIncidentServiceCreateData

`func NewIncidentServiceCreateData(type_ IncidentServiceType) *IncidentServiceCreateData`

NewIncidentServiceCreateData instantiates a new IncidentServiceCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentServiceCreateDataWithDefaults

`func NewIncidentServiceCreateDataWithDefaults() *IncidentServiceCreateData`

NewIncidentServiceCreateDataWithDefaults instantiates a new IncidentServiceCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *IncidentServiceCreateData) GetAttributes() IncidentServiceCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentServiceCreateData) GetAttributesOk() (*IncidentServiceCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentServiceCreateData) SetAttributes(v IncidentServiceCreateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentServiceCreateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentServiceCreateData) GetRelationships() IncidentServiceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentServiceCreateData) GetRelationshipsOk() (*IncidentServiceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentServiceCreateData) SetRelationships(v IncidentServiceRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentServiceCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentServiceCreateData) GetType() IncidentServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentServiceCreateData) GetTypeOk() (*IncidentServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentServiceCreateData) SetType(v IncidentServiceType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
