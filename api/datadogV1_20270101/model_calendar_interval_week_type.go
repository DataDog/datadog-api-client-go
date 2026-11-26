// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarIntervalWeekType Test-only generated schema for the dashboard versioning proof of concept.
type CalendarIntervalWeekType string

// List of CalendarIntervalWeekType.
const (
	CALENDARINTERVALWEEKTYPE_WEEK CalendarIntervalWeekType = "week"
)

var allowedCalendarIntervalWeekTypeEnumValues = []CalendarIntervalWeekType{
	CALENDARINTERVALWEEKTYPE_WEEK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CalendarIntervalWeekType) GetAllowedValues() []CalendarIntervalWeekType {
	return allowedCalendarIntervalWeekTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CalendarIntervalWeekType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CalendarIntervalWeekType(value)
	return nil
}

// NewCalendarIntervalWeekTypeFromValue returns a pointer to a valid CalendarIntervalWeekType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCalendarIntervalWeekTypeFromValue(v string) (*CalendarIntervalWeekType, error) {
	ev := CalendarIntervalWeekType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CalendarIntervalWeekType: valid values are %v", v, allowedCalendarIntervalWeekTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CalendarIntervalWeekType) IsValid() bool {
	for _, existing := range allowedCalendarIntervalWeekTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CalendarIntervalWeekType value.
func (v CalendarIntervalWeekType) Ptr() *CalendarIntervalWeekType {
	return &v
}
