// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRuleSetActionType The type of a severity modifier rule action that sets a fixed severity.
type SeverityModifierRuleSetActionType string

// List of SeverityModifierRuleSetActionType.
const (
	SEVERITYMODIFIERRULESETACTIONTYPE_SET SeverityModifierRuleSetActionType = "set"
)

var allowedSeverityModifierRuleSetActionTypeEnumValues = []SeverityModifierRuleSetActionType{
	SEVERITYMODIFIERRULESETACTIONTYPE_SET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SeverityModifierRuleSetActionType) GetAllowedValues() []SeverityModifierRuleSetActionType {
	return allowedSeverityModifierRuleSetActionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SeverityModifierRuleSetActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SeverityModifierRuleSetActionType(value)
	return nil
}

// NewSeverityModifierRuleSetActionTypeFromValue returns a pointer to a valid SeverityModifierRuleSetActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSeverityModifierRuleSetActionTypeFromValue(v string) (*SeverityModifierRuleSetActionType, error) {
	ev := SeverityModifierRuleSetActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SeverityModifierRuleSetActionType: valid values are %v", v, allowedSeverityModifierRuleSetActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SeverityModifierRuleSetActionType) IsValid() bool {
	for _, existing := range allowedSeverityModifierRuleSetActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityModifierRuleSetActionType value.
func (v SeverityModifierRuleSetActionType) Ptr() *SeverityModifierRuleSetActionType {
	return &v
}
