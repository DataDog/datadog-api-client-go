# AWSLogsAsyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]AWSLogsAsyncResponseErrors**](AWSLogsAsyncResponse_errors.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAWSLogsAsyncResponse

`func NewAWSLogsAsyncResponse() *AWSLogsAsyncResponse`

NewAWSLogsAsyncResponse instantiates a new AWSLogsAsyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSLogsAsyncResponseWithDefaults

`func NewAWSLogsAsyncResponseWithDefaults() *AWSLogsAsyncResponse`

NewAWSLogsAsyncResponseWithDefaults instantiates a new AWSLogsAsyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *AWSLogsAsyncResponse) GetErrors() []AWSLogsAsyncResponseErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AWSLogsAsyncResponse) GetErrorsOk() ([]AWSLogsAsyncResponseErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *AWSLogsAsyncResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *AWSLogsAsyncResponse) SetErrors(v []AWSLogsAsyncResponseErrors)`

SetErrors gets a reference to the given []AWSLogsAsyncResponseErrors and assigns it to the Errors field.

### GetStatus

`func (o *AWSLogsAsyncResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AWSLogsAsyncResponse) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *AWSLogsAsyncResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *AWSLogsAsyncResponse) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


