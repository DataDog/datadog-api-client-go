// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016 Test-only generated schema for the dashboard versioning proof of concept.
type TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016 string

// List of TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016.
const (
	TEXTFORMATTINGRULEREPLACEANYOFANYOFITEMTYPE4452730016_ALL TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016 = "all"
)

var allowedTextFormattingRuleReplaceAnyOfAnyOfItemType4452730016EnumValues = []TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016{
	TEXTFORMATTINGRULEREPLACEANYOFANYOFITEMTYPE4452730016_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016) GetAllowedValues() []TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016 {
	return allowedTextFormattingRuleReplaceAnyOfAnyOfItemType4452730016EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016(value)
	return nil
}

// NewTextFormattingRuleReplaceAnyOfAnyOfItemType4452730016FromValue returns a pointer to a valid TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextFormattingRuleReplaceAnyOfAnyOfItemType4452730016FromValue(v string) (*TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016, error) {
	ev := TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016: valid values are %v", v, allowedTextFormattingRuleReplaceAnyOfAnyOfItemType4452730016EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016) IsValid() bool {
	for _, existing := range allowedTextFormattingRuleReplaceAnyOfAnyOfItemType4452730016EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016 value.
func (v TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016) Ptr() *TextFormattingRuleReplaceAnyOfAnyOfItemType4452730016 {
	return &v
}
