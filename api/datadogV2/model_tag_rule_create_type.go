// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleCreateType The rule type allowed when creating a tag rule. Only `surfacing` is accepted at
// creation time.
type TagRuleCreateType string

// List of TagRuleCreateType.
const (
	TAGRULECREATETYPE_SURFACING TagRuleCreateType = "surfacing"
)

var allowedTagRuleCreateTypeEnumValues = []TagRuleCreateType{
	TAGRULECREATETYPE_SURFACING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagRuleCreateType) GetAllowedValues() []TagRuleCreateType {
	return allowedTagRuleCreateTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagRuleCreateType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagRuleCreateType(value)
	return nil
}

// NewTagRuleCreateTypeFromValue returns a pointer to a valid TagRuleCreateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagRuleCreateTypeFromValue(v string) (*TagRuleCreateType, error) {
	ev := TagRuleCreateType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagRuleCreateType: valid values are %v", v, allowedTagRuleCreateTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagRuleCreateType) IsValid() bool {
	for _, existing := range allowedTagRuleCreateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagRuleCreateType value.
func (v TagRuleCreateType) Ptr() *TagRuleCreateType {
	return &v
}
