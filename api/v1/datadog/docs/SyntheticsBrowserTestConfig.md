# SyntheticsBrowserTestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | [**[]SyntheticsAssertion**](SyntheticsAssertion.md) | Array of assertions used for the test. | [default to []]
**Request** | [**SyntheticsTestRequest**](SyntheticsTestRequest.md) |  | 
**Variables** | Pointer to [**[]SyntheticsBrowserVariable**](SyntheticsBrowserVariable.md) | Array of variables used for the test steps. | [optional] 

## Methods

### NewSyntheticsBrowserTestConfig

`func NewSyntheticsBrowserTestConfig(assertions []SyntheticsAssertion, request SyntheticsTestRequest, ) *SyntheticsBrowserTestConfig`

NewSyntheticsBrowserTestConfig instantiates a new SyntheticsBrowserTestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsBrowserTestConfigWithDefaults

`func NewSyntheticsBrowserTestConfigWithDefaults() *SyntheticsBrowserTestConfig`

NewSyntheticsBrowserTestConfigWithDefaults instantiates a new SyntheticsBrowserTestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *SyntheticsBrowserTestConfig) GetAssertions() []SyntheticsAssertion`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *SyntheticsBrowserTestConfig) GetAssertionsOk() (*[]SyntheticsAssertion, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *SyntheticsBrowserTestConfig) SetAssertions(v []SyntheticsAssertion)`

SetAssertions sets Assertions field to given value.


### GetRequest

`func (o *SyntheticsBrowserTestConfig) GetRequest() SyntheticsTestRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SyntheticsBrowserTestConfig) GetRequestOk() (*SyntheticsTestRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *SyntheticsBrowserTestConfig) SetRequest(v SyntheticsTestRequest)`

SetRequest sets Request field to given value.


### GetVariables

`func (o *SyntheticsBrowserTestConfig) GetVariables() []SyntheticsBrowserVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SyntheticsBrowserTestConfig) GetVariablesOk() (*[]SyntheticsBrowserVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SyntheticsBrowserTestConfig) SetVariables(v []SyntheticsBrowserVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SyntheticsBrowserTestConfig) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


