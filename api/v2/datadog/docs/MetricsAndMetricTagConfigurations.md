# MetricsAndMetricTagConfigurations

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Id** | Pointer to **string** | The metric name for this resource. | [optional] 
**Type** | Pointer to [**MetricTagConfigurationType**](MetricTagConfigurationType.md) |  | [optional] [default to METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS]
**Attributes** | Pointer to [**MetricTagConfigurationAttributes**](MetricTagConfigurationAttributes.md) |  | [optional] 

## Methods

### NewMetricsAndMetricTagConfigurations

`func NewMetricsAndMetricTagConfigurations() *MetricsAndMetricTagConfigurations`

NewMetricsAndMetricTagConfigurations instantiates a new MetricsAndMetricTagConfigurations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsAndMetricTagConfigurationsWithDefaults

`func NewMetricsAndMetricTagConfigurationsWithDefaults() *MetricsAndMetricTagConfigurations`

NewMetricsAndMetricTagConfigurationsWithDefaults instantiates a new MetricsAndMetricTagConfigurations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricsAndMetricTagConfigurations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsAndMetricTagConfigurations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsAndMetricTagConfigurations) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricsAndMetricTagConfigurations) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MetricsAndMetricTagConfigurations) GetType() MetricTagConfigurationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricsAndMetricTagConfigurations) GetTypeOk() (*MetricTagConfigurationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricsAndMetricTagConfigurations) SetType(v MetricTagConfigurationType)`

SetType sets Type field to given value.

### HasType

`func (o *MetricsAndMetricTagConfigurations) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *MetricsAndMetricTagConfigurations) GetAttributes() MetricTagConfigurationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricsAndMetricTagConfigurations) GetAttributesOk() (*MetricTagConfigurationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricsAndMetricTagConfigurations) SetAttributes(v MetricTagConfigurationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricsAndMetricTagConfigurations) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


