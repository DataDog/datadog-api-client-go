# LogsMetricResponseCompute

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AggregationType** | Pointer to [**LogsMetricResponseComputeAggregationType**](LogsMetricResponseComputeAggregationType.md) |  | [optional] 
**Path** | Pointer to **string** | The path to the value the log-based metric will aggregate on (only used if the aggregation type is a \&quot;distribution\&quot;). | [optional] 

## Methods

### NewLogsMetricResponseCompute

`func NewLogsMetricResponseCompute() *LogsMetricResponseCompute`

NewLogsMetricResponseCompute instantiates a new LogsMetricResponseCompute object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsMetricResponseComputeWithDefaults

`func NewLogsMetricResponseComputeWithDefaults() *LogsMetricResponseCompute`

NewLogsMetricResponseComputeWithDefaults instantiates a new LogsMetricResponseCompute object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregationType

`func (o *LogsMetricResponseCompute) GetAggregationType() LogsMetricResponseComputeAggregationType`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *LogsMetricResponseCompute) GetAggregationTypeOk() (*LogsMetricResponseComputeAggregationType, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *LogsMetricResponseCompute) SetAggregationType(v LogsMetricResponseComputeAggregationType)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *LogsMetricResponseCompute) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetPath

`func (o *LogsMetricResponseCompute) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogsMetricResponseCompute) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogsMetricResponseCompute) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogsMetricResponseCompute) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


