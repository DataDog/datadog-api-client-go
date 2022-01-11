# MetricTagConfigurationCreateData

## Properties

| Name           | Type                                                                                               | Description                        | Notes                                               |
| -------------- | -------------------------------------------------------------------------------------------------- | ---------------------------------- | --------------------------------------------------- |
| **Attributes** | Pointer to [**MetricTagConfigurationCreateAttributes**](MetricTagConfigurationCreateAttributes.md) |                                    | [optional]                                          |
| **Id**         | **string**                                                                                         | The metric name for this resource. |
| **Type**       | [**MetricTagConfigurationType**](MetricTagConfigurationType.md)                                    |                                    | [default to METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS] |

## Methods

### NewMetricTagConfigurationCreateData

`func NewMetricTagConfigurationCreateData(id string, type_ MetricTagConfigurationType) *MetricTagConfigurationCreateData`

NewMetricTagConfigurationCreateData instantiates a new MetricTagConfigurationCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricTagConfigurationCreateDataWithDefaults

`func NewMetricTagConfigurationCreateDataWithDefaults() *MetricTagConfigurationCreateData`

NewMetricTagConfigurationCreateDataWithDefaults instantiates a new MetricTagConfigurationCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *MetricTagConfigurationCreateData) GetAttributes() MetricTagConfigurationCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricTagConfigurationCreateData) GetAttributesOk() (*MetricTagConfigurationCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricTagConfigurationCreateData) SetAttributes(v MetricTagConfigurationCreateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricTagConfigurationCreateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricTagConfigurationCreateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricTagConfigurationCreateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricTagConfigurationCreateData) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *MetricTagConfigurationCreateData) GetType() MetricTagConfigurationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricTagConfigurationCreateData) GetTypeOk() (*MetricTagConfigurationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricTagConfigurationCreateData) SetType(v MetricTagConfigurationType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
