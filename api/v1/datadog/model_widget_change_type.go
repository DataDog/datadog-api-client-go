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

import (
	"fmt"
)

// WidgetChangeType Show the absolute or the relative change.
type WidgetChangeType string

// List of WidgetChangeType
const (
	WIDGETCHANGETYPE_ABSOLUTE WidgetChangeType = "absolute"
	WIDGETCHANGETYPE_RELATIVE WidgetChangeType = "relative"
)

func (v *WidgetChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetChangeType(value)
	for _, existing := range []WidgetChangeType{"absolute", "relative"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetChangeType", *v)
}

// Ptr returns reference to WidgetChangeType value
func (v WidgetChangeType) Ptr() *WidgetChangeType {
	return &v
}

type NullableWidgetChangeType struct {
	value *WidgetChangeType
	isSet bool
}

func (v NullableWidgetChangeType) Get() *WidgetChangeType {
	return v.value
}

func (v *NullableWidgetChangeType) Set(val *WidgetChangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetChangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetChangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetChangeType(val *WidgetChangeType) *NullableWidgetChangeType {
	return &NullableWidgetChangeType{value: val, isSet: true}
}

func (v NullableWidgetChangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetChangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
