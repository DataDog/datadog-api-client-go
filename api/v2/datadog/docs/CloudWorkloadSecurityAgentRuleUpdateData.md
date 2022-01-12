# CloudWorkloadSecurityAgentRuleUpdateData

## Properties

| Name           | Type                                                                                                    | Description | Notes                                                      |
| -------------- | ------------------------------------------------------------------------------------------------------- | ----------- | ---------------------------------------------------------- |
| **Attributes** | [**CloudWorkloadSecurityAgentRuleUpdateAttributes**](CloudWorkloadSecurityAgentRuleUpdateAttributes.md) |             |
| **Type**       | [**CloudWorkloadSecurityAgentRuleType**](CloudWorkloadSecurityAgentRuleType.md)                         |             | [default to CLOUDWORKLOADSECURITYAGENTRULETYPE_AGENT_RULE] |

## Methods

### NewCloudWorkloadSecurityAgentRuleUpdateData

`func NewCloudWorkloadSecurityAgentRuleUpdateData(attributes CloudWorkloadSecurityAgentRuleUpdateAttributes, type_ CloudWorkloadSecurityAgentRuleType) *CloudWorkloadSecurityAgentRuleUpdateData`

NewCloudWorkloadSecurityAgentRuleUpdateData instantiates a new CloudWorkloadSecurityAgentRuleUpdateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudWorkloadSecurityAgentRuleUpdateDataWithDefaults

`func NewCloudWorkloadSecurityAgentRuleUpdateDataWithDefaults() *CloudWorkloadSecurityAgentRuleUpdateData`

NewCloudWorkloadSecurityAgentRuleUpdateDataWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleUpdateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *CloudWorkloadSecurityAgentRuleUpdateData) GetAttributes() CloudWorkloadSecurityAgentRuleUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CloudWorkloadSecurityAgentRuleUpdateData) GetAttributesOk() (*CloudWorkloadSecurityAgentRuleUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CloudWorkloadSecurityAgentRuleUpdateData) SetAttributes(v CloudWorkloadSecurityAgentRuleUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### GetType

`func (o *CloudWorkloadSecurityAgentRuleUpdateData) GetType() CloudWorkloadSecurityAgentRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudWorkloadSecurityAgentRuleUpdateData) GetTypeOk() (*CloudWorkloadSecurityAgentRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudWorkloadSecurityAgentRuleUpdateData) SetType(v CloudWorkloadSecurityAgentRuleType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
