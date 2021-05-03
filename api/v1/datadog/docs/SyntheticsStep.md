# SyntheticsStep

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AllowFailure** | Pointer to **bool** | A boolean set to allow this step to fail. | [optional] 
**Name** | Pointer to **string** | The name of the step. | [optional] 
**Params** | Pointer to **interface{}** | The parameters of the step. | [optional] 
**Timeout** | Pointer to **int64** | The time before declaring a step failed. | [optional] 
**Type** | Pointer to [**SyntheticsStepType**](SyntheticsStepType.md) |  | [optional] 

## Methods

### NewSyntheticsStep

`func NewSyntheticsStep() *SyntheticsStep`

NewSyntheticsStep instantiates a new SyntheticsStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsStepWithDefaults

`func NewSyntheticsStepWithDefaults() *SyntheticsStep`

NewSyntheticsStepWithDefaults instantiates a new SyntheticsStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowFailure

`func (o *SyntheticsStep) GetAllowFailure() bool`

GetAllowFailure returns the AllowFailure field if non-nil, zero value otherwise.

### GetAllowFailureOk

`func (o *SyntheticsStep) GetAllowFailureOk() (*bool, bool)`

GetAllowFailureOk returns a tuple with the AllowFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFailure

`func (o *SyntheticsStep) SetAllowFailure(v bool)`

SetAllowFailure sets AllowFailure field to given value.

### HasAllowFailure

`func (o *SyntheticsStep) HasAllowFailure() bool`

HasAllowFailure returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyntheticsStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParams

`func (o *SyntheticsStep) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SyntheticsStep) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SyntheticsStep) SetParams(v interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *SyntheticsStep) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetTimeout

`func (o *SyntheticsStep) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SyntheticsStep) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SyntheticsStep) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SyntheticsStep) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsStep) GetType() SyntheticsStepType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsStep) GetTypeOk() (*SyntheticsStepType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsStep) SetType(v SyntheticsStepType)`

SetType sets Type field to given value.

### HasType

`func (o *SyntheticsStep) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


