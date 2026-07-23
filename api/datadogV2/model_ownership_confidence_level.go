// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipConfidenceLevel The ownership confidence level.
type OwnershipConfidenceLevel string

// List of OwnershipConfidenceLevel.
const (
	OWNERSHIPCONFIDENCELEVEL_HIGH   OwnershipConfidenceLevel = "high"
	OWNERSHIPCONFIDENCELEVEL_MEDIUM OwnershipConfidenceLevel = "medium"
	OWNERSHIPCONFIDENCELEVEL_LOW    OwnershipConfidenceLevel = "low"
)

var allowedOwnershipConfidenceLevelEnumValues = []OwnershipConfidenceLevel{
	OWNERSHIPCONFIDENCELEVEL_HIGH,
	OWNERSHIPCONFIDENCELEVEL_MEDIUM,
	OWNERSHIPCONFIDENCELEVEL_LOW,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OwnershipConfidenceLevel) GetAllowedValues() []OwnershipConfidenceLevel {
	return allowedOwnershipConfidenceLevelEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OwnershipConfidenceLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OwnershipConfidenceLevel(value)
	return nil
}

// NewOwnershipConfidenceLevelFromValue returns a pointer to a valid OwnershipConfidenceLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOwnershipConfidenceLevelFromValue(v string) (*OwnershipConfidenceLevel, error) {
	ev := OwnershipConfidenceLevel(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OwnershipConfidenceLevel: valid values are %v", v, allowedOwnershipConfidenceLevelEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OwnershipConfidenceLevel) IsValid() bool {
	for _, existing := range allowedOwnershipConfidenceLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipConfidenceLevel value.
func (v OwnershipConfidenceLevel) Ptr() *OwnershipConfidenceLevel {
	return &v
}
