# MetricTagConfigurationAttributes

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Aggregations** | Pointer to [**[]MetricCustomAggregation**](MetricCustomAggregation.md) | A list of queryable aggregation combinations for a count, rate, or gauge metric. By default, count and rate metrics require the (time: sum, space: sum) aggregation and Gauge metrics require the (time: avg, space: avg) aggregation. Additional time &amp; space combinations are also available:  - time: avg, space: avg - time: avg, space: max - time: avg, space: min - time: avg, space: sum - time: count, space: sum - time: max, space: max - time: min, space: min - time: sum, space: avg - time: sum, space: sum  Can only be applied to metrics that have a &#x60;metric_type&#x60; of &#x60;count&#x60;, &#x60;rate&#x60;, or &#x60;gauge&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp when the tag configuration was created. | [optional] 
**IncludePercentiles** | Pointer to **bool** | Toggle to turn on/off percentile aggregations for distribution metrics. Only present when the &#x60;metric_type&#x60; is &#x60;distribution&#x60;. | [optional] 
**MetricType** | Pointer to [**MetricTagConfigurationMetricTypes**](MetricTagConfigurationMetricTypes.md) |  | [optional] [default to METRICTAGCONFIGURATIONMETRICTYPES_GAUGE]
**ModifiedAt** | Pointer to **time.Time** | Timestamp when the tag configuration was last modified. | [optional] 
**Tags** | Pointer to **[]string** | List of tag keys on which to group. | [optional] 

## Methods

### NewMetricTagConfigurationAttributes

`func NewMetricTagConfigurationAttributes() *MetricTagConfigurationAttributes`

NewMetricTagConfigurationAttributes instantiates a new MetricTagConfigurationAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricTagConfigurationAttributesWithDefaults

`func NewMetricTagConfigurationAttributesWithDefaults() *MetricTagConfigurationAttributes`

NewMetricTagConfigurationAttributesWithDefaults instantiates a new MetricTagConfigurationAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAggregations

`func (o *MetricTagConfigurationAttributes) GetAggregations() []MetricCustomAggregation`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *MetricTagConfigurationAttributes) GetAggregationsOk() (*[]MetricCustomAggregation, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *MetricTagConfigurationAttributes) SetAggregations(v []MetricCustomAggregation)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *MetricTagConfigurationAttributes) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MetricTagConfigurationAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricTagConfigurationAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricTagConfigurationAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MetricTagConfigurationAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIncludePercentiles

`func (o *MetricTagConfigurationAttributes) GetIncludePercentiles() bool`

GetIncludePercentiles returns the IncludePercentiles field if non-nil, zero value otherwise.

### GetIncludePercentilesOk

`func (o *MetricTagConfigurationAttributes) GetIncludePercentilesOk() (*bool, bool)`

GetIncludePercentilesOk returns a tuple with the IncludePercentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePercentiles

`func (o *MetricTagConfigurationAttributes) SetIncludePercentiles(v bool)`

SetIncludePercentiles sets IncludePercentiles field to given value.

### HasIncludePercentiles

`func (o *MetricTagConfigurationAttributes) HasIncludePercentiles() bool`

HasIncludePercentiles returns a boolean if a field has been set.

### GetMetricType

`func (o *MetricTagConfigurationAttributes) GetMetricType() MetricTagConfigurationMetricTypes`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *MetricTagConfigurationAttributes) GetMetricTypeOk() (*MetricTagConfigurationMetricTypes, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *MetricTagConfigurationAttributes) SetMetricType(v MetricTagConfigurationMetricTypes)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *MetricTagConfigurationAttributes) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetModifiedAt

`func (o *MetricTagConfigurationAttributes) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *MetricTagConfigurationAttributes) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *MetricTagConfigurationAttributes) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *MetricTagConfigurationAttributes) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetTags

`func (o *MetricTagConfigurationAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricTagConfigurationAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricTagConfigurationAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricTagConfigurationAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


