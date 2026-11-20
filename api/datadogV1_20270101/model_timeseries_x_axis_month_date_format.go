// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesXAxisMonthDateFormat Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesXAxisMonthDateFormat string

// List of TimeseriesXAxisMonthDateFormat.
const (
	TIMESERIESXAXISMONTHDATEFORMAT_AUTO      TimeseriesXAxisMonthDateFormat = "auto"
	TIMESERIESXAXISMONTHDATEFORMAT_MMM       TimeseriesXAxisMonthDateFormat = "MMM"
	TIMESERIESXAXISMONTHDATEFORMAT_MMM_YYYY  TimeseriesXAxisMonthDateFormat = "MMM yyyy"
	TIMESERIESXAXISMONTHDATEFORMAT_MM_YYYY   TimeseriesXAxisMonthDateFormat = "MM/yyyy"
	TIMESERIESXAXISMONTHDATEFORMAT_MMMM_YYYY TimeseriesXAxisMonthDateFormat = "MMMM yyyy"
)

var allowedTimeseriesXAxisMonthDateFormatEnumValues = []TimeseriesXAxisMonthDateFormat{
	TIMESERIESXAXISMONTHDATEFORMAT_AUTO,
	TIMESERIESXAXISMONTHDATEFORMAT_MMM,
	TIMESERIESXAXISMONTHDATEFORMAT_MMM_YYYY,
	TIMESERIESXAXISMONTHDATEFORMAT_MM_YYYY,
	TIMESERIESXAXISMONTHDATEFORMAT_MMMM_YYYY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesXAxisMonthDateFormat) GetAllowedValues() []TimeseriesXAxisMonthDateFormat {
	return allowedTimeseriesXAxisMonthDateFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesXAxisMonthDateFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesXAxisMonthDateFormat(value)
	return nil
}

// NewTimeseriesXAxisMonthDateFormatFromValue returns a pointer to a valid TimeseriesXAxisMonthDateFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesXAxisMonthDateFormatFromValue(v string) (*TimeseriesXAxisMonthDateFormat, error) {
	ev := TimeseriesXAxisMonthDateFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesXAxisMonthDateFormat: valid values are %v", v, allowedTimeseriesXAxisMonthDateFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesXAxisMonthDateFormat) IsValid() bool {
	for _, existing := range allowedTimeseriesXAxisMonthDateFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesXAxisMonthDateFormat value.
func (v TimeseriesXAxisMonthDateFormat) Ptr() *TimeseriesXAxisMonthDateFormat {
	return &v
}
