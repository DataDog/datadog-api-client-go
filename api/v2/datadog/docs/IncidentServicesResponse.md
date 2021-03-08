# IncidentServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]IncidentServiceResponseData**](IncidentServiceResponseData.md) | An array of incident services. | 
**Included** | Pointer to [**[]IncidentServiceIncludedItems**](IncidentServiceIncludedItems.md) | Included related resources which the user requested. | [optional] [readonly] 
**Meta** | Pointer to [**IncidentServicesResponseMeta**](IncidentServicesResponseMeta.md) |  | [optional] 

## Methods

### NewIncidentServicesResponse

`func NewIncidentServicesResponse(data []IncidentServiceResponseData, ) *IncidentServicesResponse`

NewIncidentServicesResponse instantiates a new IncidentServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentServicesResponseWithDefaults

`func NewIncidentServicesResponseWithDefaults() *IncidentServicesResponse`

NewIncidentServicesResponseWithDefaults instantiates a new IncidentServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentServicesResponse) GetData() []IncidentServiceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentServicesResponse) GetDataOk() (*[]IncidentServiceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentServicesResponse) SetData(v []IncidentServiceResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentServicesResponse) GetIncluded() []IncidentServiceIncludedItems`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentServicesResponse) GetIncludedOk() (*[]IncidentServiceIncludedItems, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentServicesResponse) SetIncluded(v []IncidentServiceIncludedItems)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentServicesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *IncidentServicesResponse) GetMeta() IncidentServicesResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentServicesResponse) GetMetaOk() (*IncidentServicesResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentServicesResponse) SetMeta(v IncidentServicesResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentServicesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


