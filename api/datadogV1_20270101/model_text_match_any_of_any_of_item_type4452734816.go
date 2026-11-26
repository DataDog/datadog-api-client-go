// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextMatchAnyOfAnyOfItemType4452734816 Test-only generated schema for the dashboard versioning proof of concept.
type TextMatchAnyOfAnyOfItemType4452734816 string

// List of TextMatchAnyOfAnyOfItemType4452734816.
const (
	TEXTMATCHANYOFANYOFITEMTYPE4452734816_DOES_NOT_CONTAIN TextMatchAnyOfAnyOfItemType4452734816 = "does_not_contain"
)

var allowedTextMatchAnyOfAnyOfItemType4452734816EnumValues = []TextMatchAnyOfAnyOfItemType4452734816{
	TEXTMATCHANYOFANYOFITEMTYPE4452734816_DOES_NOT_CONTAIN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextMatchAnyOfAnyOfItemType4452734816) GetAllowedValues() []TextMatchAnyOfAnyOfItemType4452734816 {
	return allowedTextMatchAnyOfAnyOfItemType4452734816EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextMatchAnyOfAnyOfItemType4452734816) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextMatchAnyOfAnyOfItemType4452734816(value)
	return nil
}

// NewTextMatchAnyOfAnyOfItemType4452734816FromValue returns a pointer to a valid TextMatchAnyOfAnyOfItemType4452734816
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextMatchAnyOfAnyOfItemType4452734816FromValue(v string) (*TextMatchAnyOfAnyOfItemType4452734816, error) {
	ev := TextMatchAnyOfAnyOfItemType4452734816(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextMatchAnyOfAnyOfItemType4452734816: valid values are %v", v, allowedTextMatchAnyOfAnyOfItemType4452734816EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextMatchAnyOfAnyOfItemType4452734816) IsValid() bool {
	for _, existing := range allowedTextMatchAnyOfAnyOfItemType4452734816EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextMatchAnyOfAnyOfItemType4452734816 value.
func (v TextMatchAnyOfAnyOfItemType4452734816) Ptr() *TextMatchAnyOfAnyOfItemType4452734816 {
	return &v
}
