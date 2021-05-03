# MetricTagConfigurationUpdateData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**MetricTagConfigurationUpdateAttributes**](MetricTagConfigurationUpdateAttributes.md) |  | [optional] 
**Id** | **string** | The metric name for this resource. | 
**Type** | [**MetricTagConfigurationType**](MetricTagConfigurationType.md) |  | [default to METRICTAGCONFIGURATIONTYPE_MANAGE_TAGS]

## Methods

### NewMetricTagConfigurationUpdateData

`func NewMetricTagConfigurationUpdateData(id string, type_ MetricTagConfigurationType, ) *MetricTagConfigurationUpdateData`

NewMetricTagConfigurationUpdateData instantiates a new MetricTagConfigurationUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricTagConfigurationUpdateDataWithDefaults

`func NewMetricTagConfigurationUpdateDataWithDefaults() *MetricTagConfigurationUpdateData`

NewMetricTagConfigurationUpdateDataWithDefaults instantiates a new MetricTagConfigurationUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MetricTagConfigurationUpdateData) GetAttributes() MetricTagConfigurationUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricTagConfigurationUpdateData) GetAttributesOk() (*MetricTagConfigurationUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricTagConfigurationUpdateData) SetAttributes(v MetricTagConfigurationUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricTagConfigurationUpdateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricTagConfigurationUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricTagConfigurationUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricTagConfigurationUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MetricTagConfigurationUpdateData) GetType() MetricTagConfigurationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricTagConfigurationUpdateData) GetTypeOk() (*MetricTagConfigurationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricTagConfigurationUpdateData) SetType(v MetricTagConfigurationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


