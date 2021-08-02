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

// SecurityMonitoringRuleNewValueOptions Options on new value rules.
type SecurityMonitoringRuleNewValueOptions struct {
	ForgetAfter      *SecurityMonitoringRuleNewValueOptionsForgetAfter      `json:"forgetAfter,omitempty"`
	LearningDuration *SecurityMonitoringRuleNewValueOptionsLearningDuration `json:"learningDuration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
}

// NewSecurityMonitoringRuleNewValueOptions instantiates a new SecurityMonitoringRuleNewValueOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityMonitoringRuleNewValueOptions() *SecurityMonitoringRuleNewValueOptions {
	this := SecurityMonitoringRuleNewValueOptions{}
	return &this
}

// NewSecurityMonitoringRuleNewValueOptionsWithDefaults instantiates a new SecurityMonitoringRuleNewValueOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityMonitoringRuleNewValueOptionsWithDefaults() *SecurityMonitoringRuleNewValueOptions {
	this := SecurityMonitoringRuleNewValueOptions{}
	return &this
}

// GetForgetAfter returns the ForgetAfter field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleNewValueOptions) GetForgetAfter() SecurityMonitoringRuleNewValueOptionsForgetAfter {
	if o == nil || o.ForgetAfter == nil {
		var ret SecurityMonitoringRuleNewValueOptionsForgetAfter
		return ret
	}
	return *o.ForgetAfter
}

// GetForgetAfterOk returns a tuple with the ForgetAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleNewValueOptions) GetForgetAfterOk() (*SecurityMonitoringRuleNewValueOptionsForgetAfter, bool) {
	if o == nil || o.ForgetAfter == nil {
		return nil, false
	}
	return o.ForgetAfter, true
}

// HasForgetAfter returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleNewValueOptions) HasForgetAfter() bool {
	if o != nil && o.ForgetAfter != nil {
		return true
	}

	return false
}

// SetForgetAfter gets a reference to the given SecurityMonitoringRuleNewValueOptionsForgetAfter and assigns it to the ForgetAfter field.
func (o *SecurityMonitoringRuleNewValueOptions) SetForgetAfter(v SecurityMonitoringRuleNewValueOptionsForgetAfter) {
	o.ForgetAfter = &v
}

// GetLearningDuration returns the LearningDuration field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleNewValueOptions) GetLearningDuration() SecurityMonitoringRuleNewValueOptionsLearningDuration {
	if o == nil || o.LearningDuration == nil {
		var ret SecurityMonitoringRuleNewValueOptionsLearningDuration
		return ret
	}
	return *o.LearningDuration
}

// GetLearningDurationOk returns a tuple with the LearningDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleNewValueOptions) GetLearningDurationOk() (*SecurityMonitoringRuleNewValueOptionsLearningDuration, bool) {
	if o == nil || o.LearningDuration == nil {
		return nil, false
	}
	return o.LearningDuration, true
}

// HasLearningDuration returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleNewValueOptions) HasLearningDuration() bool {
	if o != nil && o.LearningDuration != nil {
		return true
	}

	return false
}

// SetLearningDuration gets a reference to the given SecurityMonitoringRuleNewValueOptionsLearningDuration and assigns it to the LearningDuration field.
func (o *SecurityMonitoringRuleNewValueOptions) SetLearningDuration(v SecurityMonitoringRuleNewValueOptionsLearningDuration) {
	o.LearningDuration = &v
}

func (o SecurityMonitoringRuleNewValueOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ForgetAfter != nil {
		toSerialize["forgetAfter"] = o.ForgetAfter
	}
	if o.LearningDuration != nil {
		toSerialize["learningDuration"] = o.LearningDuration
	}
	return json.Marshal(toSerialize)
}

func (o *SecurityMonitoringRuleNewValueOptions) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		ForgetAfter      *SecurityMonitoringRuleNewValueOptionsForgetAfter      `json:"forgetAfter,omitempty"`
		LearningDuration *SecurityMonitoringRuleNewValueOptionsLearningDuration `json:"learningDuration,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.ForgetAfter; v != nil && !v.IsValid() {
		errr := json.Unmarshal(bytes, &raw)
		if errr != nil {
			return errr
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.LearningDuration; v != nil && !v.IsValid() {
		errr := json.Unmarshal(bytes, &raw)
		if errr != nil {
			return errr
		}
		o.UnparsedObject = raw
		return nil
	}
	o.ForgetAfter = all.ForgetAfter
	o.LearningDuration = all.LearningDuration
	return nil
}

type NullableSecurityMonitoringRuleNewValueOptions struct {
	value *SecurityMonitoringRuleNewValueOptions
	isSet bool
}

func (v NullableSecurityMonitoringRuleNewValueOptions) Get() *SecurityMonitoringRuleNewValueOptions {
	return v.value
}

func (v *NullableSecurityMonitoringRuleNewValueOptions) Set(val *SecurityMonitoringRuleNewValueOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMonitoringRuleNewValueOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMonitoringRuleNewValueOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMonitoringRuleNewValueOptions(val *SecurityMonitoringRuleNewValueOptions) *NullableSecurityMonitoringRuleNewValueOptions {
	return &NullableSecurityMonitoringRuleNewValueOptions{value: val, isSet: true}
}

func (v NullableSecurityMonitoringRuleNewValueOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMonitoringRuleNewValueOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
