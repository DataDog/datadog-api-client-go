# SecurityMonitoringRuleCaseCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | A rule case contains logical operations (&#x60;&gt;&#x60;,&#x60;&gt;&#x3D;&#x60;, &#x60;&amp;&amp;&#x60;, &#x60;||&#x60;) to determine if a signal should be generated based on the event counts in the previously defined queries. | [optional] 
**Name** | Pointer to **string** | Name of the case. | [optional] 
**Notifications** | Pointer to **[]string** | Notification targets for each rule case. | [optional] 
**Status** | [**SecurityMonitoringRuleSeverity**](SecurityMonitoringRuleSeverity.md) |  | 

## Methods

### NewSecurityMonitoringRuleCaseCreate

`func NewSecurityMonitoringRuleCaseCreate(status SecurityMonitoringRuleSeverity, ) *SecurityMonitoringRuleCaseCreate`

NewSecurityMonitoringRuleCaseCreate instantiates a new SecurityMonitoringRuleCaseCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringRuleCaseCreateWithDefaults

`func NewSecurityMonitoringRuleCaseCreateWithDefaults() *SecurityMonitoringRuleCaseCreate`

NewSecurityMonitoringRuleCaseCreateWithDefaults instantiates a new SecurityMonitoringRuleCaseCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *SecurityMonitoringRuleCaseCreate) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SecurityMonitoringRuleCaseCreate) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SecurityMonitoringRuleCaseCreate) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SecurityMonitoringRuleCaseCreate) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetName

`func (o *SecurityMonitoringRuleCaseCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityMonitoringRuleCaseCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityMonitoringRuleCaseCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityMonitoringRuleCaseCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *SecurityMonitoringRuleCaseCreate) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SecurityMonitoringRuleCaseCreate) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SecurityMonitoringRuleCaseCreate) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SecurityMonitoringRuleCaseCreate) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityMonitoringRuleCaseCreate) GetStatus() SecurityMonitoringRuleSeverity`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityMonitoringRuleCaseCreate) GetStatusOk() (*SecurityMonitoringRuleSeverity, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityMonitoringRuleCaseCreate) SetStatus(v SecurityMonitoringRuleSeverity)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


