// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagStalenessStatusOverride An explicit staleness status override. Omitted when no override is set.
type FeatureFlagStalenessStatusOverride string

// List of FeatureFlagStalenessStatusOverride.
const (
	FEATUREFLAGSTALENESSSTATUSOVERRIDE_STALE     FeatureFlagStalenessStatusOverride = "STALE"
	FEATUREFLAGSTALENESSSTATUSOVERRIDE_PERMANENT FeatureFlagStalenessStatusOverride = "PERMANENT"
)

var allowedFeatureFlagStalenessStatusOverrideEnumValues = []FeatureFlagStalenessStatusOverride{
	FEATUREFLAGSTALENESSSTATUSOVERRIDE_STALE,
	FEATUREFLAGSTALENESSSTATUSOVERRIDE_PERMANENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagStalenessStatusOverride) GetAllowedValues() []FeatureFlagStalenessStatusOverride {
	return allowedFeatureFlagStalenessStatusOverrideEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagStalenessStatusOverride) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagStalenessStatusOverride(value)
	return nil
}

// NewFeatureFlagStalenessStatusOverrideFromValue returns a pointer to a valid FeatureFlagStalenessStatusOverride
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagStalenessStatusOverrideFromValue(v string) (*FeatureFlagStalenessStatusOverride, error) {
	ev := FeatureFlagStalenessStatusOverride(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagStalenessStatusOverride: valid values are %v", v, allowedFeatureFlagStalenessStatusOverrideEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagStalenessStatusOverride) IsValid() bool {
	for _, existing := range allowedFeatureFlagStalenessStatusOverrideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagStalenessStatusOverride value.
func (v FeatureFlagStalenessStatusOverride) Ptr() *FeatureFlagStalenessStatusOverride {
	return &v
}
