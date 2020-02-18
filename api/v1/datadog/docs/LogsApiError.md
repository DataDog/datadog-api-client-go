# LogsApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code identifying the error | [optional] 
**Details** | Pointer to [**[]LogsApiError**](LogsAPIError.md) | Additional error details | [optional] 
**Message** | Pointer to **string** | Error message | [optional] 

## Methods

### NewLogsApiError

`func NewLogsApiError() *LogsApiError`

NewLogsApiError instantiates a new LogsApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsApiErrorWithDefaults

`func NewLogsApiErrorWithDefaults() *LogsApiError`

NewLogsApiErrorWithDefaults instantiates a new LogsApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LogsApiError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LogsApiError) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *LogsApiError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *LogsApiError) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetDetails

`func (o *LogsApiError) GetDetails() []LogsApiError`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *LogsApiError) GetDetailsOk() ([]LogsApiError, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *LogsApiError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *LogsApiError) SetDetails(v []LogsApiError)`

SetDetails gets a reference to the given []LogsApiError and assigns it to the Details field.

### GetMessage

`func (o *LogsApiError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogsApiError) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *LogsApiError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *LogsApiError) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


