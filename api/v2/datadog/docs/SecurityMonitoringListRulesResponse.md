# SecurityMonitoringListRulesResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Data** | Pointer to [**[]SecurityMonitoringRuleResponse**](SecurityMonitoringRuleResponse.md) | Array containing the list of rules. | [optional] 
**Meta** | Pointer to [**ResponseMetaAttributes**](ResponseMetaAttributes.md) |  | [optional] 

## Methods

### NewSecurityMonitoringListRulesResponse

`func NewSecurityMonitoringListRulesResponse() *SecurityMonitoringListRulesResponse`

NewSecurityMonitoringListRulesResponse instantiates a new SecurityMonitoringListRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringListRulesResponseWithDefaults

`func NewSecurityMonitoringListRulesResponseWithDefaults() *SecurityMonitoringListRulesResponse`

NewSecurityMonitoringListRulesResponseWithDefaults instantiates a new SecurityMonitoringListRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecurityMonitoringListRulesResponse) GetData() []SecurityMonitoringRuleResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityMonitoringListRulesResponse) GetDataOk() (*[]SecurityMonitoringRuleResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityMonitoringListRulesResponse) SetData(v []SecurityMonitoringRuleResponse)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityMonitoringListRulesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *SecurityMonitoringListRulesResponse) GetMeta() ResponseMetaAttributes`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecurityMonitoringListRulesResponse) GetMetaOk() (*ResponseMetaAttributes, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecurityMonitoringListRulesResponse) SetMeta(v ResponseMetaAttributes)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecurityMonitoringListRulesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


