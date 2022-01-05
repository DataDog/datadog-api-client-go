# UsageAttributionPagination

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Limit** | Pointer to **int64** | Maximum amount of records to be returned. | [optional] 
**Offset** | Pointer to **int64** | Records to be skipped before beginning to return. | [optional] 
**SortDirection** | Pointer to **string** | Direction to sort by. | [optional] 
**SortName** | Pointer to **string** | Field to sort by. | [optional] 
**TotalNumberOfRecords** | Pointer to **int64** | Total number of records. | [optional] 

## Methods

### NewUsageAttributionPagination

`func NewUsageAttributionPagination() *UsageAttributionPagination`

NewUsageAttributionPagination instantiates a new UsageAttributionPagination object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageAttributionPaginationWithDefaults

`func NewUsageAttributionPaginationWithDefaults() *UsageAttributionPagination`

NewUsageAttributionPaginationWithDefaults instantiates a new UsageAttributionPagination object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

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

### GetOffset

`func (o *UsageAttributionPagination) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UsageAttributionPagination) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UsageAttributionPagination) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *UsageAttributionPagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortDirection

`func (o *UsageAttributionPagination) GetSortDirection() string`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *UsageAttributionPagination) GetSortDirectionOk() (*string, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *UsageAttributionPagination) SetSortDirection(v string)`

SetSortDirection sets SortDirection field to given value.

### HasSortDirection

`func (o *UsageAttributionPagination) HasSortDirection() bool`

HasSortDirection returns a boolean if a field has been set.

### GetSortName

`func (o *UsageAttributionPagination) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *UsageAttributionPagination) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *UsageAttributionPagination) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *UsageAttributionPagination) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

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


