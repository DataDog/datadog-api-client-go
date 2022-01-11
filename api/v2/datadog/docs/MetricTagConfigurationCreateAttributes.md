# MetricTagConfigurationCreateAttributes

## Properties

| Name                   | Type                                                                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Notes                                                |
| ---------------------- | ----------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------- |
| **Aggregations**       | Pointer to [**[]MetricCustomAggregation**](MetricCustomAggregation.md)        | A list of queryable aggregation combinations for a count, rate, or gauge metric. By default, count and rate metrics require the (time: sum, space: sum) aggregation and Gauge metrics require the (time: avg, space: avg) aggregation. Additional time &amp; space combinations are also available: - time: avg, space: avg - time: avg, space: max - time: avg, space: min - time: avg, space: sum - time: count, space: sum - time: max, space: max - time: min, space: min - time: sum, space: avg - time: sum, space: sum Can only be applied to metrics that have a &#x60;metric_type&#x60; of &#x60;count&#x60;, &#x60;rate&#x60;, or &#x60;gauge&#x60;. | [optional]                                           |
| **IncludePercentiles** | Pointer to **bool**                                                           | Toggle to include/exclude percentiles for a distribution metric. Defaults to false. Can only be applied to metrics that have a &#x60;metric_type&#x60; of &#x60;distribution&#x60;.                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | [optional] [default to false]                        |
| **MetricType**         | [**MetricTagConfigurationMetricTypes**](MetricTagConfigurationMetricTypes.md) |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | [default to METRICTAGCONFIGURATIONMETRICTYPES_GAUGE] |
| **Tags**               | **[]string**                                                                  | A list of tag keys that will be queryable for your metric.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | [default to []]                                      |

## Methods

### NewMetricTagConfigurationCreateAttributes

`func NewMetricTagConfigurationCreateAttributes(metricType MetricTagConfigurationMetricTypes, tags []string) *MetricTagConfigurationCreateAttributes`

NewMetricTagConfigurationCreateAttributes instantiates a new MetricTagConfigurationCreateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricTagConfigurationCreateAttributesWithDefaults

`func NewMetricTagConfigurationCreateAttributesWithDefaults() *MetricTagConfigurationCreateAttributes`

NewMetricTagConfigurationCreateAttributesWithDefaults instantiates a new MetricTagConfigurationCreateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregations

`func (o *MetricTagConfigurationCreateAttributes) GetAggregations() []MetricCustomAggregation`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *MetricTagConfigurationCreateAttributes) GetAggregationsOk() (*[]MetricCustomAggregation, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *MetricTagConfigurationCreateAttributes) SetAggregations(v []MetricCustomAggregation)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *MetricTagConfigurationCreateAttributes) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetIncludePercentiles

`func (o *MetricTagConfigurationCreateAttributes) GetIncludePercentiles() bool`

GetIncludePercentiles returns the IncludePercentiles field if non-nil, zero value otherwise.

### GetIncludePercentilesOk

`func (o *MetricTagConfigurationCreateAttributes) GetIncludePercentilesOk() (*bool, bool)`

GetIncludePercentilesOk returns a tuple with the IncludePercentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePercentiles

`func (o *MetricTagConfigurationCreateAttributes) SetIncludePercentiles(v bool)`

SetIncludePercentiles sets IncludePercentiles field to given value.

### HasIncludePercentiles

`func (o *MetricTagConfigurationCreateAttributes) HasIncludePercentiles() bool`

HasIncludePercentiles returns a boolean if a field has been set.

### GetMetricType

`func (o *MetricTagConfigurationCreateAttributes) GetMetricType() MetricTagConfigurationMetricTypes`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *MetricTagConfigurationCreateAttributes) GetMetricTypeOk() (*MetricTagConfigurationMetricTypes, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *MetricTagConfigurationCreateAttributes) SetMetricType(v MetricTagConfigurationMetricTypes)`

SetMetricType sets MetricType field to given value.

### GetTags

`func (o *MetricTagConfigurationCreateAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricTagConfigurationCreateAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricTagConfigurationCreateAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
