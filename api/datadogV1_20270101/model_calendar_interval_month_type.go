// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalendarIntervalMonthType Test-only generated schema for the dashboard versioning proof of concept.
type CalendarIntervalMonthType string

// List of CalendarIntervalMonthType.
const (
	CALENDARINTERVALMONTHTYPE_MONTH CalendarIntervalMonthType = "month"
)

var allowedCalendarIntervalMonthTypeEnumValues = []CalendarIntervalMonthType{
	CALENDARINTERVALMONTHTYPE_MONTH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CalendarIntervalMonthType) GetAllowedValues() []CalendarIntervalMonthType {
	return allowedCalendarIntervalMonthTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CalendarIntervalMonthType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CalendarIntervalMonthType(value)
	return nil
}

// NewCalendarIntervalMonthTypeFromValue returns a pointer to a valid CalendarIntervalMonthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCalendarIntervalMonthTypeFromValue(v string) (*CalendarIntervalMonthType, error) {
	ev := CalendarIntervalMonthType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CalendarIntervalMonthType: valid values are %v", v, allowedCalendarIntervalMonthTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CalendarIntervalMonthType) IsValid() bool {
	for _, existing := range allowedCalendarIntervalMonthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CalendarIntervalMonthType value.
func (v CalendarIntervalMonthType) Ptr() *CalendarIntervalMonthType {
	return &v
}
