# LogsAPIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code identifying the error | [optional] 
**Details** | Pointer to [**[]LogsAPIError**](LogsAPIError.md) | Additional error details | [optional] 
**Message** | Pointer to **string** | Error message | [optional] 

## Methods

### NewLogsAPIError

`func NewLogsAPIError() *LogsAPIError`

NewLogsAPIError instantiates a new LogsAPIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsAPIErrorWithDefaults

`func NewLogsAPIErrorWithDefaults() *LogsAPIError`

NewLogsAPIErrorWithDefaults instantiates a new LogsAPIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LogsAPIError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LogsAPIError) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *LogsAPIError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *LogsAPIError) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetDetails

`func (o *LogsAPIError) GetDetails() []LogsAPIError`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *LogsAPIError) GetDetailsOk() ([]LogsAPIError, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *LogsAPIError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *LogsAPIError) SetDetails(v []LogsAPIError)`

SetDetails gets a reference to the given []LogsAPIError and assigns it to the Details field.

### GetMessage

`func (o *LogsAPIError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogsAPIError) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *LogsAPIError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *LogsAPIError) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


