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

// TableWidgetHasSearchBar Controls the display of the search bar.
type TableWidgetHasSearchBar string

// List of TableWidgetHasSearchBar
const (
	TABLEWIDGETHASSEARCHBAR_ALWAYS TableWidgetHasSearchBar = "always"
	TABLEWIDGETHASSEARCHBAR_NEVER  TableWidgetHasSearchBar = "never"
	TABLEWIDGETHASSEARCHBAR_AUTO   TableWidgetHasSearchBar = "auto"
)

var allowedTableWidgetHasSearchBarEnumValues = []TableWidgetHasSearchBar{
	TABLEWIDGETHASSEARCHBAR_ALWAYS,
	TABLEWIDGETHASSEARCHBAR_NEVER,
	TABLEWIDGETHASSEARCHBAR_AUTO,
}

func (w *TableWidgetHasSearchBar) GetAllowedValues() []TableWidgetHasSearchBar {
	return allowedTableWidgetHasSearchBarEnumValues
}

func (v *TableWidgetHasSearchBar) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TableWidgetHasSearchBar(value)
	return nil
}

// NewTableWidgetHasSearchBarFromValue returns a pointer to a valid TableWidgetHasSearchBar
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTableWidgetHasSearchBarFromValue(v string) (*TableWidgetHasSearchBar, error) {
	ev := TableWidgetHasSearchBar(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TableWidgetHasSearchBar: valid values are %v", v, allowedTableWidgetHasSearchBarEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TableWidgetHasSearchBar) IsValid() bool {
	for _, existing := range allowedTableWidgetHasSearchBarEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TableWidgetHasSearchBar value
func (v TableWidgetHasSearchBar) Ptr() *TableWidgetHasSearchBar {
	return &v
}

type NullableTableWidgetHasSearchBar struct {
	value *TableWidgetHasSearchBar
	isSet bool
}

func (v NullableTableWidgetHasSearchBar) Get() *TableWidgetHasSearchBar {
	return v.value
}

func (v *NullableTableWidgetHasSearchBar) Set(val *TableWidgetHasSearchBar) {
	v.value = val
	v.isSet = true
}

func (v NullableTableWidgetHasSearchBar) IsSet() bool {
	return v.isSet
}

func (v *NullableTableWidgetHasSearchBar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableWidgetHasSearchBar(val *TableWidgetHasSearchBar) *NullableTableWidgetHasSearchBar {
	return &NullableTableWidgetHasSearchBar{value: val, isSet: true}
}

func (v NullableTableWidgetHasSearchBar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableWidgetHasSearchBar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
