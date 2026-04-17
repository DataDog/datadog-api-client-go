// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlagSuggestionDataType Flag suggestions resource type.
type FlagSuggestionDataType string

// List of FlagSuggestionDataType.
const (
	FLAGSUGGESTIONDATATYPE_FLAG_SUGGESTIONS FlagSuggestionDataType = "flag-suggestions"
)

var allowedFlagSuggestionDataTypeEnumValues = []FlagSuggestionDataType{
	FLAGSUGGESTIONDATATYPE_FLAG_SUGGESTIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlagSuggestionDataType) GetAllowedValues() []FlagSuggestionDataType {
	return allowedFlagSuggestionDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlagSuggestionDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlagSuggestionDataType(value)
	return nil
}

// NewFlagSuggestionDataTypeFromValue returns a pointer to a valid FlagSuggestionDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlagSuggestionDataTypeFromValue(v string) (*FlagSuggestionDataType, error) {
	ev := FlagSuggestionDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlagSuggestionDataType: valid values are %v", v, allowedFlagSuggestionDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlagSuggestionDataType) IsValid() bool {
	for _, existing := range allowedFlagSuggestionDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlagSuggestionDataType value.
func (v FlagSuggestionDataType) Ptr() *FlagSuggestionDataType {
	return &v
}
