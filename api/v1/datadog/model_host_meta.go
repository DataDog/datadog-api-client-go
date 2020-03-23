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

// HostMeta TODO.
type HostMeta struct {
	NixV *[]string `json:"nixV,omitempty"`
}

// NewHostMeta instantiates a new HostMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostMeta() *HostMeta {
	this := HostMeta{}
	return &this
}

// NewHostMetaWithDefaults instantiates a new HostMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostMetaWithDefaults() *HostMeta {
	this := HostMeta{}
	return &this
}

// GetNixV returns the NixV field value if set, zero value otherwise.
func (o *HostMeta) GetNixV() []string {
	if o == nil || o.NixV == nil {
		var ret []string
		return ret
	}
	return *o.NixV
}

// GetNixVOk returns a tuple with the NixV field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HostMeta) GetNixVOk() ([]string, bool) {
	if o == nil || o.NixV == nil {
		var ret []string
		return ret, false
	}
	return *o.NixV, true
}

// HasNixV returns a boolean if a field has been set.
func (o *HostMeta) HasNixV() bool {
	if o != nil && o.NixV != nil {
		return true
	}

	return false
}

// SetNixV gets a reference to the given []string and assigns it to the NixV field.
func (o *HostMeta) SetNixV(v []string) {
	o.NixV = &v
}

func (o HostMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NixV != nil {
		toSerialize["nixV"] = o.NixV
	}
	return json.Marshal(toSerialize)
}

type NullableHostMeta struct {
	value *HostMeta
	isSet bool
}

func (v NullableHostMeta) Get() *HostMeta {
	return v.value
}

func (v NullableHostMeta) Set(val *HostMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableHostMeta) IsSet() bool {
	return v.isSet
}

func (v NullableHostMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostMeta(val *HostMeta) *NullableHostMeta {
	return &NullableHostMeta{value: val, isSet: true}
}

func (v NullableHostMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
