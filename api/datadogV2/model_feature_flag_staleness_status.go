// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagStalenessStatus The computed staleness status of a feature flag.
type FeatureFlagStalenessStatus string

// List of FeatureFlagStalenessStatus.
const (
	FEATUREFLAGSTALENESSSTATUS_ACTIVE    FeatureFlagStalenessStatus = "ACTIVE"
	FEATUREFLAGSTALENESSSTATUS_STALE     FeatureFlagStalenessStatus = "STALE"
	FEATUREFLAGSTALENESSSTATUS_PERMANENT FeatureFlagStalenessStatus = "PERMANENT"
)

var allowedFeatureFlagStalenessStatusEnumValues = []FeatureFlagStalenessStatus{
	FEATUREFLAGSTALENESSSTATUS_ACTIVE,
	FEATUREFLAGSTALENESSSTATUS_STALE,
	FEATUREFLAGSTALENESSSTATUS_PERMANENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagStalenessStatus) GetAllowedValues() []FeatureFlagStalenessStatus {
	return allowedFeatureFlagStalenessStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagStalenessStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagStalenessStatus(value)
	return nil
}

// NewFeatureFlagStalenessStatusFromValue returns a pointer to a valid FeatureFlagStalenessStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagStalenessStatusFromValue(v string) (*FeatureFlagStalenessStatus, error) {
	ev := FeatureFlagStalenessStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagStalenessStatus: valid values are %v", v, allowedFeatureFlagStalenessStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagStalenessStatus) IsValid() bool {
	for _, existing := range allowedFeatureFlagStalenessStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagStalenessStatus value.
func (v FeatureFlagStalenessStatus) Ptr() *FeatureFlagStalenessStatus {
	return &v
}
