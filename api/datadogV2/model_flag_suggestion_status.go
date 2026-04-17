// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlagSuggestionStatus The status of a flag suggestion.
type FlagSuggestionStatus string

// List of FlagSuggestionStatus.
const (
	FLAGSUGGESTIONSTATUS_PENDING  FlagSuggestionStatus = "pending"
	FLAGSUGGESTIONSTATUS_REJECTED FlagSuggestionStatus = "rejected"
	FLAGSUGGESTIONSTATUS_APPROVED FlagSuggestionStatus = "approved"
)

var allowedFlagSuggestionStatusEnumValues = []FlagSuggestionStatus{
	FLAGSUGGESTIONSTATUS_PENDING,
	FLAGSUGGESTIONSTATUS_REJECTED,
	FLAGSUGGESTIONSTATUS_APPROVED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlagSuggestionStatus) GetAllowedValues() []FlagSuggestionStatus {
	return allowedFlagSuggestionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlagSuggestionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlagSuggestionStatus(value)
	return nil
}

// NewFlagSuggestionStatusFromValue returns a pointer to a valid FlagSuggestionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlagSuggestionStatusFromValue(v string) (*FlagSuggestionStatus, error) {
	ev := FlagSuggestionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlagSuggestionStatus: valid values are %v", v, allowedFlagSuggestionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlagSuggestionStatus) IsValid() bool {
	for _, existing := range allowedFlagSuggestionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlagSuggestionStatus value.
func (v FlagSuggestionStatus) Ptr() *FlagSuggestionStatus {
	return &v
}
