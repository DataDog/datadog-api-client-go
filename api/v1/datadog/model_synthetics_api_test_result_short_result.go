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

// SyntheticsAPITestResultShortResult Result of the last API test run.
type SyntheticsAPITestResultShortResult struct {
	Timings *SyntheticsTiming `json:"timings,omitempty"`
}

// NewSyntheticsAPITestResultShortResult instantiates a new SyntheticsAPITestResultShortResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsAPITestResultShortResult() *SyntheticsAPITestResultShortResult {
	this := SyntheticsAPITestResultShortResult{}
	return &this
}

// NewSyntheticsAPITestResultShortResultWithDefaults instantiates a new SyntheticsAPITestResultShortResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsAPITestResultShortResultWithDefaults() *SyntheticsAPITestResultShortResult {
	this := SyntheticsAPITestResultShortResult{}
	return &this
}

// GetTimings returns the Timings field value if set, zero value otherwise.
func (o *SyntheticsAPITestResultShortResult) GetTimings() SyntheticsTiming {
	if o == nil || o.Timings == nil {
		var ret SyntheticsTiming
		return ret
	}
	return *o.Timings
}

// GetTimingsOk returns a tuple with the Timings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultShortResult) GetTimingsOk() (*SyntheticsTiming, bool) {
	if o == nil || o.Timings == nil {
		return nil, false
	}
	return o.Timings, true
}

// HasTimings returns a boolean if a field has been set.
func (o *SyntheticsAPITestResultShortResult) HasTimings() bool {
	if o != nil && o.Timings != nil {
		return true
	}

	return false
}

// SetTimings gets a reference to the given SyntheticsTiming and assigns it to the Timings field.
func (o *SyntheticsAPITestResultShortResult) SetTimings(v SyntheticsTiming) {
	o.Timings = &v
}

func (o SyntheticsAPITestResultShortResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timings != nil {
		toSerialize["timings"] = o.Timings
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsAPITestResultShortResult struct {
	value *SyntheticsAPITestResultShortResult
	isSet bool
}

func (v NullableSyntheticsAPITestResultShortResult) Get() *SyntheticsAPITestResultShortResult {
	return v.value
}

func (v *NullableSyntheticsAPITestResultShortResult) Set(val *SyntheticsAPITestResultShortResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAPITestResultShortResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAPITestResultShortResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAPITestResultShortResult(val *SyntheticsAPITestResultShortResult) *NullableSyntheticsAPITestResultShortResult {
	return &NullableSyntheticsAPITestResultShortResult{value: val, isSet: true}
}

func (v NullableSyntheticsAPITestResultShortResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAPITestResultShortResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


