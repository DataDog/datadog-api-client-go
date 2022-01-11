# NotebooksResponsePage

## Properties

| Name                   | Type                 | Description                                                                                                                                 | Notes      |
| ---------------------- | -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **TotalCount**         | Pointer to **int64** | The total number of notebooks that would be returned if the request was not filtered by &#x60;start&#x60; and &#x60;count&#x60; parameters. | [optional] |
| **TotalFilteredCount** | Pointer to **int64** | The total number of notebooks returned.                                                                                                     | [optional] |

## Methods

### NewNotebooksResponsePage

`func NewNotebooksResponsePage() *NotebooksResponsePage`

NewNotebooksResponsePage instantiates a new NotebooksResponsePage object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNotebooksResponsePageWithDefaults

`func NewNotebooksResponsePageWithDefaults() *NotebooksResponsePage`

NewNotebooksResponsePageWithDefaults instantiates a new NotebooksResponsePage object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTotalCount

`func (o *NotebooksResponsePage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *NotebooksResponsePage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *NotebooksResponsePage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *NotebooksResponsePage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalFilteredCount

`func (o *NotebooksResponsePage) GetTotalFilteredCount() int64`

GetTotalFilteredCount returns the TotalFilteredCount field if non-nil, zero value otherwise.

### GetTotalFilteredCountOk

`func (o *NotebooksResponsePage) GetTotalFilteredCountOk() (*int64, bool)`

GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFilteredCount

`func (o *NotebooksResponsePage) SetTotalFilteredCount(v int64)`

SetTotalFilteredCount sets TotalFilteredCount field to given value.

### HasTotalFilteredCount

`func (o *NotebooksResponsePage) HasTotalFilteredCount() bool`

HasTotalFilteredCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
