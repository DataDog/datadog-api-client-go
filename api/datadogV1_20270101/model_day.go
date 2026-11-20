// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Day Test-only generated schema for the dashboard versioning proof of concept.
type Day string

// List of Day.
const (
	DAY_MONDAY    Day = "monday"
	DAY_TUESDAY   Day = "tuesday"
	DAY_WEDNESDAY Day = "wednesday"
	DAY_THURSDAY  Day = "thursday"
	DAY_FRIDAY    Day = "friday"
	DAY_SATURDAY  Day = "saturday"
	DAY_SUNDAY    Day = "sunday"
)

var allowedDayEnumValues = []Day{
	DAY_MONDAY,
	DAY_TUESDAY,
	DAY_WEDNESDAY,
	DAY_THURSDAY,
	DAY_FRIDAY,
	DAY_SATURDAY,
	DAY_SUNDAY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Day) GetAllowedValues() []Day {
	return allowedDayEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Day) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Day(value)
	return nil
}

// NewDayFromValue returns a pointer to a valid Day
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDayFromValue(v string) (*Day, error) {
	ev := Day(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Day: valid values are %v", v, allowedDayEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Day) IsValid() bool {
	for _, existing := range allowedDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Day value.
func (v Day) Ptr() *Day {
	return &v
}
