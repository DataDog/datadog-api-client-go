# IncidentTeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**IncidentTeamResponseData**](IncidentTeamResponseData.md) |  | 
**Included** | Pointer to [**[]IncidentTeamIncludedItems**](IncidentTeamIncludedItems.md) | Included objects from relationships. | [optional] [readonly] 

## Methods

### NewIncidentTeamResponse

`func NewIncidentTeamResponse(data IncidentTeamResponseData, ) *IncidentTeamResponse`

NewIncidentTeamResponse instantiates a new IncidentTeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTeamResponseWithDefaults

`func NewIncidentTeamResponseWithDefaults() *IncidentTeamResponse`

NewIncidentTeamResponseWithDefaults instantiates a new IncidentTeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentTeamResponse) GetData() IncidentTeamResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentTeamResponse) GetDataOk() (*IncidentTeamResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentTeamResponse) SetData(v IncidentTeamResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentTeamResponse) GetIncluded() []IncidentTeamIncludedItems`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentTeamResponse) GetIncludedOk() (*[]IncidentTeamIncludedItems, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentTeamResponse) SetIncluded(v []IncidentTeamIncludedItems)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentTeamResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


