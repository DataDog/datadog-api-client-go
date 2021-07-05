# IncidentServiceUpdateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**IncidentServiceUpdateAttributes**](IncidentServiceUpdateAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The incident service&#39;s ID. | [optional] 
**Relationships** | Pointer to [**IncidentServiceRelationships**](IncidentServiceRelationships.md) |  | [optional] 
**Type** | [**IncidentServiceType**](IncidentServiceType.md) |  | [default to INCIDENTSERVICETYPE_SERVICES]

## Methods

### NewIncidentServiceUpdateData

`func NewIncidentServiceUpdateData(type_ IncidentServiceType) *IncidentServiceUpdateData`

NewIncidentServiceUpdateData instantiates a new IncidentServiceUpdateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentServiceUpdateDataWithDefaults

`func NewIncidentServiceUpdateDataWithDefaults() *IncidentServiceUpdateData`

NewIncidentServiceUpdateDataWithDefaults instantiates a new IncidentServiceUpdateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *IncidentServiceUpdateData) GetAttributes() IncidentServiceUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentServiceUpdateData) GetAttributesOk() (*IncidentServiceUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentServiceUpdateData) SetAttributes(v IncidentServiceUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentServiceUpdateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentServiceUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentServiceUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentServiceUpdateData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentServiceUpdateData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentServiceUpdateData) GetRelationships() IncidentServiceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentServiceUpdateData) GetRelationshipsOk() (*IncidentServiceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentServiceUpdateData) SetRelationships(v IncidentServiceRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentServiceUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentServiceUpdateData) GetType() IncidentServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentServiceUpdateData) GetTypeOk() (*IncidentServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentServiceUpdateData) SetType(v IncidentServiceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


