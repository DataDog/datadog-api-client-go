// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OccurrenceFilterValue Test-only generated schema for the dashboard versioning proof of concept.
type OccurrenceFilterValue string

// List of OccurrenceFilterValue.
const (
	OCCURRENCEFILTERVALUE_ONE   OccurrenceFilterValue = "1"
	OCCURRENCEFILTERVALUE_TWO   OccurrenceFilterValue = "2"
	OCCURRENCEFILTERVALUE_THREE OccurrenceFilterValue = "3"
	OCCURRENCEFILTERVALUE_FOUR  OccurrenceFilterValue = "4"
	OCCURRENCEFILTERVALUE_FIVE  OccurrenceFilterValue = "5"
)

var allowedOccurrenceFilterValueEnumValues = []OccurrenceFilterValue{
	OCCURRENCEFILTERVALUE_ONE,
	OCCURRENCEFILTERVALUE_TWO,
	OCCURRENCEFILTERVALUE_THREE,
	OCCURRENCEFILTERVALUE_FOUR,
	OCCURRENCEFILTERVALUE_FIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OccurrenceFilterValue) GetAllowedValues() []OccurrenceFilterValue {
	return allowedOccurrenceFilterValueEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OccurrenceFilterValue) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OccurrenceFilterValue(value)
	return nil
}

// NewOccurrenceFilterValueFromValue returns a pointer to a valid OccurrenceFilterValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOccurrenceFilterValueFromValue(v string) (*OccurrenceFilterValue, error) {
	ev := OccurrenceFilterValue(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OccurrenceFilterValue: valid values are %v", v, allowedOccurrenceFilterValueEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OccurrenceFilterValue) IsValid() bool {
	for _, existing := range allowedOccurrenceFilterValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OccurrenceFilterValue value.
func (v OccurrenceFilterValue) Ptr() *OccurrenceFilterValue {
	return &v
}
