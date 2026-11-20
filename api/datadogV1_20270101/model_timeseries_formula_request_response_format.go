// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesFormulaRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesFormulaRequestResponseFormat string

// List of TimeseriesFormulaRequestResponseFormat.
const (
	TIMESERIESFORMULAREQUESTRESPONSEFORMAT_TIMESERIES TimeseriesFormulaRequestResponseFormat = "timeseries"
)

var allowedTimeseriesFormulaRequestResponseFormatEnumValues = []TimeseriesFormulaRequestResponseFormat{
	TIMESERIESFORMULAREQUESTRESPONSEFORMAT_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesFormulaRequestResponseFormat) GetAllowedValues() []TimeseriesFormulaRequestResponseFormat {
	return allowedTimeseriesFormulaRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesFormulaRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesFormulaRequestResponseFormat(value)
	return nil
}

// NewTimeseriesFormulaRequestResponseFormatFromValue returns a pointer to a valid TimeseriesFormulaRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesFormulaRequestResponseFormatFromValue(v string) (*TimeseriesFormulaRequestResponseFormat, error) {
	ev := TimeseriesFormulaRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesFormulaRequestResponseFormat: valid values are %v", v, allowedTimeseriesFormulaRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesFormulaRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedTimeseriesFormulaRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesFormulaRequestResponseFormat value.
func (v TimeseriesFormulaRequestResponseFormat) Ptr() *TimeseriesFormulaRequestResponseFormat {
	return &v
}
