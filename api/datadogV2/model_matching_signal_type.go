// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MatchingSignalType The type of the resource. The value should always be `matching_signal`.
type MatchingSignalType string

// List of MatchingSignalType.
const (
	MATCHINGSIGNALTYPE_MATCHING_SIGNAL MatchingSignalType = "matching_signal"
)

var allowedMatchingSignalTypeEnumValues = []MatchingSignalType{
	MATCHINGSIGNALTYPE_MATCHING_SIGNAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MatchingSignalType) GetAllowedValues() []MatchingSignalType {
	return allowedMatchingSignalTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MatchingSignalType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MatchingSignalType(value)
	return nil
}

// NewMatchingSignalTypeFromValue returns a pointer to a valid MatchingSignalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMatchingSignalTypeFromValue(v string) (*MatchingSignalType, error) {
	ev := MatchingSignalType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MatchingSignalType: valid values are %v", v, allowedMatchingSignalTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MatchingSignalType) IsValid() bool {
	for _, existing := range allowedMatchingSignalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MatchingSignalType value.
func (v MatchingSignalType) Ptr() *MatchingSignalType {
	return &v
}
