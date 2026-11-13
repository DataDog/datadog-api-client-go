// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagSuggestionAttributesCurrentStatus The legacy review status. Use `lifecycle_status` to distinguish approval from application.
type FeatureFlagSuggestionAttributesCurrentStatus string

// List of FeatureFlagSuggestionAttributesCurrentStatus.
const (
	FEATUREFLAGSUGGESTIONATTRIBUTESCURRENTSTATUS_PENDING  FeatureFlagSuggestionAttributesCurrentStatus = "pending"
	FEATUREFLAGSUGGESTIONATTRIBUTESCURRENTSTATUS_APPROVED FeatureFlagSuggestionAttributesCurrentStatus = "approved"
	FEATUREFLAGSUGGESTIONATTRIBUTESCURRENTSTATUS_REJECTED FeatureFlagSuggestionAttributesCurrentStatus = "rejected"
)

var allowedFeatureFlagSuggestionAttributesCurrentStatusEnumValues = []FeatureFlagSuggestionAttributesCurrentStatus{
	FEATUREFLAGSUGGESTIONATTRIBUTESCURRENTSTATUS_PENDING,
	FEATUREFLAGSUGGESTIONATTRIBUTESCURRENTSTATUS_APPROVED,
	FEATUREFLAGSUGGESTIONATTRIBUTESCURRENTSTATUS_REJECTED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagSuggestionAttributesCurrentStatus) GetAllowedValues() []FeatureFlagSuggestionAttributesCurrentStatus {
	return allowedFeatureFlagSuggestionAttributesCurrentStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagSuggestionAttributesCurrentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagSuggestionAttributesCurrentStatus(value)
	return nil
}

// NewFeatureFlagSuggestionAttributesCurrentStatusFromValue returns a pointer to a valid FeatureFlagSuggestionAttributesCurrentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagSuggestionAttributesCurrentStatusFromValue(v string) (*FeatureFlagSuggestionAttributesCurrentStatus, error) {
	ev := FeatureFlagSuggestionAttributesCurrentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagSuggestionAttributesCurrentStatus: valid values are %v", v, allowedFeatureFlagSuggestionAttributesCurrentStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagSuggestionAttributesCurrentStatus) IsValid() bool {
	for _, existing := range allowedFeatureFlagSuggestionAttributesCurrentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagSuggestionAttributesCurrentStatus value.
func (v FeatureFlagSuggestionAttributesCurrentStatus) Ptr() *FeatureFlagSuggestionAttributesCurrentStatus {
	return &v
}
