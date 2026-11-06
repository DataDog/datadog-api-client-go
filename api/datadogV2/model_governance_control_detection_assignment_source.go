// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlDetectionAssignmentSource How the detection's current assignment was determined. Possible values are `auto_resolved`, `manual`, `reassigned`, and `cleared`. `null` when not set.
type GovernanceControlDetectionAssignmentSource string

// List of GovernanceControlDetectionAssignmentSource.
const (
	GOVERNANCECONTROLDETECTIONASSIGNMENTSOURCE_AUTO_RESOLVED GovernanceControlDetectionAssignmentSource = "auto_resolved"
	GOVERNANCECONTROLDETECTIONASSIGNMENTSOURCE_MANUAL        GovernanceControlDetectionAssignmentSource = "manual"
	GOVERNANCECONTROLDETECTIONASSIGNMENTSOURCE_REASSIGNED    GovernanceControlDetectionAssignmentSource = "reassigned"
	GOVERNANCECONTROLDETECTIONASSIGNMENTSOURCE_CLEARED       GovernanceControlDetectionAssignmentSource = "cleared"
)

var allowedGovernanceControlDetectionAssignmentSourceEnumValues = []GovernanceControlDetectionAssignmentSource{
	GOVERNANCECONTROLDETECTIONASSIGNMENTSOURCE_AUTO_RESOLVED,
	GOVERNANCECONTROLDETECTIONASSIGNMENTSOURCE_MANUAL,
	GOVERNANCECONTROLDETECTIONASSIGNMENTSOURCE_REASSIGNED,
	GOVERNANCECONTROLDETECTIONASSIGNMENTSOURCE_CLEARED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceControlDetectionAssignmentSource) GetAllowedValues() []GovernanceControlDetectionAssignmentSource {
	return allowedGovernanceControlDetectionAssignmentSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceControlDetectionAssignmentSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceControlDetectionAssignmentSource(value)
	return nil
}

// NewGovernanceControlDetectionAssignmentSourceFromValue returns a pointer to a valid GovernanceControlDetectionAssignmentSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceControlDetectionAssignmentSourceFromValue(v string) (*GovernanceControlDetectionAssignmentSource, error) {
	ev := GovernanceControlDetectionAssignmentSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceControlDetectionAssignmentSource: valid values are %v", v, allowedGovernanceControlDetectionAssignmentSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceControlDetectionAssignmentSource) IsValid() bool {
	for _, existing := range allowedGovernanceControlDetectionAssignmentSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceControlDetectionAssignmentSource value.
func (v GovernanceControlDetectionAssignmentSource) Ptr() *GovernanceControlDetectionAssignmentSource {
	return &v
}

// NullableGovernanceControlDetectionAssignmentSource handles when a null is used for GovernanceControlDetectionAssignmentSource.
type NullableGovernanceControlDetectionAssignmentSource struct {
	value *GovernanceControlDetectionAssignmentSource
	isSet bool
}

// Get returns the associated value.
func (v NullableGovernanceControlDetectionAssignmentSource) Get() *GovernanceControlDetectionAssignmentSource {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableGovernanceControlDetectionAssignmentSource) Set(val *GovernanceControlDetectionAssignmentSource) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableGovernanceControlDetectionAssignmentSource) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableGovernanceControlDetectionAssignmentSource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableGovernanceControlDetectionAssignmentSource initializes the struct as if Set has been called.
func NewNullableGovernanceControlDetectionAssignmentSource(val *GovernanceControlDetectionAssignmentSource) *NullableGovernanceControlDetectionAssignmentSource {
	return &NullableGovernanceControlDetectionAssignmentSource{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableGovernanceControlDetectionAssignmentSource) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableGovernanceControlDetectionAssignmentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
