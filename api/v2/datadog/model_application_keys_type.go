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

// ApplicationKeysType Application Keys resource type.
type ApplicationKeysType string

// List of ApplicationKeysType
const (
	APPLICATIONKEYSTYPE_APPLICATION_KEYS ApplicationKeysType = "application_keys"
)

var allowedApplicationKeysTypeEnumValues = []ApplicationKeysType{
	"application_keys",
}

func (v *ApplicationKeysType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplicationKeysType(value)
	for _, existing := range allowedApplicationKeysTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplicationKeysType", value)
}

// NewApplicationKeysTypeFromValue returns a pointer to a valid ApplicationKeysType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplicationKeysTypeFromValue(v string) (*ApplicationKeysType, error) {
	ev := ApplicationKeysType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplicationKeysType: valid values are %v", v, allowedApplicationKeysTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplicationKeysType) IsValid() bool {
	for _, existing := range allowedApplicationKeysTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApplicationKeysType value
func (v ApplicationKeysType) Ptr() *ApplicationKeysType {
	return &v
}

type NullableApplicationKeysType struct {
	value *ApplicationKeysType
	isSet bool
}

func (v NullableApplicationKeysType) Get() *ApplicationKeysType {
	return v.value
}

func (v *NullableApplicationKeysType) Set(val *ApplicationKeysType) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationKeysType) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationKeysType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationKeysType(val *ApplicationKeysType) *NullableApplicationKeysType {
	return &NullableApplicationKeysType{value: val, isSet: true}
}

func (v NullableApplicationKeysType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationKeysType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
