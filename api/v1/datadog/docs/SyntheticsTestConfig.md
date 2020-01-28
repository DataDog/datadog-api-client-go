# SyntheticsTestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | Pointer to [**[]SyntheticsAssertion**](SyntheticsAssertion.md) |  | [default to []]
**Request** | Pointer to [**SyntheticsTestRequest**](SyntheticsTestRequest.md) |  | 
**Variables** | Pointer to [**[]SyntheticsBrowserVariable**](SyntheticsBrowserVariable.md) |  | [optional] 

## Methods

### GetAssertions

`func (o *SyntheticsTestConfig) GetAssertions() []SyntheticsAssertion`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *SyntheticsTestConfig) GetAssertionsOk() ([]SyntheticsAssertion, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssertions

`func (o *SyntheticsTestConfig) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.

### SetAssertions

`func (o *SyntheticsTestConfig) SetAssertions(v []SyntheticsAssertion)`

SetAssertions gets a reference to the given []SyntheticsAssertion and assigns it to the Assertions field.

### GetRequest

`func (o *SyntheticsTestConfig) GetRequest() SyntheticsTestRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SyntheticsTestConfig) GetRequestOk() (SyntheticsTestRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequest

`func (o *SyntheticsTestConfig) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequest

`func (o *SyntheticsTestConfig) SetRequest(v SyntheticsTestRequest)`

SetRequest gets a reference to the given SyntheticsTestRequest and assigns it to the Request field.

### GetVariables

`func (o *SyntheticsTestConfig) GetVariables() []SyntheticsBrowserVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SyntheticsTestConfig) GetVariablesOk() ([]SyntheticsBrowserVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVariables

`func (o *SyntheticsTestConfig) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariables

`func (o *SyntheticsTestConfig) SetVariables(v []SyntheticsBrowserVariable)`

SetVariables gets a reference to the given []SyntheticsBrowserVariable and assigns it to the Variables field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


