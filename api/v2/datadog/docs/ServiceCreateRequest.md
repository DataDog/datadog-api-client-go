# ServiceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ServiceCreateData**](ServiceCreateData.md) |  | 

## Methods

### NewServiceCreateRequest

`func NewServiceCreateRequest(data ServiceCreateData, ) *ServiceCreateRequest`

NewServiceCreateRequest instantiates a new ServiceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCreateRequestWithDefaults

`func NewServiceCreateRequestWithDefaults() *ServiceCreateRequest`

NewServiceCreateRequestWithDefaults instantiates a new ServiceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceCreateRequest) GetData() ServiceCreateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceCreateRequest) GetDataOk() (*ServiceCreateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceCreateRequest) SetData(v ServiceCreateData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


