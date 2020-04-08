# SyntheticsAssertion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to [**SyntheticsAssertionOperator**](SyntheticsAssertionOperator.md) |  | 
**Property** | Pointer to **string** | TODO. | [optional] 
**Target** | Pointer to [**interface{}**](.md) | TODO. | [optional] 
**Type** | Pointer to [**SyntheticsAssertionType**](SyntheticsAssertionType.md) |  | 

## Methods

### NewSyntheticsAssertion

`func NewSyntheticsAssertion(operator SyntheticsAssertionOperator, type_ SyntheticsAssertionType, ) *SyntheticsAssertion`

NewSyntheticsAssertion instantiates a new SyntheticsAssertion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAssertionWithDefaults

`func NewSyntheticsAssertionWithDefaults() *SyntheticsAssertion`

NewSyntheticsAssertionWithDefaults instantiates a new SyntheticsAssertion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SyntheticsAssertion) GetOperator() SyntheticsAssertionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SyntheticsAssertion) GetOperatorOk() (*SyntheticsAssertionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SyntheticsAssertion) SetOperator(v SyntheticsAssertionOperator)`

SetOperator sets Operator field to given value.


### GetProperty

`func (o *SyntheticsAssertion) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SyntheticsAssertion) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SyntheticsAssertion) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SyntheticsAssertion) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetTarget

`func (o *SyntheticsAssertion) GetTarget() interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SyntheticsAssertion) GetTargetOk() (*interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SyntheticsAssertion) SetTarget(v interface{})`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SyntheticsAssertion) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsAssertion) GetType() SyntheticsAssertionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsAssertion) GetTypeOk() (*SyntheticsAssertionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsAssertion) SetType(v SyntheticsAssertionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


