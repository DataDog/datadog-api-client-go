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

// WidgetLineType Type of lines displayed.
type WidgetLineType string

// List of WidgetLineType
const (
	WIDGETLINETYPE_DASHED WidgetLineType = "dashed"
	WIDGETLINETYPE_DOTTED WidgetLineType = "dotted"
	WIDGETLINETYPE_SOLID WidgetLineType = "solid"
)

func (v *WidgetLineType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetLineType(value)
	for _, existing := range []WidgetLineType{ "dashed", "dotted", "solid",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetLineType", *v)
}

// Ptr returns reference to WidgetLineType value
func (v WidgetLineType) Ptr() *WidgetLineType {
	return &v
}

type NullableWidgetLineType struct {
	value *WidgetLineType
	isSet bool
}

func (v NullableWidgetLineType) Get() *WidgetLineType {
	return v.value
}

func (v *NullableWidgetLineType) Set(val *WidgetLineType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetLineType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetLineType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetLineType(val *WidgetLineType) *NullableWidgetLineType {
	return &NullableWidgetLineType{value: val, isSet: true}
}

func (v NullableWidgetLineType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetLineType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

