// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetAdvancedQueryRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetAdvancedQueryRequestResponseFormat string

// List of TimeseriesWidgetAdvancedQueryRequestResponseFormat.
const (
	TIMESERIESWIDGETADVANCEDQUERYREQUESTRESPONSEFORMAT_TIMESERIES TimeseriesWidgetAdvancedQueryRequestResponseFormat = "timeseries"
)

var allowedTimeseriesWidgetAdvancedQueryRequestResponseFormatEnumValues = []TimeseriesWidgetAdvancedQueryRequestResponseFormat{
	TIMESERIESWIDGETADVANCEDQUERYREQUESTRESPONSEFORMAT_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetAdvancedQueryRequestResponseFormat) GetAllowedValues() []TimeseriesWidgetAdvancedQueryRequestResponseFormat {
	return allowedTimeseriesWidgetAdvancedQueryRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetAdvancedQueryRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetAdvancedQueryRequestResponseFormat(value)
	return nil
}

// NewTimeseriesWidgetAdvancedQueryRequestResponseFormatFromValue returns a pointer to a valid TimeseriesWidgetAdvancedQueryRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetAdvancedQueryRequestResponseFormatFromValue(v string) (*TimeseriesWidgetAdvancedQueryRequestResponseFormat, error) {
	ev := TimeseriesWidgetAdvancedQueryRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetAdvancedQueryRequestResponseFormat: valid values are %v", v, allowedTimeseriesWidgetAdvancedQueryRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetAdvancedQueryRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetAdvancedQueryRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetAdvancedQueryRequestResponseFormat value.
func (v TimeseriesWidgetAdvancedQueryRequestResponseFormat) Ptr() *TimeseriesWidgetAdvancedQueryRequestResponseFormat {
	return &v
}
