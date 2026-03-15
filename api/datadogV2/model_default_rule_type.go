// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DefaultRuleType The JSON:API type for default rules.
type DefaultRuleType string

// List of DefaultRuleType.
const (
	DEFAULTRULETYPE_DEFAULT_RULE DefaultRuleType = "default-rule"
)

var allowedDefaultRuleTypeEnumValues = []DefaultRuleType{
	DEFAULTRULETYPE_DEFAULT_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DefaultRuleType) GetAllowedValues() []DefaultRuleType {
	return allowedDefaultRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DefaultRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DefaultRuleType(value)
	return nil
}

// NewDefaultRuleTypeFromValue returns a pointer to a valid DefaultRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDefaultRuleTypeFromValue(v string) (*DefaultRuleType, error) {
	ev := DefaultRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DefaultRuleType: valid values are %v", v, allowedDefaultRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DefaultRuleType) IsValid() bool {
	for _, existing := range allowedDefaultRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DefaultRuleType value.
func (v DefaultRuleType) Ptr() *DefaultRuleType {
	return &v
}
