# UsageLambdaHour

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FuncCount** | Pointer to **int64** | Contains the number of different functions for each region and AWS account. | [optional] 
**Hour** | Pointer to **time.Time** | The hour for the usage. | [optional] 
**InvocationsSum** | Pointer to **int64** | Contains the sum of invocations of all functions. | [optional] 

## Methods

### NewUsageLambdaHour

`func NewUsageLambdaHour() *UsageLambdaHour`

NewUsageLambdaHour instantiates a new UsageLambdaHour object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLambdaHourWithDefaults

`func NewUsageLambdaHourWithDefaults() *UsageLambdaHour`

NewUsageLambdaHourWithDefaults instantiates a new UsageLambdaHour object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuncCount

`func (o *UsageLambdaHour) GetFuncCount() int64`

GetFuncCount returns the FuncCount field if non-nil, zero value otherwise.

### GetFuncCountOk

`func (o *UsageLambdaHour) GetFuncCountOk() (*int64, bool)`

GetFuncCountOk returns a tuple with the FuncCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncCount

`func (o *UsageLambdaHour) SetFuncCount(v int64)`

SetFuncCount sets FuncCount field to given value.

### HasFuncCount

`func (o *UsageLambdaHour) HasFuncCount() bool`

HasFuncCount returns a boolean if a field has been set.

### GetHour

`func (o *UsageLambdaHour) GetHour() time.Time`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *UsageLambdaHour) GetHourOk() (*time.Time, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *UsageLambdaHour) SetHour(v time.Time)`

SetHour sets Hour field to given value.

### HasHour

`func (o *UsageLambdaHour) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetInvocationsSum

`func (o *UsageLambdaHour) GetInvocationsSum() int64`

GetInvocationsSum returns the InvocationsSum field if non-nil, zero value otherwise.

### GetInvocationsSumOk

`func (o *UsageLambdaHour) GetInvocationsSumOk() (*int64, bool)`

GetInvocationsSumOk returns a tuple with the InvocationsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationsSum

`func (o *UsageLambdaHour) SetInvocationsSum(v int64)`

SetInvocationsSum sets InvocationsSum field to given value.

### HasInvocationsSum

`func (o *UsageLambdaHour) HasInvocationsSum() bool`

HasInvocationsSum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


