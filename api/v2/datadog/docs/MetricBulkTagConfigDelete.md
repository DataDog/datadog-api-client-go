# MetricBulkTagConfigDelete

## Properties

| Name           | Type                                                                                         | Description                                  | Notes                                                     |
| -------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------- | --------------------------------------------------------- |
| **Attributes** | Pointer to [**MetricBulkTagConfigDeleteAttributes**](MetricBulkTagConfigDeleteAttributes.md) |                                              | [optional]                                                |
| **Id**         | **string**                                                                                   | A text prefix to match against metric names. |
| **Type**       | [**MetricBulkConfigureTagsType**](MetricBulkConfigureTagsType.md)                            |                                              | [default to METRICBULKCONFIGURETAGSTYPE_BULK_MANAGE_TAGS] |

## Methods

### NewMetricBulkTagConfigDelete

`func NewMetricBulkTagConfigDelete(id string, type_ MetricBulkConfigureTagsType) *MetricBulkTagConfigDelete`

NewMetricBulkTagConfigDelete instantiates a new MetricBulkTagConfigDelete object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricBulkTagConfigDeleteWithDefaults

`func NewMetricBulkTagConfigDeleteWithDefaults() *MetricBulkTagConfigDelete`

NewMetricBulkTagConfigDeleteWithDefaults instantiates a new MetricBulkTagConfigDelete object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *MetricBulkTagConfigDelete) GetAttributes() MetricBulkTagConfigDeleteAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricBulkTagConfigDelete) GetAttributesOk() (*MetricBulkTagConfigDeleteAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricBulkTagConfigDelete) SetAttributes(v MetricBulkTagConfigDeleteAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricBulkTagConfigDelete) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricBulkTagConfigDelete) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricBulkTagConfigDelete) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricBulkTagConfigDelete) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *MetricBulkTagConfigDelete) GetType() MetricBulkConfigureTagsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricBulkTagConfigDelete) GetTypeOk() (*MetricBulkConfigureTagsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricBulkTagConfigDelete) SetType(v MetricBulkConfigureTagsType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
