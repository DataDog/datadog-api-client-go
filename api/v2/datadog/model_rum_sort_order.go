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


// RUMSortOrder The order to use, ascending or descending.
type RUMSortOrder string

// List of RUMSortOrder
const (
	RUMSORTORDER_ASCENDING RUMSortOrder = "asc"
	RUMSORTORDER_DESCENDING RUMSortOrder = "desc"
)

var allowedRUMSortOrderEnumValues = []RUMSortOrder{
	RUMSORTORDER_ASCENDING,
	RUMSORTORDER_DESCENDING,
}

func (w *RUMSortOrder) GetAllowedValues() []RUMSortOrder {
	return allowedRUMSortOrderEnumValues
}

func (v *RUMSortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMSortOrder(value)
	return nil
}

// NewRUMSortOrderFromValue returns a pointer to a valid RUMSortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRUMSortOrderFromValue(v string) (*RUMSortOrder, error) {
	ev := RUMSortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RUMSortOrder: valid values are %v", v, allowedRUMSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RUMSortOrder) IsValid() bool {
	for _, existing := range allowedRUMSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMSortOrder value
func (v RUMSortOrder) Ptr() *RUMSortOrder {
	return &v
}

type NullableRUMSortOrder struct {
	value *RUMSortOrder
	isSet bool
}

func (v NullableRUMSortOrder) Get() *RUMSortOrder {
	return v.value
}

func (v *NullableRUMSortOrder) Set(val *RUMSortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableRUMSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableRUMSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRUMSortOrder(val *RUMSortOrder) *NullableRUMSortOrder {
	return &NullableRUMSortOrder{value: val, isSet: true}
}

func (v NullableRUMSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRUMSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
