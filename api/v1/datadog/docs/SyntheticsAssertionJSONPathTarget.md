# SyntheticsAssertionJSONPathTarget

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Operator** | [**SyntheticsAssertionJSONPathOperator**](SyntheticsAssertionJSONPathOperator.md) |  | 
**Property** | Pointer to **string** | The associated assertion property. | [optional] 
**Target** | Pointer to [**SyntheticsAssertionJSONPathTargetTarget**](SyntheticsAssertionJSONPathTargetTarget.md) |  | [optional] 
**Type** | [**SyntheticsAssertionType**](SyntheticsAssertionType.md) |  | 

## Methods

### NewSyntheticsAssertionJSONPathTarget

`func NewSyntheticsAssertionJSONPathTarget(operator SyntheticsAssertionJSONPathOperator, type_ SyntheticsAssertionType, ) *SyntheticsAssertionJSONPathTarget`

NewSyntheticsAssertionJSONPathTarget instantiates a new SyntheticsAssertionJSONPathTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAssertionJSONPathTargetWithDefaults

`func NewSyntheticsAssertionJSONPathTargetWithDefaults() *SyntheticsAssertionJSONPathTarget`

NewSyntheticsAssertionJSONPathTargetWithDefaults instantiates a new SyntheticsAssertionJSONPathTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SyntheticsAssertionJSONPathTarget) GetOperator() SyntheticsAssertionJSONPathOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SyntheticsAssertionJSONPathTarget) GetOperatorOk() (*SyntheticsAssertionJSONPathOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SyntheticsAssertionJSONPathTarget) SetOperator(v SyntheticsAssertionJSONPathOperator)`

SetOperator sets Operator field to given value.


### GetProperty

`func (o *SyntheticsAssertionJSONPathTarget) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SyntheticsAssertionJSONPathTarget) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SyntheticsAssertionJSONPathTarget) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SyntheticsAssertionJSONPathTarget) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetTarget

`func (o *SyntheticsAssertionJSONPathTarget) GetTarget() SyntheticsAssertionJSONPathTargetTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SyntheticsAssertionJSONPathTarget) GetTargetOk() (*SyntheticsAssertionJSONPathTargetTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SyntheticsAssertionJSONPathTarget) SetTarget(v SyntheticsAssertionJSONPathTargetTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SyntheticsAssertionJSONPathTarget) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsAssertionJSONPathTarget) GetType() SyntheticsAssertionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsAssertionJSONPathTarget) GetTypeOk() (*SyntheticsAssertionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsAssertionJSONPathTarget) SetType(v SyntheticsAssertionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


