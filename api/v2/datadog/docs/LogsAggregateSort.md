# LogsAggregateSort

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Aggregation** | Pointer to [**LogsAggregationFunction**](LogsAggregationFunction.md) |  | [optional] 
**Metric** | Pointer to **string** | The metric to sort by (only used for &#x60;type&#x3D;measure&#x60;) | [optional] 
**Order** | Pointer to [**LogsSortOrder**](LogsSortOrder.md) |  | [optional] 
**Type** | Pointer to [**LogsAggregateSortType**](LogsAggregateSortType.md) |  | [optional] [default to LOGSAGGREGATESORTTYPE_ALPHABETICAL]

## Methods

### NewLogsAggregateSort

`func NewLogsAggregateSort() *LogsAggregateSort`

NewLogsAggregateSort instantiates a new LogsAggregateSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsAggregateSortWithDefaults

`func NewLogsAggregateSortWithDefaults() *LogsAggregateSort`

NewLogsAggregateSortWithDefaults instantiates a new LogsAggregateSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *LogsAggregateSort) GetAggregation() LogsAggregationFunction`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *LogsAggregateSort) GetAggregationOk() (*LogsAggregationFunction, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *LogsAggregateSort) SetAggregation(v LogsAggregationFunction)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *LogsAggregateSort) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetMetric

`func (o *LogsAggregateSort) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *LogsAggregateSort) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *LogsAggregateSort) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *LogsAggregateSort) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetOrder

`func (o *LogsAggregateSort) GetOrder() LogsSortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *LogsAggregateSort) GetOrderOk() (*LogsSortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *LogsAggregateSort) SetOrder(v LogsSortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *LogsAggregateSort) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *LogsAggregateSort) GetType() LogsAggregateSortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsAggregateSort) GetTypeOk() (*LogsAggregateSortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsAggregateSort) SetType(v LogsAggregateSortType)`

SetType sets Type field to given value.

### HasType

`func (o *LogsAggregateSort) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


