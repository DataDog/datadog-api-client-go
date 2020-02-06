# HttpLogError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | 
**Message** | Pointer to **string** |  | 

## Methods

### NewHttpLogError

`func NewHttpLogError(code int32, message string, ) *HttpLogError`

NewHttpLogError instantiates a new HttpLogError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpLogErrorWithDefaults

`func NewHttpLogErrorWithDefaults() *HttpLogError`

NewHttpLogErrorWithDefaults instantiates a new HttpLogError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *HttpLogError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HttpLogError) GetCodeOk() (int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *HttpLogError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *HttpLogError) SetCode(v int32)`

SetCode gets a reference to the given int32 and assigns it to the Code field.

### GetMessage

`func (o *HttpLogError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HttpLogError) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *HttpLogError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *HttpLogError) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


