// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlatSortByAnyOfAnyOfItemProperty4446002432 Test-only generated schema for the dashboard versioning proof of concept.
type FlatSortByAnyOfAnyOfItemProperty4446002432 string

// List of FlatSortByAnyOfAnyOfItemProperty4446002432.
const (
	FLATSORTBYANYOFANYOFITEMPROPERTY4446002432_LABEL FlatSortByAnyOfAnyOfItemProperty4446002432 = "label"
)

var allowedFlatSortByAnyOfAnyOfItemProperty4446002432EnumValues = []FlatSortByAnyOfAnyOfItemProperty4446002432{
	FLATSORTBYANYOFANYOFITEMPROPERTY4446002432_LABEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlatSortByAnyOfAnyOfItemProperty4446002432) GetAllowedValues() []FlatSortByAnyOfAnyOfItemProperty4446002432 {
	return allowedFlatSortByAnyOfAnyOfItemProperty4446002432EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlatSortByAnyOfAnyOfItemProperty4446002432) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlatSortByAnyOfAnyOfItemProperty4446002432(value)
	return nil
}

// NewFlatSortByAnyOfAnyOfItemProperty4446002432FromValue returns a pointer to a valid FlatSortByAnyOfAnyOfItemProperty4446002432
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlatSortByAnyOfAnyOfItemProperty4446002432FromValue(v string) (*FlatSortByAnyOfAnyOfItemProperty4446002432, error) {
	ev := FlatSortByAnyOfAnyOfItemProperty4446002432(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlatSortByAnyOfAnyOfItemProperty4446002432: valid values are %v", v, allowedFlatSortByAnyOfAnyOfItemProperty4446002432EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlatSortByAnyOfAnyOfItemProperty4446002432) IsValid() bool {
	for _, existing := range allowedFlatSortByAnyOfAnyOfItemProperty4446002432EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlatSortByAnyOfAnyOfItemProperty4446002432 value.
func (v FlatSortByAnyOfAnyOfItemProperty4446002432) Ptr() *FlatSortByAnyOfAnyOfItemProperty4446002432 {
	return &v
}
