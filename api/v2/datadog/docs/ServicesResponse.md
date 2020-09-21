# ServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ServiceResponseData**](ServiceResponseData.md) | An array of services. | 
**Included** | Pointer to [**[]ServiceIncludedItems**](ServiceIncludedItems.md) | Included related resources which the user requested. | [optional] [readonly] 
**Meta** | Pointer to [**ServicesResponseMeta**](ServicesResponse_meta.md) |  | [optional] 

## Methods

### NewServicesResponse

`func NewServicesResponse(data []ServiceResponseData, ) *ServicesResponse`

NewServicesResponse instantiates a new ServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesResponseWithDefaults

`func NewServicesResponseWithDefaults() *ServicesResponse`

NewServicesResponseWithDefaults instantiates a new ServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServicesResponse) GetData() []ServiceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServicesResponse) GetDataOk() (*[]ServiceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServicesResponse) SetData(v []ServiceResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *ServicesResponse) GetIncluded() []ServiceIncludedItems`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ServicesResponse) GetIncludedOk() (*[]ServiceIncludedItems, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ServicesResponse) SetIncluded(v []ServiceIncludedItems)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ServicesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *ServicesResponse) GetMeta() ServicesResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServicesResponse) GetMetaOk() (*ServicesResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServicesResponse) SetMeta(v ServicesResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServicesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


