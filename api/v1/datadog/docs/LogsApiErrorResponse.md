# LogsApiErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**LogsApiError**](LogsAPIError.md) |  | [optional] 

## Methods

### NewLogsApiErrorResponse

`func NewLogsApiErrorResponse() *LogsApiErrorResponse`

NewLogsApiErrorResponse instantiates a new LogsApiErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsApiErrorResponseWithDefaults

`func NewLogsApiErrorResponseWithDefaults() *LogsApiErrorResponse`

NewLogsApiErrorResponseWithDefaults instantiates a new LogsApiErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *LogsApiErrorResponse) GetError() LogsApiError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LogsApiErrorResponse) GetErrorOk() (LogsApiError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *LogsApiErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *LogsApiErrorResponse) SetError(v LogsApiError)`

SetError gets a reference to the given LogsApiError and assigns it to the Error field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


