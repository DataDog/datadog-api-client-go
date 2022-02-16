# MetricBulkTagConfigCreate

## Properties

| Name           | Type                                                                                         | Description                                  | Notes                                                     |
| -------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------- | --------------------------------------------------------- |
| **Attributes** | Pointer to [**MetricBulkTagConfigCreateAttributes**](MetricBulkTagConfigCreateAttributes.md) |                                              | [optional]                                                |
| **Id**         | **string**                                                                                   | A text prefix to match against metric names. |
| **Type**       | [**MetricBulkConfigureTagsType**](MetricBulkConfigureTagsType.md)                            |                                              | [default to METRICBULKCONFIGURETAGSTYPE_BULK_MANAGE_TAGS] |

## Methods

### NewMetricBulkTagConfigCreate

`func NewMetricBulkTagConfigCreate(id string, type_ MetricBulkConfigureTagsType) *MetricBulkTagConfigCreate`

NewMetricBulkTagConfigCreate instantiates a new MetricBulkTagConfigCreate object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMetricBulkTagConfigCreateWithDefaults

`func NewMetricBulkTagConfigCreateWithDefaults() *MetricBulkTagConfigCreate`

NewMetricBulkTagConfigCreateWithDefaults instantiates a new MetricBulkTagConfigCreate object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *MetricBulkTagConfigCreate) GetAttributes() MetricBulkTagConfigCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricBulkTagConfigCreate) GetAttributesOk() (*MetricBulkTagConfigCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricBulkTagConfigCreate) SetAttributes(v MetricBulkTagConfigCreateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MetricBulkTagConfigCreate) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *MetricBulkTagConfigCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricBulkTagConfigCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricBulkTagConfigCreate) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *MetricBulkTagConfigCreate) GetType() MetricBulkConfigureTagsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricBulkTagConfigCreate) GetTypeOk() (*MetricBulkConfigureTagsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricBulkTagConfigCreate) SetType(v MetricBulkConfigureTagsType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
