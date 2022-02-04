# MonitorFormulaAndFunctionEventQueryDefinitionCompute

## Properties

| Name            | Type                                                                                          | Description                      | Notes      |
| --------------- | --------------------------------------------------------------------------------------------- | -------------------------------- | ---------- |
| **Aggregation** | [**MonitorFormulaAndFunctionEventAggregation**](MonitorFormulaAndFunctionEventAggregation.md) |                                  |
| **Interval**    | Pointer to **int64**                                                                          | A time interval in milliseconds. | [optional] |
| **Metric**      | Pointer to **string**                                                                         | Measurable attribute to compute. | [optional] |

## Methods

### NewMonitorFormulaAndFunctionEventQueryDefinitionCompute

`func NewMonitorFormulaAndFunctionEventQueryDefinitionCompute(aggregation MonitorFormulaAndFunctionEventAggregation) *MonitorFormulaAndFunctionEventQueryDefinitionCompute`

NewMonitorFormulaAndFunctionEventQueryDefinitionCompute instantiates a new MonitorFormulaAndFunctionEventQueryDefinitionCompute object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorFormulaAndFunctionEventQueryDefinitionComputeWithDefaults

`func NewMonitorFormulaAndFunctionEventQueryDefinitionComputeWithDefaults() *MonitorFormulaAndFunctionEventQueryDefinitionCompute`

NewMonitorFormulaAndFunctionEventQueryDefinitionComputeWithDefaults instantiates a new MonitorFormulaAndFunctionEventQueryDefinitionCompute object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregation

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) GetAggregation() MonitorFormulaAndFunctionEventAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) GetAggregationOk() (*MonitorFormulaAndFunctionEventAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) SetAggregation(v MonitorFormulaAndFunctionEventAggregation)`

SetAggregation sets Aggregation field to given value.

### GetInterval

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetric

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MonitorFormulaAndFunctionEventQueryDefinitionCompute) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
