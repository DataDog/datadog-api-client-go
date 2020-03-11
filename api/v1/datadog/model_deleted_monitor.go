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

// DeletedMonitor Response from the DeleteMonitor call
type DeletedMonitor struct {
	// ID of the deleted monitor
	DeletedMonitorId *int64 `json:"deleted_monitor_id,omitempty"`
}

// NewDeletedMonitor instantiates a new DeletedMonitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletedMonitor() *DeletedMonitor {
	this := DeletedMonitor{}
	return &this
}

// NewDeletedMonitorWithDefaults instantiates a new DeletedMonitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedMonitorWithDefaults() *DeletedMonitor {
	this := DeletedMonitor{}
	return &this
}

// GetDeletedMonitorId returns the DeletedMonitorId field value if set, zero value otherwise.
func (o *DeletedMonitor) GetDeletedMonitorId() int64 {
	if o == nil || o.DeletedMonitorId == nil {
		var ret int64
		return ret
	}
	return *o.DeletedMonitorId
}

// GetDeletedMonitorIdOk returns a tuple with the DeletedMonitorId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeletedMonitor) GetDeletedMonitorIdOk() (int64, bool) {
	if o == nil || o.DeletedMonitorId == nil {
		var ret int64
		return ret, false
	}
	return *o.DeletedMonitorId, true
}

// HasDeletedMonitorId returns a boolean if a field has been set.
func (o *DeletedMonitor) HasDeletedMonitorId() bool {
	if o != nil && o.DeletedMonitorId != nil {
		return true
	}

	return false
}

// SetDeletedMonitorId gets a reference to the given int64 and assigns it to the DeletedMonitorId field.
func (o *DeletedMonitor) SetDeletedMonitorId(v int64) {
	o.DeletedMonitorId = &v
}

func (o DeletedMonitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeletedMonitorId != nil {
		toSerialize["deleted_monitor_id"] = o.DeletedMonitorId
	}
	return json.Marshal(toSerialize)
}

type NullableDeletedMonitor struct {
	value *DeletedMonitor
	isSet bool
}

func (v NullableDeletedMonitor) Get() *DeletedMonitor {
	return v.value
}

func (v NullableDeletedMonitor) Set(val *DeletedMonitor) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletedMonitor) IsSet() bool {
	return v.isSet
}

func (v NullableDeletedMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletedMonitor(val *DeletedMonitor) *NullableDeletedMonitor {
	return &NullableDeletedMonitor{value: val, isSet: true}
}

func (v NullableDeletedMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletedMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
