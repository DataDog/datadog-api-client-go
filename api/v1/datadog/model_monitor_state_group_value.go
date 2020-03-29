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

// MonitorStateGroupValue TODO.
type MonitorStateGroupValue struct {
	// TODO.
	FromTs *int64 `json:"from_ts,omitempty"`
	// TODO.
	Left *float64 `json:"left,omitempty"`
	// TODO.
	Right *float64 `json:"right,omitempty"`
	// TODO.
	ToTs *int64 `json:"to_ts,omitempty"`
	// TODO.
	Value *float64 `json:"value,omitempty"`
}

// NewMonitorStateGroupValue instantiates a new MonitorStateGroupValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorStateGroupValue() *MonitorStateGroupValue {
	this := MonitorStateGroupValue{}
	return &this
}

// NewMonitorStateGroupValueWithDefaults instantiates a new MonitorStateGroupValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorStateGroupValueWithDefaults() *MonitorStateGroupValue {
	this := MonitorStateGroupValue{}
	return &this
}

// GetFromTs returns the FromTs field value if set, zero value otherwise.
func (o *MonitorStateGroupValue) GetFromTs() int64 {
	if o == nil || o.FromTs == nil {
		var ret int64
		return ret
	}
	return *o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroupValue) GetFromTsOk() (int64, bool) {
	if o == nil || o.FromTs == nil {
		var ret int64
		return ret, false
	}
	return *o.FromTs, true
}

// HasFromTs returns a boolean if a field has been set.
func (o *MonitorStateGroupValue) HasFromTs() bool {
	if o != nil && o.FromTs != nil {
		return true
	}

	return false
}

// SetFromTs gets a reference to the given int64 and assigns it to the FromTs field.
func (o *MonitorStateGroupValue) SetFromTs(v int64) {
	o.FromTs = &v
}

// GetLeft returns the Left field value if set, zero value otherwise.
func (o *MonitorStateGroupValue) GetLeft() float64 {
	if o == nil || o.Left == nil {
		var ret float64
		return ret
	}
	return *o.Left
}

// GetLeftOk returns a tuple with the Left field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroupValue) GetLeftOk() (float64, bool) {
	if o == nil || o.Left == nil {
		var ret float64
		return ret, false
	}
	return *o.Left, true
}

// HasLeft returns a boolean if a field has been set.
func (o *MonitorStateGroupValue) HasLeft() bool {
	if o != nil && o.Left != nil {
		return true
	}

	return false
}

// SetLeft gets a reference to the given float64 and assigns it to the Left field.
func (o *MonitorStateGroupValue) SetLeft(v float64) {
	o.Left = &v
}

// GetRight returns the Right field value if set, zero value otherwise.
func (o *MonitorStateGroupValue) GetRight() float64 {
	if o == nil || o.Right == nil {
		var ret float64
		return ret
	}
	return *o.Right
}

// GetRightOk returns a tuple with the Right field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroupValue) GetRightOk() (float64, bool) {
	if o == nil || o.Right == nil {
		var ret float64
		return ret, false
	}
	return *o.Right, true
}

// HasRight returns a boolean if a field has been set.
func (o *MonitorStateGroupValue) HasRight() bool {
	if o != nil && o.Right != nil {
		return true
	}

	return false
}

// SetRight gets a reference to the given float64 and assigns it to the Right field.
func (o *MonitorStateGroupValue) SetRight(v float64) {
	o.Right = &v
}

// GetToTs returns the ToTs field value if set, zero value otherwise.
func (o *MonitorStateGroupValue) GetToTs() int64 {
	if o == nil || o.ToTs == nil {
		var ret int64
		return ret
	}
	return *o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroupValue) GetToTsOk() (int64, bool) {
	if o == nil || o.ToTs == nil {
		var ret int64
		return ret, false
	}
	return *o.ToTs, true
}

// HasToTs returns a boolean if a field has been set.
func (o *MonitorStateGroupValue) HasToTs() bool {
	if o != nil && o.ToTs != nil {
		return true
	}

	return false
}

// SetToTs gets a reference to the given int64 and assigns it to the ToTs field.
func (o *MonitorStateGroupValue) SetToTs(v int64) {
	o.ToTs = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MonitorStateGroupValue) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MonitorStateGroupValue) GetValueOk() (float64, bool) {
	if o == nil || o.Value == nil {
		var ret float64
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MonitorStateGroupValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *MonitorStateGroupValue) SetValue(v float64) {
	o.Value = &v
}

func (o MonitorStateGroupValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromTs != nil {
		toSerialize["from_ts"] = o.FromTs
	}
	if o.Left != nil {
		toSerialize["left"] = o.Left
	}
	if o.Right != nil {
		toSerialize["right"] = o.Right
	}
	if o.ToTs != nil {
		toSerialize["to_ts"] = o.ToTs
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorStateGroupValue struct {
	value *MonitorStateGroupValue
	isSet bool
}

func (v NullableMonitorStateGroupValue) Get() *MonitorStateGroupValue {
	return v.value
}

func (v NullableMonitorStateGroupValue) Set(val *MonitorStateGroupValue) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorStateGroupValue) IsSet() bool {
	return v.isSet
}

func (v NullableMonitorStateGroupValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorStateGroupValue(val *MonitorStateGroupValue) *NullableMonitorStateGroupValue {
	return &NullableMonitorStateGroupValue{value: val, isSet: true}
}

func (v NullableMonitorStateGroupValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorStateGroupValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
