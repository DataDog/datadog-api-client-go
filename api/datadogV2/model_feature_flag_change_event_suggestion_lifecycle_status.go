// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventSuggestionLifecycleStatus Where the change request is in its lifecycle.
type FeatureFlagChangeEventSuggestionLifecycleStatus string

// List of FeatureFlagChangeEventSuggestionLifecycleStatus.
const (
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_PENDING   FeatureFlagChangeEventSuggestionLifecycleStatus = "pending"
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_APPROVED  FeatureFlagChangeEventSuggestionLifecycleStatus = "approved"
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_APPLIED   FeatureFlagChangeEventSuggestionLifecycleStatus = "applied"
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_REJECTED  FeatureFlagChangeEventSuggestionLifecycleStatus = "rejected"
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_DISCARDED FeatureFlagChangeEventSuggestionLifecycleStatus = "discarded"
)

var allowedFeatureFlagChangeEventSuggestionLifecycleStatusEnumValues = []FeatureFlagChangeEventSuggestionLifecycleStatus{
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_PENDING,
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_APPROVED,
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_APPLIED,
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_REJECTED,
	FEATUREFLAGCHANGEEVENTSUGGESTIONLIFECYCLESTATUS_DISCARDED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagChangeEventSuggestionLifecycleStatus) GetAllowedValues() []FeatureFlagChangeEventSuggestionLifecycleStatus {
	return allowedFeatureFlagChangeEventSuggestionLifecycleStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagChangeEventSuggestionLifecycleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagChangeEventSuggestionLifecycleStatus(value)
	return nil
}

// NewFeatureFlagChangeEventSuggestionLifecycleStatusFromValue returns a pointer to a valid FeatureFlagChangeEventSuggestionLifecycleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagChangeEventSuggestionLifecycleStatusFromValue(v string) (*FeatureFlagChangeEventSuggestionLifecycleStatus, error) {
	ev := FeatureFlagChangeEventSuggestionLifecycleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagChangeEventSuggestionLifecycleStatus: valid values are %v", v, allowedFeatureFlagChangeEventSuggestionLifecycleStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagChangeEventSuggestionLifecycleStatus) IsValid() bool {
	for _, existing := range allowedFeatureFlagChangeEventSuggestionLifecycleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagChangeEventSuggestionLifecycleStatus value.
func (v FeatureFlagChangeEventSuggestionLifecycleStatus) Ptr() *FeatureFlagChangeEventSuggestionLifecycleStatus {
	return &v
}
