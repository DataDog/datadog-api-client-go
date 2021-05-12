# MetricDistinctVolume

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**MetricDistinctVolumeAttributes**](MetricDistinctVolumeAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The metric name for this resource. | [optional] 
**Type** | Pointer to [**MetricDistinctVolumeType**](MetricDistinctVolumeType.md) |  | [optional] [default to METRICDISTINCTVOLUMETYPE_DISTINCT_METRIC_VOLUMES]

## Methods

### NewMetricDistinctVolume

`func NewMetricDistinctVolume() *MetricDistinctVolume`

NewMetricDistinctVolume instantiates a new MetricDistinctVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDistinctVolumeWithDefaults

`func NewMetricDistinctVolumeWithDefaults() *MetricDistinctVolume`

NewMetricDistinctVolumeWithDefaults instantiates a new MetricDistinctVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *MetricDistinctVolume) GetAttributes() MetricDistinctVolumeAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricDistinctVolume) GetAttributesOk() (*MetricDistinctVolumeAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricDistinctVolume) SetAttributes(v MetricDistinctVolumeAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricDistinctVolume) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricDistinctVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricDistinctVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricDistinctVolume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricDistinctVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MetricDistinctVolume) GetType() MetricDistinctVolumeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricDistinctVolume) GetTypeOk() (*MetricDistinctVolumeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricDistinctVolume) SetType(v MetricDistinctVolumeType)`

SetType sets Type field to given value.

### HasType

`func (o *MetricDistinctVolume) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


