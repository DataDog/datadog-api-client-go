// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextMatchAnyOfAnyOfItemType4452732896 Test-only generated schema for the dashboard versioning proof of concept.
type TextMatchAnyOfAnyOfItemType4452732896 string

// List of TextMatchAnyOfAnyOfItemType4452732896.
const (
	TEXTMATCHANYOFANYOFITEMTYPE4452732896_IS_NOT TextMatchAnyOfAnyOfItemType4452732896 = "is_not"
)

var allowedTextMatchAnyOfAnyOfItemType4452732896EnumValues = []TextMatchAnyOfAnyOfItemType4452732896{
	TEXTMATCHANYOFANYOFITEMTYPE4452732896_IS_NOT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextMatchAnyOfAnyOfItemType4452732896) GetAllowedValues() []TextMatchAnyOfAnyOfItemType4452732896 {
	return allowedTextMatchAnyOfAnyOfItemType4452732896EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextMatchAnyOfAnyOfItemType4452732896) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextMatchAnyOfAnyOfItemType4452732896(value)
	return nil
}

// NewTextMatchAnyOfAnyOfItemType4452732896FromValue returns a pointer to a valid TextMatchAnyOfAnyOfItemType4452732896
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextMatchAnyOfAnyOfItemType4452732896FromValue(v string) (*TextMatchAnyOfAnyOfItemType4452732896, error) {
	ev := TextMatchAnyOfAnyOfItemType4452732896(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextMatchAnyOfAnyOfItemType4452732896: valid values are %v", v, allowedTextMatchAnyOfAnyOfItemType4452732896EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextMatchAnyOfAnyOfItemType4452732896) IsValid() bool {
	for _, existing := range allowedTextMatchAnyOfAnyOfItemType4452732896EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextMatchAnyOfAnyOfItemType4452732896 value.
func (v TextMatchAnyOfAnyOfItemType4452732896) Ptr() *TextMatchAnyOfAnyOfItemType4452732896 {
	return &v
}
