// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextMatchAnyOfAnyOfItemType4452733856 Test-only generated schema for the dashboard versioning proof of concept.
type TextMatchAnyOfAnyOfItemType4452733856 string

// List of TextMatchAnyOfAnyOfItemType4452733856.
const (
	TEXTMATCHANYOFANYOFITEMTYPE4452733856_CONTAINS TextMatchAnyOfAnyOfItemType4452733856 = "contains"
)

var allowedTextMatchAnyOfAnyOfItemType4452733856EnumValues = []TextMatchAnyOfAnyOfItemType4452733856{
	TEXTMATCHANYOFANYOFITEMTYPE4452733856_CONTAINS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextMatchAnyOfAnyOfItemType4452733856) GetAllowedValues() []TextMatchAnyOfAnyOfItemType4452733856 {
	return allowedTextMatchAnyOfAnyOfItemType4452733856EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextMatchAnyOfAnyOfItemType4452733856) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextMatchAnyOfAnyOfItemType4452733856(value)
	return nil
}

// NewTextMatchAnyOfAnyOfItemType4452733856FromValue returns a pointer to a valid TextMatchAnyOfAnyOfItemType4452733856
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextMatchAnyOfAnyOfItemType4452733856FromValue(v string) (*TextMatchAnyOfAnyOfItemType4452733856, error) {
	ev := TextMatchAnyOfAnyOfItemType4452733856(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextMatchAnyOfAnyOfItemType4452733856: valid values are %v", v, allowedTextMatchAnyOfAnyOfItemType4452733856EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextMatchAnyOfAnyOfItemType4452733856) IsValid() bool {
	for _, existing := range allowedTextMatchAnyOfAnyOfItemType4452733856EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextMatchAnyOfAnyOfItemType4452733856 value.
func (v TextMatchAnyOfAnyOfItemType4452733856) Ptr() *TextMatchAnyOfAnyOfItemType4452733856 {
	return &v
}
