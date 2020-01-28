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

// SyntheticsTestConfig struct for SyntheticsTestConfig
type SyntheticsTestConfig struct {
	Assertions []SyntheticsAssertion        `json:"assertions"`
	Request    SyntheticsTestRequest        `json:"request"`
	Variables  *[]SyntheticsBrowserVariable `json:"variables,omitempty"`
}

// GetAssertions returns the Assertions field value
func (o *SyntheticsTestConfig) GetAssertions() []SyntheticsAssertion {
	if o == nil {
		var ret []SyntheticsAssertion
		return ret
	}

	return o.Assertions
}

// SetAssertions sets field value
func (o *SyntheticsTestConfig) SetAssertions(v []SyntheticsAssertion) {
	o.Assertions = v
}

// GetRequest returns the Request field value
func (o *SyntheticsTestConfig) GetRequest() SyntheticsTestRequest {
	if o == nil {
		var ret SyntheticsTestRequest
		return ret
	}

	return o.Request
}

// SetRequest sets field value
func (o *SyntheticsTestConfig) SetRequest(v SyntheticsTestRequest) {
	o.Request = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *SyntheticsTestConfig) GetVariables() []SyntheticsBrowserVariable {
	if o == nil || o.Variables == nil {
		var ret []SyntheticsBrowserVariable
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestConfig) GetVariablesOk() ([]SyntheticsBrowserVariable, bool) {
	if o == nil || o.Variables == nil {
		var ret []SyntheticsBrowserVariable
		return ret, false
	}
	return *o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *SyntheticsTestConfig) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []SyntheticsBrowserVariable and assigns it to the Variables field.
func (o *SyntheticsTestConfig) SetVariables(v []SyntheticsBrowserVariable) {
	o.Variables = &v
}

type NullableSyntheticsTestConfig struct {
	Value        SyntheticsTestConfig
	ExplicitNull bool
}

func (v NullableSyntheticsTestConfig) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsTestConfig) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
