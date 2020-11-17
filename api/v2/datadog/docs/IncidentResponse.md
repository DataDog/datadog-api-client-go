# IncidentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**IncidentResponseData**](IncidentResponseData.md) |  | 
**Included** | Pointer to [**[]IncidentResponseIncludedItem**](IncidentResponseIncludedItem.md) | Included related resources that the user requested. | [optional] [readonly] 

## Methods

### NewIncidentResponse

`func NewIncidentResponse(data IncidentResponseData, ) *IncidentResponse`

NewIncidentResponse instantiates a new IncidentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentResponseWithDefaults

`func NewIncidentResponseWithDefaults() *IncidentResponse`

NewIncidentResponseWithDefaults instantiates a new IncidentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentResponse) GetData() IncidentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentResponse) GetDataOk() (*IncidentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentResponse) SetData(v IncidentResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentResponse) GetIncluded() []IncidentResponseIncludedItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentResponse) GetIncludedOk() (*[]IncidentResponseIncludedItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentResponse) SetIncluded(v []IncidentResponseIncludedItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


