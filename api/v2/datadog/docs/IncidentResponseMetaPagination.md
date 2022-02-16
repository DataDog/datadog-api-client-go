# IncidentResponseMetaPagination

## Properties

| Name           | Type                 | Description                                                                                                 | Notes      |
| -------------- | -------------------- | ----------------------------------------------------------------------------------------------------------- | ---------- |
| **NextOffset** | Pointer to **int64** | The index of the first element in the next page of results. Equal to page size added to the current offset. | [optional] |
| **Offset**     | Pointer to **int64** | The index of the first element in the results.                                                              | [optional] |
| **Size**       | Pointer to **int64** | Maximum size of pages to return.                                                                            | [optional] |

## Methods

### NewIncidentResponseMetaPagination

`func NewIncidentResponseMetaPagination() *IncidentResponseMetaPagination`

NewIncidentResponseMetaPagination instantiates a new IncidentResponseMetaPagination object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentResponseMetaPaginationWithDefaults

`func NewIncidentResponseMetaPaginationWithDefaults() *IncidentResponseMetaPagination`

NewIncidentResponseMetaPaginationWithDefaults instantiates a new IncidentResponseMetaPagination object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetNextOffset

`func (o *IncidentResponseMetaPagination) GetNextOffset() int64`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *IncidentResponseMetaPagination) GetNextOffsetOk() (*int64, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *IncidentResponseMetaPagination) SetNextOffset(v int64)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *IncidentResponseMetaPagination) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetOffset

`func (o *IncidentResponseMetaPagination) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IncidentResponseMetaPagination) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IncidentResponseMetaPagination) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IncidentResponseMetaPagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSize

`func (o *IncidentResponseMetaPagination) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IncidentResponseMetaPagination) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IncidentResponseMetaPagination) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *IncidentResponseMetaPagination) HasSize() bool`

HasSize returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
