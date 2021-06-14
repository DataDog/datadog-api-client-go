# UsageAttributionPagination

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Limit** | Pointer to **int64** | Maximum amount of records to be returned. | [optional] 
**NextRecordId** | Pointer to **string** | The cursor to use to get the next results, if any. To make the next request, use the same parameters with the addition of this next_record_id. | [optional] 
**TotalNumberOfRecords** | Pointer to **int64** | Total number of records. (deprecated after May 1st, 2021) | [optional] 

## Methods

### NewUsageAttributionPagination

`func NewUsageAttributionPagination() *UsageAttributionPagination`

NewUsageAttributionPagination instantiates a new UsageAttributionPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAttributionPaginationWithDefaults

`func NewUsageAttributionPaginationWithDefaults() *UsageAttributionPagination`

NewUsageAttributionPaginationWithDefaults instantiates a new UsageAttributionPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *UsageAttributionPagination) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UsageAttributionPagination) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UsageAttributionPagination) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UsageAttributionPagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextRecordId

`func (o *UsageAttributionPagination) GetNextRecordId() string`

GetNextRecordId returns the NextRecordId field if non-nil, zero value otherwise.

### GetNextRecordIdOk

`func (o *UsageAttributionPagination) GetNextRecordIdOk() (*string, bool)`

GetNextRecordIdOk returns a tuple with the NextRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecordId

`func (o *UsageAttributionPagination) SetNextRecordId(v string)`

SetNextRecordId sets NextRecordId field to given value.

### HasNextRecordId

`func (o *UsageAttributionPagination) HasNextRecordId() bool`

HasNextRecordId returns a boolean if a field has been set.

### GetTotalNumberOfRecords

`func (o *UsageAttributionPagination) GetTotalNumberOfRecords() int64`

GetTotalNumberOfRecords returns the TotalNumberOfRecords field if non-nil, zero value otherwise.

### GetTotalNumberOfRecordsOk

`func (o *UsageAttributionPagination) GetTotalNumberOfRecordsOk() (*int64, bool)`

GetTotalNumberOfRecordsOk returns a tuple with the TotalNumberOfRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfRecords

`func (o *UsageAttributionPagination) SetTotalNumberOfRecords(v int64)`

SetTotalNumberOfRecords sets TotalNumberOfRecords field to given value.

### HasTotalNumberOfRecords

`func (o *UsageAttributionPagination) HasTotalNumberOfRecords() bool`

HasTotalNumberOfRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


