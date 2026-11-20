// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PowerpackState Test-only generated schema for the dashboard versioning proof of concept.
type PowerpackState string

// List of PowerpackState.
const (
	POWERPACKSTATE_ACTIVE   PowerpackState = "active"
	POWERPACKSTATE_ARCHIVED PowerpackState = "archived"
)

var allowedPowerpackStateEnumValues = []PowerpackState{
	POWERPACKSTATE_ACTIVE,
	POWERPACKSTATE_ARCHIVED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PowerpackState) GetAllowedValues() []PowerpackState {
	return allowedPowerpackStateEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PowerpackState) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PowerpackState(value)
	return nil
}

// NewPowerpackStateFromValue returns a pointer to a valid PowerpackState
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPowerpackStateFromValue(v string) (*PowerpackState, error) {
	ev := PowerpackState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PowerpackState: valid values are %v", v, allowedPowerpackStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PowerpackState) IsValid() bool {
	for _, existing := range allowedPowerpackStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerpackState value.
func (v PowerpackState) Ptr() *PowerpackState {
	return &v
}
