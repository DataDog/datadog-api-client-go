// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StabilityLevel The stability level for action filtering.
type StabilityLevel string

// List of StabilityLevel.
const (
	STABILITYLEVEL_UNSPECIFIED StabilityLevel = "UNSPECIFIED"
	STABILITYLEVEL_DEV         StabilityLevel = "DEV"
	STABILITYLEVEL_BETA        StabilityLevel = "BETA"
	STABILITYLEVEL_STABLE      StabilityLevel = "STABLE"
)

var allowedStabilityLevelEnumValues = []StabilityLevel{
	STABILITYLEVEL_UNSPECIFIED,
	STABILITYLEVEL_DEV,
	STABILITYLEVEL_BETA,
	STABILITYLEVEL_STABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *StabilityLevel) GetAllowedValues() []StabilityLevel {
	return allowedStabilityLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StabilityLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StabilityLevel(value)
	return nil
}

// NewStabilityLevelFromValue returns a pointer to a valid StabilityLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStabilityLevelFromValue(v string) (*StabilityLevel, error) {
	ev := StabilityLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StabilityLevel: valid values are %v", v, allowedStabilityLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StabilityLevel) IsValid() bool {
	for _, existing := range allowedStabilityLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StabilityLevel value.
func (v StabilityLevel) Ptr() *StabilityLevel {
	return &v
}
