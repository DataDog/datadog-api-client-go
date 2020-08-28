# SecurityMonitoringRuleCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | A rule case contains logical operations (&#x60;&gt;&#x60;,&#x60;&gt;&#x3D;&#x60;, &#x60;&amp;&amp;&#x60;, &#x60;||&#x60;) to determine if a signal should be generated based on the event counts in the previously defined queries. | [optional] 
**Name** | Pointer to **string** | Name of the case. | [optional] 
**Notifications** | Pointer to **[]string** | Notification targets for each rule case. | [optional] 
**Status** | Pointer to [**SecurityMonitoringRuleSeverity**](SecurityMonitoringRuleSeverity.md) |  | [optional] 

## Methods

### NewSecurityMonitoringRuleCase

`func NewSecurityMonitoringRuleCase() *SecurityMonitoringRuleCase`

NewSecurityMonitoringRuleCase instantiates a new SecurityMonitoringRuleCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringRuleCaseWithDefaults

`func NewSecurityMonitoringRuleCaseWithDefaults() *SecurityMonitoringRuleCase`

NewSecurityMonitoringRuleCaseWithDefaults instantiates a new SecurityMonitoringRuleCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *SecurityMonitoringRuleCase) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SecurityMonitoringRuleCase) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SecurityMonitoringRuleCase) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SecurityMonitoringRuleCase) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetName

`func (o *SecurityMonitoringRuleCase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityMonitoringRuleCase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityMonitoringRuleCase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityMonitoringRuleCase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *SecurityMonitoringRuleCase) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SecurityMonitoringRuleCase) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SecurityMonitoringRuleCase) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SecurityMonitoringRuleCase) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityMonitoringRuleCase) GetStatus() SecurityMonitoringRuleSeverity`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityMonitoringRuleCase) GetStatusOk() (*SecurityMonitoringRuleSeverity, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityMonitoringRuleCase) SetStatus(v SecurityMonitoringRuleSeverity)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityMonitoringRuleCase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


