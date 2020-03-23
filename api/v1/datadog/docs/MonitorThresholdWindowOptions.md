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

`func (o *MonitorThresholdWindowOptions) GetRecoveryWindow() NullableString`

GetRecoveryWindow returns the RecoveryWindow field if non-nil, zero value otherwise.

### GetRecoveryWindowOk

`func (o *MonitorThresholdWindowOptions) GetRecoveryWindowOk() (NullableString, bool)`

GetRecoveryWindowOk returns a tuple with the RecoveryWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecoveryWindow

`func (o *MonitorThresholdWindowOptions) HasRecoveryWindow() bool`

HasRecoveryWindow returns a boolean if a field has been set.

### SetRecoveryWindow

`func (o *MonitorThresholdWindowOptions) SetRecoveryWindow(v NullableString)`

SetRecoveryWindow gets a reference to the given NullableString and assigns it to the RecoveryWindow field.

### SetRecoveryWindowExplicitNull

`func (o *MonitorThresholdWindowOptions) SetRecoveryWindowExplicitNull(b bool)`

SetRecoveryWindowExplicitNull (un)sets RecoveryWindow to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RecoveryWindow value is set to nil even if false is passed
### GetTriggerWindow

`func (o *MonitorThresholdWindowOptions) GetTriggerWindow() NullableString`

GetTriggerWindow returns the TriggerWindow field if non-nil, zero value otherwise.

### GetTriggerWindowOk

`func (o *MonitorThresholdWindowOptions) GetTriggerWindowOk() (NullableString, bool)`

GetTriggerWindowOk returns a tuple with the TriggerWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggerWindow

`func (o *MonitorThresholdWindowOptions) HasTriggerWindow() bool`

HasTriggerWindow returns a boolean if a field has been set.

### SetTriggerWindow

`func (o *MonitorThresholdWindowOptions) SetTriggerWindow(v NullableString)`

SetTriggerWindow gets a reference to the given NullableString and assigns it to the TriggerWindow field.

### SetTriggerWindowExplicitNull

`func (o *MonitorThresholdWindowOptions) SetTriggerWindowExplicitNull(b bool)`

SetTriggerWindowExplicitNull (un)sets TriggerWindow to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TriggerWindow value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


