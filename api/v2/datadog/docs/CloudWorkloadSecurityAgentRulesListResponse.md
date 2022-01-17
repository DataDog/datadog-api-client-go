# CloudWorkloadSecurityAgentRulesListResponse

## Properties

| Name     | Type                                                                                         | Description                    | Notes      |
| -------- | -------------------------------------------------------------------------------------------- | ------------------------------ | ---------- |
| **Data** | Pointer to [**[]CloudWorkloadSecurityAgentRuleData**](CloudWorkloadSecurityAgentRuleData.md) | A list of Agent rules objects. | [optional] |

## Methods

### NewCloudWorkloadSecurityAgentRulesListResponse

`func NewCloudWorkloadSecurityAgentRulesListResponse() *CloudWorkloadSecurityAgentRulesListResponse`

NewCloudWorkloadSecurityAgentRulesListResponse instantiates a new CloudWorkloadSecurityAgentRulesListResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudWorkloadSecurityAgentRulesListResponseWithDefaults

`func NewCloudWorkloadSecurityAgentRulesListResponseWithDefaults() *CloudWorkloadSecurityAgentRulesListResponse`

NewCloudWorkloadSecurityAgentRulesListResponseWithDefaults instantiates a new CloudWorkloadSecurityAgentRulesListResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *CloudWorkloadSecurityAgentRulesListResponse) GetData() []CloudWorkloadSecurityAgentRuleData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudWorkloadSecurityAgentRulesListResponse) GetDataOk() (*[]CloudWorkloadSecurityAgentRuleData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudWorkloadSecurityAgentRulesListResponse) SetData(v []CloudWorkloadSecurityAgentRuleData)`

SetData sets Data field to given value.

### HasData

`func (o *CloudWorkloadSecurityAgentRulesListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
