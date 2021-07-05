# LogsGroupBy

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Facet** | **string** | The name of the facet to use (required) | 
**Histogram** | Pointer to [**LogsGroupByHistogram**](LogsGroupByHistogram.md) |  | [optional] 
**Limit** | Pointer to **int64** | The maximum buckets to return for this group by | [optional] [default to 10]
**Missing** | Pointer to [**LogsGroupByMissing**](LogsGroupByMissing.md) |  | [optional] 
**Sort** | Pointer to [**LogsAggregateSort**](LogsAggregateSort.md) |  | [optional] 
**Total** | Pointer to [**LogsGroupByTotal**](LogsGroupByTotal.md) |  | [optional] 

## Methods

### NewLogsGroupBy

`func NewLogsGroupBy(facet string) *LogsGroupBy`

NewLogsGroupBy instantiates a new LogsGroupBy object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsGroupByWithDefaults

`func NewLogsGroupByWithDefaults() *LogsGroupBy`

NewLogsGroupByWithDefaults instantiates a new LogsGroupBy object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFacet

`func (o *LogsGroupBy) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *LogsGroupBy) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *LogsGroupBy) SetFacet(v string)`

SetFacet sets Facet field to given value.


### GetHistogram

`func (o *LogsGroupBy) GetHistogram() LogsGroupByHistogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *LogsGroupBy) GetHistogramOk() (*LogsGroupByHistogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *LogsGroupBy) SetHistogram(v LogsGroupByHistogram)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *LogsGroupBy) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetLimit

`func (o *LogsGroupBy) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LogsGroupBy) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LogsGroupBy) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LogsGroupBy) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMissing

`func (o *LogsGroupBy) GetMissing() LogsGroupByMissing`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *LogsGroupBy) GetMissingOk() (*LogsGroupByMissing, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *LogsGroupBy) SetMissing(v LogsGroupByMissing)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *LogsGroupBy) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetSort

`func (o *LogsGroupBy) GetSort() LogsAggregateSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LogsGroupBy) GetSortOk() (*LogsAggregateSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LogsGroupBy) SetSort(v LogsAggregateSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LogsGroupBy) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTotal

`func (o *LogsGroupBy) GetTotal() LogsGroupByTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LogsGroupBy) GetTotalOk() (*LogsGroupByTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LogsGroupBy) SetTotal(v LogsGroupByTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LogsGroupBy) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


