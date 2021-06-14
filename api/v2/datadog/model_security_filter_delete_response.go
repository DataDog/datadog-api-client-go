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

// SecurityFilterDeleteResponse Response object which only includes the metadata.
type SecurityFilterDeleteResponse struct {
	Meta *SecurityFilterMeta `json:"meta,omitempty"`
}

// NewSecurityFilterDeleteResponse instantiates a new SecurityFilterDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityFilterDeleteResponse() *SecurityFilterDeleteResponse {
	this := SecurityFilterDeleteResponse{}
	return &this
}

// NewSecurityFilterDeleteResponseWithDefaults instantiates a new SecurityFilterDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityFilterDeleteResponseWithDefaults() *SecurityFilterDeleteResponse {
	this := SecurityFilterDeleteResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SecurityFilterDeleteResponse) GetMeta() SecurityFilterMeta {
	if o == nil || o.Meta == nil {
		var ret SecurityFilterMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFilterDeleteResponse) GetMetaOk() (*SecurityFilterMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SecurityFilterDeleteResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given SecurityFilterMeta and assigns it to the Meta field.
func (o *SecurityFilterDeleteResponse) SetMeta(v SecurityFilterMeta) {
	o.Meta = &v
}

func (o SecurityFilterDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityFilterDeleteResponse struct {
	value *SecurityFilterDeleteResponse
	isSet bool
}

func (v NullableSecurityFilterDeleteResponse) Get() *SecurityFilterDeleteResponse {
	return v.value
}

func (v *NullableSecurityFilterDeleteResponse) Set(val *SecurityFilterDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityFilterDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityFilterDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityFilterDeleteResponse(val *SecurityFilterDeleteResponse) *NullableSecurityFilterDeleteResponse {
	return &NullableSecurityFilterDeleteResponse{value: val, isSet: true}
}

func (v NullableSecurityFilterDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityFilterDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
