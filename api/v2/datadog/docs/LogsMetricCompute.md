# LogsMetricCompute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationType** | [**LogsMetricComputeAggregationType**](LogsMetricComputeAggregationType.md) |  | 
**Path** | Pointer to **string** | The path to the value the log-based metric will aggregate on (only used if the aggregation type is a \&quot;distribution\&quot;). | [optional] 

## Methods

### NewLogsMetricCompute

`func NewLogsMetricCompute(aggregationType LogsMetricComputeAggregationType, ) *LogsMetricCompute`

NewLogsMetricCompute instantiates a new LogsMetricCompute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsMetricComputeWithDefaults

`func NewLogsMetricComputeWithDefaults() *LogsMetricCompute`

NewLogsMetricComputeWithDefaults instantiates a new LogsMetricCompute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationType

`func (o *LogsMetricCompute) GetAggregationType() LogsMetricComputeAggregationType`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *LogsMetricCompute) GetAggregationTypeOk() (*LogsMetricComputeAggregationType, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *LogsMetricCompute) SetAggregationType(v LogsMetricComputeAggregationType)`

SetAggregationType sets AggregationType field to given value.


### GetPath

`func (o *LogsMetricCompute) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LogsMetricCompute) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LogsMetricCompute) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LogsMetricCompute) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


