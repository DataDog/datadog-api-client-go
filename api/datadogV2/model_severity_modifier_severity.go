// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierSeverity The severity to assign to matched findings. `info_none` is not supported for the `iac_misconfiguration`, `runtime_code_vulnerability`, `secret`, or `static_code_vulnerability` finding types.
type SeverityModifierSeverity string

// List of SeverityModifierSeverity.
const (
	SEVERITYMODIFIERSEVERITY_INFO_NONE SeverityModifierSeverity = "info_none"
	SEVERITYMODIFIERSEVERITY_LOW       SeverityModifierSeverity = "low"
	SEVERITYMODIFIERSEVERITY_MEDIUM    SeverityModifierSeverity = "medium"
	SEVERITYMODIFIERSEVERITY_HIGH      SeverityModifierSeverity = "high"
	SEVERITYMODIFIERSEVERITY_CRITICAL  SeverityModifierSeverity = "critical"
)

var allowedSeverityModifierSeverityEnumValues = []SeverityModifierSeverity{
	SEVERITYMODIFIERSEVERITY_INFO_NONE,
	SEVERITYMODIFIERSEVERITY_LOW,
	SEVERITYMODIFIERSEVERITY_MEDIUM,
	SEVERITYMODIFIERSEVERITY_HIGH,
	SEVERITYMODIFIERSEVERITY_CRITICAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeverityModifierSeverity) GetAllowedValues() []SeverityModifierSeverity {
	return allowedSeverityModifierSeverityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeverityModifierSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeverityModifierSeverity(value)
	return nil
}

// NewSeverityModifierSeverityFromValue returns a pointer to a valid SeverityModifierSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityModifierSeverityFromValue(v string) (*SeverityModifierSeverity, error) {
	ev := SeverityModifierSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeverityModifierSeverity: valid values are %v", v, allowedSeverityModifierSeverityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeverityModifierSeverity) IsValid() bool {
	for _, existing := range allowedSeverityModifierSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityModifierSeverity value.
func (v SeverityModifierSeverity) Ptr() *SeverityModifierSeverity {
	return &v
}
