// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetDdsqlRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetDdsqlRequestResponseFormat string

// List of TimeseriesWidgetDdsqlRequestResponseFormat.
const (
	TIMESERIESWIDGETDDSQLREQUESTRESPONSEFORMAT_TIMESERIES TimeseriesWidgetDdsqlRequestResponseFormat = "timeseries"
)

var allowedTimeseriesWidgetDdsqlRequestResponseFormatEnumValues = []TimeseriesWidgetDdsqlRequestResponseFormat{
	TIMESERIESWIDGETDDSQLREQUESTRESPONSEFORMAT_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetDdsqlRequestResponseFormat) GetAllowedValues() []TimeseriesWidgetDdsqlRequestResponseFormat {
	return allowedTimeseriesWidgetDdsqlRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetDdsqlRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetDdsqlRequestResponseFormat(value)
	return nil
}

// NewTimeseriesWidgetDdsqlRequestResponseFormatFromValue returns a pointer to a valid TimeseriesWidgetDdsqlRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetDdsqlRequestResponseFormatFromValue(v string) (*TimeseriesWidgetDdsqlRequestResponseFormat, error) {
	ev := TimeseriesWidgetDdsqlRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetDdsqlRequestResponseFormat: valid values are %v", v, allowedTimeseriesWidgetDdsqlRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetDdsqlRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetDdsqlRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetDdsqlRequestResponseFormat value.
func (v TimeseriesWidgetDdsqlRequestResponseFormat) Ptr() *TimeseriesWidgetDdsqlRequestResponseFormat {
	return &v
}
