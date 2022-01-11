# MonthlyUsageAttributionPagination

## Properties

| Name             | Type                  | Description                                                                                                                                               | Notes      |
| ---------------- | --------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **NextRecordId** | Pointer to **string** | The cursor to use to get the next results, if any. To make the next request, use the same parameters with the addition of the &#x60;next_record_id&#x60;. | [optional] |

## Methods

### NewMonthlyUsageAttributionPagination

`func NewMonthlyUsageAttributionPagination() *MonthlyUsageAttributionPagination`

NewMonthlyUsageAttributionPagination instantiates a new MonthlyUsageAttributionPagination object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonthlyUsageAttributionPaginationWithDefaults

`func NewMonthlyUsageAttributionPaginationWithDefaults() *MonthlyUsageAttributionPagination`

NewMonthlyUsageAttributionPaginationWithDefaults instantiates a new MonthlyUsageAttributionPagination object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetNextRecordId

`func (o *MonthlyUsageAttributionPagination) GetNextRecordId() string`

GetNextRecordId returns the NextRecordId field if non-nil, zero value otherwise.

### GetNextRecordIdOk

`func (o *MonthlyUsageAttributionPagination) GetNextRecordIdOk() (*string, bool)`

GetNextRecordIdOk returns a tuple with the NextRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecordId

`func (o *MonthlyUsageAttributionPagination) SetNextRecordId(v string)`

SetNextRecordId sets NextRecordId field to given value.

### HasNextRecordId

`func (o *MonthlyUsageAttributionPagination) HasNextRecordId() bool`

HasNextRecordId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
