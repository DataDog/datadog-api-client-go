# MetricVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**MetricIngestedIndexedVolumeAttributes**](MetricIngestedIndexedVolumeAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The metric name for this resource. | [optional] 
**Type** | Pointer to [**MetricIngestedIndexedVolumeType**](MetricIngestedIndexedVolumeType.md) |  | [optional] [default to "metric_volumes"]

## Methods

### NewMetricVolumes

`func NewMetricVolumes() *MetricVolumes`

NewMetricVolumes instantiates a new MetricVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricVolumesWithDefaults

`func NewMetricVolumesWithDefaults() *MetricVolumes`

NewMetricVolumesWithDefaults instantiates a new MetricVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MetricVolumes) GetAttributes() MetricIngestedIndexedVolumeAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricVolumes) GetAttributesOk() (*MetricIngestedIndexedVolumeAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricVolumes) SetAttributes(v MetricIngestedIndexedVolumeAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricVolumes) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricVolumes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricVolumes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricVolumes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricVolumes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MetricVolumes) GetType() MetricIngestedIndexedVolumeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricVolumes) GetTypeOk() (*MetricIngestedIndexedVolumeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricVolumes) SetType(v MetricIngestedIndexedVolumeType)`

SetType sets Type field to given value.

### HasType

`func (o *MetricVolumes) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


