# FormulaAndFunctionEventQueryDefinitionCompute

## Properties

| Name            | Type                                                                            | Description                      | Notes      |
| --------------- | ------------------------------------------------------------------------------- | -------------------------------- | ---------- |
| **Aggregation** | [**FormulaAndFunctionEventAggregation**](FormulaAndFunctionEventAggregation.md) |                                  |
| **Interval**    | Pointer to **int64**                                                            | A time interval in milliseconds. | [optional] |
| **Metric**      | Pointer to **string**                                                           | Measurable attribute to compute. | [optional] |

## Methods

### NewFormulaAndFunctionEventQueryDefinitionCompute

`func NewFormulaAndFunctionEventQueryDefinitionCompute(aggregation FormulaAndFunctionEventAggregation) *FormulaAndFunctionEventQueryDefinitionCompute`

NewFormulaAndFunctionEventQueryDefinitionCompute instantiates a new FormulaAndFunctionEventQueryDefinitionCompute object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFormulaAndFunctionEventQueryDefinitionComputeWithDefaults

`func NewFormulaAndFunctionEventQueryDefinitionComputeWithDefaults() *FormulaAndFunctionEventQueryDefinitionCompute`

NewFormulaAndFunctionEventQueryDefinitionComputeWithDefaults instantiates a new FormulaAndFunctionEventQueryDefinitionCompute object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregation

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) GetAggregation() FormulaAndFunctionEventAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) GetAggregationOk() (*FormulaAndFunctionEventAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) SetAggregation(v FormulaAndFunctionEventAggregation)`

SetAggregation sets Aggregation field to given value.

### GetInterval

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetric

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *FormulaAndFunctionEventQueryDefinitionCompute) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
