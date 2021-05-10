# IncidentResponseData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**IncidentResponseAttributes**](IncidentResponseAttributes.md) |  | [optional] 
**Id** | **string** | The incident&#39;s ID. | 
**Relationships** | Pointer to [**IncidentResponseRelationships**](IncidentResponseRelationships.md) |  | [optional] 
**Type** | [**IncidentType**](IncidentType.md) |  | [default to INCIDENTTYPE_INCIDENTS]

## Methods

### NewIncidentResponseData

`func NewIncidentResponseData(id string, type_ IncidentType, ) *IncidentResponseData`

NewIncidentResponseData instantiates a new IncidentResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentResponseDataWithDefaults

`func NewIncidentResponseDataWithDefaults() *IncidentResponseData`

NewIncidentResponseDataWithDefaults instantiates a new IncidentResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentResponseData) GetAttributes() IncidentResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentResponseData) GetAttributesOk() (*IncidentResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentResponseData) SetAttributes(v IncidentResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *IncidentResponseData) GetRelationships() IncidentResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentResponseData) GetRelationshipsOk() (*IncidentResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentResponseData) SetRelationships(v IncidentResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentResponseData) GetType() IncidentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentResponseData) GetTypeOk() (*IncidentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentResponseData) SetType(v IncidentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


