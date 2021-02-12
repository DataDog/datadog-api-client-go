# MetricTagConfigurationUpdateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludePercentiles** | Pointer to **bool** | Toggle to include/exclude percentiles for a distribution metric. Defaults to false. Can only be applied to metrics that have a &#x60;metric_type&#x60; of &#x60;distribution&#x60;. | [optional] [default to false]
**Tags** | Pointer to **[]string** | A list of tag keys that will be queryable for your metric. | [optional] [default to []]

## Methods

### NewMetricTagConfigurationUpdateAttributes

`func NewMetricTagConfigurationUpdateAttributes() *MetricTagConfigurationUpdateAttributes`

NewMetricTagConfigurationUpdateAttributes instantiates a new MetricTagConfigurationUpdateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricTagConfigurationUpdateAttributesWithDefaults

`func NewMetricTagConfigurationUpdateAttributesWithDefaults() *MetricTagConfigurationUpdateAttributes`

NewMetricTagConfigurationUpdateAttributesWithDefaults instantiates a new MetricTagConfigurationUpdateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludePercentiles

`func (o *MetricTagConfigurationUpdateAttributes) GetIncludePercentiles() bool`

GetIncludePercentiles returns the IncludePercentiles field if non-nil, zero value otherwise.

### GetIncludePercentilesOk

`func (o *MetricTagConfigurationUpdateAttributes) GetIncludePercentilesOk() (*bool, bool)`

GetIncludePercentilesOk returns a tuple with the IncludePercentiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePercentiles

`func (o *MetricTagConfigurationUpdateAttributes) SetIncludePercentiles(v bool)`

SetIncludePercentiles sets IncludePercentiles field to given value.

### HasIncludePercentiles

`func (o *MetricTagConfigurationUpdateAttributes) HasIncludePercentiles() bool`

HasIncludePercentiles returns a boolean if a field has been set.

### GetTags

`func (o *MetricTagConfigurationUpdateAttributes) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricTagConfigurationUpdateAttributes) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricTagConfigurationUpdateAttributes) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricTagConfigurationUpdateAttributes) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


