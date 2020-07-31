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

// SyntheticsTestOptionsMonitorOptions Object containing the options for a Synthetic test as a monitor (e.g. renotification).
type SyntheticsTestOptionsMonitorOptions struct {
	// Time interval before renotifying if the test is still failing (in minutes).
	RenotifyInterval *int64 `json:"renotify_interval,omitempty"`
}

// NewSyntheticsTestOptionsMonitorOptions instantiates a new SyntheticsTestOptionsMonitorOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTestOptionsMonitorOptions() *SyntheticsTestOptionsMonitorOptions {
	this := SyntheticsTestOptionsMonitorOptions{}
	return &this
}

// NewSyntheticsTestOptionsMonitorOptionsWithDefaults instantiates a new SyntheticsTestOptionsMonitorOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTestOptionsMonitorOptionsWithDefaults() *SyntheticsTestOptionsMonitorOptions {
	this := SyntheticsTestOptionsMonitorOptions{}
	return &this
}

// GetRenotifyInterval returns the RenotifyInterval field value if set, zero value otherwise.
func (o *SyntheticsTestOptionsMonitorOptions) GetRenotifyInterval() int64 {
	if o == nil || o.RenotifyInterval == nil {
		var ret int64
		return ret
	}
	return *o.RenotifyInterval
}

// GetRenotifyIntervalOk returns a tuple with the RenotifyInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestOptionsMonitorOptions) GetRenotifyIntervalOk() (*int64, bool) {
	if o == nil || o.RenotifyInterval == nil {
		return nil, false
	}
	return o.RenotifyInterval, true
}

// HasRenotifyInterval returns a boolean if a field has been set.
func (o *SyntheticsTestOptionsMonitorOptions) HasRenotifyInterval() bool {
	if o != nil && o.RenotifyInterval != nil {
		return true
	}

	return false
}

// SetRenotifyInterval gets a reference to the given int64 and assigns it to the RenotifyInterval field.
func (o *SyntheticsTestOptionsMonitorOptions) SetRenotifyInterval(v int64) {
	o.RenotifyInterval = &v
}

func (o SyntheticsTestOptionsMonitorOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RenotifyInterval != nil {
		toSerialize["renotify_interval"] = o.RenotifyInterval
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsTestOptionsMonitorOptions struct {
	value *SyntheticsTestOptionsMonitorOptions
	isSet bool
}

func (v NullableSyntheticsTestOptionsMonitorOptions) Get() *SyntheticsTestOptionsMonitorOptions {
	return v.value
}

func (v *NullableSyntheticsTestOptionsMonitorOptions) Set(val *SyntheticsTestOptionsMonitorOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestOptionsMonitorOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsTestOptionsMonitorOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestOptionsMonitorOptions(val *SyntheticsTestOptionsMonitorOptions) *NullableSyntheticsTestOptionsMonitorOptions {
	return &NullableSyntheticsTestOptionsMonitorOptions{value: val, isSet: true}
}

func (v NullableSyntheticsTestOptionsMonitorOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestOptionsMonitorOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
