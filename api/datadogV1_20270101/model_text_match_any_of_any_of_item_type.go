// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TextMatchAnyOfAnyOfItemType Test-only generated schema for the dashboard versioning proof of concept.
type TextMatchAnyOfAnyOfItemType string

// List of TextMatchAnyOfAnyOfItemType.
const (
	TEXTMATCHANYOFANYOFITEMTYPE_IS TextMatchAnyOfAnyOfItemType = "is"
)

var allowedTextMatchAnyOfAnyOfItemTypeEnumValues = []TextMatchAnyOfAnyOfItemType{
	TEXTMATCHANYOFANYOFITEMTYPE_IS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TextMatchAnyOfAnyOfItemType) GetAllowedValues() []TextMatchAnyOfAnyOfItemType {
	return allowedTextMatchAnyOfAnyOfItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TextMatchAnyOfAnyOfItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TextMatchAnyOfAnyOfItemType(value)
	return nil
}

// NewTextMatchAnyOfAnyOfItemTypeFromValue returns a pointer to a valid TextMatchAnyOfAnyOfItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTextMatchAnyOfAnyOfItemTypeFromValue(v string) (*TextMatchAnyOfAnyOfItemType, error) {
	ev := TextMatchAnyOfAnyOfItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TextMatchAnyOfAnyOfItemType: valid values are %v", v, allowedTextMatchAnyOfAnyOfItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TextMatchAnyOfAnyOfItemType) IsValid() bool {
	for _, existing := range allowedTextMatchAnyOfAnyOfItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TextMatchAnyOfAnyOfItemType value.
func (v TextMatchAnyOfAnyOfItemType) Ptr() *TextMatchAnyOfAnyOfItemType {
	return &v
}
