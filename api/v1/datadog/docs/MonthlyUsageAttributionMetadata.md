# MonthlyUsageAttributionMetadata

## Properties

| Name           | Type                                                                                     | Description                       | Notes      |
| -------------- | ---------------------------------------------------------------------------------------- | --------------------------------- | ---------- |
| **Aggregates** | Pointer to [**[]UsageAttributionAggregatesBody**](UsageAttributionAggregatesBody.md)     | An array of available aggregates. | [optional] |
| **Pagination** | Pointer to [**MonthlyUsageAttributionPagination**](MonthlyUsageAttributionPagination.md) |                                   | [optional] |

## Methods

### NewMonthlyUsageAttributionMetadata

`func NewMonthlyUsageAttributionMetadata() *MonthlyUsageAttributionMetadata`

NewMonthlyUsageAttributionMetadata instantiates a new MonthlyUsageAttributionMetadata object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonthlyUsageAttributionMetadataWithDefaults

`func NewMonthlyUsageAttributionMetadataWithDefaults() *MonthlyUsageAttributionMetadata`

NewMonthlyUsageAttributionMetadataWithDefaults instantiates a new MonthlyUsageAttributionMetadata object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregates

`func (o *MonthlyUsageAttributionMetadata) GetAggregates() []UsageAttributionAggregatesBody`

GetAggregates returns the Aggregates field if non-nil, zero value otherwise.

### GetAggregatesOk

`func (o *MonthlyUsageAttributionMetadata) GetAggregatesOk() (*[]UsageAttributionAggregatesBody, bool)`

GetAggregatesOk returns a tuple with the Aggregates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregates

`func (o *MonthlyUsageAttributionMetadata) SetAggregates(v []UsageAttributionAggregatesBody)`

SetAggregates sets Aggregates field to given value.

### HasAggregates

`func (o *MonthlyUsageAttributionMetadata) HasAggregates() bool`

HasAggregates returns a boolean if a field has been set.

### GetPagination

`func (o *MonthlyUsageAttributionMetadata) GetPagination() MonthlyUsageAttributionPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *MonthlyUsageAttributionMetadata) GetPaginationOk() (*MonthlyUsageAttributionPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *MonthlyUsageAttributionMetadata) SetPagination(v MonthlyUsageAttributionPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *MonthlyUsageAttributionMetadata) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
