# SecurityMonitoringRuleOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationWindow** | Pointer to [**SecurityMonitoringRuleEvaluationWindow**](SecurityMonitoringRuleEvaluationWindow.md) |  | [optional] 
**KeepAlive** | Pointer to [**SecurityMonitoringRuleKeepAlive**](SecurityMonitoringRuleKeepAlive.md) |  | [optional] 
**MaxSignalDuration** | Pointer to [**SecurityMonitoringRuleMaxSignalDuration**](SecurityMonitoringRuleMaxSignalDuration.md) |  | [optional] 

## Methods

### NewSecurityMonitoringRuleOptions

`func NewSecurityMonitoringRuleOptions() *SecurityMonitoringRuleOptions`

NewSecurityMonitoringRuleOptions instantiates a new SecurityMonitoringRuleOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMonitoringRuleOptionsWithDefaults

`func NewSecurityMonitoringRuleOptionsWithDefaults() *SecurityMonitoringRuleOptions`

NewSecurityMonitoringRuleOptionsWithDefaults instantiates a new SecurityMonitoringRuleOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluationWindow

`func (o *SecurityMonitoringRuleOptions) GetEvaluationWindow() SecurityMonitoringRuleEvaluationWindow`

GetEvaluationWindow returns the EvaluationWindow field if non-nil, zero value otherwise.

### GetEvaluationWindowOk

`func (o *SecurityMonitoringRuleOptions) GetEvaluationWindowOk() (*SecurityMonitoringRuleEvaluationWindow, bool)`

GetEvaluationWindowOk returns a tuple with the EvaluationWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationWindow

`func (o *SecurityMonitoringRuleOptions) SetEvaluationWindow(v SecurityMonitoringRuleEvaluationWindow)`

SetEvaluationWindow sets EvaluationWindow field to given value.

### HasEvaluationWindow

`func (o *SecurityMonitoringRuleOptions) HasEvaluationWindow() bool`

HasEvaluationWindow returns a boolean if a field has been set.

### GetKeepAlive

`func (o *SecurityMonitoringRuleOptions) GetKeepAlive() SecurityMonitoringRuleKeepAlive`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *SecurityMonitoringRuleOptions) GetKeepAliveOk() (*SecurityMonitoringRuleKeepAlive, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *SecurityMonitoringRuleOptions) SetKeepAlive(v SecurityMonitoringRuleKeepAlive)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *SecurityMonitoringRuleOptions) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetMaxSignalDuration

`func (o *SecurityMonitoringRuleOptions) GetMaxSignalDuration() SecurityMonitoringRuleMaxSignalDuration`

GetMaxSignalDuration returns the MaxSignalDuration field if non-nil, zero value otherwise.

### GetMaxSignalDurationOk

`func (o *SecurityMonitoringRuleOptions) GetMaxSignalDurationOk() (*SecurityMonitoringRuleMaxSignalDuration, bool)`

GetMaxSignalDurationOk returns a tuple with the MaxSignalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSignalDuration

`func (o *SecurityMonitoringRuleOptions) SetMaxSignalDuration(v SecurityMonitoringRuleMaxSignalDuration)`

SetMaxSignalDuration sets MaxSignalDuration field to given value.

### HasMaxSignalDuration

`func (o *SecurityMonitoringRuleOptions) HasMaxSignalDuration() bool`

HasMaxSignalDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


