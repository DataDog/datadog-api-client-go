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

// HostMuteSettings Combination of settings to mute a host.
type HostMuteSettings struct {
	// POSIX timestamp in seconds when the host is unmuted. If omitted, the host remains muted until explicitly unmuted.
	End *int64 `json:"end,omitempty"`
	// Message to associate with the muting of this host.
	Message *string `json:"message,omitempty"`
	// If true and the host is already muted, replaces existing host mute settings.
	Override *bool `json:"override,omitempty"`
}

// NewHostMuteSettings instantiates a new HostMuteSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMuteSettings() *HostMuteSettings {
	this := HostMuteSettings{}
	return &this
}

// NewHostMuteSettingsWithDefaults instantiates a new HostMuteSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMuteSettingsWithDefaults() *HostMuteSettings {
	this := HostMuteSettings{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *HostMuteSettings) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMuteSettings) GetEndOk() (int64, bool) {
	if o == nil || o.End == nil {
		var ret int64
		return ret, false
	}
	return *o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *HostMuteSettings) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *HostMuteSettings) SetEnd(v int64) {
	o.End = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HostMuteSettings) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMuteSettings) GetMessageOk() (string, bool) {
	if o == nil || o.Message == nil {
		var ret string
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HostMuteSettings) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HostMuteSettings) SetMessage(v string) {
	o.Message = &v
}

// GetOverride returns the Override field value if set, zero value otherwise.
func (o *HostMuteSettings) GetOverride() bool {
	if o == nil || o.Override == nil {
		var ret bool
		return ret
	}
	return *o.Override
}

// GetOverrideOk returns a tuple with the Override field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMuteSettings) GetOverrideOk() (bool, bool) {
	if o == nil || o.Override == nil {
		var ret bool
		return ret, false
	}
	return *o.Override, true
}

// HasOverride returns a boolean if a field has been set.
func (o *HostMuteSettings) HasOverride() bool {
	if o != nil && o.Override != nil {
		return true
	}

	return false
}

// SetOverride gets a reference to the given bool and assigns it to the Override field.
func (o *HostMuteSettings) SetOverride(v bool) {
	o.Override = &v
}

func (o HostMuteSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Override != nil {
		toSerialize["override"] = o.Override
	}
	return json.Marshal(toSerialize)
}

type NullableHostMuteSettings struct {
	value *HostMuteSettings
	isSet bool
}

func (v NullableHostMuteSettings) Get() *HostMuteSettings {
	return v.value
}

func (v NullableHostMuteSettings) Set(val *HostMuteSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMuteSettings) IsSet() bool {
	return v.isSet
}

func (v NullableHostMuteSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMuteSettings(val *HostMuteSettings) *NullableHostMuteSettings {
	return &NullableHostMuteSettings{value: val, isSet: true}
}

func (v NullableHostMuteSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMuteSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
