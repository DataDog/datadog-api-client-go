# HTTPLogErrors

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Errors** | Pointer to [**[]HTTPLogError**](HTTPLogError.md) | Structured errors. | [optional] 

## Methods

### NewHTTPLogErrors

`func NewHTTPLogErrors() *HTTPLogErrors`

NewHTTPLogErrors instantiates a new HTTPLogErrors object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHTTPLogErrorsWithDefaults

`func NewHTTPLogErrorsWithDefaults() *HTTPLogErrors`

NewHTTPLogErrorsWithDefaults instantiates a new HTTPLogErrors object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetErrors

`func (o *HTTPLogErrors) GetErrors() []HTTPLogError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *HTTPLogErrors) GetErrorsOk() (*[]HTTPLogError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *HTTPLogErrors) SetErrors(v []HTTPLogError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *HTTPLogErrors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


