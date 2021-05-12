# FormulaAndFunctionEventQueryGroupBySort

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Aggregation** | [**FormulaAndFunctionEventAggregation**](FormulaAndFunctionEventAggregation.md) |  | 
**Metric** | Pointer to **string** | Metric used for sorting group by results. | [optional] 
**Order** | Pointer to [**QuerySortOrder**](QuerySortOrder.md) |  | [optional] [default to QUERYSORTORDER_DESC]

## Methods

### NewFormulaAndFunctionEventQueryGroupBySort

`func NewFormulaAndFunctionEventQueryGroupBySort(aggregation FormulaAndFunctionEventAggregation, ) *FormulaAndFunctionEventQueryGroupBySort`

NewFormulaAndFunctionEventQueryGroupBySort instantiates a new FormulaAndFunctionEventQueryGroupBySort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormulaAndFunctionEventQueryGroupBySortWithDefaults

`func NewFormulaAndFunctionEventQueryGroupBySortWithDefaults() *FormulaAndFunctionEventQueryGroupBySort`

NewFormulaAndFunctionEventQueryGroupBySortWithDefaults instantiates a new FormulaAndFunctionEventQueryGroupBySort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *FormulaAndFunctionEventQueryGroupBySort) GetAggregation() FormulaAndFunctionEventAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *FormulaAndFunctionEventQueryGroupBySort) GetAggregationOk() (*FormulaAndFunctionEventAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *FormulaAndFunctionEventQueryGroupBySort) SetAggregation(v FormulaAndFunctionEventAggregation)`

SetAggregation sets Aggregation field to given value.


### GetMetric

`func (o *FormulaAndFunctionEventQueryGroupBySort) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *FormulaAndFunctionEventQueryGroupBySort) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *FormulaAndFunctionEventQueryGroupBySort) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *FormulaAndFunctionEventQueryGroupBySort) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetOrder

`func (o *FormulaAndFunctionEventQueryGroupBySort) GetOrder() QuerySortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FormulaAndFunctionEventQueryGroupBySort) GetOrderOk() (*QuerySortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FormulaAndFunctionEventQueryGroupBySort) SetOrder(v QuerySortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FormulaAndFunctionEventQueryGroupBySort) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


