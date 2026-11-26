// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesXAxisDayDateFormat Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesXAxisDayDateFormat string

// List of TimeseriesXAxisDayDateFormat.
const (
	TIMESERIESXAXISDAYDATEFORMAT_MMM_D        TimeseriesXAxisDayDateFormat = "MMM d"
	TIMESERIESXAXISDAYDATEFORMAT_MMM_D__YYYY  TimeseriesXAxisDayDateFormat = "MMM d, yyyy"
	TIMESERIESXAXISDAYDATEFORMAT_MM_DD_YYYY   TimeseriesXAxisDayDateFormat = "MM/dd/yyyy"
	TIMESERIESXAXISDAYDATEFORMAT_DD_MM_YYYY   TimeseriesXAxisDayDateFormat = "dd/MM/yyyy"
	TIMESERIESXAXISDAYDATEFORMAT_MMMM_D__YYYY TimeseriesXAxisDayDateFormat = "MMMM d, yyyy"
	TIMESERIESXAXISDAYDATEFORMAT_AUTO         TimeseriesXAxisDayDateFormat = "auto"
)

var allowedTimeseriesXAxisDayDateFormatEnumValues = []TimeseriesXAxisDayDateFormat{
	TIMESERIESXAXISDAYDATEFORMAT_MMM_D,
	TIMESERIESXAXISDAYDATEFORMAT_MMM_D__YYYY,
	TIMESERIESXAXISDAYDATEFORMAT_MM_DD_YYYY,
	TIMESERIESXAXISDAYDATEFORMAT_DD_MM_YYYY,
	TIMESERIESXAXISDAYDATEFORMAT_MMMM_D__YYYY,
	TIMESERIESXAXISDAYDATEFORMAT_AUTO,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesXAxisDayDateFormat) GetAllowedValues() []TimeseriesXAxisDayDateFormat {
	return allowedTimeseriesXAxisDayDateFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesXAxisDayDateFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesXAxisDayDateFormat(value)
	return nil
}

// NewTimeseriesXAxisDayDateFormatFromValue returns a pointer to a valid TimeseriesXAxisDayDateFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesXAxisDayDateFormatFromValue(v string) (*TimeseriesXAxisDayDateFormat, error) {
	ev := TimeseriesXAxisDayDateFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesXAxisDayDateFormat: valid values are %v", v, allowedTimeseriesXAxisDayDateFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesXAxisDayDateFormat) IsValid() bool {
	for _, existing := range allowedTimeseriesXAxisDayDateFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesXAxisDayDateFormat value.
func (v TimeseriesXAxisDayDateFormat) Ptr() *TimeseriesXAxisDayDateFormat {
	return &v
}
