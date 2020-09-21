# ServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ServiceResponseData**](ServiceResponseData.md) |  | 
**Included** | Pointer to [**[]ServiceIncludedItems**](ServiceIncludedItems.md) | Included objects from relationships. | [optional] [readonly] 

## Methods

### NewServiceResponse

`func NewServiceResponse(data ServiceResponseData, ) *ServiceResponse`

NewServiceResponse instantiates a new ServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseWithDefaults

`func NewServiceResponseWithDefaults() *ServiceResponse`

NewServiceResponseWithDefaults instantiates a new ServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceResponse) GetData() ServiceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceResponse) GetDataOk() (*ServiceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceResponse) SetData(v ServiceResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *ServiceResponse) GetIncluded() []ServiceIncludedItems`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *ServiceResponse) GetIncludedOk() (*[]ServiceIncludedItems, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *ServiceResponse) SetIncluded(v []ServiceIncludedItems)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *ServiceResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


