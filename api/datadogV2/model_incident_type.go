// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"encoding/json"
	"fmt"

)


// IncidentType Incident resource type.
type IncidentType string

// List of IncidentType.
const (
	INCIDENTTYPE_INCIDENTS IncidentType = "incidents"
)

var allowedIncidentTypeEnumValues = []IncidentType{
	INCIDENTTYPE_INCIDENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentType) GetAllowedValues() []IncidentType {
	return allowedIncidentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentType(value)
	return nil
}

// NewIncidentTypeFromValue returns a pointer to a valid IncidentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentTypeFromValue(v string) (*IncidentType, error) {
	ev := IncidentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentType: valid values are %v", v, allowedIncidentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentType) IsValid() bool {
	for _, existing := range allowedIncidentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentType value.
func (v IncidentType) Ptr() *IncidentType {
	return &v
}

// NullableIncidentType handles when a null is used for IncidentType.
type NullableIncidentType struct {
	value *IncidentType
	isSet bool
}

// Get returns the associated value.
func (v NullableIncidentType) Get() *IncidentType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableIncidentType) Set(val *IncidentType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableIncidentType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableIncidentType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncidentType initializes the struct as if Set has been called.
func NewNullableIncidentType(val *IncidentType) *NullableIncidentType {
	return &NullableIncidentType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableIncidentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableIncidentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
