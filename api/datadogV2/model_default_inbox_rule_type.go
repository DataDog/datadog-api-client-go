// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DefaultInboxRuleType The JSON:API type for default inbox rules.
type DefaultInboxRuleType string

// List of DefaultInboxRuleType.
const (
	DEFAULTINBOXRULETYPE_DEFAULT_INBOX_RULES DefaultInboxRuleType = "default_inbox_rules"
)

var allowedDefaultInboxRuleTypeEnumValues = []DefaultInboxRuleType{
	DEFAULTINBOXRULETYPE_DEFAULT_INBOX_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DefaultInboxRuleType) GetAllowedValues() []DefaultInboxRuleType {
	return allowedDefaultInboxRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DefaultInboxRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DefaultInboxRuleType(value)
	return nil
}

// NewDefaultInboxRuleTypeFromValue returns a pointer to a valid DefaultInboxRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDefaultInboxRuleTypeFromValue(v string) (*DefaultInboxRuleType, error) {
	ev := DefaultInboxRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DefaultInboxRuleType: valid values are %v", v, allowedDefaultInboxRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DefaultInboxRuleType) IsValid() bool {
	for _, existing := range allowedDefaultInboxRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DefaultInboxRuleType value.
func (v DefaultInboxRuleType) Ptr() *DefaultInboxRuleType {
	return &v
}
