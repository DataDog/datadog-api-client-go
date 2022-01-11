# MonitorOptionsAggregation

## Properties

| Name        | Type                  | Description                         | Notes      |
| ----------- | --------------------- | ----------------------------------- | ---------- |
| **GroupBy** | Pointer to **string** | Group to break down the monitor on. | [optional] |
| **Metric**  | Pointer to **string** | Metric name used in the monitor.    | [optional] |
| **Type**    | Pointer to **string** | Metric type used in the monitor.    | [optional] |

## Methods

### NewMonitorOptionsAggregation

`func NewMonitorOptionsAggregation() *MonitorOptionsAggregation`

NewMonitorOptionsAggregation instantiates a new MonitorOptionsAggregation object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorOptionsAggregationWithDefaults

`func NewMonitorOptionsAggregationWithDefaults() *MonitorOptionsAggregation`

NewMonitorOptionsAggregationWithDefaults instantiates a new MonitorOptionsAggregation object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetGroupBy

`func (o *MonitorOptionsAggregation) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MonitorOptionsAggregation) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MonitorOptionsAggregation) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MonitorOptionsAggregation) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetMetric

`func (o *MonitorOptionsAggregation) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MonitorOptionsAggregation) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MonitorOptionsAggregation) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MonitorOptionsAggregation) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetType

`func (o *MonitorOptionsAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorOptionsAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorOptionsAggregation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitorOptionsAggregation) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
