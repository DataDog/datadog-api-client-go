# SyntheticsTestConfig

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Assertions** | Pointer to [**[]SyntheticsAssertion**](SyntheticsAssertion.md) | Array of assertions used for the test. | [optional] [default to []]
**ConfigVariables** | Pointer to [**[]SyntheticsConfigVariable**](SyntheticsConfigVariable.md) | API tests only - array of variables used for the test. | [optional] 
**Request** | Pointer to [**SyntheticsTestRequest**](SyntheticsTestRequest.md) |  | [optional] 
**Variables** | Pointer to [**[]SyntheticsBrowserVariable**](SyntheticsBrowserVariable.md) | Browser tests only - array of variables used for the test steps. | [optional] 

## Methods

### NewSyntheticsTestConfig

`func NewSyntheticsTestConfig() *SyntheticsTestConfig`

NewSyntheticsTestConfig instantiates a new SyntheticsTestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsTestConfigWithDefaults

`func NewSyntheticsTestConfigWithDefaults() *SyntheticsTestConfig`

NewSyntheticsTestConfigWithDefaults instantiates a new SyntheticsTestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *SyntheticsTestConfig) GetAssertions() []SyntheticsAssertion`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *SyntheticsTestConfig) GetAssertionsOk() (*[]SyntheticsAssertion, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *SyntheticsTestConfig) SetAssertions(v []SyntheticsAssertion)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *SyntheticsTestConfig) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.

### GetConfigVariables

`func (o *SyntheticsTestConfig) GetConfigVariables() []SyntheticsConfigVariable`

GetConfigVariables returns the ConfigVariables field if non-nil, zero value otherwise.

### GetConfigVariablesOk

`func (o *SyntheticsTestConfig) GetConfigVariablesOk() (*[]SyntheticsConfigVariable, bool)`

GetConfigVariablesOk returns a tuple with the ConfigVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigVariables

`func (o *SyntheticsTestConfig) SetConfigVariables(v []SyntheticsConfigVariable)`

SetConfigVariables sets ConfigVariables field to given value.

### HasConfigVariables

`func (o *SyntheticsTestConfig) HasConfigVariables() bool`

HasConfigVariables returns a boolean if a field has been set.

### GetRequest

`func (o *SyntheticsTestConfig) GetRequest() SyntheticsTestRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SyntheticsTestConfig) GetRequestOk() (*SyntheticsTestRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *SyntheticsTestConfig) SetRequest(v SyntheticsTestRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *SyntheticsTestConfig) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetVariables

`func (o *SyntheticsTestConfig) GetVariables() []SyntheticsBrowserVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SyntheticsTestConfig) GetVariablesOk() (*[]SyntheticsBrowserVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SyntheticsTestConfig) SetVariables(v []SyntheticsBrowserVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SyntheticsTestConfig) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


