# MetricBulkTagConfigStatus

## Properties

| Name           | Type                                                                                         | Description                                  | Notes                                                     |
| -------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------- | --------------------------------------------------------- |
| **Attributes** | Pointer to [**MetricBulkTagConfigStatusAttributes**](MetricBulkTagConfigStatusAttributes.md) |                                              | [optional]                                                |
| **Id**         | **string**                                                                                   | A text prefix to match against metric names. |
| **Type**       | [**MetricBulkConfigureTagsType**](MetricBulkConfigureTagsType.md)                            |                                              | [default to METRICBULKCONFIGURETAGSTYPE_BULK_MANAGE_TAGS] |

## Methods

### NewMetricBulkTagConfigStatus

`func NewMetricBulkTagConfigStatus(id string, type_ MetricBulkConfigureTagsType) *MetricBulkTagConfigStatus`

NewMetricBulkTagConfigStatus instantiates a new MetricBulkTagConfigStatus object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricBulkTagConfigStatusWithDefaults

`func NewMetricBulkTagConfigStatusWithDefaults() *MetricBulkTagConfigStatus`

NewMetricBulkTagConfigStatusWithDefaults instantiates a new MetricBulkTagConfigStatus object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *MetricBulkTagConfigStatus) GetAttributes() MetricBulkTagConfigStatusAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricBulkTagConfigStatus) GetAttributesOk() (*MetricBulkTagConfigStatusAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricBulkTagConfigStatus) SetAttributes(v MetricBulkTagConfigStatusAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricBulkTagConfigStatus) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricBulkTagConfigStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricBulkTagConfigStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricBulkTagConfigStatus) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *MetricBulkTagConfigStatus) GetType() MetricBulkConfigureTagsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricBulkTagConfigStatus) GetTypeOk() (*MetricBulkConfigureTagsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricBulkTagConfigStatus) SetType(v MetricBulkConfigureTagsType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
