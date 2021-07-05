# SLOListResponseMetadataPage

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**TotalCount** | Pointer to **int64** | The total number of resources that could be retrieved ignoring the parameters and filters in the request. | [optional] 
**TotalFilteredCount** | Pointer to **int64** | The total number of resources that match the parameters and filters in the request. This attribute can be used by a client to determine the total number of pages. | [optional] 

## Methods

### NewSLOListResponseMetadataPage

`func NewSLOListResponseMetadataPage() *SLOListResponseMetadataPage`

NewSLOListResponseMetadataPage instantiates a new SLOListResponseMetadataPage object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSLOListResponseMetadataPageWithDefaults

`func NewSLOListResponseMetadataPageWithDefaults() *SLOListResponseMetadataPage`

NewSLOListResponseMetadataPageWithDefaults instantiates a new SLOListResponseMetadataPage object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetTotalCount

`func (o *SLOListResponseMetadataPage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SLOListResponseMetadataPage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SLOListResponseMetadataPage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SLOListResponseMetadataPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalFilteredCount

`func (o *SLOListResponseMetadataPage) GetTotalFilteredCount() int64`

GetTotalFilteredCount returns the TotalFilteredCount field if non-nil, zero value otherwise.

### GetTotalFilteredCountOk

`func (o *SLOListResponseMetadataPage) GetTotalFilteredCountOk() (*int64, bool)`

GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFilteredCount

`func (o *SLOListResponseMetadataPage) SetTotalFilteredCount(v int64)`

SetTotalFilteredCount sets TotalFilteredCount field to given value.

### HasTotalFilteredCount

`func (o *SLOListResponseMetadataPage) HasTotalFilteredCount() bool`

HasTotalFilteredCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


