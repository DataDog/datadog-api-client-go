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

// SyntheticsAPIStepSubtype The subtype of the Synthetic multistep API test step, currently only supporting `http`.
type SyntheticsAPIStepSubtype string

// List of SyntheticsAPIStepSubtype
const (
	SYNTHETICSAPISTEPSUBTYPE_HTTP SyntheticsAPIStepSubtype = "http"
)

var allowedSyntheticsAPIStepSubtypeEnumValues = []SyntheticsAPIStepSubtype{
	"http",
}

func (v *SyntheticsAPIStepSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticsAPIStepSubtype(value)
	for _, existing := range allowedSyntheticsAPIStepSubtypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticsAPIStepSubtype", value)
}

// NewSyntheticsAPIStepSubtypeFromValue returns a pointer to a valid SyntheticsAPIStepSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsAPIStepSubtypeFromValue(v string) (*SyntheticsAPIStepSubtype, error) {
	ev := SyntheticsAPIStepSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsAPIStepSubtype: valid values are %v", v, allowedSyntheticsAPIStepSubtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsAPIStepSubtype) IsValid() bool {
	for _, existing := range allowedSyntheticsAPIStepSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsAPIStepSubtype value
func (v SyntheticsAPIStepSubtype) Ptr() *SyntheticsAPIStepSubtype {
	return &v
}

type NullableSyntheticsAPIStepSubtype struct {
	value *SyntheticsAPIStepSubtype
	isSet bool
}

func (v NullableSyntheticsAPIStepSubtype) Get() *SyntheticsAPIStepSubtype {
	return v.value
}

func (v *NullableSyntheticsAPIStepSubtype) Set(val *SyntheticsAPIStepSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAPIStepSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAPIStepSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAPIStepSubtype(val *SyntheticsAPIStepSubtype) *NullableSyntheticsAPIStepSubtype {
	return &NullableSyntheticsAPIStepSubtype{value: val, isSet: true}
}

func (v NullableSyntheticsAPIStepSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAPIStepSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
