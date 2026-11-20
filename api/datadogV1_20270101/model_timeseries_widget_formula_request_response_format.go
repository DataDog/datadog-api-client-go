// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetFormulaRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetFormulaRequestResponseFormat string

// List of TimeseriesWidgetFormulaRequestResponseFormat.
const (
	TIMESERIESWIDGETFORMULAREQUESTRESPONSEFORMAT_TIMESERIES TimeseriesWidgetFormulaRequestResponseFormat = "timeseries"
)

var allowedTimeseriesWidgetFormulaRequestResponseFormatEnumValues = []TimeseriesWidgetFormulaRequestResponseFormat{
	TIMESERIESWIDGETFORMULAREQUESTRESPONSEFORMAT_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesWidgetFormulaRequestResponseFormat) GetAllowedValues() []TimeseriesWidgetFormulaRequestResponseFormat {
	return allowedTimeseriesWidgetFormulaRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesWidgetFormulaRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesWidgetFormulaRequestResponseFormat(value)
	return nil
}

// NewTimeseriesWidgetFormulaRequestResponseFormatFromValue returns a pointer to a valid TimeseriesWidgetFormulaRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesWidgetFormulaRequestResponseFormatFromValue(v string) (*TimeseriesWidgetFormulaRequestResponseFormat, error) {
	ev := TimeseriesWidgetFormulaRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesWidgetFormulaRequestResponseFormat: valid values are %v", v, allowedTimeseriesWidgetFormulaRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesWidgetFormulaRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedTimeseriesWidgetFormulaRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesWidgetFormulaRequestResponseFormat value.
func (v TimeseriesWidgetFormulaRequestResponseFormat) Ptr() *TimeseriesWidgetFormulaRequestResponseFormat {
	return &v
}
