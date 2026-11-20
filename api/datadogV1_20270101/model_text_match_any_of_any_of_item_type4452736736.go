// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextMatchAnyOfAnyOfItemType4452736736 Test-only generated schema for the dashboard versioning proof of concept.
type TextMatchAnyOfAnyOfItemType4452736736 string

// List of TextMatchAnyOfAnyOfItemType4452736736.
const (
	TEXTMATCHANYOFANYOFITEMTYPE4452736736_ENDS_WITH TextMatchAnyOfAnyOfItemType4452736736 = "ends_with"
)

var allowedTextMatchAnyOfAnyOfItemType4452736736EnumValues = []TextMatchAnyOfAnyOfItemType4452736736{
	TEXTMATCHANYOFANYOFITEMTYPE4452736736_ENDS_WITH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextMatchAnyOfAnyOfItemType4452736736) GetAllowedValues() []TextMatchAnyOfAnyOfItemType4452736736 {
	return allowedTextMatchAnyOfAnyOfItemType4452736736EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextMatchAnyOfAnyOfItemType4452736736) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextMatchAnyOfAnyOfItemType4452736736(value)
	return nil
}

// NewTextMatchAnyOfAnyOfItemType4452736736FromValue returns a pointer to a valid TextMatchAnyOfAnyOfItemType4452736736
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextMatchAnyOfAnyOfItemType4452736736FromValue(v string) (*TextMatchAnyOfAnyOfItemType4452736736, error) {
	ev := TextMatchAnyOfAnyOfItemType4452736736(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextMatchAnyOfAnyOfItemType4452736736: valid values are %v", v, allowedTextMatchAnyOfAnyOfItemType4452736736EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextMatchAnyOfAnyOfItemType4452736736) IsValid() bool {
	for _, existing := range allowedTextMatchAnyOfAnyOfItemType4452736736EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextMatchAnyOfAnyOfItemType4452736736 value.
func (v TextMatchAnyOfAnyOfItemType4452736736) Ptr() *TextMatchAnyOfAnyOfItemType4452736736 {
	return &v
}
