# HTTPLogError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | TODO. | 
**Message** | Pointer to **string** | TODO. | 

## Methods

### NewHTTPLogError

`func NewHTTPLogError(code int32, message string, ) *HTTPLogError`

NewHTTPLogError instantiates a new HTTPLogError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPLogErrorWithDefaults

`func NewHTTPLogErrorWithDefaults() *HTTPLogError`

NewHTTPLogErrorWithDefaults instantiates a new HTTPLogError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *HTTPLogError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HTTPLogError) GetCodeOk() (int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *HTTPLogError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *HTTPLogError) SetCode(v int32)`

SetCode gets a reference to the given int32 and assigns it to the Code field.

### GetMessage

`func (o *HTTPLogError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HTTPLogError) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *HTTPLogError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *HTTPLogError) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


