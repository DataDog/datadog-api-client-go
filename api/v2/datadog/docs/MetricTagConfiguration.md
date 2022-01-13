# MetricTagConfiguration

## Properties

| Name           | Type                                                                                   | Description                        | Notes                                                          |
| -------------- | -------------------------------------------------------------------------------------- | ---------------------------------- | -------------------------------------------------------------- |
| **Attributes** | Pointer to [**MetricTagConfigurationAttributes**](MetricTagConfigurationAttributes.md) |                                    | [optional]                                                     |
| **Id**         | Pointer to **string**                                                                  | The metric name for this resource. | [optional]                                                     |
| **Type**       | Pointer to [**MetricTagConfigurationType**](MetricTagConfigurationType.md)             |                                    | [optional] [default to METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS] |

## Methods

### NewMetricTagConfiguration

`func NewMetricTagConfiguration() *MetricTagConfiguration`

NewMetricTagConfiguration instantiates a new MetricTagConfiguration object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricTagConfigurationWithDefaults

`func NewMetricTagConfigurationWithDefaults() *MetricTagConfiguration`

NewMetricTagConfigurationWithDefaults instantiates a new MetricTagConfiguration object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *MetricTagConfiguration) GetAttributes() MetricTagConfigurationAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricTagConfiguration) GetAttributesOk() (*MetricTagConfigurationAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricTagConfiguration) SetAttributes(v MetricTagConfigurationAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricTagConfiguration) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricTagConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricTagConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricTagConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricTagConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MetricTagConfiguration) GetType() MetricTagConfigurationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricTagConfiguration) GetTypeOk() (*MetricTagConfigurationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricTagConfiguration) SetType(v MetricTagConfigurationType)`

SetType sets Type field to given value.

### HasType

`func (o *MetricTagConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
