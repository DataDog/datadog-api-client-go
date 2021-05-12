# LogsListRequest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Index** | Pointer to **string** | The log index on which the request is performed. For multi-index organizations, the default is all live indexes. Historical indexes of rehydrated logs must be specified. | [optional] 
**Limit** | Pointer to **int32** | Number of logs return in the response. | [optional] 
**Query** | Pointer to **string** | The search query - following the log search syntax. | [optional] 
**Sort** | Pointer to [**LogsSort**](LogsSort.md) |  | [optional] 
**StartAt** | Pointer to **string** | Hash identifier of the first log to return in the list, available in a log &#x60;id&#x60; attribute. This parameter is used for the pagination feature.  **Note**: This parameter is ignored if the corresponding log is out of the scope of the specified time window. | [optional] 
**Time** | [**LogsListRequestTime**](LogsListRequestTime.md) |  | 

## Methods

### NewLogsListRequest

`func NewLogsListRequest(time LogsListRequestTime, ) *LogsListRequest`

NewLogsListRequest instantiates a new LogsListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsListRequestWithDefaults

`func NewLogsListRequestWithDefaults() *LogsListRequest`

NewLogsListRequestWithDefaults instantiates a new LogsListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *LogsListRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *LogsListRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *LogsListRequest) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *LogsListRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLimit

`func (o *LogsListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LogsListRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LogsListRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LogsListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetQuery

`func (o *LogsListRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsListRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogsListRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *LogsListRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSort

`func (o *LogsListRequest) GetSort() LogsSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LogsListRequest) GetSortOk() (*LogsSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LogsListRequest) SetSort(v LogsSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LogsListRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartAt

`func (o *LogsListRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *LogsListRequest) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *LogsListRequest) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *LogsListRequest) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTime

`func (o *LogsListRequest) GetTime() LogsListRequestTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LogsListRequest) GetTimeOk() (*LogsListRequestTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *LogsListRequest) SetTime(v LogsListRequestTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


