# ServicesResponseMetaPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextOffset** | Pointer to **int64** | The index of the first element in the next page of results. Equal to page size added to the current offset. | [optional] 
**Offset** | Pointer to **int64** | The index of the first element in the results. | [optional] 
**Size** | Pointer to **int64** | Maximum size of pages to return. | [optional] 

## Methods

### NewServicesResponseMetaPagination

`func NewServicesResponseMetaPagination() *ServicesResponseMetaPagination`

NewServicesResponseMetaPagination instantiates a new ServicesResponseMetaPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesResponseMetaPaginationWithDefaults

`func NewServicesResponseMetaPaginationWithDefaults() *ServicesResponseMetaPagination`

NewServicesResponseMetaPaginationWithDefaults instantiates a new ServicesResponseMetaPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextOffset

`func (o *ServicesResponseMetaPagination) GetNextOffset() int64`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ServicesResponseMetaPagination) GetNextOffsetOk() (*int64, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ServicesResponseMetaPagination) SetNextOffset(v int64)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ServicesResponseMetaPagination) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetOffset

`func (o *ServicesResponseMetaPagination) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ServicesResponseMetaPagination) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ServicesResponseMetaPagination) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ServicesResponseMetaPagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSize

`func (o *ServicesResponseMetaPagination) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ServicesResponseMetaPagination) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ServicesResponseMetaPagination) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ServicesResponseMetaPagination) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


