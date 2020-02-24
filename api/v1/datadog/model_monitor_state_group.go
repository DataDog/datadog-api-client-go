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

// MonitorStateGroup struct for MonitorStateGroup
type MonitorStateGroup struct {
	LastDataTs      *int64                  `json:"last_data_ts,omitempty"`
	LastNodataTs    *int64                  `json:"last_nodata_ts,omitempty"`
	LastNotifiedTs  *int64                  `json:"last_notified_ts,omitempty"`
	LastResolvedTs  *int64                  `json:"last_resolved_ts,omitempty"`
	LastTriggeredTs *int64                  `json:"last_triggered_ts,omitempty"`
	Message         *string                 `json:"message,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Status          *MonitorOverallStates   `json:"status,omitempty"`
	TriggeringValue *MonitorStateGroupValue `json:"triggering_value,omitempty"`
}

// NewMonitorStateGroup instantiates a new MonitorStateGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorStateGroup() *MonitorStateGroup {
	this := MonitorStateGroup{}
	return &this
}

// NewMonitorStateGroupWithDefaults instantiates a new MonitorStateGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorStateGroupWithDefaults() *MonitorStateGroup {
	this := MonitorStateGroup{}
	return &this
}

// GetLastDataTs returns the LastDataTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastDataTs() int64 {
	if o == nil || o.LastDataTs == nil {
		var ret int64
		return ret
	}
	return *o.LastDataTs
}

// GetLastDataTsOk returns a tuple with the LastDataTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastDataTsOk() (int64, bool) {
	if o == nil || o.LastDataTs == nil {
		var ret int64
		return ret, false
	}
	return *o.LastDataTs, true
}

// HasLastDataTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastDataTs() bool {
	if o != nil && o.LastDataTs != nil {
		return true
	}

	return false
}

// SetLastDataTs gets a reference to the given int64 and assigns it to the LastDataTs field.
func (o *MonitorStateGroup) SetLastDataTs(v int64) {
	o.LastDataTs = &v
}

// GetLastNodataTs returns the LastNodataTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastNodataTs() int64 {
	if o == nil || o.LastNodataTs == nil {
		var ret int64
		return ret
	}
	return *o.LastNodataTs
}

// GetLastNodataTsOk returns a tuple with the LastNodataTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastNodataTsOk() (int64, bool) {
	if o == nil || o.LastNodataTs == nil {
		var ret int64
		return ret, false
	}
	return *o.LastNodataTs, true
}

// HasLastNodataTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastNodataTs() bool {
	if o != nil && o.LastNodataTs != nil {
		return true
	}

	return false
}

// SetLastNodataTs gets a reference to the given int64 and assigns it to the LastNodataTs field.
func (o *MonitorStateGroup) SetLastNodataTs(v int64) {
	o.LastNodataTs = &v
}

// GetLastNotifiedTs returns the LastNotifiedTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastNotifiedTs() int64 {
	if o == nil || o.LastNotifiedTs == nil {
		var ret int64
		return ret
	}
	return *o.LastNotifiedTs
}

// GetLastNotifiedTsOk returns a tuple with the LastNotifiedTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastNotifiedTsOk() (int64, bool) {
	if o == nil || o.LastNotifiedTs == nil {
		var ret int64
		return ret, false
	}
	return *o.LastNotifiedTs, true
}

// HasLastNotifiedTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastNotifiedTs() bool {
	if o != nil && o.LastNotifiedTs != nil {
		return true
	}

	return false
}

// SetLastNotifiedTs gets a reference to the given int64 and assigns it to the LastNotifiedTs field.
func (o *MonitorStateGroup) SetLastNotifiedTs(v int64) {
	o.LastNotifiedTs = &v
}

// GetLastResolvedTs returns the LastResolvedTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastResolvedTs() int64 {
	if o == nil || o.LastResolvedTs == nil {
		var ret int64
		return ret
	}
	return *o.LastResolvedTs
}

// GetLastResolvedTsOk returns a tuple with the LastResolvedTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastResolvedTsOk() (int64, bool) {
	if o == nil || o.LastResolvedTs == nil {
		var ret int64
		return ret, false
	}
	return *o.LastResolvedTs, true
}

// HasLastResolvedTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastResolvedTs() bool {
	if o != nil && o.LastResolvedTs != nil {
		return true
	}

	return false
}

// SetLastResolvedTs gets a reference to the given int64 and assigns it to the LastResolvedTs field.
func (o *MonitorStateGroup) SetLastResolvedTs(v int64) {
	o.LastResolvedTs = &v
}

// GetLastTriggeredTs returns the LastTriggeredTs field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetLastTriggeredTs() int64 {
	if o == nil || o.LastTriggeredTs == nil {
		var ret int64
		return ret
	}
	return *o.LastTriggeredTs
}

// GetLastTriggeredTsOk returns a tuple with the LastTriggeredTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetLastTriggeredTsOk() (int64, bool) {
	if o == nil || o.LastTriggeredTs == nil {
		var ret int64
		return ret, false
	}
	return *o.LastTriggeredTs, true
}

// HasLastTriggeredTs returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasLastTriggeredTs() bool {
	if o != nil && o.LastTriggeredTs != nil {
		return true
	}

	return false
}

// SetLastTriggeredTs gets a reference to the given int64 and assigns it to the LastTriggeredTs field.
func (o *MonitorStateGroup) SetLastTriggeredTs(v int64) {
	o.LastTriggeredTs = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MonitorStateGroup) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MonitorStateGroup) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetStatus() MonitorOverallStates {
	if o == nil || o.Status == nil {
		var ret MonitorOverallStates
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetStatusOk() (MonitorOverallStates, bool) {
	if o == nil || o.Status == nil {
		var ret MonitorOverallStates
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MonitorOverallStates and assigns it to the Status field.
func (o *MonitorStateGroup) SetStatus(v MonitorOverallStates) {
	o.Status = &v
}

// GetTriggeringValue returns the TriggeringValue field value if set, zero value otherwise.
func (o *MonitorStateGroup) GetTriggeringValue() MonitorStateGroupValue {
	if o == nil || o.TriggeringValue == nil {
		var ret MonitorStateGroupValue
		return ret
	}
	return *o.TriggeringValue
}

// GetTriggeringValueOk returns a tuple with the TriggeringValue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroup) GetTriggeringValueOk() (MonitorStateGroupValue, bool) {
	if o == nil || o.TriggeringValue == nil {
		var ret MonitorStateGroupValue
		return ret, false
	}
	return *o.TriggeringValue, true
}

// HasTriggeringValue returns a boolean if a field has been set.
func (o *MonitorStateGroup) HasTriggeringValue() bool {
	if o != nil && o.TriggeringValue != nil {
		return true
	}

	return false
}

// SetTriggeringValue gets a reference to the given MonitorStateGroupValue and assigns it to the TriggeringValue field.
func (o *MonitorStateGroup) SetTriggeringValue(v MonitorStateGroupValue) {
	o.TriggeringValue = &v
}

func (o MonitorStateGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastDataTs != nil {
		toSerialize["last_data_ts"] = o.LastDataTs
	}
	if o.LastNodataTs != nil {
		toSerialize["last_nodata_ts"] = o.LastNodataTs
	}
	if o.LastNotifiedTs != nil {
		toSerialize["last_notified_ts"] = o.LastNotifiedTs
	}
	if o.LastResolvedTs != nil {
		toSerialize["last_resolved_ts"] = o.LastResolvedTs
	}
	if o.LastTriggeredTs != nil {
		toSerialize["last_triggered_ts"] = o.LastTriggeredTs
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TriggeringValue != nil {
		toSerialize["triggering_value"] = o.TriggeringValue
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorStateGroup struct {
	value *MonitorStateGroup
	isSet bool
}

func (v NullableMonitorStateGroup) Get() *MonitorStateGroup {
	return v.value
}

func (v NullableMonitorStateGroup) Set(val *MonitorStateGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorStateGroup) IsSet() bool {
	return v.isSet
}

func (v NullableMonitorStateGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorStateGroup(val *MonitorStateGroup) *NullableMonitorStateGroup {
	return &NullableMonitorStateGroup{value: val, isSet: true}
}

func (v NullableMonitorStateGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorStateGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
