# MonitorThresholdWindowOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryWindow** | Pointer to **NullableString** | Describes how long an anomalous metric must be normal before the alert recovers. | [optional] 
**TriggerWindow** | Pointer to **NullableString** | Describes how long a metric must be anomalous before an alert triggers. | [optional] 

## Methods

### NewMonitorThresholdWindowOptions

`func NewMonitorThresholdWindowOptions() *MonitorThresholdWindowOptions`

NewMonitorThresholdWindowOptions instantiates a new MonitorThresholdWindowOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorThresholdWindowOptionsWithDefaults

`func NewMonitorThresholdWindowOptionsWithDefaults() *MonitorThresholdWindowOptions`

NewMonitorThresholdWindowOptionsWithDefaults instantiates a new MonitorThresholdWindowOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryWindow

`func (o *MonitorThresholdWindowOptions) GetRecoveryWindow() string`

GetRecoveryWindow returns the RecoveryWindow field if non-nil, zero value otherwise.

### GetRecoveryWindowOk

`func (o *MonitorThresholdWindowOptions) GetRecoveryWindowOk() (*string, bool)`

GetRecoveryWindowOk returns a tuple with the RecoveryWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryWindow

`func (o *MonitorThresholdWindowOptions) SetRecoveryWindow(v string)`

SetRecoveryWindow sets RecoveryWindow field to given value.

### HasRecoveryWindow

`func (o *MonitorThresholdWindowOptions) HasRecoveryWindow() bool`

HasRecoveryWindow returns a boolean if a field has been set.

### SetRecoveryWindowNil

`func (o *MonitorThresholdWindowOptions) SetRecoveryWindowNil(b bool)`

 SetRecoveryWindowNil sets the value for RecoveryWindow to be an explicit nil

### UnsetRecoveryWindow
`func (o *MonitorThresholdWindowOptions) UnsetRecoveryWindow()`

UnsetRecoveryWindow ensures that no value is present for RecoveryWindow, not even an explicit nil
### GetTriggerWindow

`func (o *MonitorThresholdWindowOptions) GetTriggerWindow() string`

GetTriggerWindow returns the TriggerWindow field if non-nil, zero value otherwise.

### GetTriggerWindowOk

`func (o *MonitorThresholdWindowOptions) GetTriggerWindowOk() (*string, bool)`

GetTriggerWindowOk returns a tuple with the TriggerWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerWindow

`func (o *MonitorThresholdWindowOptions) SetTriggerWindow(v string)`

SetTriggerWindow sets TriggerWindow field to given value.

### HasTriggerWindow

`func (o *MonitorThresholdWindowOptions) HasTriggerWindow() bool`

HasTriggerWindow returns a boolean if a field has been set.

### SetTriggerWindowNil

`func (o *MonitorThresholdWindowOptions) SetTriggerWindowNil(b bool)`

 SetTriggerWindowNil sets the value for TriggerWindow to be an explicit nil

### UnsetTriggerWindow
`func (o *MonitorThresholdWindowOptions) UnsetTriggerWindow()`

UnsetTriggerWindow ensures that no value is present for TriggerWindow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


