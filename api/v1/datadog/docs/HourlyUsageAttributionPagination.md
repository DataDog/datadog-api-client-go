# HourlyUsageAttributionPagination

## Properties

| Name             | Type                  | Description                                                                                                                        | Notes      |
| ---------------- | --------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **NextRecordId** | Pointer to **string** | The cursor to get the next results (if any). To make the next request, use the same parameters and add &#x60;next_record_id&#x60;. | [optional] |

## Methods

### NewHourlyUsageAttributionPagination

`func NewHourlyUsageAttributionPagination() *HourlyUsageAttributionPagination`

NewHourlyUsageAttributionPagination instantiates a new HourlyUsageAttributionPagination object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHourlyUsageAttributionPaginationWithDefaults

`func NewHourlyUsageAttributionPaginationWithDefaults() *HourlyUsageAttributionPagination`

NewHourlyUsageAttributionPaginationWithDefaults instantiates a new HourlyUsageAttributionPagination object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetNextRecordId

`func (o *HourlyUsageAttributionPagination) GetNextRecordId() string`

GetNextRecordId returns the NextRecordId field if non-nil, zero value otherwise.

### GetNextRecordIdOk

`func (o *HourlyUsageAttributionPagination) GetNextRecordIdOk() (*string, bool)`

GetNextRecordIdOk returns a tuple with the NextRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecordId

`func (o *HourlyUsageAttributionPagination) SetNextRecordId(v string)`

SetNextRecordId sets NextRecordId field to given value.

### HasNextRecordId

`func (o *HourlyUsageAttributionPagination) HasNextRecordId() bool`

HasNextRecordId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
