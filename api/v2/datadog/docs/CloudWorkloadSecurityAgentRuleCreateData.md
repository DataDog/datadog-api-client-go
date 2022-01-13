# CloudWorkloadSecurityAgentRuleCreateData

## Properties

| Name           | Type                                                                                                    | Description | Notes                                                      |
| -------------- | ------------------------------------------------------------------------------------------------------- | ----------- | ---------------------------------------------------------- |
| **Attributes** | [**CloudWorkloadSecurityAgentRuleCreateAttributes**](CloudWorkloadSecurityAgentRuleCreateAttributes.md) |             |
| **Type**       | [**CloudWorkloadSecurityAgentRuleType**](CloudWorkloadSecurityAgentRuleType.md)                         |             | [default to CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE] |

## Methods

### NewCloudWorkloadSecurityAgentRuleCreateData

`func NewCloudWorkloadSecurityAgentRuleCreateData(attributes CloudWorkloadSecurityAgentRuleCreateAttributes, type_ CloudWorkloadSecurityAgentRuleType) *CloudWorkloadSecurityAgentRuleCreateData`

NewCloudWorkloadSecurityAgentRuleCreateData instantiates a new CloudWorkloadSecurityAgentRuleCreateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudWorkloadSecurityAgentRuleCreateDataWithDefaults

`func NewCloudWorkloadSecurityAgentRuleCreateDataWithDefaults() *CloudWorkloadSecurityAgentRuleCreateData`

NewCloudWorkloadSecurityAgentRuleCreateDataWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleCreateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *CloudWorkloadSecurityAgentRuleCreateData) GetAttributes() CloudWorkloadSecurityAgentRuleCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CloudWorkloadSecurityAgentRuleCreateData) GetAttributesOk() (*CloudWorkloadSecurityAgentRuleCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CloudWorkloadSecurityAgentRuleCreateData) SetAttributes(v CloudWorkloadSecurityAgentRuleCreateAttributes)`

SetAttributes sets Attributes field to given value.

### GetType

`func (o *CloudWorkloadSecurityAgentRuleCreateData) GetType() CloudWorkloadSecurityAgentRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudWorkloadSecurityAgentRuleCreateData) GetTypeOk() (*CloudWorkloadSecurityAgentRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudWorkloadSecurityAgentRuleCreateData) SetType(v CloudWorkloadSecurityAgentRuleType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
