/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// SyntheticsAPITestResultFullCheck Object describing the API test configuration.
type SyntheticsAPITestResultFullCheck struct {
	Config SyntheticsTestConfig `json:"config"`
}

// NewSyntheticsAPITestResultFullCheck instantiates a new SyntheticsAPITestResultFullCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsAPITestResultFullCheck(config SyntheticsTestConfig) *SyntheticsAPITestResultFullCheck {
	this := SyntheticsAPITestResultFullCheck{}
	this.Config = config
	return &this
}

// NewSyntheticsAPITestResultFullCheckWithDefaults instantiates a new SyntheticsAPITestResultFullCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsAPITestResultFullCheckWithDefaults() *SyntheticsAPITestResultFullCheck {
	this := SyntheticsAPITestResultFullCheck{}
	return &this
}

// GetConfig returns the Config field value
func (o *SyntheticsAPITestResultFullCheck) GetConfig() SyntheticsTestConfig {
	if o == nil {
		var ret SyntheticsTestConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAPITestResultFullCheck) GetConfigOk() (*SyntheticsTestConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *SyntheticsAPITestResultFullCheck) SetConfig(v SyntheticsTestConfig) {
	o.Config = v
}

func (o SyntheticsAPITestResultFullCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsAPITestResultFullCheck) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Config *SyntheticsTestConfig `json:"config"`
	}{}
	all := struct {
		Config SyntheticsTestConfig `json:"config"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Config == nil {
		return fmt.Errorf("Required field config missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Config = all.Config
	return nil
}

type NullableSyntheticsAPITestResultFullCheck struct {
	value *SyntheticsAPITestResultFullCheck
	isSet bool
}

func (v NullableSyntheticsAPITestResultFullCheck) Get() *SyntheticsAPITestResultFullCheck {
	return v.value
}

func (v *NullableSyntheticsAPITestResultFullCheck) Set(val *SyntheticsAPITestResultFullCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAPITestResultFullCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAPITestResultFullCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAPITestResultFullCheck(val *SyntheticsAPITestResultFullCheck) *NullableSyntheticsAPITestResultFullCheck {
	return &NullableSyntheticsAPITestResultFullCheck{value: val, isSet: true}
}

func (v NullableSyntheticsAPITestResultFullCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAPITestResultFullCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
