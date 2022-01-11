# IncidentServiceResponseData

## Properties

| Name              | Type                                                                                     | Description                    | Notes                                     |
| ----------------- | ---------------------------------------------------------------------------------------- | ------------------------------ | ----------------------------------------- |
| **Attributes**    | Pointer to [**IncidentServiceResponseAttributes**](IncidentServiceResponseAttributes.md) |                                | [optional]                                |
| **Id**            | **string**                                                                               | The incident service&#39;s ID. |
| **Relationships** | Pointer to [**IncidentServiceRelationships**](IncidentServiceRelationships.md)           |                                | [optional]                                |
| **Type**          | [**IncidentServiceType**](IncidentServiceType.md)                                        |                                | [default to INCIDENTSERVICETYPE_SERVICES] |

## Methods

### NewIncidentServiceResponseData

`func NewIncidentServiceResponseData(id string, type_ IncidentServiceType) *IncidentServiceResponseData`

NewIncidentServiceResponseData instantiates a new IncidentServiceResponseData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentServiceResponseDataWithDefaults

`func NewIncidentServiceResponseDataWithDefaults() *IncidentServiceResponseData`

NewIncidentServiceResponseDataWithDefaults instantiates a new IncidentServiceResponseData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *IncidentServiceResponseData) GetAttributes() IncidentServiceResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentServiceResponseData) GetAttributesOk() (*IncidentServiceResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentServiceResponseData) SetAttributes(v IncidentServiceResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentServiceResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentServiceResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentServiceResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentServiceResponseData) SetId(v string)`

SetId sets Id field to given value.

### GetRelationships

`func (o *IncidentServiceResponseData) GetRelationships() IncidentServiceRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentServiceResponseData) GetRelationshipsOk() (*IncidentServiceRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentServiceResponseData) SetRelationships(v IncidentServiceRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentServiceResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentServiceResponseData) GetType() IncidentServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentServiceResponseData) GetTypeOk() (*IncidentServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentServiceResponseData) SetType(v IncidentServiceType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
