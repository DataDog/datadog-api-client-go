# IncidentTeamsResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | [**[]IncidentTeamResponseData**](IncidentTeamResponseData.md) | An array of incident teams. | 
**Included** | Pointer to [**[]IncidentTeamIncludedItems**](IncidentTeamIncludedItems.md) | Included related resources which the user requested. | [optional] [readonly] 
**Meta** | Pointer to [**IncidentServicesResponseMeta**](IncidentServicesResponseMeta.md) |  | [optional] 

## Methods

### NewIncidentTeamsResponse

`func NewIncidentTeamsResponse(data []IncidentTeamResponseData) *IncidentTeamsResponse`

NewIncidentTeamsResponse instantiates a new IncidentTeamsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentTeamsResponseWithDefaults

`func NewIncidentTeamsResponseWithDefaults() *IncidentTeamsResponse`

NewIncidentTeamsResponseWithDefaults instantiates a new IncidentTeamsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *IncidentTeamsResponse) GetData() []IncidentTeamResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentTeamsResponse) GetDataOk() (*[]IncidentTeamResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentTeamsResponse) SetData(v []IncidentTeamResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentTeamsResponse) GetIncluded() []IncidentTeamIncludedItems`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentTeamsResponse) GetIncludedOk() (*[]IncidentTeamIncludedItems, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentTeamsResponse) SetIncluded(v []IncidentTeamIncludedItems)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentTeamsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *IncidentTeamsResponse) GetMeta() IncidentServicesResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentTeamsResponse) GetMetaOk() (*IncidentServicesResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentTeamsResponse) SetMeta(v IncidentServicesResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentTeamsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


