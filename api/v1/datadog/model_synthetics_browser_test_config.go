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

// SyntheticsBrowserTestConfig Configuration object for a Synthetic browser test.
type SyntheticsBrowserTestConfig struct {
	// Array of assertions used for the test.
	Assertions []SyntheticsAssertion `json:"assertions"`
	Request    SyntheticsTestRequest `json:"request"`
	// Array of variables used for the test steps.
	Variables *[]SyntheticsBrowserVariable `json:"variables,omitempty"`
}

// NewSyntheticsBrowserTestConfig instantiates a new SyntheticsBrowserTestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsBrowserTestConfig(assertions []SyntheticsAssertion, request SyntheticsTestRequest) *SyntheticsBrowserTestConfig {
	this := SyntheticsBrowserTestConfig{}
	this.Assertions = assertions
	this.Request = request
	return &this
}

// NewSyntheticsBrowserTestConfigWithDefaults instantiates a new SyntheticsBrowserTestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsBrowserTestConfigWithDefaults() *SyntheticsBrowserTestConfig {
	this := SyntheticsBrowserTestConfig{}
	return &this
}

// GetAssertions returns the Assertions field value
func (o *SyntheticsBrowserTestConfig) GetAssertions() []SyntheticsAssertion {
	if o == nil {
		var ret []SyntheticsAssertion
		return ret
	}

	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestConfig) GetAssertionsOk() (*[]SyntheticsAssertion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assertions, true
}

// SetAssertions sets field value
func (o *SyntheticsBrowserTestConfig) SetAssertions(v []SyntheticsAssertion) {
	o.Assertions = v
}

// GetRequest returns the Request field value
func (o *SyntheticsBrowserTestConfig) GetRequest() SyntheticsTestRequest {
	if o == nil {
		var ret SyntheticsTestRequest
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestConfig) GetRequestOk() (*SyntheticsTestRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *SyntheticsBrowserTestConfig) SetRequest(v SyntheticsTestRequest) {
	o.Request = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *SyntheticsBrowserTestConfig) GetVariables() []SyntheticsBrowserVariable {
	if o == nil || o.Variables == nil {
		var ret []SyntheticsBrowserVariable
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsBrowserTestConfig) GetVariablesOk() (*[]SyntheticsBrowserVariable, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *SyntheticsBrowserTestConfig) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []SyntheticsBrowserVariable and assigns it to the Variables field.
func (o *SyntheticsBrowserTestConfig) SetVariables(v []SyntheticsBrowserVariable) {
	o.Variables = &v
}

func (o SyntheticsBrowserTestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assertions"] = o.Assertions
	}
	if true {
		toSerialize["request"] = o.Request
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsBrowserTestConfig struct {
	value *SyntheticsBrowserTestConfig
	isSet bool
}

func (v NullableSyntheticsBrowserTestConfig) Get() *SyntheticsBrowserTestConfig {
	return v.value
}

func (v *NullableSyntheticsBrowserTestConfig) Set(val *SyntheticsBrowserTestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBrowserTestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBrowserTestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBrowserTestConfig(val *SyntheticsBrowserTestConfig) *NullableSyntheticsBrowserTestConfig {
	return &NullableSyntheticsBrowserTestConfig{value: val, isSet: true}
}

func (v NullableSyntheticsBrowserTestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBrowserTestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
