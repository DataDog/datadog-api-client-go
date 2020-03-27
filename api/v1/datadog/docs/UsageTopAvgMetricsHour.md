# UsageTopAvgMetricsHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgMetricHour** | Pointer to **int64** | Average number of timeseries per hour in which the metric occurs. | [optional] 
**MaxMetricHour** | Pointer to **int64** | Maximum number of timeseries per hour in which the metric occurs. | [optional] 
**MetricCategory** | Pointer to [**UsageMetricCategory**](UsageMetricCategory.md) |  | [optional] 
**MetricName** | Pointer to **string** | Contains the custom metric name. | [optional] 

## Methods

### NewUsageTopAvgMetricsHour

`func NewUsageTopAvgMetricsHour() *UsageTopAvgMetricsHour`

NewUsageTopAvgMetricsHour instantiates a new UsageTopAvgMetricsHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageTopAvgMetricsHourWithDefaults

`func NewUsageTopAvgMetricsHourWithDefaults() *UsageTopAvgMetricsHour`

NewUsageTopAvgMetricsHourWithDefaults instantiates a new UsageTopAvgMetricsHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgMetricHour

`func (o *UsageTopAvgMetricsHour) GetAvgMetricHour() int64`

GetAvgMetricHour returns the AvgMetricHour field if non-nil, zero value otherwise.

### GetAvgMetricHourOk

`func (o *UsageTopAvgMetricsHour) GetAvgMetricHourOk() (*int64, bool)`

GetAvgMetricHourOk returns a tuple with the AvgMetricHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMetricHour

`func (o *UsageTopAvgMetricsHour) SetAvgMetricHour(v int64)`

SetAvgMetricHour sets AvgMetricHour field to given value.

### HasAvgMetricHour

`func (o *UsageTopAvgMetricsHour) HasAvgMetricHour() bool`

HasAvgMetricHour returns a boolean if a field has been set.

### GetMaxMetricHour

`func (o *UsageTopAvgMetricsHour) GetMaxMetricHour() int64`

GetMaxMetricHour returns the MaxMetricHour field if non-nil, zero value otherwise.

### GetMaxMetricHourOk

`func (o *UsageTopAvgMetricsHour) GetMaxMetricHourOk() (*int64, bool)`

GetMaxMetricHourOk returns a tuple with the MaxMetricHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMetricHour

`func (o *UsageTopAvgMetricsHour) SetMaxMetricHour(v int64)`

SetMaxMetricHour sets MaxMetricHour field to given value.

### HasMaxMetricHour

`func (o *UsageTopAvgMetricsHour) HasMaxMetricHour() bool`

HasMaxMetricHour returns a boolean if a field has been set.

### GetMetricCategory

`func (o *UsageTopAvgMetricsHour) GetMetricCategory() UsageMetricCategory`

GetMetricCategory returns the MetricCategory field if non-nil, zero value otherwise.

### GetMetricCategoryOk

`func (o *UsageTopAvgMetricsHour) GetMetricCategoryOk() (*UsageMetricCategory, bool)`

GetMetricCategoryOk returns a tuple with the MetricCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCategory

`func (o *UsageTopAvgMetricsHour) SetMetricCategory(v UsageMetricCategory)`

SetMetricCategory sets MetricCategory field to given value.

### HasMetricCategory

`func (o *UsageTopAvgMetricsHour) HasMetricCategory() bool`

HasMetricCategory returns a boolean if a field has been set.

### GetMetricName

`func (o *UsageTopAvgMetricsHour) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UsageTopAvgMetricsHour) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UsageTopAvgMetricsHour) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *UsageTopAvgMetricsHour) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


