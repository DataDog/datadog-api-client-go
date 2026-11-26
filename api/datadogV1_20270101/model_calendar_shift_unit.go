// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarShiftUnit Test-only generated schema for the dashboard versioning proof of concept.
type CalendarShiftUnit string

// List of CalendarShiftUnit.
const (
	CALENDARSHIFTUNIT_D  CalendarShiftUnit = "d"
	CALENDARSHIFTUNIT_W  CalendarShiftUnit = "w"
	CALENDARSHIFTUNIT_MO CalendarShiftUnit = "mo"
)

var allowedCalendarShiftUnitEnumValues = []CalendarShiftUnit{
	CALENDARSHIFTUNIT_D,
	CALENDARSHIFTUNIT_W,
	CALENDARSHIFTUNIT_MO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CalendarShiftUnit) GetAllowedValues() []CalendarShiftUnit {
	return allowedCalendarShiftUnitEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CalendarShiftUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CalendarShiftUnit(value)
	return nil
}

// NewCalendarShiftUnitFromValue returns a pointer to a valid CalendarShiftUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCalendarShiftUnitFromValue(v string) (*CalendarShiftUnit, error) {
	ev := CalendarShiftUnit(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CalendarShiftUnit: valid values are %v", v, allowedCalendarShiftUnitEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CalendarShiftUnit) IsValid() bool {
	for _, existing := range allowedCalendarShiftUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CalendarShiftUnit value.
func (v CalendarShiftUnit) Ptr() *CalendarShiftUnit {
	return &v
}
