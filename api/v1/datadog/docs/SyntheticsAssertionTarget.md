# SyntheticsAssertionTarget

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Operator** | [**SyntheticsAssertionOperator**](SyntheticsAssertionOperator.md) |  | 
**Property** | Pointer to **string** | The associated assertion property. | [optional] 
**Target** | Pointer to **interface{}** | Value used by the operator. | [optional] 
**Type** | [**SyntheticsAssertionType**](SyntheticsAssertionType.md) |  | 

## Methods

### NewSyntheticsAssertionTarget

`func NewSyntheticsAssertionTarget(operator SyntheticsAssertionOperator, type_ SyntheticsAssertionType, ) *SyntheticsAssertionTarget`

NewSyntheticsAssertionTarget instantiates a new SyntheticsAssertionTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAssertionTargetWithDefaults

`func NewSyntheticsAssertionTargetWithDefaults() *SyntheticsAssertionTarget`

NewSyntheticsAssertionTargetWithDefaults instantiates a new SyntheticsAssertionTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SyntheticsAssertionTarget) GetOperator() SyntheticsAssertionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SyntheticsAssertionTarget) GetOperatorOk() (*SyntheticsAssertionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SyntheticsAssertionTarget) SetOperator(v SyntheticsAssertionOperator)`

SetOperator sets Operator field to given value.


### GetProperty

`func (o *SyntheticsAssertionTarget) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SyntheticsAssertionTarget) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SyntheticsAssertionTarget) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SyntheticsAssertionTarget) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetTarget

`func (o *SyntheticsAssertionTarget) GetTarget() interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SyntheticsAssertionTarget) GetTargetOk() (*interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SyntheticsAssertionTarget) SetTarget(v interface{})`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SyntheticsAssertionTarget) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsAssertionTarget) GetType() SyntheticsAssertionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsAssertionTarget) GetTypeOk() (*SyntheticsAssertionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsAssertionTarget) SetType(v SyntheticsAssertionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


