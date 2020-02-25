# LogsQueryCompute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** |  | 
**Facet** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **int64** |  | [optional] 

## Methods

### NewLogsQueryCompute

`func NewLogsQueryCompute(aggregation string, ) *LogsQueryCompute`

NewLogsQueryCompute instantiates a new LogsQueryCompute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsQueryComputeWithDefaults

`func NewLogsQueryComputeWithDefaults() *LogsQueryCompute`

NewLogsQueryComputeWithDefaults instantiates a new LogsQueryCompute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *LogsQueryCompute) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *LogsQueryCompute) GetAggregationOk() (string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAggregation

`func (o *LogsQueryCompute) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### SetAggregation

`func (o *LogsQueryCompute) SetAggregation(v string)`

SetAggregation gets a reference to the given string and assigns it to the Aggregation field.

### GetFacet

`func (o *LogsQueryCompute) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *LogsQueryCompute) GetFacetOk() (string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFacet

`func (o *LogsQueryCompute) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### SetFacet

`func (o *LogsQueryCompute) SetFacet(v string)`

SetFacet gets a reference to the given string and assigns it to the Facet field.

### GetInterval

`func (o *LogsQueryCompute) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *LogsQueryCompute) GetIntervalOk() (int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInterval

`func (o *LogsQueryCompute) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetInterval

`func (o *LogsQueryCompute) SetInterval(v int64)`

SetInterval gets a reference to the given int64 and assigns it to the Interval field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


