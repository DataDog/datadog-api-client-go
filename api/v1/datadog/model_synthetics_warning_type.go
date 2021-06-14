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

// SyntheticsWarningType User locator used.
type SyntheticsWarningType string

// List of SyntheticsWarningType
const (
	SYNTHETICSWARNINGTYPE_USER_LOCATOR SyntheticsWarningType = "user_locator"
)

var allowedSyntheticsWarningTypeEnumValues = []SyntheticsWarningType{
	"user_locator",
}

func (w *SyntheticsWarningType) GetAllowedValues() []SyntheticsWarningType {
	return allowedSyntheticsWarningTypeEnumValues
}

func (v *SyntheticsWarningType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticsWarningType(value)
	for _, existing := range allowedSyntheticsWarningTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticsWarningType", value)
}

// NewSyntheticsWarningTypeFromValue returns a pointer to a valid SyntheticsWarningType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticsWarningTypeFromValue(v string) (*SyntheticsWarningType, error) {
	ev := SyntheticsWarningType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticsWarningType: valid values are %v", v, allowedSyntheticsWarningTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticsWarningType) IsValid() bool {
	for _, existing := range allowedSyntheticsWarningTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsWarningType value
func (v SyntheticsWarningType) Ptr() *SyntheticsWarningType {
	return &v
}

type NullableSyntheticsWarningType struct {
	value *SyntheticsWarningType
	isSet bool
}

func (v NullableSyntheticsWarningType) Get() *SyntheticsWarningType {
	return v.value
}

func (v *NullableSyntheticsWarningType) Set(val *SyntheticsWarningType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsWarningType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsWarningType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsWarningType(val *SyntheticsWarningType) *NullableSyntheticsWarningType {
	return &NullableSyntheticsWarningType{value: val, isSet: true}
}

func (v NullableSyntheticsWarningType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsWarningType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
