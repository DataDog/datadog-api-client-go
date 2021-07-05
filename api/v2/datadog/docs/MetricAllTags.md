# MetricAllTags

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**MetricAllTagsAttributes**](MetricAllTagsAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The metric name for this resource. | [optional] 
**Type** | Pointer to [**MetricType**](MetricType.md) |  | [optional] [default to METRICTYPE_METRICS]

## Methods

### NewMetricAllTags

`func NewMetricAllTags() *MetricAllTags`

NewMetricAllTags instantiates a new MetricAllTags object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricAllTagsWithDefaults

`func NewMetricAllTagsWithDefaults() *MetricAllTags`

NewMetricAllTagsWithDefaults instantiates a new MetricAllTags object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *MetricAllTags) GetAttributes() MetricAllTagsAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricAllTags) GetAttributesOk() (*MetricAllTagsAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricAllTags) SetAttributes(v MetricAllTagsAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricAllTags) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricAllTags) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricAllTags) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricAllTags) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetricAllTags) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MetricAllTags) GetType() MetricType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricAllTags) GetTypeOk() (*MetricType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricAllTags) SetType(v MetricType)`

SetType sets Type field to given value.

### HasType

`func (o *MetricAllTags) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


