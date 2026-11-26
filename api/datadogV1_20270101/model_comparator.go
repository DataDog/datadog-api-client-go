// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Comparator Test-only generated schema for the dashboard versioning proof of concept.
type Comparator string

// List of Comparator.
const (
	COMPARATOR_EQ    Comparator = "="
	COMPARATOR_GT    Comparator = ">"
	COMPARATOR_GT_EQ Comparator = ">="
	COMPARATOR_LT    Comparator = "<"
	COMPARATOR_LT_EQ Comparator = "<="
)

var allowedComparatorEnumValues = []Comparator{
	COMPARATOR_EQ,
	COMPARATOR_GT,
	COMPARATOR_GT_EQ,
	COMPARATOR_LT,
	COMPARATOR_LT_EQ,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Comparator) GetAllowedValues() []Comparator {
	return allowedComparatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Comparator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Comparator(value)
	return nil
}

// NewComparatorFromValue returns a pointer to a valid Comparator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewComparatorFromValue(v string) (*Comparator, error) {
	ev := Comparator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Comparator: valid values are %v", v, allowedComparatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Comparator) IsValid() bool {
	for _, existing := range allowedComparatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Comparator value.
func (v Comparator) Ptr() *Comparator {
	return &v
}
