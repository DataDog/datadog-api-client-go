# SyntheticsAPITestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | [**[]SyntheticsAssertion**](SyntheticsAssertion.md) | Array of assertions used for the test. | [default to []]
**ConfigVariables** | Pointer to [**[]SyntheticsConfigVariable**](SyntheticsConfigVariable.md) | Array of variables used for the test. | [optional] 
**Request** | [**SyntheticsTestRequest**](SyntheticsTestRequest.md) |  | 

## Methods

### NewSyntheticsAPITestConfig

`func NewSyntheticsAPITestConfig(assertions []SyntheticsAssertion, request SyntheticsTestRequest, ) *SyntheticsAPITestConfig`

NewSyntheticsAPITestConfig instantiates a new SyntheticsAPITestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsAPITestConfigWithDefaults

`func NewSyntheticsAPITestConfigWithDefaults() *SyntheticsAPITestConfig`

NewSyntheticsAPITestConfigWithDefaults instantiates a new SyntheticsAPITestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *SyntheticsAPITestConfig) GetAssertions() []SyntheticsAssertion`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *SyntheticsAPITestConfig) GetAssertionsOk() (*[]SyntheticsAssertion, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *SyntheticsAPITestConfig) SetAssertions(v []SyntheticsAssertion)`

SetAssertions sets Assertions field to given value.


### GetConfigVariables

`func (o *SyntheticsAPITestConfig) GetConfigVariables() []SyntheticsConfigVariable`

GetConfigVariables returns the ConfigVariables field if non-nil, zero value otherwise.

### GetConfigVariablesOk

`func (o *SyntheticsAPITestConfig) GetConfigVariablesOk() (*[]SyntheticsConfigVariable, bool)`

GetConfigVariablesOk returns a tuple with the ConfigVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigVariables

`func (o *SyntheticsAPITestConfig) SetConfigVariables(v []SyntheticsConfigVariable)`

SetConfigVariables sets ConfigVariables field to given value.

### HasConfigVariables

`func (o *SyntheticsAPITestConfig) HasConfigVariables() bool`

HasConfigVariables returns a boolean if a field has been set.

### GetRequest

`func (o *SyntheticsAPITestConfig) GetRequest() SyntheticsTestRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SyntheticsAPITestConfig) GetRequestOk() (*SyntheticsTestRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *SyntheticsAPITestConfig) SetRequest(v SyntheticsTestRequest)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


