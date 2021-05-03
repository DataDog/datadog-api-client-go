# LogsCompute

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Aggregation** | [**LogsAggregationFunction**](LogsAggregationFunction.md) |  | 
**Interval** | Pointer to **string** | The time buckets&#39; size (only used for type&#x3D;timeseries) Defaults to a resolution of 150 points | [optional] 
**Metric** | Pointer to **string** | The metric to use | [optional] 
**Type** | Pointer to [**LogsComputeType**](LogsComputeType.md) |  | [optional] [default to LOGSCOMPUTETYPE_TOTAL]

## Methods

### NewLogsCompute

`func NewLogsCompute(aggregation LogsAggregationFunction, ) *LogsCompute`

NewLogsCompute instantiates a new LogsCompute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsComputeWithDefaults

`func NewLogsComputeWithDefaults() *LogsCompute`

NewLogsComputeWithDefaults instantiates a new LogsCompute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *LogsCompute) GetAggregation() LogsAggregationFunction`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *LogsCompute) GetAggregationOk() (*LogsAggregationFunction, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *LogsCompute) SetAggregation(v LogsAggregationFunction)`

SetAggregation sets Aggregation field to given value.


### GetInterval

`func (o *LogsCompute) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *LogsCompute) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *LogsCompute) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *LogsCompute) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetric

`func (o *LogsCompute) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *LogsCompute) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *LogsCompute) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *LogsCompute) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetType

`func (o *LogsCompute) GetType() LogsComputeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsCompute) GetTypeOk() (*LogsComputeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsCompute) SetType(v LogsComputeType)`

SetType sets Type field to given value.

### HasType

`func (o *LogsCompute) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


