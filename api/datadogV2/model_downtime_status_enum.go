// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DowntimeStatusEnum The current status of the downtime.
type DowntimeStatusEnum string

// List of DowntimeStatusEnum.
const (
	DOWNTIMESTATUSENUM_ACTIVE    DowntimeStatusEnum = "active"
	DOWNTIMESTATUSENUM_CANCELED  DowntimeStatusEnum = "canceled"
	DOWNTIMESTATUSENUM_ENDED     DowntimeStatusEnum = "ended"
	DOWNTIMESTATUSENUM_SCHEDULED DowntimeStatusEnum = "scheduled"
)

var allowedDowntimeStatusEnumEnumValues = []DowntimeStatusEnum{
	DOWNTIMESTATUSENUM_ACTIVE,
	DOWNTIMESTATUSENUM_CANCELED,
	DOWNTIMESTATUSENUM_ENDED,
	DOWNTIMESTATUSENUM_SCHEDULED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeStatusEnum) GetAllowedValues() []DowntimeStatusEnum {
	return allowedDowntimeStatusEnumEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeStatusEnum(value)
	return nil
}

// NewDowntimeStatusEnumFromValue returns a pointer to a valid DowntimeStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeStatusEnumFromValue(v string) (*DowntimeStatusEnum, error) {
	ev := DowntimeStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeStatusEnum: valid values are %v", v, allowedDowntimeStatusEnumEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeStatusEnum) IsValid() bool {
	for _, existing := range allowedDowntimeStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeStatusEnum value.
func (v DowntimeStatusEnum) Ptr() *DowntimeStatusEnum {
	return &v
}

// NullableDowntimeStatusEnum handles when a null is used for DowntimeStatusEnum.
type NullableDowntimeStatusEnum struct {
	value *DowntimeStatusEnum
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeStatusEnum) Get() *DowntimeStatusEnum {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeStatusEnum) Set(val *DowntimeStatusEnum) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeStatusEnum) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDowntimeStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeStatusEnum initializes the struct as if Set has been called.
func NewNullableDowntimeStatusEnum(val *DowntimeStatusEnum) *NullableDowntimeStatusEnum {
	return &NullableDowntimeStatusEnum{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
