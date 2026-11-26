// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlatSortByAnyOfAnyOfItemProperty Test-only generated schema for the dashboard versioning proof of concept.
type FlatSortByAnyOfAnyOfItemProperty string

// List of FlatSortByAnyOfAnyOfItemProperty.
const (
	FLATSORTBYANYOFANYOFITEMPROPERTY_VALUE FlatSortByAnyOfAnyOfItemProperty = "value"
)

var allowedFlatSortByAnyOfAnyOfItemPropertyEnumValues = []FlatSortByAnyOfAnyOfItemProperty{
	FLATSORTBYANYOFANYOFITEMPROPERTY_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlatSortByAnyOfAnyOfItemProperty) GetAllowedValues() []FlatSortByAnyOfAnyOfItemProperty {
	return allowedFlatSortByAnyOfAnyOfItemPropertyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlatSortByAnyOfAnyOfItemProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlatSortByAnyOfAnyOfItemProperty(value)
	return nil
}

// NewFlatSortByAnyOfAnyOfItemPropertyFromValue returns a pointer to a valid FlatSortByAnyOfAnyOfItemProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlatSortByAnyOfAnyOfItemPropertyFromValue(v string) (*FlatSortByAnyOfAnyOfItemProperty, error) {
	ev := FlatSortByAnyOfAnyOfItemProperty(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlatSortByAnyOfAnyOfItemProperty: valid values are %v", v, allowedFlatSortByAnyOfAnyOfItemPropertyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlatSortByAnyOfAnyOfItemProperty) IsValid() bool {
	for _, existing := range allowedFlatSortByAnyOfAnyOfItemPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlatSortByAnyOfAnyOfItemProperty value.
func (v FlatSortByAnyOfAnyOfItemProperty) Ptr() *FlatSortByAnyOfAnyOfItemProperty {
	return &v
}
