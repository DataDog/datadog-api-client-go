# LogsListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** | For multi-index organizations, the log index in which the request is performed. | [optional] 
**Limit** | Pointer to **int32** | Number of logs return in the response. | [optional] 
**Query** | Pointer to **string** | The search query - following the Log search syntax. | 
**Sort** | Pointer to [**LogsSort**](LogsSort.md) |  | [optional] 
**StartAt** | Pointer to **string** | Hash identifier of the first log to return in the list, available in a log &#x60;id&#x60; attribute. This parameter is used for the pagination feature. **Note**: this parameter is ignored if the corresponding log is out of the scope of the specified time window. | [optional] 
**Time** | Pointer to [**LogsListRequestTime**](LogsListRequest_time.md) |  | 

## Methods

### NewLogsListRequest

`func NewLogsListRequest(query string, time LogsListRequestTime, ) *LogsListRequest`

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

`func (o *LogsListRequest) GetIndexOk() (string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndex

`func (o *LogsListRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndex

`func (o *LogsListRequest) SetIndex(v string)`

SetIndex gets a reference to the given string and assigns it to the Index field.

### GetLimit

`func (o *LogsListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LogsListRequest) GetLimitOk() (int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *LogsListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *LogsListRequest) SetLimit(v int32)`

SetLimit gets a reference to the given int32 and assigns it to the Limit field.

### GetQuery

`func (o *LogsListRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogsListRequest) GetQueryOk() (string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *LogsListRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *LogsListRequest) SetQuery(v string)`

SetQuery gets a reference to the given string and assigns it to the Query field.

### GetSort

`func (o *LogsListRequest) GetSort() LogsSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LogsListRequest) GetSortOk() (LogsSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSort

`func (o *LogsListRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSort

`func (o *LogsListRequest) SetSort(v LogsSort)`

SetSort gets a reference to the given LogsSort and assigns it to the Sort field.

### GetStartAt

`func (o *LogsListRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *LogsListRequest) GetStartAtOk() (string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartAt

`func (o *LogsListRequest) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAt

`func (o *LogsListRequest) SetStartAt(v string)`

SetStartAt gets a reference to the given string and assigns it to the StartAt field.

### GetTime

`func (o *LogsListRequest) GetTime() LogsListRequestTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *LogsListRequest) GetTimeOk() (LogsListRequestTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *LogsListRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *LogsListRequest) SetTime(v LogsListRequestTime)`

SetTime gets a reference to the given LogsListRequestTime and assigns it to the Time field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


