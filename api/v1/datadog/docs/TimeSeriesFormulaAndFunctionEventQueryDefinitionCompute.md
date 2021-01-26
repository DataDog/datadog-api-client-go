# TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | [**FormulaAndFunctionEventAggregation**](FormulaAndFunctionEventAggregation.md) |  | 
**Interval** | Pointer to **int64** | A time interval in milliseconds. | [optional] 
**Metric** | Pointer to **string** | Measurable attribute to compute. | [optional] 

## Methods

### NewTimeSeriesFormulaAndFunctionEventQueryDefinitionCompute

`func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionCompute(aggregation FormulaAndFunctionEventAggregation, ) *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute`

NewTimeSeriesFormulaAndFunctionEventQueryDefinitionCompute instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesFormulaAndFunctionEventQueryDefinitionComputeWithDefaults

`func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionComputeWithDefaults() *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute`

NewTimeSeriesFormulaAndFunctionEventQueryDefinitionComputeWithDefaults instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) GetAggregation() FormulaAndFunctionEventAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) GetAggregationOk() (*FormulaAndFunctionEventAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) SetAggregation(v FormulaAndFunctionEventAggregation)`

SetAggregation sets Aggregation field to given value.


### GetInterval

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionCompute) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


