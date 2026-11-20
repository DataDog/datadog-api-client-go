// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextFormattingRuleReplaceAnyOfAnyOfItemType Test-only generated schema for the dashboard versioning proof of concept.
type TextFormattingRuleReplaceAnyOfAnyOfItemType string

// List of TextFormattingRuleReplaceAnyOfAnyOfItemType.
const (
	TEXTFORMATTINGRULEREPLACEANYOFANYOFITEMTYPE_SUBSTRING TextFormattingRuleReplaceAnyOfAnyOfItemType = "substring"
)

var allowedTextFormattingRuleReplaceAnyOfAnyOfItemTypeEnumValues = []TextFormattingRuleReplaceAnyOfAnyOfItemType{
	TEXTFORMATTINGRULEREPLACEANYOFANYOFITEMTYPE_SUBSTRING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextFormattingRuleReplaceAnyOfAnyOfItemType) GetAllowedValues() []TextFormattingRuleReplaceAnyOfAnyOfItemType {
	return allowedTextFormattingRuleReplaceAnyOfAnyOfItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextFormattingRuleReplaceAnyOfAnyOfItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextFormattingRuleReplaceAnyOfAnyOfItemType(value)
	return nil
}

// NewTextFormattingRuleReplaceAnyOfAnyOfItemTypeFromValue returns a pointer to a valid TextFormattingRuleReplaceAnyOfAnyOfItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextFormattingRuleReplaceAnyOfAnyOfItemTypeFromValue(v string) (*TextFormattingRuleReplaceAnyOfAnyOfItemType, error) {
	ev := TextFormattingRuleReplaceAnyOfAnyOfItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextFormattingRuleReplaceAnyOfAnyOfItemType: valid values are %v", v, allowedTextFormattingRuleReplaceAnyOfAnyOfItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextFormattingRuleReplaceAnyOfAnyOfItemType) IsValid() bool {
	for _, existing := range allowedTextFormattingRuleReplaceAnyOfAnyOfItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextFormattingRuleReplaceAnyOfAnyOfItemType value.
func (v TextFormattingRuleReplaceAnyOfAnyOfItemType) Ptr() *TextFormattingRuleReplaceAnyOfAnyOfItemType {
	return &v
}
