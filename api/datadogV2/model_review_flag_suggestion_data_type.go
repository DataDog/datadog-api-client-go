// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReviewFlagSuggestionDataType The resource type for reviewing feature flag suggestions.
type ReviewFlagSuggestionDataType string

// List of ReviewFlagSuggestionDataType.
const (
	REVIEWFLAGSUGGESTIONDATATYPE_FLAG_SUGGESTION_EVENTS ReviewFlagSuggestionDataType = "flag-suggestion-events"
)

var allowedReviewFlagSuggestionDataTypeEnumValues = []ReviewFlagSuggestionDataType{
	REVIEWFLAGSUGGESTIONDATATYPE_FLAG_SUGGESTION_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReviewFlagSuggestionDataType) GetAllowedValues() []ReviewFlagSuggestionDataType {
	return allowedReviewFlagSuggestionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReviewFlagSuggestionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReviewFlagSuggestionDataType(value)
	return nil
}

// NewReviewFlagSuggestionDataTypeFromValue returns a pointer to a valid ReviewFlagSuggestionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReviewFlagSuggestionDataTypeFromValue(v string) (*ReviewFlagSuggestionDataType, error) {
	ev := ReviewFlagSuggestionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReviewFlagSuggestionDataType: valid values are %v", v, allowedReviewFlagSuggestionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReviewFlagSuggestionDataType) IsValid() bool {
	for _, existing := range allowedReviewFlagSuggestionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReviewFlagSuggestionDataType value.
func (v ReviewFlagSuggestionDataType) Ptr() *ReviewFlagSuggestionDataType {
	return &v
}
