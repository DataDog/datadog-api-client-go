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

// WidgetGrouping The kind of grouping to use.
type WidgetGrouping string

// List of WidgetGrouping
const (
	WIDGETGROUPING_CHECK WidgetGrouping = "check"
	WIDGETGROUPING_CLUSTER WidgetGrouping = "cluster"
)

func (v *WidgetGrouping) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetGrouping(value)
	for _, existing := range []WidgetGrouping{ "check", "cluster",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetGrouping", *v)
}

// Ptr returns reference to WidgetGrouping value
func (v WidgetGrouping) Ptr() *WidgetGrouping {
	return &v
}

type NullableWidgetGrouping struct {
	value *WidgetGrouping
	isSet bool
}

func (v NullableWidgetGrouping) Get() *WidgetGrouping {
	return v.value
}

func (v *NullableWidgetGrouping) Set(val *WidgetGrouping) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetGrouping) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetGrouping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetGrouping(val *WidgetGrouping) *NullableWidgetGrouping {
	return &NullableWidgetGrouping{value: val, isSet: true}
}

func (v NullableWidgetGrouping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetGrouping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

