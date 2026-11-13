// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagSuggestionDataType The resource type for feature flag suggestions.
type FeatureFlagSuggestionDataType string

// List of FeatureFlagSuggestionDataType.
const (
	FEATUREFLAGSUGGESTIONDATATYPE_FLAG_SUGGESTIONS FeatureFlagSuggestionDataType = "flag-suggestions"
)

var allowedFeatureFlagSuggestionDataTypeEnumValues = []FeatureFlagSuggestionDataType{
	FEATUREFLAGSUGGESTIONDATATYPE_FLAG_SUGGESTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagSuggestionDataType) GetAllowedValues() []FeatureFlagSuggestionDataType {
	return allowedFeatureFlagSuggestionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagSuggestionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagSuggestionDataType(value)
	return nil
}

// NewFeatureFlagSuggestionDataTypeFromValue returns a pointer to a valid FeatureFlagSuggestionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagSuggestionDataTypeFromValue(v string) (*FeatureFlagSuggestionDataType, error) {
	ev := FeatureFlagSuggestionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagSuggestionDataType: valid values are %v", v, allowedFeatureFlagSuggestionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagSuggestionDataType) IsValid() bool {
	for _, existing := range allowedFeatureFlagSuggestionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagSuggestionDataType value.
func (v FeatureFlagSuggestionDataType) Ptr() *FeatureFlagSuggestionDataType {
	return &v
}
