# IncidentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentUpdateAttributes**](IncidentUpdateAttributes.md) |  | [optional] 
**Id** | **string** | The team&#39;s ID. | 
**Relationships** | Pointer to [**IncidentUpdateRelationships**](IncidentUpdateRelationships.md) |  | [optional] 
**Type** | [**IncidentType**](IncidentType.md) |  | [default to INCIDENTTYPE_INCIDENTS]

## Methods

### NewIncidentUpdateData

`func NewIncidentUpdateData(id string, type_ IncidentType, ) *IncidentUpdateData`

NewIncidentUpdateData instantiates a new IncidentUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentUpdateDataWithDefaults

`func NewIncidentUpdateDataWithDefaults() *IncidentUpdateData`

NewIncidentUpdateDataWithDefaults instantiates a new IncidentUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentUpdateData) GetAttributes() IncidentUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentUpdateData) GetAttributesOk() (*IncidentUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentUpdateData) SetAttributes(v IncidentUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentUpdateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *IncidentUpdateData) GetRelationships() IncidentUpdateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentUpdateData) GetRelationshipsOk() (*IncidentUpdateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentUpdateData) SetRelationships(v IncidentUpdateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentUpdateData) GetType() IncidentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentUpdateData) GetTypeOk() (*IncidentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentUpdateData) SetType(v IncidentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


