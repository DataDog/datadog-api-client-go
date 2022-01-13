# LogsMetricUpdateData

## Properties

| Name           | Type                                                            | Description | Notes                                    |
| -------------- | --------------------------------------------------------------- | ----------- | ---------------------------------------- |
| **Attributes** | [**LogsMetricUpdateAttributes**](LogsMetricUpdateAttributes.md) |             |
| **Type**       | [**LogsMetricType**](LogsMetricType.md)                         |             | [default to LOGSMETRICTYPE_LOGS_METRICS] |

## Methods

### NewLogsMetricUpdateData

`func NewLogsMetricUpdateData(attributes LogsMetricUpdateAttributes, type_ LogsMetricType) *LogsMetricUpdateData`

NewLogsMetricUpdateData instantiates a new LogsMetricUpdateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsMetricUpdateDataWithDefaults

`func NewLogsMetricUpdateDataWithDefaults() *LogsMetricUpdateData`

NewLogsMetricUpdateDataWithDefaults instantiates a new LogsMetricUpdateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *LogsMetricUpdateData) GetAttributes() LogsMetricUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogsMetricUpdateData) GetAttributesOk() (*LogsMetricUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogsMetricUpdateData) SetAttributes(v LogsMetricUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### GetType

`func (o *LogsMetricUpdateData) GetType() LogsMetricType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsMetricUpdateData) GetTypeOk() (*LogsMetricType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsMetricUpdateData) SetType(v LogsMetricType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
