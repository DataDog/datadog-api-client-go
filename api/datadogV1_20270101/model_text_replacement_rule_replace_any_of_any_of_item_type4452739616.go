// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616 Test-only generated schema for the dashboard versioning proof of concept.
type TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616 string

// List of TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616.
const (
	TEXTREPLACEMENTRULEREPLACEANYOFANYOFITEMTYPE4452739616_ALL TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616 = "all"
)

var allowedTextReplacementRuleReplaceAnyOfAnyOfItemType4452739616EnumValues = []TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616{
	TEXTREPLACEMENTRULEREPLACEANYOFANYOFITEMTYPE4452739616_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616) GetAllowedValues() []TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616 {
	return allowedTextReplacementRuleReplaceAnyOfAnyOfItemType4452739616EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616(value)
	return nil
}

// NewTextReplacementRuleReplaceAnyOfAnyOfItemType4452739616FromValue returns a pointer to a valid TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextReplacementRuleReplaceAnyOfAnyOfItemType4452739616FromValue(v string) (*TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616, error) {
	ev := TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616: valid values are %v", v, allowedTextReplacementRuleReplaceAnyOfAnyOfItemType4452739616EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616) IsValid() bool {
	for _, existing := range allowedTextReplacementRuleReplaceAnyOfAnyOfItemType4452739616EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616 value.
func (v TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616) Ptr() *TextReplacementRuleReplaceAnyOfAnyOfItemType4452739616 {
	return &v
}
