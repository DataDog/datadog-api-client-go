# TimeSeriesFormulaAndFunctionEventQueryGroupBySort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | [**FormulaAndFunctionEventAggregation**](FormulaAndFunctionEventAggregation.md) |  | 
**Metric** | Pointer to **string** | Metric used for sorting group by results. | [optional] 
**Order** | Pointer to [**QuerySortOrder**](QuerySortOrder.md) |  | [optional] [default to "desc"]

## Methods

### NewTimeSeriesFormulaAndFunctionEventQueryGroupBySort

`func NewTimeSeriesFormulaAndFunctionEventQueryGroupBySort(aggregation FormulaAndFunctionEventAggregation, ) *TimeSeriesFormulaAndFunctionEventQueryGroupBySort`

NewTimeSeriesFormulaAndFunctionEventQueryGroupBySort instantiates a new TimeSeriesFormulaAndFunctionEventQueryGroupBySort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesFormulaAndFunctionEventQueryGroupBySortWithDefaults

`func NewTimeSeriesFormulaAndFunctionEventQueryGroupBySortWithDefaults() *TimeSeriesFormulaAndFunctionEventQueryGroupBySort`

NewTimeSeriesFormulaAndFunctionEventQueryGroupBySortWithDefaults instantiates a new TimeSeriesFormulaAndFunctionEventQueryGroupBySort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) GetAggregation() FormulaAndFunctionEventAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) GetAggregationOk() (*FormulaAndFunctionEventAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) SetAggregation(v FormulaAndFunctionEventAggregation)`

SetAggregation sets Aggregation field to given value.


### GetMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetOrder

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) GetOrder() QuerySortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) GetOrderOk() (*QuerySortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) SetOrder(v QuerySortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TimeSeriesFormulaAndFunctionEventQueryGroupBySort) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


