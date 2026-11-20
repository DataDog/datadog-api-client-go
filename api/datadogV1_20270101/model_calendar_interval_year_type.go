// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarIntervalYearType Test-only generated schema for the dashboard versioning proof of concept.
type CalendarIntervalYearType string

// List of CalendarIntervalYearType.
const (
	CALENDARINTERVALYEARTYPE_YEAR CalendarIntervalYearType = "year"
)

var allowedCalendarIntervalYearTypeEnumValues = []CalendarIntervalYearType{
	CALENDARINTERVALYEARTYPE_YEAR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CalendarIntervalYearType) GetAllowedValues() []CalendarIntervalYearType {
	return allowedCalendarIntervalYearTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CalendarIntervalYearType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CalendarIntervalYearType(value)
	return nil
}

// NewCalendarIntervalYearTypeFromValue returns a pointer to a valid CalendarIntervalYearType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCalendarIntervalYearTypeFromValue(v string) (*CalendarIntervalYearType, error) {
	ev := CalendarIntervalYearType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CalendarIntervalYearType: valid values are %v", v, allowedCalendarIntervalYearTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CalendarIntervalYearType) IsValid() bool {
	for _, existing := range allowedCalendarIntervalYearTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CalendarIntervalYearType value.
func (v CalendarIntervalYearType) Ptr() *CalendarIntervalYearType {
	return &v
}
