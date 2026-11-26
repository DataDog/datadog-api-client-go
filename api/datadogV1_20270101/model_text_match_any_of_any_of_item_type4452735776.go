// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextMatchAnyOfAnyOfItemType4452735776 Test-only generated schema for the dashboard versioning proof of concept.
type TextMatchAnyOfAnyOfItemType4452735776 string

// List of TextMatchAnyOfAnyOfItemType4452735776.
const (
	TEXTMATCHANYOFANYOFITEMTYPE4452735776_STARTS_WITH TextMatchAnyOfAnyOfItemType4452735776 = "starts_with"
)

var allowedTextMatchAnyOfAnyOfItemType4452735776EnumValues = []TextMatchAnyOfAnyOfItemType4452735776{
	TEXTMATCHANYOFANYOFITEMTYPE4452735776_STARTS_WITH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextMatchAnyOfAnyOfItemType4452735776) GetAllowedValues() []TextMatchAnyOfAnyOfItemType4452735776 {
	return allowedTextMatchAnyOfAnyOfItemType4452735776EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextMatchAnyOfAnyOfItemType4452735776) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextMatchAnyOfAnyOfItemType4452735776(value)
	return nil
}

// NewTextMatchAnyOfAnyOfItemType4452735776FromValue returns a pointer to a valid TextMatchAnyOfAnyOfItemType4452735776
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextMatchAnyOfAnyOfItemType4452735776FromValue(v string) (*TextMatchAnyOfAnyOfItemType4452735776, error) {
	ev := TextMatchAnyOfAnyOfItemType4452735776(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextMatchAnyOfAnyOfItemType4452735776: valid values are %v", v, allowedTextMatchAnyOfAnyOfItemType4452735776EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextMatchAnyOfAnyOfItemType4452735776) IsValid() bool {
	for _, existing := range allowedTextMatchAnyOfAnyOfItemType4452735776EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextMatchAnyOfAnyOfItemType4452735776 value.
func (v TextMatchAnyOfAnyOfItemType4452735776) Ptr() *TextMatchAnyOfAnyOfItemType4452735776 {
	return &v
}
