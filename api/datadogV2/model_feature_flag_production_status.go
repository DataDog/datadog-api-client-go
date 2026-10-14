// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagProductionStatus The status of a feature flag across production environments.
type FeatureFlagProductionStatus string

// List of FeatureFlagProductionStatus.
const (
	FEATUREFLAGPRODUCTIONSTATUS_ENABLED    FeatureFlagProductionStatus = "ENABLED"
	FEATUREFLAGPRODUCTIONSTATUS_DISABLED   FeatureFlagProductionStatus = "DISABLED"
	FEATUREFLAGPRODUCTIONSTATUS_OVERRIDDEN FeatureFlagProductionStatus = "OVERRIDDEN"
)

var allowedFeatureFlagProductionStatusEnumValues = []FeatureFlagProductionStatus{
	FEATUREFLAGPRODUCTIONSTATUS_ENABLED,
	FEATUREFLAGPRODUCTIONSTATUS_DISABLED,
	FEATUREFLAGPRODUCTIONSTATUS_OVERRIDDEN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagProductionStatus) GetAllowedValues() []FeatureFlagProductionStatus {
	return allowedFeatureFlagProductionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagProductionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagProductionStatus(value)
	return nil
}

// NewFeatureFlagProductionStatusFromValue returns a pointer to a valid FeatureFlagProductionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagProductionStatusFromValue(v string) (*FeatureFlagProductionStatus, error) {
	ev := FeatureFlagProductionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagProductionStatus: valid values are %v", v, allowedFeatureFlagProductionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagProductionStatus) IsValid() bool {
	for _, existing := range allowedFeatureFlagProductionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagProductionStatus value.
func (v FeatureFlagProductionStatus) Ptr() *FeatureFlagProductionStatus {
	return &v
}
