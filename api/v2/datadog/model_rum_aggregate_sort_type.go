/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

// RUMAggregateSortType The type of sorting algorithm.
type RUMAggregateSortType string

// List of RUMAggregateSortType
const (
	RUMAGGREGATESORTTYPE_ALPHABETICAL RUMAggregateSortType = "alphabetical"
	RUMAGGREGATESORTTYPE_MEASURE      RUMAggregateSortType = "measure"
)

var allowedRUMAggregateSortTypeEnumValues = []RUMAggregateSortType{
	RUMAGGREGATESORTTYPE_ALPHABETICAL,
	RUMAGGREGATESORTTYPE_MEASURE,
}

func (w *RUMAggregateSortType) GetAllowedValues() []RUMAggregateSortType {
	return allowedRUMAggregateSortTypeEnumValues
}

func (v *RUMAggregateSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMAggregateSortType(value)
	return nil
}

// NewRUMAggregateSortTypeFromValue returns a pointer to a valid RUMAggregateSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRUMAggregateSortTypeFromValue(v string) (*RUMAggregateSortType, error) {
	ev := RUMAggregateSortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RUMAggregateSortType: valid values are %v", v, allowedRUMAggregateSortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RUMAggregateSortType) IsValid() bool {
	for _, existing := range allowedRUMAggregateSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMAggregateSortType value
func (v RUMAggregateSortType) Ptr() *RUMAggregateSortType {
	return &v
}

type NullableRUMAggregateSortType struct {
	value *RUMAggregateSortType
	isSet bool
}

func (v NullableRUMAggregateSortType) Get() *RUMAggregateSortType {
	return v.value
}

func (v *NullableRUMAggregateSortType) Set(val *RUMAggregateSortType) {
	v.value = val
	v.isSet = true
}

func (v NullableRUMAggregateSortType) IsSet() bool {
	return v.isSet
}

func (v *NullableRUMAggregateSortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRUMAggregateSortType(val *RUMAggregateSortType) *NullableRUMAggregateSortType {
	return &NullableRUMAggregateSortType{value: val, isSet: true}
}

func (v NullableRUMAggregateSortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRUMAggregateSortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
