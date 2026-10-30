// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventType The resource type.
type FeatureFlagChangeEventType string

// List of FeatureFlagChangeEventType.
const (
	FEATUREFLAGCHANGEEVENTTYPE_CHANGE_EVENTS FeatureFlagChangeEventType = "change-events"
)

var allowedFeatureFlagChangeEventTypeEnumValues = []FeatureFlagChangeEventType{
	FEATUREFLAGCHANGEEVENTTYPE_CHANGE_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagChangeEventType) GetAllowedValues() []FeatureFlagChangeEventType {
	return allowedFeatureFlagChangeEventTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagChangeEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagChangeEventType(value)
	return nil
}

// NewFeatureFlagChangeEventTypeFromValue returns a pointer to a valid FeatureFlagChangeEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagChangeEventTypeFromValue(v string) (*FeatureFlagChangeEventType, error) {
	ev := FeatureFlagChangeEventType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagChangeEventType: valid values are %v", v, allowedFeatureFlagChangeEventTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagChangeEventType) IsValid() bool {
	for _, existing := range allowedFeatureFlagChangeEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagChangeEventType value.
func (v FeatureFlagChangeEventType) Ptr() *FeatureFlagChangeEventType {
	return &v
}
