# AwsLogsAsyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]AwsLogsAsyncResponseErrors**](AWSLogsAsyncResponse_errors.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### GetErrors

`func (o *AwsLogsAsyncResponse) GetErrors() []AwsLogsAsyncResponseErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AwsLogsAsyncResponse) GetErrorsOk() ([]AwsLogsAsyncResponseErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *AwsLogsAsyncResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *AwsLogsAsyncResponse) SetErrors(v []AwsLogsAsyncResponseErrors)`

SetErrors gets a reference to the given []AwsLogsAsyncResponseErrors and assigns it to the Errors field.

### GetStatus

`func (o *AwsLogsAsyncResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsLogsAsyncResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *AwsLogsAsyncResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *AwsLogsAsyncResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


