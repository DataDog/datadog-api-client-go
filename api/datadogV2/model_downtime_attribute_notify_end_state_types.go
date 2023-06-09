// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DowntimeAttributeNotifyEndStateTypes State that will trigger a monitor notification when the `notify_end_types` action occurs.
type DowntimeAttributeNotifyEndStateTypes string

// List of DowntimeAttributeNotifyEndStateTypes.
const (
	DOWNTIMEATTRIBUTENOTIFYENDSTATETYPES_ALERT   DowntimeAttributeNotifyEndStateTypes = "alert"
	DOWNTIMEATTRIBUTENOTIFYENDSTATETYPES_NO_DATA DowntimeAttributeNotifyEndStateTypes = "no data"
	DOWNTIMEATTRIBUTENOTIFYENDSTATETYPES_WARN    DowntimeAttributeNotifyEndStateTypes = "warn"
)

var allowedDowntimeAttributeNotifyEndStateTypesEnumValues = []DowntimeAttributeNotifyEndStateTypes{
	DOWNTIMEATTRIBUTENOTIFYENDSTATETYPES_ALERT,
	DOWNTIMEATTRIBUTENOTIFYENDSTATETYPES_NO_DATA,
	DOWNTIMEATTRIBUTENOTIFYENDSTATETYPES_WARN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeAttributeNotifyEndStateTypes) GetAllowedValues() []DowntimeAttributeNotifyEndStateTypes {
	return allowedDowntimeAttributeNotifyEndStateTypesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeAttributeNotifyEndStateTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeAttributeNotifyEndStateTypes(value)
	return nil
}

// NewDowntimeAttributeNotifyEndStateTypesFromValue returns a pointer to a valid DowntimeAttributeNotifyEndStateTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeAttributeNotifyEndStateTypesFromValue(v string) (*DowntimeAttributeNotifyEndStateTypes, error) {
	ev := DowntimeAttributeNotifyEndStateTypes(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeAttributeNotifyEndStateTypes: valid values are %v", v, allowedDowntimeAttributeNotifyEndStateTypesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeAttributeNotifyEndStateTypes) IsValid() bool {
	for _, existing := range allowedDowntimeAttributeNotifyEndStateTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeAttributeNotifyEndStateTypes value.
func (v DowntimeAttributeNotifyEndStateTypes) Ptr() *DowntimeAttributeNotifyEndStateTypes {
	return &v
}

// NullableDowntimeAttributeNotifyEndStateTypes handles when a null is used for DowntimeAttributeNotifyEndStateTypes.
type NullableDowntimeAttributeNotifyEndStateTypes struct {
	value *DowntimeAttributeNotifyEndStateTypes
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeAttributeNotifyEndStateTypes) Get() *DowntimeAttributeNotifyEndStateTypes {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeAttributeNotifyEndStateTypes) Set(val *DowntimeAttributeNotifyEndStateTypes) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeAttributeNotifyEndStateTypes) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDowntimeAttributeNotifyEndStateTypes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeAttributeNotifyEndStateTypes initializes the struct as if Set has been called.
func NewNullableDowntimeAttributeNotifyEndStateTypes(val *DowntimeAttributeNotifyEndStateTypes) *NullableDowntimeAttributeNotifyEndStateTypes {
	return &NullableDowntimeAttributeNotifyEndStateTypes{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeAttributeNotifyEndStateTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeAttributeNotifyEndStateTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
