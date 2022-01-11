# LogsAggregateResponse

## Properties

| Name     | Type                                                                     | Description | Notes      |
| -------- | ------------------------------------------------------------------------ | ----------- | ---------- |
| **Data** | Pointer to [**LogsAggregateResponseData**](LogsAggregateResponseData.md) |             | [optional] |
| **Meta** | Pointer to [**LogsResponseMetadata**](LogsResponseMetadata.md)           |             | [optional] |

## Methods

### NewLogsAggregateResponse

`func NewLogsAggregateResponse() *LogsAggregateResponse`

NewLogsAggregateResponse instantiates a new LogsAggregateResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogsAggregateResponseWithDefaults

`func NewLogsAggregateResponseWithDefaults() *LogsAggregateResponse`

NewLogsAggregateResponseWithDefaults instantiates a new LogsAggregateResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *LogsAggregateResponse) GetData() LogsAggregateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LogsAggregateResponse) GetDataOk() (*LogsAggregateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LogsAggregateResponse) SetData(v LogsAggregateResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *LogsAggregateResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *LogsAggregateResponse) GetMeta() LogsResponseMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LogsAggregateResponse) GetMetaOk() (*LogsResponseMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LogsAggregateResponse) SetMeta(v LogsResponseMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LogsAggregateResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
