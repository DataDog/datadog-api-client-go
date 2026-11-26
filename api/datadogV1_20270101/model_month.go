// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Month Test-only generated schema for the dashboard versioning proof of concept.
type Month string

// List of Month.
const (
	MONTH_JANUARY   Month = "january"
	MONTH_FEBRUARY  Month = "february"
	MONTH_MARCH     Month = "march"
	MONTH_APRIL     Month = "april"
	MONTH_MAY       Month = "may"
	MONTH_JUNE      Month = "june"
	MONTH_JULY      Month = "july"
	MONTH_AUGUST    Month = "august"
	MONTH_SEPTEMBER Month = "september"
	MONTH_OCTOBER   Month = "october"
	MONTH_NOVEMBER  Month = "november"
	MONTH_DECEMBER  Month = "december"
)

var allowedMonthEnumValues = []Month{
	MONTH_JANUARY,
	MONTH_FEBRUARY,
	MONTH_MARCH,
	MONTH_APRIL,
	MONTH_MAY,
	MONTH_JUNE,
	MONTH_JULY,
	MONTH_AUGUST,
	MONTH_SEPTEMBER,
	MONTH_OCTOBER,
	MONTH_NOVEMBER,
	MONTH_DECEMBER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Month) GetAllowedValues() []Month {
	return allowedMonthEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Month) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Month(value)
	return nil
}

// NewMonthFromValue returns a pointer to a valid Month
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonthFromValue(v string) (*Month, error) {
	ev := Month(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Month: valid values are %v", v, allowedMonthEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Month) IsValid() bool {
	for _, existing := range allowedMonthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Month value.
func (v Month) Ptr() *Month {
	return &v
}
