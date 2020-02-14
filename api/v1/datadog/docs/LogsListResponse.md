# LogsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]Log**](Log.md) |  | [optional] 
**NextLogId** | Pointer to **string** | Hash identifier of the next log to return in the list. This parameter is used for the pagination feature. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

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

`func (o *LogsListResponse) GetLogsOk() ([]Log, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogs

`func (o *LogsListResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogs

`func (o *LogsListResponse) SetLogs(v []Log)`

SetLogs gets a reference to the given []Log and assigns it to the Logs field.

### GetNextLogId

`func (o *LogsListResponse) GetNextLogId() string`

GetNextLogId returns the NextLogId field if non-nil, zero value otherwise.

### GetNextLogIdOk

`func (o *LogsListResponse) GetNextLogIdOk() (string, bool)`

GetNextLogIdOk returns a tuple with the NextLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextLogId

`func (o *LogsListResponse) HasNextLogId() bool`

HasNextLogId returns a boolean if a field has been set.

### SetNextLogId

`func (o *LogsListResponse) SetNextLogId(v string)`

SetNextLogId gets a reference to the given string and assigns it to the NextLogId field.

### GetStatus

`func (o *LogsListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogsListResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *LogsListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *LogsListResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


