# MetricIngestedIndexedVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**MetricIngestedIndexedVolumeAttributes**](MetricIngestedIndexedVolumeAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The metric name for this resource. | [optional] 
**Type** | Pointer to [**MetricIngestedIndexedVolumeType**](MetricIngestedIndexedVolumeType.md) |  | [optional] [default to METRICINGESTEDINDEXEDVOLUMETYPE_METRIC_VOLUMES]

## Methods

### NewMetricIngestedIndexedVolume

`func NewMetricIngestedIndexedVolume() *MetricIngestedIndexedVolume`

NewMetricIngestedIndexedVolume instantiates a new MetricIngestedIndexedVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricIngestedIndexedVolumeWithDefaults

`func NewMetricIngestedIndexedVolumeWithDefaults() *MetricIngestedIndexedVolume`

NewMetricIngestedIndexedVolumeWithDefaults instantiates a new MetricIngestedIndexedVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MetricIngestedIndexedVolume) GetAttributes() MetricIngestedIndexedVolumeAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricIngestedIndexedVolume) GetAttributesOk() (*MetricIngestedIndexedVolumeAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricIngestedIndexedVolume) SetAttributes(v MetricIngestedIndexedVolumeAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricIngestedIndexedVolume) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricIngestedIndexedVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricIngestedIndexedVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricIngestedIndexedVolume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricIngestedIndexedVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MetricIngestedIndexedVolume) GetType() MetricIngestedIndexedVolumeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricIngestedIndexedVolume) GetTypeOk() (*MetricIngestedIndexedVolumeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricIngestedIndexedVolume) SetType(v MetricIngestedIndexedVolumeType)`

SetType sets Type field to given value.

### HasType

`func (o *MetricIngestedIndexedVolume) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


