# SLOBulkDeleteResponseData

## Properties

| Name        | Type                    | Description                                                                                                                                                                                  | Notes      |
| ----------- | ----------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **Deleted** | Pointer to **[]string** | An array of service level objective object IDs that indicates which objects that were completely deleted.                                                                                    | [optional] |
| **Updated** | Pointer to **[]string** | An array of service level objective object IDs that indicates which objects that were modified (objects for which at least one threshold was deleted, but that were not completely deleted). | [optional] |

## Methods

### NewSLOBulkDeleteResponseData

`func NewSLOBulkDeleteResponseData() *SLOBulkDeleteResponseData`

NewSLOBulkDeleteResponseData instantiates a new SLOBulkDeleteResponseData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOBulkDeleteResponseDataWithDefaults

`func NewSLOBulkDeleteResponseDataWithDefaults() *SLOBulkDeleteResponseData`

NewSLOBulkDeleteResponseDataWithDefaults instantiates a new SLOBulkDeleteResponseData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDeleted

`func (o *SLOBulkDeleteResponseData) GetDeleted() []string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SLOBulkDeleteResponseData) GetDeletedOk() (*[]string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SLOBulkDeleteResponseData) SetDeleted(v []string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *SLOBulkDeleteResponseData) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUpdated

`func (o *SLOBulkDeleteResponseData) GetUpdated() []string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SLOBulkDeleteResponseData) GetUpdatedOk() (*[]string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SLOBulkDeleteResponseData) SetUpdated(v []string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SLOBulkDeleteResponseData) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
