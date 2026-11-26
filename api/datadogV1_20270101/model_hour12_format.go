// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Hour12Format Test-only generated schema for the dashboard versioning proof of concept.
type Hour12Format string

// List of Hour12Format.
const (
	HOUR12FORMAT_TWELVE_AM Hour12Format = "12am"
	HOUR12FORMAT_ONE_AM    Hour12Format = "1am"
	HOUR12FORMAT_TWO_AM    Hour12Format = "2am"
	HOUR12FORMAT_THREE_AM  Hour12Format = "3am"
	HOUR12FORMAT_FOUR_AM   Hour12Format = "4am"
	HOUR12FORMAT_FIVE_AM   Hour12Format = "5am"
	HOUR12FORMAT_SIX_AM    Hour12Format = "6am"
	HOUR12FORMAT_SEVEN_AM  Hour12Format = "7am"
	HOUR12FORMAT_EIGHT_AM  Hour12Format = "8am"
	HOUR12FORMAT_NINE_AM   Hour12Format = "9am"
	HOUR12FORMAT_TEN_AM    Hour12Format = "10am"
	HOUR12FORMAT_ELEVEN_AM Hour12Format = "11am"
	HOUR12FORMAT_TWELVE_PM Hour12Format = "12pm"
	HOUR12FORMAT_ONE_PM    Hour12Format = "1pm"
	HOUR12FORMAT_TWO_PM    Hour12Format = "2pm"
	HOUR12FORMAT_THREE_PM  Hour12Format = "3pm"
	HOUR12FORMAT_FOUR_PM   Hour12Format = "4pm"
	HOUR12FORMAT_FIVE_PM   Hour12Format = "5pm"
	HOUR12FORMAT_SIX_PM    Hour12Format = "6pm"
	HOUR12FORMAT_SEVEN_PM  Hour12Format = "7pm"
	HOUR12FORMAT_EIGHT_PM  Hour12Format = "8pm"
	HOUR12FORMAT_NINE_PM   Hour12Format = "9pm"
	HOUR12FORMAT_TEN_PM    Hour12Format = "10pm"
	HOUR12FORMAT_ELEVEN_PM Hour12Format = "11pm"
)

var allowedHour12FormatEnumValues = []Hour12Format{
	HOUR12FORMAT_TWELVE_AM,
	HOUR12FORMAT_ONE_AM,
	HOUR12FORMAT_TWO_AM,
	HOUR12FORMAT_THREE_AM,
	HOUR12FORMAT_FOUR_AM,
	HOUR12FORMAT_FIVE_AM,
	HOUR12FORMAT_SIX_AM,
	HOUR12FORMAT_SEVEN_AM,
	HOUR12FORMAT_EIGHT_AM,
	HOUR12FORMAT_NINE_AM,
	HOUR12FORMAT_TEN_AM,
	HOUR12FORMAT_ELEVEN_AM,
	HOUR12FORMAT_TWELVE_PM,
	HOUR12FORMAT_ONE_PM,
	HOUR12FORMAT_TWO_PM,
	HOUR12FORMAT_THREE_PM,
	HOUR12FORMAT_FOUR_PM,
	HOUR12FORMAT_FIVE_PM,
	HOUR12FORMAT_SIX_PM,
	HOUR12FORMAT_SEVEN_PM,
	HOUR12FORMAT_EIGHT_PM,
	HOUR12FORMAT_NINE_PM,
	HOUR12FORMAT_TEN_PM,
	HOUR12FORMAT_ELEVEN_PM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *Hour12Format) GetAllowedValues() []Hour12Format {
	return allowedHour12FormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Hour12Format) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Hour12Format(value)
	return nil
}

// NewHour12FormatFromValue returns a pointer to a valid Hour12Format
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHour12FormatFromValue(v string) (*Hour12Format, error) {
	ev := Hour12Format(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Hour12Format: valid values are %v", v, allowedHour12FormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Hour12Format) IsValid() bool {
	for _, existing := range allowedHour12FormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Hour12Format value.
func (v Hour12Format) Ptr() *Hour12Format {
	return &v
}
