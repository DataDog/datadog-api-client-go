// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarIntervalDayType Test-only generated schema for the dashboard versioning proof of concept.
type CalendarIntervalDayType string

// List of CalendarIntervalDayType.
const (
	CALENDARINTERVALDAYTYPE_DAY CalendarIntervalDayType = "day"
)

var allowedCalendarIntervalDayTypeEnumValues = []CalendarIntervalDayType{
	CALENDARINTERVALDAYTYPE_DAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CalendarIntervalDayType) GetAllowedValues() []CalendarIntervalDayType {
	return allowedCalendarIntervalDayTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CalendarIntervalDayType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CalendarIntervalDayType(value)
	return nil
}

// NewCalendarIntervalDayTypeFromValue returns a pointer to a valid CalendarIntervalDayType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCalendarIntervalDayTypeFromValue(v string) (*CalendarIntervalDayType, error) {
	ev := CalendarIntervalDayType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CalendarIntervalDayType: valid values are %v", v, allowedCalendarIntervalDayTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CalendarIntervalDayType) IsValid() bool {
	for _, existing := range allowedCalendarIntervalDayTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CalendarIntervalDayType value.
func (v CalendarIntervalDayType) Ptr() *CalendarIntervalDayType {
	return &v
}
