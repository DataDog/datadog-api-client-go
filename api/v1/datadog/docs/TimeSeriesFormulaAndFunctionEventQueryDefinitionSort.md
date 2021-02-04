# TimeSeriesFormulaAndFunctionEventQueryDefinitionSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | [**FormulaAndFunctionEventAggregation**](FormulaAndFunctionEventAggregation.md) |  | 
**Metric** | Pointer to **string** | Metric used for sorting group by results. | [optional] 
**Order** | Pointer to [**QuerySortOrder**](QuerySortOrder.md) |  | [optional] [default to "desc"]
**Type** | Pointer to [**FormulaAndFunctionEventsSortType**](FormulaAndFunctionEventsSortType.md) |  | [optional] 

## Methods

### NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSort

`func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSort(aggregation FormulaAndFunctionEventAggregation, ) *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort`

NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSort instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSortWithDefaults

`func NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSortWithDefaults() *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort`

NewTimeSeriesFormulaAndFunctionEventQueryDefinitionSortWithDefaults instantiates a new TimeSeriesFormulaAndFunctionEventQueryDefinitionSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetAggregation() FormulaAndFunctionEventAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetAggregationOk() (*FormulaAndFunctionEventAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) SetAggregation(v FormulaAndFunctionEventAggregation)`

SetAggregation sets Aggregation field to given value.


### GetMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetOrder

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetOrder() QuerySortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetOrderOk() (*QuerySortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) SetOrder(v QuerySortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetType() FormulaAndFunctionEventsSortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) GetTypeOk() (*FormulaAndFunctionEventsSortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) SetType(v FormulaAndFunctionEventsSortType)`

SetType sets Type field to given value.

### HasType

`func (o *TimeSeriesFormulaAndFunctionEventQueryDefinitionSort) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


