# Pagination

## Properties

| Name                   | Type                 | Description                                    | Notes      |
| ---------------------- | -------------------- | ---------------------------------------------- | ---------- |
| **TotalCount**         | Pointer to **int64** | Total count.                                   | [optional] |
| **TotalFilteredCount** | Pointer to **int64** | Total count of elements matched by the filter. | [optional] |

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTotalCount

`func (o *Pagination) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Pagination) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Pagination) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Pagination) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalFilteredCount

`func (o *Pagination) GetTotalFilteredCount() int64`

GetTotalFilteredCount returns the TotalFilteredCount field if non-nil, zero value otherwise.

### GetTotalFilteredCountOk

`func (o *Pagination) GetTotalFilteredCountOk() (*int64, bool)`

GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFilteredCount

`func (o *Pagination) SetTotalFilteredCount(v int64)`

SetTotalFilteredCount sets TotalFilteredCount field to given value.

### HasTotalFilteredCount

`func (o *Pagination) HasTotalFilteredCount() bool`

HasTotalFilteredCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
