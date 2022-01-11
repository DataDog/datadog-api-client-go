# LogQueryDefinitionGroupBy

## Properties

| Name      | Type                                                                             | Description                           | Notes      |
| --------- | -------------------------------------------------------------------------------- | ------------------------------------- | ---------- |
| **Facet** | **string**                                                                       | Facet name.                           |
| **Limit** | Pointer to **int64**                                                             | Maximum number of items in the group. | [optional] |
| **Sort**  | Pointer to [**LogQueryDefinitionGroupBySort**](LogQueryDefinitionGroupBySort.md) |                                       | [optional] |

## Methods

### NewLogQueryDefinitionGroupBy

`func NewLogQueryDefinitionGroupBy(facet string) *LogQueryDefinitionGroupBy`

NewLogQueryDefinitionGroupBy instantiates a new LogQueryDefinitionGroupBy object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogQueryDefinitionGroupByWithDefaults

`func NewLogQueryDefinitionGroupByWithDefaults() *LogQueryDefinitionGroupBy`

NewLogQueryDefinitionGroupByWithDefaults instantiates a new LogQueryDefinitionGroupBy object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFacet

`func (o *LogQueryDefinitionGroupBy) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *LogQueryDefinitionGroupBy) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *LogQueryDefinitionGroupBy) SetFacet(v string)`

SetFacet sets Facet field to given value.

### GetLimit

`func (o *LogQueryDefinitionGroupBy) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LogQueryDefinitionGroupBy) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LogQueryDefinitionGroupBy) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LogQueryDefinitionGroupBy) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSort

`func (o *LogQueryDefinitionGroupBy) GetSort() LogQueryDefinitionGroupBySort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LogQueryDefinitionGroupBy) GetSortOk() (*LogQueryDefinitionGroupBySort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LogQueryDefinitionGroupBy) SetSort(v LogQueryDefinitionGroupBySort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LogQueryDefinitionGroupBy) HasSort() bool`

HasSort returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
