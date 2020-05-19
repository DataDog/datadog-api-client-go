# SecurityMonitoringRuleCreatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cases** | Pointer to [**[]SecurityMonitoringRuleCase**](SecurityMonitoringRuleCase.md) | Cases for generating signals. | 
**Enabled** | Pointer to **bool** | Whether the rule is enabled. | 
**Message** | Pointer to **string** | Message for generated signals. | 
**Name** | Pointer to **string** | The name of the rule | 
**Options** | Pointer to [**SecurityMonitoringRuleOptions**](SecurityMonitoringRuleOptions.md) |  | 
**Queries** | Pointer to [**[]SecurityMonitoringRuleQuery**](SecurityMonitoringRuleQuery.md) | Queries for selecting logs which are part of the rule. | 
**Tags** | Pointer to **[]string** | Tags for generated signals. | 

## Methods

### NewSecurityMonitoringRuleCreatePayload

`func NewSecurityMonitoringRuleCreatePayload(cases []SecurityMonitoringRuleCase, enabled bool, message string, name string, options SecurityMonitoringRuleOptions, queries []SecurityMonitoringRuleQuery, tags []string, ) *SecurityMonitoringRuleCreatePayload`

NewSecurityMonitoringRuleCreatePayload instantiates a new SecurityMonitoringRuleCreatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringRuleCreatePayloadWithDefaults

`func NewSecurityMonitoringRuleCreatePayloadWithDefaults() *SecurityMonitoringRuleCreatePayload`

NewSecurityMonitoringRuleCreatePayloadWithDefaults instantiates a new SecurityMonitoringRuleCreatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCases

`func (o *SecurityMonitoringRuleCreatePayload) GetCases() []SecurityMonitoringRuleCase`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *SecurityMonitoringRuleCreatePayload) GetCasesOk() (*[]SecurityMonitoringRuleCase, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *SecurityMonitoringRuleCreatePayload) SetCases(v []SecurityMonitoringRuleCase)`

SetCases sets Cases field to given value.


### GetEnabled

`func (o *SecurityMonitoringRuleCreatePayload) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityMonitoringRuleCreatePayload) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityMonitoringRuleCreatePayload) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMessage

`func (o *SecurityMonitoringRuleCreatePayload) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityMonitoringRuleCreatePayload) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityMonitoringRuleCreatePayload) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetName

`func (o *SecurityMonitoringRuleCreatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityMonitoringRuleCreatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityMonitoringRuleCreatePayload) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *SecurityMonitoringRuleCreatePayload) GetOptions() SecurityMonitoringRuleOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SecurityMonitoringRuleCreatePayload) GetOptionsOk() (*SecurityMonitoringRuleOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SecurityMonitoringRuleCreatePayload) SetOptions(v SecurityMonitoringRuleOptions)`

SetOptions sets Options field to given value.


### GetQueries

`func (o *SecurityMonitoringRuleCreatePayload) GetQueries() []SecurityMonitoringRuleQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *SecurityMonitoringRuleCreatePayload) GetQueriesOk() (*[]SecurityMonitoringRuleQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *SecurityMonitoringRuleCreatePayload) SetQueries(v []SecurityMonitoringRuleQuery)`

SetQueries sets Queries field to given value.


### GetTags

`func (o *SecurityMonitoringRuleCreatePayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityMonitoringRuleCreatePayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityMonitoringRuleCreatePayload) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


