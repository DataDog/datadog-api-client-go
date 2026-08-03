// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestImpactLevel The impact level of the flaky test, derived from its impact score.
type FlakyTestImpactLevel string

// List of FlakyTestImpactLevel.
const (
	FLAKYTESTIMPACTLEVEL_LOW    FlakyTestImpactLevel = "low"
	FLAKYTESTIMPACTLEVEL_MEDIUM FlakyTestImpactLevel = "medium"
	FLAKYTESTIMPACTLEVEL_HIGH   FlakyTestImpactLevel = "high"
)

var allowedFlakyTestImpactLevelEnumValues = []FlakyTestImpactLevel{
	FLAKYTESTIMPACTLEVEL_LOW,
	FLAKYTESTIMPACTLEVEL_MEDIUM,
	FLAKYTESTIMPACTLEVEL_HIGH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlakyTestImpactLevel) GetAllowedValues() []FlakyTestImpactLevel {
	return allowedFlakyTestImpactLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlakyTestImpactLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlakyTestImpactLevel(value)
	return nil
}

// NewFlakyTestImpactLevelFromValue returns a pointer to a valid FlakyTestImpactLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlakyTestImpactLevelFromValue(v string) (*FlakyTestImpactLevel, error) {
	ev := FlakyTestImpactLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlakyTestImpactLevel: valid values are %v", v, allowedFlakyTestImpactLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlakyTestImpactLevel) IsValid() bool {
	for _, existing := range allowedFlakyTestImpactLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlakyTestImpactLevel value.
func (v FlakyTestImpactLevel) Ptr() *FlakyTestImpactLevel {
	return &v
}
