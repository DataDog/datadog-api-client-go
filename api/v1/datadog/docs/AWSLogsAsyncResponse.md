# AWSLogsAsyncResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Errors** | Pointer to [**[]AWSLogsAsyncError**](AWSLogsAsyncError.md) | List of errors. | [optional] 
**Status** | Pointer to **string** | Status of the properties. | [optional] 

## Methods

### NewAWSLogsAsyncResponse

`func NewAWSLogsAsyncResponse() *AWSLogsAsyncResponse`

NewAWSLogsAsyncResponse instantiates a new AWSLogsAsyncResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAWSLogsAsyncResponseWithDefaults

`func NewAWSLogsAsyncResponseWithDefaults() *AWSLogsAsyncResponse`

NewAWSLogsAsyncResponseWithDefaults instantiates a new AWSLogsAsyncResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetErrors

`func (o *AWSLogsAsyncResponse) GetErrors() []AWSLogsAsyncError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AWSLogsAsyncResponse) GetErrorsOk() (*[]AWSLogsAsyncError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AWSLogsAsyncResponse) SetErrors(v []AWSLogsAsyncError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AWSLogsAsyncResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *AWSLogsAsyncResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AWSLogsAsyncResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AWSLogsAsyncResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AWSLogsAsyncResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


