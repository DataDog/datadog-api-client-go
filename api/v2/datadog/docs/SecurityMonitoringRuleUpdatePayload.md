# SecurityMonitoringRuleUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cases** | Pointer to [**[]SecurityMonitoringRuleCase**](SecurityMonitoringRuleCase.md) | Cases for generating signals. | [optional] 
**Enabled** | Pointer to **bool** | Whether the rule is enabled. | [optional] 
**Message** | Pointer to **string** | Message for generated signals. | [optional] 
**Name** | Pointer to **string** | Name of the rule. | [optional] 
**Options** | Pointer to [**SecurityMonitoringRuleOptions**](SecurityMonitoringRuleOptions.md) |  | [optional] 
**Queries** | Pointer to [**[]SecurityMonitoringRuleQuery**](SecurityMonitoringRuleQuery.md) | Queries for selecting logs which are part of the rule. | [optional] 
**Tags** | Pointer to **[]string** | Tags for generated signals. | [optional] 

## Methods

### NewSecurityMonitoringRuleUpdatePayload

`func NewSecurityMonitoringRuleUpdatePayload() *SecurityMonitoringRuleUpdatePayload`

NewSecurityMonitoringRuleUpdatePayload instantiates a new SecurityMonitoringRuleUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringRuleUpdatePayloadWithDefaults

`func NewSecurityMonitoringRuleUpdatePayloadWithDefaults() *SecurityMonitoringRuleUpdatePayload`

NewSecurityMonitoringRuleUpdatePayloadWithDefaults instantiates a new SecurityMonitoringRuleUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCases

`func (o *SecurityMonitoringRuleUpdatePayload) GetCases() []SecurityMonitoringRuleCase`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *SecurityMonitoringRuleUpdatePayload) GetCasesOk() (*[]SecurityMonitoringRuleCase, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *SecurityMonitoringRuleUpdatePayload) SetCases(v []SecurityMonitoringRuleCase)`

SetCases sets Cases field to given value.

### HasCases

`func (o *SecurityMonitoringRuleUpdatePayload) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetEnabled

`func (o *SecurityMonitoringRuleUpdatePayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityMonitoringRuleUpdatePayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityMonitoringRuleUpdatePayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityMonitoringRuleUpdatePayload) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMessage

`func (o *SecurityMonitoringRuleUpdatePayload) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityMonitoringRuleUpdatePayload) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityMonitoringRuleUpdatePayload) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SecurityMonitoringRuleUpdatePayload) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *SecurityMonitoringRuleUpdatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityMonitoringRuleUpdatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityMonitoringRuleUpdatePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityMonitoringRuleUpdatePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *SecurityMonitoringRuleUpdatePayload) GetOptions() SecurityMonitoringRuleOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SecurityMonitoringRuleUpdatePayload) GetOptionsOk() (*SecurityMonitoringRuleOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SecurityMonitoringRuleUpdatePayload) SetOptions(v SecurityMonitoringRuleOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SecurityMonitoringRuleUpdatePayload) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetQueries

`func (o *SecurityMonitoringRuleUpdatePayload) GetQueries() []SecurityMonitoringRuleQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SecurityMonitoringRuleUpdatePayload) GetQueriesOk() (*[]SecurityMonitoringRuleQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SecurityMonitoringRuleUpdatePayload) SetQueries(v []SecurityMonitoringRuleQuery)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *SecurityMonitoringRuleUpdatePayload) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetTags

`func (o *SecurityMonitoringRuleUpdatePayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityMonitoringRuleUpdatePayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityMonitoringRuleUpdatePayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecurityMonitoringRuleUpdatePayload) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


