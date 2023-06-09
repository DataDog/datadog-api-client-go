// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DowntimeUserType User resource type.
type DowntimeUserType string

// List of DowntimeUserType.
const (
	DOWNTIMEUSERTYPE_USERS DowntimeUserType = "users"
)

var allowedDowntimeUserTypeEnumValues = []DowntimeUserType{
	DOWNTIMEUSERTYPE_USERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeUserType) GetAllowedValues() []DowntimeUserType {
	return allowedDowntimeUserTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeUserType(value)
	return nil
}

// NewDowntimeUserTypeFromValue returns a pointer to a valid DowntimeUserType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeUserTypeFromValue(v string) (*DowntimeUserType, error) {
	ev := DowntimeUserType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeUserType: valid values are %v", v, allowedDowntimeUserTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeUserType) IsValid() bool {
	for _, existing := range allowedDowntimeUserTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeUserType value.
func (v DowntimeUserType) Ptr() *DowntimeUserType {
	return &v
}

// NullableDowntimeUserType handles when a null is used for DowntimeUserType.
type NullableDowntimeUserType struct {
	value *DowntimeUserType
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeUserType) Get() *DowntimeUserType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeUserType) Set(val *DowntimeUserType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeUserType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDowntimeUserType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeUserType initializes the struct as if Set has been called.
func NewNullableDowntimeUserType(val *DowntimeUserType) *NullableDowntimeUserType {
	return &NullableDowntimeUserType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
