# IncidentServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**IncidentServiceResponseData**](IncidentServiceResponseData.md) |  | 
**Included** | Pointer to [**[]IncidentServiceIncludedItems**](IncidentServiceIncludedItems.md) | Included objects from relationships. | [optional] [readonly] 

## Methods

### NewIncidentServiceResponse

`func NewIncidentServiceResponse(data IncidentServiceResponseData, ) *IncidentServiceResponse`

NewIncidentServiceResponse instantiates a new IncidentServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentServiceResponseWithDefaults

`func NewIncidentServiceResponseWithDefaults() *IncidentServiceResponse`

NewIncidentServiceResponseWithDefaults instantiates a new IncidentServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentServiceResponse) GetData() IncidentServiceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentServiceResponse) GetDataOk() (*IncidentServiceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentServiceResponse) SetData(v IncidentServiceResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentServiceResponse) GetIncluded() []IncidentServiceIncludedItems`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentServiceResponse) GetIncludedOk() (*[]IncidentServiceIncludedItems, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentServiceResponse) SetIncluded(v []IncidentServiceIncludedItems)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentServiceResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


