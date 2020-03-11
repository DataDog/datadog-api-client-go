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

// WidgetSort the model 'WidgetSort'
type WidgetSort string

// List of WidgetSort
const (
	WIDGETSORT_ASCENDING  WidgetSort = "asc"
	WIDGETSORT_DESCENDING WidgetSort = "desc"
)

// Ptr returns reference to WidgetSort value
func (v WidgetSort) Ptr() *WidgetSort {
	return &v
}

type NullableWidgetSort struct {
	value *WidgetSort
	isSet bool
}

func (v NullableWidgetSort) Get() *WidgetSort {
	return v.value
}

func (v NullableWidgetSort) Set(val *WidgetSort) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetSort) IsSet() bool {
	return v.isSet
}

func (v NullableWidgetSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetSort(val *WidgetSort) *NullableWidgetSort {
	return &NullableWidgetSort{value: val, isSet: true}
}

func (v NullableWidgetSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
