# MonitorThresholds

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Critical** | Pointer to **float64** | The monitor &#x60;CRITICAL&#x60; threshold. | [optional] 
**CriticalRecovery** | Pointer to **NullableFloat64** | The monitor &#x60;CRITICAL&#x60; recovery threshold. | [optional] 
**Ok** | Pointer to **NullableFloat64** | The monitor &#x60;OK&#x60; threshold. | [optional] 
**Unknown** | Pointer to **NullableFloat64** | The monitor UNKNOWN threshold. | [optional] 
**Warning** | Pointer to **NullableFloat64** | The monitor &#x60;WARNING&#x60; threshold. | [optional] 
**WarningRecovery** | Pointer to **NullableFloat64** | The monitor &#x60;WARNING&#x60; recovery threshold. | [optional] 

## Methods

### NewMonitorThresholds

`func NewMonitorThresholds() *MonitorThresholds`

NewMonitorThresholds instantiates a new MonitorThresholds object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMonitorThresholdsWithDefaults

`func NewMonitorThresholdsWithDefaults() *MonitorThresholds`

NewMonitorThresholdsWithDefaults instantiates a new MonitorThresholds object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCritical

`func (o *MonitorThresholds) GetCritical() float64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *MonitorThresholds) GetCriticalOk() (*float64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *MonitorThresholds) SetCritical(v float64)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *MonitorThresholds) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetCriticalRecovery

`func (o *MonitorThresholds) GetCriticalRecovery() float64`

GetCriticalRecovery returns the CriticalRecovery field if non-nil, zero value otherwise.

### GetCriticalRecoveryOk

`func (o *MonitorThresholds) GetCriticalRecoveryOk() (*float64, bool)`

GetCriticalRecoveryOk returns a tuple with the CriticalRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalRecovery

`func (o *MonitorThresholds) SetCriticalRecovery(v float64)`

SetCriticalRecovery sets CriticalRecovery field to given value.

### HasCriticalRecovery

`func (o *MonitorThresholds) HasCriticalRecovery() bool`

HasCriticalRecovery returns a boolean if a field has been set.

### SetCriticalRecoveryNil

`func (o *MonitorThresholds) SetCriticalRecoveryNil(b bool)`

 SetCriticalRecoveryNil sets the value for CriticalRecovery to be an explicit nil

### UnsetCriticalRecovery
`func (o *MonitorThresholds) UnsetCriticalRecovery()`

UnsetCriticalRecovery ensures that no value is present for CriticalRecovery, not even an explicit nil
### GetOk

`func (o *MonitorThresholds) GetOk() float64`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *MonitorThresholds) GetOkOk() (*float64, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *MonitorThresholds) SetOk(v float64)`

SetOk sets Ok field to given value.

### HasOk

`func (o *MonitorThresholds) HasOk() bool`

HasOk returns a boolean if a field has been set.

### SetOkNil

`func (o *MonitorThresholds) SetOkNil(b bool)`

 SetOkNil sets the value for Ok to be an explicit nil

### UnsetOk
`func (o *MonitorThresholds) UnsetOk()`

UnsetOk ensures that no value is present for Ok, not even an explicit nil
### GetUnknown

`func (o *MonitorThresholds) GetUnknown() float64`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *MonitorThresholds) GetUnknownOk() (*float64, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *MonitorThresholds) SetUnknown(v float64)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *MonitorThresholds) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.

### SetUnknownNil

`func (o *MonitorThresholds) SetUnknownNil(b bool)`

 SetUnknownNil sets the value for Unknown to be an explicit nil

### UnsetUnknown
`func (o *MonitorThresholds) UnsetUnknown()`

UnsetUnknown ensures that no value is present for Unknown, not even an explicit nil
### GetWarning

`func (o *MonitorThresholds) GetWarning() float64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *MonitorThresholds) GetWarningOk() (*float64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *MonitorThresholds) SetWarning(v float64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *MonitorThresholds) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarningNil

`func (o *MonitorThresholds) SetWarningNil(b bool)`

 SetWarningNil sets the value for Warning to be an explicit nil

### UnsetWarning
`func (o *MonitorThresholds) UnsetWarning()`

UnsetWarning ensures that no value is present for Warning, not even an explicit nil
### GetWarningRecovery

`func (o *MonitorThresholds) GetWarningRecovery() float64`

GetWarningRecovery returns the WarningRecovery field if non-nil, zero value otherwise.

### GetWarningRecoveryOk

`func (o *MonitorThresholds) GetWarningRecoveryOk() (*float64, bool)`

GetWarningRecoveryOk returns a tuple with the WarningRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningRecovery

`func (o *MonitorThresholds) SetWarningRecovery(v float64)`

SetWarningRecovery sets WarningRecovery field to given value.

### HasWarningRecovery

`func (o *MonitorThresholds) HasWarningRecovery() bool`

HasWarningRecovery returns a boolean if a field has been set.

### SetWarningRecoveryNil

`func (o *MonitorThresholds) SetWarningRecoveryNil(b bool)`

 SetWarningRecoveryNil sets the value for WarningRecovery to be an explicit nil

### UnsetWarningRecovery
`func (o *MonitorThresholds) UnsetWarningRecovery()`

UnsetWarningRecovery ensures that no value is present for WarningRecovery, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


