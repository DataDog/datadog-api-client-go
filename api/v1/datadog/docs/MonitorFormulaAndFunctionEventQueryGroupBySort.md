# MonitorFormulaAndFunctionEventQueryGroupBySort

## Properties

| Name            | Type                                                                                          | Description                               | Notes                                       |
| --------------- | --------------------------------------------------------------------------------------------- | ----------------------------------------- | ------------------------------------------- |
| **Aggregation** | [**MonitorFormulaAndFunctionEventAggregation**](MonitorFormulaAndFunctionEventAggregation.md) |                                           |
| **Metric**      | Pointer to **string**                                                                         | Metric used for sorting group by results. | [optional]                                  |
| **Order**       | Pointer to [**QuerySortOrder**](QuerySortOrder.md)                                            |                                           | [optional] [default to QUERYSORTORDER_DESC] |

## Methods

### NewMonitorFormulaAndFunctionEventQueryGroupBySort

`func NewMonitorFormulaAndFunctionEventQueryGroupBySort(aggregation MonitorFormulaAndFunctionEventAggregation) *MonitorFormulaAndFunctionEventQueryGroupBySort`

NewMonitorFormulaAndFunctionEventQueryGroupBySort instantiates a new MonitorFormulaAndFunctionEventQueryGroupBySort object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorFormulaAndFunctionEventQueryGroupBySortWithDefaults

`func NewMonitorFormulaAndFunctionEventQueryGroupBySortWithDefaults() *MonitorFormulaAndFunctionEventQueryGroupBySort`

NewMonitorFormulaAndFunctionEventQueryGroupBySortWithDefaults instantiates a new MonitorFormulaAndFunctionEventQueryGroupBySort object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregation

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) GetAggregation() MonitorFormulaAndFunctionEventAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) GetAggregationOk() (*MonitorFormulaAndFunctionEventAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) SetAggregation(v MonitorFormulaAndFunctionEventAggregation)`

SetAggregation sets Aggregation field to given value.

### GetMetric

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetOrder

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) GetOrder() QuerySortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) GetOrderOk() (*QuerySortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) SetOrder(v QuerySortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MonitorFormulaAndFunctionEventQueryGroupBySort) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
