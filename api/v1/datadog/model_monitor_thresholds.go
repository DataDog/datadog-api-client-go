/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// MonitorThresholds struct for MonitorThresholds
type MonitorThresholds struct {
	Critical         *float64 `json:"critical,omitempty"`
	CriticalRecovery *float64 `json:"critical_recovery,omitempty"`
	Ok               *float64 `json:"ok,omitempty"`
	Unknown          *float64 `json:"unknown,omitempty"`
	Warning          *float64 `json:"warning,omitempty"`
	WarningRecovery  *float64 `json:"warning_recovery,omitempty"`
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
func (o *MonitorThresholds) GetCriticalRecovery() float64 {
	if o == nil || o.CriticalRecovery == nil {
		var ret float64
		return ret
	}
	return *o.CriticalRecovery
}

// GetCriticalRecoveryOk returns a tuple with the CriticalRecovery field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetCriticalRecoveryOk() (float64, bool) {
	if o == nil || o.CriticalRecovery == nil {
		var ret float64
		return ret, false
	}
	return *o.CriticalRecovery, true
}

// HasCriticalRecovery returns a boolean if a field has been set.
func (o *MonitorThresholds) HasCriticalRecovery() bool {
	if o != nil && o.CriticalRecovery != nil {
		return true
	}

	return false
}

// SetCriticalRecovery gets a reference to the given float64 and assigns it to the CriticalRecovery field.
func (o *MonitorThresholds) SetCriticalRecovery(v float64) {
	o.CriticalRecovery = &v
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *MonitorThresholds) GetOk() float64 {
	if o == nil || o.Ok == nil {
		var ret float64
		return ret
	}
	return *o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetOkOk() (float64, bool) {
	if o == nil || o.Ok == nil {
		var ret float64
		return ret, false
	}
	return *o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *MonitorThresholds) HasOk() bool {
	if o != nil && o.Ok != nil {
		return true
	}

	return false
}

// SetOk gets a reference to the given float64 and assigns it to the Ok field.
func (o *MonitorThresholds) SetOk(v float64) {
	o.Ok = &v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *MonitorThresholds) GetUnknown() float64 {
	if o == nil || o.Unknown == nil {
		var ret float64
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetUnknownOk() (float64, bool) {
	if o == nil || o.Unknown == nil {
		var ret float64
		return ret, false
	}
	return *o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *MonitorThresholds) HasUnknown() bool {
	if o != nil && o.Unknown != nil {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given float64 and assigns it to the Unknown field.
func (o *MonitorThresholds) SetUnknown(v float64) {
	o.Unknown = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *MonitorThresholds) GetWarning() float64 {
	if o == nil || o.Warning == nil {
		var ret float64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetWarningOk() (float64, bool) {
	if o == nil || o.Warning == nil {
		var ret float64
		return ret, false
	}
	return *o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *MonitorThresholds) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given float64 and assigns it to the Warning field.
func (o *MonitorThresholds) SetWarning(v float64) {
	o.Warning = &v
}

// GetWarningRecovery returns the WarningRecovery field value if set, zero value otherwise.
func (o *MonitorThresholds) GetWarningRecovery() float64 {
	if o == nil || o.WarningRecovery == nil {
		var ret float64
		return ret
	}
	return *o.WarningRecovery
}

// GetWarningRecoveryOk returns a tuple with the WarningRecovery field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorThresholds) GetWarningRecoveryOk() (float64, bool) {
	if o == nil || o.WarningRecovery == nil {
		var ret float64
		return ret, false
	}
	return *o.WarningRecovery, true
}

// HasWarningRecovery returns a boolean if a field has been set.
func (o *MonitorThresholds) HasWarningRecovery() bool {
	if o != nil && o.WarningRecovery != nil {
		return true
	}

	return false
}

// SetWarningRecovery gets a reference to the given float64 and assigns it to the WarningRecovery field.
func (o *MonitorThresholds) SetWarningRecovery(v float64) {
	o.WarningRecovery = &v
}

type NullableMonitorThresholds struct {
	Value        MonitorThresholds
	ExplicitNull bool
}

func (v NullableMonitorThresholds) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMonitorThresholds) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
