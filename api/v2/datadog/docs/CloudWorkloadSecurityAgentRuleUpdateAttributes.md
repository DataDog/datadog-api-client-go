# CloudWorkloadSecurityAgentRuleUpdateAttributes

## Properties

| Name            | Type                  | Description                            | Notes      |
| --------------- | --------------------- | -------------------------------------- | ---------- |
| **Description** | Pointer to **string** | The description of the Agent rule.     | [optional] |
| **Enabled**     | Pointer to **bool**   | Whether the Agent rule is enabled.     | [optional] |
| **Expression**  | Pointer to **string** | The SECL expression of the Agent rule. | [optional] |

## Methods

### NewCloudWorkloadSecurityAgentRuleUpdateAttributes

`func NewCloudWorkloadSecurityAgentRuleUpdateAttributes() *CloudWorkloadSecurityAgentRuleUpdateAttributes`

NewCloudWorkloadSecurityAgentRuleUpdateAttributes instantiates a new CloudWorkloadSecurityAgentRuleUpdateAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudWorkloadSecurityAgentRuleUpdateAttributesWithDefaults

`func NewCloudWorkloadSecurityAgentRuleUpdateAttributesWithDefaults() *CloudWorkloadSecurityAgentRuleUpdateAttributes`

NewCloudWorkloadSecurityAgentRuleUpdateAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleUpdateAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDescription

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpression

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *CloudWorkloadSecurityAgentRuleUpdateAttributes) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
