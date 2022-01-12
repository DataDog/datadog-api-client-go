# CloudWorkloadSecurityAgentRuleData

## Properties

| Name           | Type                                                                                                   | Description               | Notes                                                                 |
| -------------- | ------------------------------------------------------------------------------------------------------ | ------------------------- | --------------------------------------------------------------------- |
| **Attributes** | Pointer to [**CloudWorkloadSecurityAgentRuleAttributes**](CloudWorkloadSecurityAgentRuleAttributes.md) |                           | [optional]                                                            |
| **Id**         | Pointer to **string**                                                                                  | The ID of the Agent rule. | [optional]                                                            |
| **Type**       | Pointer to [**CloudWorkloadSecurityAgentRuleType**](CloudWorkloadSecurityAgentRuleType.md)             |                           | [optional] [default to CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE] |

## Methods

### NewCloudWorkloadSecurityAgentRuleData

`func NewCloudWorkloadSecurityAgentRuleData() *CloudWorkloadSecurityAgentRuleData`

NewCloudWorkloadSecurityAgentRuleData instantiates a new CloudWorkloadSecurityAgentRuleData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudWorkloadSecurityAgentRuleDataWithDefaults

`func NewCloudWorkloadSecurityAgentRuleDataWithDefaults() *CloudWorkloadSecurityAgentRuleData`

NewCloudWorkloadSecurityAgentRuleDataWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *CloudWorkloadSecurityAgentRuleData) GetAttributes() CloudWorkloadSecurityAgentRuleAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CloudWorkloadSecurityAgentRuleData) GetAttributesOk() (*CloudWorkloadSecurityAgentRuleAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CloudWorkloadSecurityAgentRuleData) SetAttributes(v CloudWorkloadSecurityAgentRuleAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CloudWorkloadSecurityAgentRuleData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *CloudWorkloadSecurityAgentRuleData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudWorkloadSecurityAgentRuleData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudWorkloadSecurityAgentRuleData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudWorkloadSecurityAgentRuleData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CloudWorkloadSecurityAgentRuleData) GetType() CloudWorkloadSecurityAgentRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudWorkloadSecurityAgentRuleData) GetTypeOk() (*CloudWorkloadSecurityAgentRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudWorkloadSecurityAgentRuleData) SetType(v CloudWorkloadSecurityAgentRuleType)`

SetType sets Type field to given value.

### HasType

`func (o *CloudWorkloadSecurityAgentRuleData) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
