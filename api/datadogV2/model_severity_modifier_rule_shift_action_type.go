// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRuleShiftActionType The type of a severity modifier rule action that shifts the severity by one rank.
type SeverityModifierRuleShiftActionType string

// List of SeverityModifierRuleShiftActionType.
const (
	SEVERITYMODIFIERRULESHIFTACTIONTYPE_SHIFT SeverityModifierRuleShiftActionType = "shift"
)

var allowedSeverityModifierRuleShiftActionTypeEnumValues = []SeverityModifierRuleShiftActionType{
	SEVERITYMODIFIERRULESHIFTACTIONTYPE_SHIFT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeverityModifierRuleShiftActionType) GetAllowedValues() []SeverityModifierRuleShiftActionType {
	return allowedSeverityModifierRuleShiftActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeverityModifierRuleShiftActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeverityModifierRuleShiftActionType(value)
	return nil
}

// NewSeverityModifierRuleShiftActionTypeFromValue returns a pointer to a valid SeverityModifierRuleShiftActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityModifierRuleShiftActionTypeFromValue(v string) (*SeverityModifierRuleShiftActionType, error) {
	ev := SeverityModifierRuleShiftActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeverityModifierRuleShiftActionType: valid values are %v", v, allowedSeverityModifierRuleShiftActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeverityModifierRuleShiftActionType) IsValid() bool {
	for _, existing := range allowedSeverityModifierRuleShiftActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityModifierRuleShiftActionType value.
func (v SeverityModifierRuleShiftActionType) Ptr() *SeverityModifierRuleShiftActionType {
	return &v
}
