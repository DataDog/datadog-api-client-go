# LogsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]Log**](Log.md) | Array of logs matching the request and the &#x60;nextLogId&#x60; if sent. | [optional] 
**NextLogId** | Pointer to **string** | Hash identifier of the next log to return in the list. This parameter is used for the pagination feature. | [optional] 
**Status** | Pointer to **string** | Status of the response. | [optional] 

## Methods

### NewLogsListResponse

`func NewLogsListResponse() *LogsListResponse`

NewLogsListResponse instantiates a new LogsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsListResponseWithDefaults

`func NewLogsListResponseWithDefaults() *LogsListResponse`

NewLogsListResponseWithDefaults instantiates a new LogsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *LogsListResponse) GetLogs() []Log`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *LogsListResponse) GetLogsOk() (*[]Log, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *LogsListResponse) SetLogs(v []Log)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *LogsListResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetNextLogId

`func (o *LogsListResponse) GetNextLogId() string`

GetNextLogId returns the NextLogId field if non-nil, zero value otherwise.

### GetNextLogIdOk

`func (o *LogsListResponse) GetNextLogIdOk() (*string, bool)`

GetNextLogIdOk returns a tuple with the NextLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLogId

`func (o *LogsListResponse) SetNextLogId(v string)`

SetNextLogId sets NextLogId field to given value.

### HasNextLogId

`func (o *LogsListResponse) HasNextLogId() bool`

HasNextLogId returns a boolean if a field has been set.

### GetStatus

`func (o *LogsListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogsListResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogsListResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogsListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


