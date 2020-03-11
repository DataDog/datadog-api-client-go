/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// MonitorThresholds struct for MonitorThresholds
type MonitorThresholds struct {
	Critical         *float64        `json:"critical,omitempty"`
	CriticalRecovery NullableFloat64 `json:"critical_recovery,omitempty"`
	Ok               NullableFloat64 `json:"ok,omitempty"`
	Unknown          NullableFloat64 `json:"unknown,omitempty"`
	Warning          NullableFloat64 `json:"warning,omitempty"`
	WarningRecovery  NullableFloat64 `json:"warning_recovery,omitempty"`
}

// NewMonitorThresholds instantiates a new MonitorThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorThresholds() *MonitorThresholds {
	this := MonitorThresholds{}
	return &this
}

// NewMonitorThresholdsWithDefaults instantiates a new MonitorThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorThresholdsWithDefaults() *MonitorThresholds {
	this := MonitorThresholds{}
	return &this
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *MonitorThresholds) GetCritical() float64 {
	if o == nil || o.Critical == nil {
		var ret float64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetCriticalOk() (float64, bool) {
	if o == nil || o.Critical == nil {
		var ret float64
		return ret, false
	}
	return *o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *MonitorThresholds) HasCritical() bool {
	if o != nil && o.Critical != nil {
		return true
	}

	return false
}

// SetCritical gets a reference to the given float64 and assigns it to the Critical field.
func (o *MonitorThresholds) SetCritical(v float64) {
	o.Critical = &v
}

// GetCriticalRecovery returns the CriticalRecovery field value if set, zero value otherwise.
func (o *MonitorThresholds) GetCriticalRecovery() NullableFloat64 {
	if o == nil {
		var ret NullableFloat64
		return ret
	}
	return o.CriticalRecovery
}

// GetCriticalRecoveryOk returns a tuple with the CriticalRecovery field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetCriticalRecoveryOk() (NullableFloat64, bool) {
	if o == nil {
		var ret NullableFloat64
		return ret, false
	}
	return o.CriticalRecovery, o.CriticalRecovery.IsSet()
}

// HasCriticalRecovery returns a boolean if a field has been set.
func (o *MonitorThresholds) HasCriticalRecovery() bool {
	if o != nil && o.CriticalRecovery.IsSet() {
		return true
	}

	return false
}

// SetCriticalRecovery gets a reference to the given NullableFloat64 and assigns it to the CriticalRecovery field.
func (o *MonitorThresholds) SetCriticalRecovery(v NullableFloat64) {
	o.CriticalRecovery = v
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *MonitorThresholds) GetOk() NullableFloat64 {
	if o == nil {
		var ret NullableFloat64
		return ret
	}
	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetOkOk() (NullableFloat64, bool) {
	if o == nil {
		var ret NullableFloat64
		return ret, false
	}
	return o.Ok, o.Ok.IsSet()
}

// HasOk returns a boolean if a field has been set.
func (o *MonitorThresholds) HasOk() bool {
	if o != nil && o.Ok.IsSet() {
		return true
	}

	return false
}

// SetOk gets a reference to the given NullableFloat64 and assigns it to the Ok field.
func (o *MonitorThresholds) SetOk(v NullableFloat64) {
	o.Ok = v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *MonitorThresholds) GetUnknown() NullableFloat64 {
	if o == nil {
		var ret NullableFloat64
		return ret
	}
	return o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetUnknownOk() (NullableFloat64, bool) {
	if o == nil {
		var ret NullableFloat64
		return ret, false
	}
	return o.Unknown, o.Unknown.IsSet()
}

// HasUnknown returns a boolean if a field has been set.
func (o *MonitorThresholds) HasUnknown() bool {
	if o != nil && o.Unknown.IsSet() {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given NullableFloat64 and assigns it to the Unknown field.
func (o *MonitorThresholds) SetUnknown(v NullableFloat64) {
	o.Unknown = v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *MonitorThresholds) GetWarning() NullableFloat64 {
	if o == nil {
		var ret NullableFloat64
		return ret
	}
	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetWarningOk() (NullableFloat64, bool) {
	if o == nil {
		var ret NullableFloat64
		return ret, false
	}
	return o.Warning, o.Warning.IsSet()
}

// HasWarning returns a boolean if a field has been set.
func (o *MonitorThresholds) HasWarning() bool {
	if o != nil && o.Warning.IsSet() {
		return true
	}

	return false
}

// SetWarning gets a reference to the given NullableFloat64 and assigns it to the Warning field.
func (o *MonitorThresholds) SetWarning(v NullableFloat64) {
	o.Warning = v
}

// GetWarningRecovery returns the WarningRecovery field value if set, zero value otherwise.
func (o *MonitorThresholds) GetWarningRecovery() NullableFloat64 {
	if o == nil {
		var ret NullableFloat64
		return ret
	}
	return o.WarningRecovery
}

// GetWarningRecoveryOk returns a tuple with the WarningRecovery field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetWarningRecoveryOk() (NullableFloat64, bool) {
	if o == nil {
		var ret NullableFloat64
		return ret, false
	}
	return o.WarningRecovery, o.WarningRecovery.IsSet()
}

// HasWarningRecovery returns a boolean if a field has been set.
func (o *MonitorThresholds) HasWarningRecovery() bool {
	if o != nil && o.WarningRecovery.IsSet() {
		return true
	}

	return false
}

// SetWarningRecovery gets a reference to the given NullableFloat64 and assigns it to the WarningRecovery field.
func (o *MonitorThresholds) SetWarningRecovery(v NullableFloat64) {
	o.WarningRecovery = v
}

func (o MonitorThresholds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Critical != nil {
		toSerialize["critical"] = o.Critical
	}
	if o.CriticalRecovery.IsSet() {
		toSerialize["critical_recovery"] = o.CriticalRecovery.Get()
	}
	if o.Ok.IsSet() {
		toSerialize["ok"] = o.Ok.Get()
	}
	if o.Unknown.IsSet() {
		toSerialize["unknown"] = o.Unknown.Get()
	}
	if o.Warning.IsSet() {
		toSerialize["warning"] = o.Warning.Get()
	}
	if o.WarningRecovery.IsSet() {
		toSerialize["warning_recovery"] = o.WarningRecovery.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorThresholds struct {
	value *MonitorThresholds
	isSet bool
}

func (v NullableMonitorThresholds) Get() *MonitorThresholds {
	return v.value
}

func (v NullableMonitorThresholds) Set(val *MonitorThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorThresholds) IsSet() bool {
	return v.isSet
}

func (v NullableMonitorThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorThresholds(val *MonitorThresholds) *NullableMonitorThresholds {
	return &NullableMonitorThresholds{value: val, isSet: true}
}

func (v NullableMonitorThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
