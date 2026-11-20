// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OccurrenceFilterOperator Test-only generated schema for the dashboard versioning proof of concept.
type OccurrenceFilterOperator string

// List of OccurrenceFilterOperator.
const (
	OCCURRENCEFILTEROPERATOR_EQ OccurrenceFilterOperator = "="
)

var allowedOccurrenceFilterOperatorEnumValues = []OccurrenceFilterOperator{
	OCCURRENCEFILTEROPERATOR_EQ,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OccurrenceFilterOperator) GetAllowedValues() []OccurrenceFilterOperator {
	return allowedOccurrenceFilterOperatorEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OccurrenceFilterOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OccurrenceFilterOperator(value)
	return nil
}

// NewOccurrenceFilterOperatorFromValue returns a pointer to a valid OccurrenceFilterOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOccurrenceFilterOperatorFromValue(v string) (*OccurrenceFilterOperator, error) {
	ev := OccurrenceFilterOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OccurrenceFilterOperator: valid values are %v", v, allowedOccurrenceFilterOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OccurrenceFilterOperator) IsValid() bool {
	for _, existing := range allowedOccurrenceFilterOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OccurrenceFilterOperator value.
func (v OccurrenceFilterOperator) Ptr() *OccurrenceFilterOperator {
	return &v
}
