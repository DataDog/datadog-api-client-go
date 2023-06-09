// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DowntimeResourceType Downtime resource type.
type DowntimeResourceType string

// List of DowntimeResourceType.
const (
	DOWNTIMERESOURCETYPE_DOWNTIME DowntimeResourceType = "downtime"
)

var allowedDowntimeResourceTypeEnumValues = []DowntimeResourceType{
	DOWNTIMERESOURCETYPE_DOWNTIME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeResourceType) GetAllowedValues() []DowntimeResourceType {
	return allowedDowntimeResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeResourceType(value)
	return nil
}

// NewDowntimeResourceTypeFromValue returns a pointer to a valid DowntimeResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeResourceTypeFromValue(v string) (*DowntimeResourceType, error) {
	ev := DowntimeResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeResourceType: valid values are %v", v, allowedDowntimeResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeResourceType) IsValid() bool {
	for _, existing := range allowedDowntimeResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeResourceType value.
func (v DowntimeResourceType) Ptr() *DowntimeResourceType {
	return &v
}

// NullableDowntimeResourceType handles when a null is used for DowntimeResourceType.
type NullableDowntimeResourceType struct {
	value *DowntimeResourceType
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeResourceType) Get() *DowntimeResourceType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeResourceType) Set(val *DowntimeResourceType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeResourceType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDowntimeResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeResourceType initializes the struct as if Set has been called.
func NewNullableDowntimeResourceType(val *DowntimeResourceType) *NullableDowntimeResourceType {
	return &NullableDowntimeResourceType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
