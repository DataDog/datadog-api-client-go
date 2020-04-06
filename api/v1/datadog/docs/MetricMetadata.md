# MetricMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Metric description. | [optional] 
**Integration** | Pointer to **string** | Name of the integration that sent the metric if applicable. | [optional] [readonly] 
**PerUnit** | Pointer to **string** | Per unit of the metric such as &#x60;second&#x60; in &#x60;bytes per second&#x60;. | [optional] 
**ShortName** | Pointer to **string** | A more human-readable and abbreviated version of the metric name. | [optional] 
**StatsdInterval** | Pointer to **int64** | Statsd flush interval of the metric in seconds if applicable. | [optional] 
**Type** | Pointer to **string** | Metric type such as &#x60;gauge&#x60; or &#x60;rate&#x60;. | [optional] 
**Unit** | Pointer to **string** | Primary unit of the metric such as &#x60;byte&#x60; or &#x60;operation&#x60;. | [optional] 

## Methods

### NewMetricMetadata

`func NewMetricMetadata() *MetricMetadata`

NewMetricMetadata instantiates a new MetricMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricMetadataWithDefaults

`func NewMetricMetadataWithDefaults() *MetricMetadata`

NewMetricMetadataWithDefaults instantiates a new MetricMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MetricMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIntegration

`func (o *MetricMetadata) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *MetricMetadata) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *MetricMetadata) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *MetricMetadata) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetPerUnit

`func (o *MetricMetadata) GetPerUnit() string`

GetPerUnit returns the PerUnit field if non-nil, zero value otherwise.

### GetPerUnitOk

`func (o *MetricMetadata) GetPerUnitOk() (*string, bool)`

GetPerUnitOk returns a tuple with the PerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUnit

`func (o *MetricMetadata) SetPerUnit(v string)`

SetPerUnit sets PerUnit field to given value.

### HasPerUnit

`func (o *MetricMetadata) HasPerUnit() bool`

HasPerUnit returns a boolean if a field has been set.

### GetShortName

`func (o *MetricMetadata) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *MetricMetadata) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *MetricMetadata) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *MetricMetadata) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetStatsdInterval

`func (o *MetricMetadata) GetStatsdInterval() int64`

GetStatsdInterval returns the StatsdInterval field if non-nil, zero value otherwise.

### GetStatsdIntervalOk

`func (o *MetricMetadata) GetStatsdIntervalOk() (*int64, bool)`

GetStatsdIntervalOk returns a tuple with the StatsdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsdInterval

`func (o *MetricMetadata) SetStatsdInterval(v int64)`

SetStatsdInterval sets StatsdInterval field to given value.

### HasStatsdInterval

`func (o *MetricMetadata) HasStatsdInterval() bool`

HasStatsdInterval returns a boolean if a field has been set.

### GetType

`func (o *MetricMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetricMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *MetricMetadata) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricMetadata) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricMetadata) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricMetadata) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


