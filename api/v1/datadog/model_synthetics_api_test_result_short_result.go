/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// SyntheticsApiTestResultShortResult struct for SyntheticsApiTestResultShortResult
type SyntheticsApiTestResultShortResult struct {
	Timings *SyntheticsTiming `json:"timings,omitempty"`
}

// GetTimings returns the Timings field value if set, zero value otherwise.
func (o *SyntheticsApiTestResultShortResult) GetTimings() SyntheticsTiming {
	if o == nil || o.Timings == nil {
		var ret SyntheticsTiming
		return ret
	}
	return *o.Timings
}

// GetTimingsOk returns a tuple with the Timings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsApiTestResultShortResult) GetTimingsOk() (SyntheticsTiming, bool) {
	if o == nil || o.Timings == nil {
		var ret SyntheticsTiming
		return ret, false
	}
	return *o.Timings, true
}

// HasTimings returns a boolean if a field has been set.
func (o *SyntheticsApiTestResultShortResult) HasTimings() bool {
	if o != nil && o.Timings != nil {
		return true
	}

	return false
}

// SetTimings gets a reference to the given SyntheticsTiming and assigns it to the Timings field.
func (o *SyntheticsApiTestResultShortResult) SetTimings(v SyntheticsTiming) {
	o.Timings = &v
}

type NullableSyntheticsApiTestResultShortResult struct {
	Value        SyntheticsApiTestResultShortResult
	ExplicitNull bool
}

func (v NullableSyntheticsApiTestResultShortResult) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsApiTestResultShortResult) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
