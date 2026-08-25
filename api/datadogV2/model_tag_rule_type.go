// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleType How the rule is enforced. `blocking` rejects telemetry that violates the rule.
// `surfacing` only highlights non-compliant telemetry without blocking it.
type TagRuleType string

// List of TagRuleType.
const (
	TAGRULETYPE_BLOCKING  TagRuleType = "blocking"
	TAGRULETYPE_SURFACING TagRuleType = "surfacing"
)

var allowedTagRuleTypeEnumValues = []TagRuleType{
	TAGRULETYPE_BLOCKING,
	TAGRULETYPE_SURFACING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagRuleType) GetAllowedValues() []TagRuleType {
	return allowedTagRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagRuleType(value)
	return nil
}

// NewTagRuleTypeFromValue returns a pointer to a valid TagRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagRuleTypeFromValue(v string) (*TagRuleType, error) {
	ev := TagRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagRuleType: valid values are %v", v, allowedTagRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagRuleType) IsValid() bool {
	for _, existing := range allowedTagRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagRuleType value.
func (v TagRuleType) Ptr() *TagRuleType {
	return &v
}
