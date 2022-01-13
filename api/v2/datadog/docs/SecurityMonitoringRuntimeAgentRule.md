# SecurityMonitoringRuntimeAgentRule

## Properties

| Name            | Type                  | Description                                                                                                                                                                                                                                                                                  | Notes      |
| --------------- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **AgentRuleId** | Pointer to **string** | The Agent rule ID. Must be unique within the rule.                                                                                                                                                                                                                                           | [optional] |
| **Expression**  | Pointer to **string** | A Runtime Security expression determines what activity should be collected by the Datadog Agent. These logical expressions can use predefined operators and attributes. Tags cannot be used in Runtime Security expressions. Instead, allow or deny based on tags under the advanced option. | [optional] |

## Methods

### NewSecurityMonitoringRuntimeAgentRule

`func NewSecurityMonitoringRuntimeAgentRule() *SecurityMonitoringRuntimeAgentRule`

NewSecurityMonitoringRuntimeAgentRule instantiates a new SecurityMonitoringRuntimeAgentRule object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityMonitoringRuntimeAgentRuleWithDefaults

`func NewSecurityMonitoringRuntimeAgentRuleWithDefaults() *SecurityMonitoringRuntimeAgentRule`

NewSecurityMonitoringRuntimeAgentRuleWithDefaults instantiates a new SecurityMonitoringRuntimeAgentRule object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAgentRuleId

`func (o *SecurityMonitoringRuntimeAgentRule) GetAgentRuleId() string`

GetAgentRuleId returns the AgentRuleId field if non-nil, zero value otherwise.

### GetAgentRuleIdOk

`func (o *SecurityMonitoringRuntimeAgentRule) GetAgentRuleIdOk() (*string, bool)`

GetAgentRuleIdOk returns a tuple with the AgentRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRuleId

`func (o *SecurityMonitoringRuntimeAgentRule) SetAgentRuleId(v string)`

SetAgentRuleId sets AgentRuleId field to given value.

### HasAgentRuleId

`func (o *SecurityMonitoringRuntimeAgentRule) HasAgentRuleId() bool`

HasAgentRuleId returns a boolean if a field has been set.

### GetExpression

`func (o *SecurityMonitoringRuntimeAgentRule) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SecurityMonitoringRuntimeAgentRule) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SecurityMonitoringRuntimeAgentRule) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *SecurityMonitoringRuntimeAgentRule) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
