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

// UsageIncidentManagementHour Incident management usage for a given organization for a given hour.
type UsageIncidentManagementHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the total number monthly active users from the start of the given hour's month until the given hour.
	MonthlyActiveUsers *int64 `json:"monthly_active_users,omitempty"`
}

// NewUsageIncidentManagementHour instantiates a new UsageIncidentManagementHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageIncidentManagementHour() *UsageIncidentManagementHour {
	this := UsageIncidentManagementHour{}
	return &this
}

// NewUsageIncidentManagementHourWithDefaults instantiates a new UsageIncidentManagementHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageIncidentManagementHourWithDefaults() *UsageIncidentManagementHour {
	this := UsageIncidentManagementHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageIncidentManagementHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIncidentManagementHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageIncidentManagementHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageIncidentManagementHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetMonthlyActiveUsers returns the MonthlyActiveUsers field value if set, zero value otherwise.
func (o *UsageIncidentManagementHour) GetMonthlyActiveUsers() int64 {
	if o == nil || o.MonthlyActiveUsers == nil {
		var ret int64
		return ret
	}
	return *o.MonthlyActiveUsers
}

// GetMonthlyActiveUsersOk returns a tuple with the MonthlyActiveUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIncidentManagementHour) GetMonthlyActiveUsersOk() (*int64, bool) {
	if o == nil || o.MonthlyActiveUsers == nil {
		return nil, false
	}
	return o.MonthlyActiveUsers, true
}

// HasMonthlyActiveUsers returns a boolean if a field has been set.
func (o *UsageIncidentManagementHour) HasMonthlyActiveUsers() bool {
	if o != nil && o.MonthlyActiveUsers != nil {
		return true
	}

	return false
}

// SetMonthlyActiveUsers gets a reference to the given int64 and assigns it to the MonthlyActiveUsers field.
func (o *UsageIncidentManagementHour) SetMonthlyActiveUsers(v int64) {
	o.MonthlyActiveUsers = &v
}

func (o UsageIncidentManagementHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.MonthlyActiveUsers != nil {
		toSerialize["monthly_active_users"] = o.MonthlyActiveUsers
	}
	return json.Marshal(toSerialize)
}

type NullableUsageIncidentManagementHour struct {
	value *UsageIncidentManagementHour
	isSet bool
}

func (v NullableUsageIncidentManagementHour) Get() *UsageIncidentManagementHour {
	return v.value
}

func (v *NullableUsageIncidentManagementHour) Set(val *UsageIncidentManagementHour) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageIncidentManagementHour) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageIncidentManagementHour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageIncidentManagementHour(val *UsageIncidentManagementHour) *NullableUsageIncidentManagementHour {
	return &NullableUsageIncidentManagementHour{value: val, isSet: true}
}

func (v NullableUsageIncidentManagementHour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageIncidentManagementHour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
