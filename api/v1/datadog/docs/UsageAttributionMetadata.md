# UsageAttributionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregates** | Pointer to [**[]UsageAttributionAggregatesBody**](UsageAttributionAggregatesBody.md) | An array of available aggregates. | [optional] 
**Pagination** | Pointer to [**UsageAttributionPagination**](UsageAttributionPagination.md) |  | [optional] 

## Methods

### NewUsageAttributionMetadata

`func NewUsageAttributionMetadata() *UsageAttributionMetadata`

NewUsageAttributionMetadata instantiates a new UsageAttributionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAttributionMetadataWithDefaults

`func NewUsageAttributionMetadataWithDefaults() *UsageAttributionMetadata`

NewUsageAttributionMetadataWithDefaults instantiates a new UsageAttributionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregates

`func (o *UsageAttributionMetadata) GetAggregates() []UsageAttributionAggregatesBody`

GetAggregates returns the Aggregates field if non-nil, zero value otherwise.

### GetAggregatesOk

`func (o *UsageAttributionMetadata) GetAggregatesOk() (*[]UsageAttributionAggregatesBody, bool)`

GetAggregatesOk returns a tuple with the Aggregates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregates

`func (o *UsageAttributionMetadata) SetAggregates(v []UsageAttributionAggregatesBody)`

SetAggregates sets Aggregates field to given value.

### HasAggregates

`func (o *UsageAttributionMetadata) HasAggregates() bool`

HasAggregates returns a boolean if a field has been set.

### GetPagination

`func (o *UsageAttributionMetadata) GetPagination() UsageAttributionPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *UsageAttributionMetadata) GetPaginationOk() (*UsageAttributionPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *UsageAttributionMetadata) SetPagination(v UsageAttributionPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *UsageAttributionMetadata) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


