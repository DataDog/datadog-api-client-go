// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeSpan Test-only generated schema for the dashboard versioning proof of concept.
type TimeSpan string

// List of TimeSpan.
const (
	TIMESPAN_ONE_M     TimeSpan = "1m"
	TIMESPAN_FIVE_M    TimeSpan = "5m"
	TIMESPAN_TEN_M     TimeSpan = "10m"
	TIMESPAN_FIFTEEN_M TimeSpan = "15m"
	TIMESPAN_THIRTY_M  TimeSpan = "30m"
	TIMESPAN_ONE_H     TimeSpan = "1h"
	TIMESPAN_FOUR_H    TimeSpan = "4h"
	TIMESPAN_ONE_D     TimeSpan = "1d"
	TIMESPAN_TWO_D     TimeSpan = "2d"
	TIMESPAN_ONE_W     TimeSpan = "1w"
	TIMESPAN_ONE_MO    TimeSpan = "1mo"
	TIMESPAN_THREE_MO  TimeSpan = "3mo"
	TIMESPAN_SIX_MO    TimeSpan = "6mo"
	TIMESPAN_ONE_Y     TimeSpan = "1y"
)

var allowedTimeSpanEnumValues = []TimeSpan{
	TIMESPAN_ONE_M,
	TIMESPAN_FIVE_M,
	TIMESPAN_TEN_M,
	TIMESPAN_FIFTEEN_M,
	TIMESPAN_THIRTY_M,
	TIMESPAN_ONE_H,
	TIMESPAN_FOUR_H,
	TIMESPAN_ONE_D,
	TIMESPAN_TWO_D,
	TIMESPAN_ONE_W,
	TIMESPAN_ONE_MO,
	TIMESPAN_THREE_MO,
	TIMESPAN_SIX_MO,
	TIMESPAN_ONE_Y,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeSpan) GetAllowedValues() []TimeSpan {
	return allowedTimeSpanEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeSpan) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeSpan(value)
	return nil
}

// NewTimeSpanFromValue returns a pointer to a valid TimeSpan
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeSpanFromValue(v string) (*TimeSpan, error) {
	ev := TimeSpan(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeSpan: valid values are %v", v, allowedTimeSpanEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeSpan) IsValid() bool {
	for _, existing := range allowedTimeSpanEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeSpan value.
func (v TimeSpan) Ptr() *TimeSpan {
	return &v
}
