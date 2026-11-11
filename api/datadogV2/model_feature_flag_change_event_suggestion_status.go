// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventSuggestionStatus The review status of the change request.
type FeatureFlagChangeEventSuggestionStatus string

// List of FeatureFlagChangeEventSuggestionStatus.
const (
	FEATUREFLAGCHANGEEVENTSUGGESTIONSTATUS_PENDING  FeatureFlagChangeEventSuggestionStatus = "pending"
	FEATUREFLAGCHANGEEVENTSUGGESTIONSTATUS_APPROVED FeatureFlagChangeEventSuggestionStatus = "approved"
	FEATUREFLAGCHANGEEVENTSUGGESTIONSTATUS_REJECTED FeatureFlagChangeEventSuggestionStatus = "rejected"
)

var allowedFeatureFlagChangeEventSuggestionStatusEnumValues = []FeatureFlagChangeEventSuggestionStatus{
	FEATUREFLAGCHANGEEVENTSUGGESTIONSTATUS_PENDING,
	FEATUREFLAGCHANGEEVENTSUGGESTIONSTATUS_APPROVED,
	FEATUREFLAGCHANGEEVENTSUGGESTIONSTATUS_REJECTED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagChangeEventSuggestionStatus) GetAllowedValues() []FeatureFlagChangeEventSuggestionStatus {
	return allowedFeatureFlagChangeEventSuggestionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagChangeEventSuggestionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagChangeEventSuggestionStatus(value)
	return nil
}

// NewFeatureFlagChangeEventSuggestionStatusFromValue returns a pointer to a valid FeatureFlagChangeEventSuggestionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagChangeEventSuggestionStatusFromValue(v string) (*FeatureFlagChangeEventSuggestionStatus, error) {
	ev := FeatureFlagChangeEventSuggestionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagChangeEventSuggestionStatus: valid values are %v", v, allowedFeatureFlagChangeEventSuggestionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagChangeEventSuggestionStatus) IsValid() bool {
	for _, existing := range allowedFeatureFlagChangeEventSuggestionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagChangeEventSuggestionStatus value.
func (v FeatureFlagChangeEventSuggestionStatus) Ptr() *FeatureFlagChangeEventSuggestionStatus {
	return &v
}
