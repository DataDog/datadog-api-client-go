# LogsMetricResponseData

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Attributes** | Pointer to [**LogsMetricResponseAttributes**](LogsMetricResponseAttributes.md) |  | [optional] 
**Id** | Pointer to **string** | The name of the log-based metric. | [optional] 
**Type** | Pointer to [**LogsMetricType**](LogsMetricType.md) |  | [optional] [default to LOGSMETRICTYPE_LOGS_METRICS]

## Methods

### NewLogsMetricResponseData

`func NewLogsMetricResponseData() *LogsMetricResponseData`

NewLogsMetricResponseData instantiates a new LogsMetricResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsMetricResponseDataWithDefaults

`func NewLogsMetricResponseDataWithDefaults() *LogsMetricResponseData`

NewLogsMetricResponseDataWithDefaults instantiates a new LogsMetricResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *LogsMetricResponseData) GetAttributes() LogsMetricResponseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogsMetricResponseData) GetAttributesOk() (*LogsMetricResponseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogsMetricResponseData) SetAttributes(v LogsMetricResponseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LogsMetricResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *LogsMetricResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogsMetricResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogsMetricResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogsMetricResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LogsMetricResponseData) GetType() LogsMetricType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogsMetricResponseData) GetTypeOk() (*LogsMetricType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogsMetricResponseData) SetType(v LogsMetricType)`

SetType sets Type field to given value.

### HasType

`func (o *LogsMetricResponseData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


