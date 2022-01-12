# CloudWorkloadSecurityAgentRuleAttributes

## Properties

| Name             | Type                                                                                                                 | Description                                                      | Notes      |
| ---------------- | -------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- | ---------- |
| **Category**     | Pointer to **string**                                                                                                | The category of the Agent rule.                                  | [optional] |
| **CreationDate** | Pointer to **int64**                                                                                                 | When the Agent rule was created, timestamp in milliseconds.      | [optional] |
| **Creator**      | Pointer to [**CloudWorkloadSecurityAgentRuleCreatorAttributes**](CloudWorkloadSecurityAgentRuleCreatorAttributes.md) |                                                                  | [optional] |
| **DefaultRule**  | Pointer to **bool**                                                                                                  | Whether the rule is included by default.                         | [optional] |
| **Description**  | Pointer to **string**                                                                                                | The description of the Agent rule.                               | [optional] |
| **Enabled**      | Pointer to **bool**                                                                                                  | Whether the Agent rule is enabled.                               | [optional] |
| **Expression**   | Pointer to **string**                                                                                                | The SECL expression of the Agent rule.                           | [optional] |
| **Name**         | Pointer to **string**                                                                                                | The name of the Agent rule.                                      | [optional] |
| **UpdatedAt**    | Pointer to **int64**                                                                                                 | When the Agent rule was last updated, timestamp in milliseconds. | [optional] |
| **Updater**      | Pointer to [**CloudWorkloadSecurityAgentRuleUpdaterAttributes**](CloudWorkloadSecurityAgentRuleUpdaterAttributes.md) |                                                                  | [optional] |
| **Version**      | Pointer to **int64**                                                                                                 | The version of the Agent rule.                                   | [optional] |

## Methods

### NewCloudWorkloadSecurityAgentRuleAttributes

`func NewCloudWorkloadSecurityAgentRuleAttributes() *CloudWorkloadSecurityAgentRuleAttributes`

NewCloudWorkloadSecurityAgentRuleAttributes instantiates a new CloudWorkloadSecurityAgentRuleAttributes object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudWorkloadSecurityAgentRuleAttributesWithDefaults

`func NewCloudWorkloadSecurityAgentRuleAttributesWithDefaults() *CloudWorkloadSecurityAgentRuleAttributes`

NewCloudWorkloadSecurityAgentRuleAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleAttributes object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCategory

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreationDate

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCreator

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreator() CloudWorkloadSecurityAgentRuleCreatorAttributes`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetCreatorOk() (*CloudWorkloadSecurityAgentRuleCreatorAttributes, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetCreator(v CloudWorkloadSecurityAgentRuleCreatorAttributes)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDefaultRule

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDefaultRule() bool`

GetDefaultRule returns the DefaultRule field if non-nil, zero value otherwise.

### GetDefaultRuleOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDefaultRuleOk() (*bool, bool)`

GetDefaultRuleOk returns a tuple with the DefaultRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRule

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetDefaultRule(v bool)`

SetDefaultRule sets DefaultRule field to given value.

### HasDefaultRule

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasDefaultRule() bool`

HasDefaultRule returns a boolean if a field has been set.

### GetDescription

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpression

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetName

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdater

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdater() CloudWorkloadSecurityAgentRuleUpdaterAttributes`

GetUpdater returns the Updater field if non-nil, zero value otherwise.

### GetUpdaterOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetUpdaterOk() (*CloudWorkloadSecurityAgentRuleUpdaterAttributes, bool)`

GetUpdaterOk returns a tuple with the Updater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdater

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetUpdater(v CloudWorkloadSecurityAgentRuleUpdaterAttributes)`

SetUpdater sets Updater field to given value.

### HasUpdater

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasUpdater() bool`

HasUpdater returns a boolean if a field has been set.

### GetVersion

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudWorkloadSecurityAgentRuleAttributes) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudWorkloadSecurityAgentRuleAttributes) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudWorkloadSecurityAgentRuleAttributes) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
