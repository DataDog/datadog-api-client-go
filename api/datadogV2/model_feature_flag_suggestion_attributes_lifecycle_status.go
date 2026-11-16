// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagSuggestionAttributesLifecycleStatus The lifecycle status of the suggestion.
type FeatureFlagSuggestionAttributesLifecycleStatus string

// List of FeatureFlagSuggestionAttributesLifecycleStatus.
const (
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_PENDING   FeatureFlagSuggestionAttributesLifecycleStatus = "pending"
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_APPROVED  FeatureFlagSuggestionAttributesLifecycleStatus = "approved"
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_APPLIED   FeatureFlagSuggestionAttributesLifecycleStatus = "applied"
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_REJECTED  FeatureFlagSuggestionAttributesLifecycleStatus = "rejected"
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_DISCARDED FeatureFlagSuggestionAttributesLifecycleStatus = "discarded"
)

var allowedFeatureFlagSuggestionAttributesLifecycleStatusEnumValues = []FeatureFlagSuggestionAttributesLifecycleStatus{
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_PENDING,
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_APPROVED,
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_APPLIED,
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_REJECTED,
	FEATUREFLAGSUGGESTIONATTRIBUTESLIFECYCLESTATUS_DISCARDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagSuggestionAttributesLifecycleStatus) GetAllowedValues() []FeatureFlagSuggestionAttributesLifecycleStatus {
	return allowedFeatureFlagSuggestionAttributesLifecycleStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagSuggestionAttributesLifecycleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagSuggestionAttributesLifecycleStatus(value)
	return nil
}

// NewFeatureFlagSuggestionAttributesLifecycleStatusFromValue returns a pointer to a valid FeatureFlagSuggestionAttributesLifecycleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagSuggestionAttributesLifecycleStatusFromValue(v string) (*FeatureFlagSuggestionAttributesLifecycleStatus, error) {
	ev := FeatureFlagSuggestionAttributesLifecycleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagSuggestionAttributesLifecycleStatus: valid values are %v", v, allowedFeatureFlagSuggestionAttributesLifecycleStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagSuggestionAttributesLifecycleStatus) IsValid() bool {
	for _, existing := range allowedFeatureFlagSuggestionAttributesLifecycleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagSuggestionAttributesLifecycleStatus value.
func (v FeatureFlagSuggestionAttributesLifecycleStatus) Ptr() *FeatureFlagSuggestionAttributesLifecycleStatus {
	return &v
}
