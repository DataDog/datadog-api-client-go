# CloudWorkloadSecurityAgentRuleCreateAttributes

## Properties

| Name            | Type                  | Description                            | Notes      |
| --------------- | --------------------- | -------------------------------------- | ---------- |
| **Description** | Pointer to **string** | The description of the Agent rule.     | [optional] |
| **Enabled**     | Pointer to **bool**   | Whether the Agent rule is enabled.     | [optional] |
| **Expression**  | **string**            | The SECL expression of the Agent rule. |
| **Name**        | **string**            | The name of the Agent rule.            |

## Methods

### NewCloudWorkloadSecurityAgentRuleCreateAttributes

`func NewCloudWorkloadSecurityAgentRuleCreateAttributes(expression string, name string) *CloudWorkloadSecurityAgentRuleCreateAttributes`

NewCloudWorkloadSecurityAgentRuleCreateAttributes instantiates a new CloudWorkloadSecurityAgentRuleCreateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudWorkloadSecurityAgentRuleCreateAttributesWithDefaults

`func NewCloudWorkloadSecurityAgentRuleCreateAttributesWithDefaults() *CloudWorkloadSecurityAgentRuleCreateAttributes`

NewCloudWorkloadSecurityAgentRuleCreateAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleCreateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDescription

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpression

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetExpression(v string)`

SetExpression sets Expression field to given value.

### GetName

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudWorkloadSecurityAgentRuleCreateAttributes) SetName(v string)`

SetName sets Name field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
