/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"time"
)

// UsageAnalyzedLogsHour The number of analyzed logs for each hour for a given organization.
type UsageAnalyzedLogsHour struct {
	// Contains the number of analyzed logs.
	AnalyzedLogs *int64 `json:"analyzed_logs,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
}

// NewUsageAnalyzedLogsHour instantiates a new UsageAnalyzedLogsHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageAnalyzedLogsHour() *UsageAnalyzedLogsHour {
	this := UsageAnalyzedLogsHour{}
	return &this
}

// NewUsageAnalyzedLogsHourWithDefaults instantiates a new UsageAnalyzedLogsHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageAnalyzedLogsHourWithDefaults() *UsageAnalyzedLogsHour {
	this := UsageAnalyzedLogsHour{}
	return &this
}

// GetAnalyzedLogs returns the AnalyzedLogs field value if set, zero value otherwise.
func (o *UsageAnalyzedLogsHour) GetAnalyzedLogs() int64 {
	if o == nil || o.AnalyzedLogs == nil {
		var ret int64
		return ret
	}
	return *o.AnalyzedLogs
}

// GetAnalyzedLogsOk returns a tuple with the AnalyzedLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAnalyzedLogsHour) GetAnalyzedLogsOk() (*int64, bool) {
	if o == nil || o.AnalyzedLogs == nil {
		return nil, false
	}
	return o.AnalyzedLogs, true
}

// HasAnalyzedLogs returns a boolean if a field has been set.
func (o *UsageAnalyzedLogsHour) HasAnalyzedLogs() bool {
	if o != nil && o.AnalyzedLogs != nil {
		return true
	}

	return false
}

// SetAnalyzedLogs gets a reference to the given int64 and assigns it to the AnalyzedLogs field.
func (o *UsageAnalyzedLogsHour) SetAnalyzedLogs(v int64) {
	o.AnalyzedLogs = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageAnalyzedLogsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAnalyzedLogsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageAnalyzedLogsHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageAnalyzedLogsHour) SetHour(v time.Time) {
	o.Hour = &v
}

func (o UsageAnalyzedLogsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnalyzedLogs != nil {
		toSerialize["analyzed_logs"] = o.AnalyzedLogs
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	return json.Marshal(toSerialize)
}

type NullableUsageAnalyzedLogsHour struct {
	value *UsageAnalyzedLogsHour
	isSet bool
}

func (v NullableUsageAnalyzedLogsHour) Get() *UsageAnalyzedLogsHour {
	return v.value
}

func (v *NullableUsageAnalyzedLogsHour) Set(val *UsageAnalyzedLogsHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageAnalyzedLogsHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageAnalyzedLogsHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageAnalyzedLogsHour(val *UsageAnalyzedLogsHour) *NullableUsageAnalyzedLogsHour {
	return &NullableUsageAnalyzedLogsHour{value: val, isSet: true}
}

func (v NullableUsageAnalyzedLogsHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageAnalyzedLogsHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


