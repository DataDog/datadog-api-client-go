# MonitorThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Critical** | Pointer to **float64** | The monitor CRITICAL threshold. | [optional] 
**CriticalRecovery** | Pointer to **NullableFloat64** | The monitor CRITICAL recovery threshold. | [optional] 
**Ok** | Pointer to **NullableFloat64** | The monitor OK threshold. | [optional] 
**Unknown** | Pointer to **NullableFloat64** | TODO. | [optional] 
**Warning** | Pointer to **NullableFloat64** | The monitor WARNING threshold. | [optional] 
**WarningRecovery** | Pointer to **NullableFloat64** | The monitor WARNING recovery threshold. | [optional] 

## Methods

### NewMonitorThresholds

`func NewMonitorThresholds() *MonitorThresholds`

NewMonitorThresholds instantiates a new MonitorThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorThresholdsWithDefaults

`func NewMonitorThresholdsWithDefaults() *MonitorThresholds`

NewMonitorThresholdsWithDefaults instantiates a new MonitorThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCritical

`func (o *MonitorThresholds) GetCritical() float64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *MonitorThresholds) GetCriticalOk() (float64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCritical

`func (o *MonitorThresholds) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### SetCritical

`func (o *MonitorThresholds) SetCritical(v float64)`

SetCritical gets a reference to the given float64 and assigns it to the Critical field.

### GetCriticalRecovery

`func (o *MonitorThresholds) GetCriticalRecovery() NullableFloat64`

GetCriticalRecovery returns the CriticalRecovery field if non-nil, zero value otherwise.

### GetCriticalRecoveryOk

`func (o *MonitorThresholds) GetCriticalRecoveryOk() (NullableFloat64, bool)`

GetCriticalRecoveryOk returns a tuple with the CriticalRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCriticalRecovery

`func (o *MonitorThresholds) HasCriticalRecovery() bool`

HasCriticalRecovery returns a boolean if a field has been set.

### SetCriticalRecovery

`func (o *MonitorThresholds) SetCriticalRecovery(v NullableFloat64)`

SetCriticalRecovery gets a reference to the given NullableFloat64 and assigns it to the CriticalRecovery field.

### SetCriticalRecoveryExplicitNull

`func (o *MonitorThresholds) SetCriticalRecoveryExplicitNull(b bool)`

SetCriticalRecoveryExplicitNull (un)sets CriticalRecovery to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CriticalRecovery value is set to nil even if false is passed
### GetOk

`func (o *MonitorThresholds) GetOk() NullableFloat64`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *MonitorThresholds) GetOkOk() (NullableFloat64, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOk

`func (o *MonitorThresholds) HasOk() bool`

HasOk returns a boolean if a field has been set.

### SetOk

`func (o *MonitorThresholds) SetOk(v NullableFloat64)`

SetOk gets a reference to the given NullableFloat64 and assigns it to the Ok field.

### SetOkExplicitNull

`func (o *MonitorThresholds) SetOkExplicitNull(b bool)`

SetOkExplicitNull (un)sets Ok to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Ok value is set to nil even if false is passed
### GetUnknown

`func (o *MonitorThresholds) GetUnknown() NullableFloat64`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *MonitorThresholds) GetUnknownOk() (NullableFloat64, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnknown

`func (o *MonitorThresholds) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.

### SetUnknown

`func (o *MonitorThresholds) SetUnknown(v NullableFloat64)`

SetUnknown gets a reference to the given NullableFloat64 and assigns it to the Unknown field.

### SetUnknownExplicitNull

`func (o *MonitorThresholds) SetUnknownExplicitNull(b bool)`

SetUnknownExplicitNull (un)sets Unknown to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Unknown value is set to nil even if false is passed
### GetWarning

`func (o *MonitorThresholds) GetWarning() NullableFloat64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *MonitorThresholds) GetWarningOk() (NullableFloat64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWarning

`func (o *MonitorThresholds) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarning

`func (o *MonitorThresholds) SetWarning(v NullableFloat64)`

SetWarning gets a reference to the given NullableFloat64 and assigns it to the Warning field.

### SetWarningExplicitNull

`func (o *MonitorThresholds) SetWarningExplicitNull(b bool)`

SetWarningExplicitNull (un)sets Warning to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Warning value is set to nil even if false is passed
### GetWarningRecovery

`func (o *MonitorThresholds) GetWarningRecovery() NullableFloat64`

GetWarningRecovery returns the WarningRecovery field if non-nil, zero value otherwise.

### GetWarningRecoveryOk

`func (o *MonitorThresholds) GetWarningRecoveryOk() (NullableFloat64, bool)`

GetWarningRecoveryOk returns a tuple with the WarningRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWarningRecovery

`func (o *MonitorThresholds) HasWarningRecovery() bool`

HasWarningRecovery returns a boolean if a field has been set.

### SetWarningRecovery

`func (o *MonitorThresholds) SetWarningRecovery(v NullableFloat64)`

SetWarningRecovery gets a reference to the given NullableFloat64 and assigns it to the WarningRecovery field.

### SetWarningRecoveryExplicitNull

`func (o *MonitorThresholds) SetWarningRecoveryExplicitNull(b bool)`

SetWarningRecoveryExplicitNull (un)sets WarningRecovery to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WarningRecovery value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


