// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierSeverityDelta The direction in which to shift the severity of matched findings by one rank.
type SeverityModifierSeverityDelta string

// List of SeverityModifierSeverityDelta.
const (
	SEVERITYMODIFIERSEVERITYDELTA_UP_ONE   SeverityModifierSeverityDelta = "up_one"
	SEVERITYMODIFIERSEVERITYDELTA_DOWN_ONE SeverityModifierSeverityDelta = "down_one"
)

var allowedSeverityModifierSeverityDeltaEnumValues = []SeverityModifierSeverityDelta{
	SEVERITYMODIFIERSEVERITYDELTA_UP_ONE,
	SEVERITYMODIFIERSEVERITYDELTA_DOWN_ONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeverityModifierSeverityDelta) GetAllowedValues() []SeverityModifierSeverityDelta {
	return allowedSeverityModifierSeverityDeltaEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeverityModifierSeverityDelta) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeverityModifierSeverityDelta(value)
	return nil
}

// NewSeverityModifierSeverityDeltaFromValue returns a pointer to a valid SeverityModifierSeverityDelta
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityModifierSeverityDeltaFromValue(v string) (*SeverityModifierSeverityDelta, error) {
	ev := SeverityModifierSeverityDelta(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeverityModifierSeverityDelta: valid values are %v", v, allowedSeverityModifierSeverityDeltaEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeverityModifierSeverityDelta) IsValid() bool {
	for _, existing := range allowedSeverityModifierSeverityDeltaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityModifierSeverityDelta value.
func (v SeverityModifierSeverityDelta) Ptr() *SeverityModifierSeverityDelta {
	return &v
}
