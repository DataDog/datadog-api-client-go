// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesProjectionType Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesProjectionType string

// List of TimeseriesProjectionType.
const (
	TIMESERIESPROJECTIONTYPE_TIMESERIES TimeseriesProjectionType = "timeseries"
)

var allowedTimeseriesProjectionTypeEnumValues = []TimeseriesProjectionType{
	TIMESERIESPROJECTIONTYPE_TIMESERIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesProjectionType) GetAllowedValues() []TimeseriesProjectionType {
	return allowedTimeseriesProjectionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesProjectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesProjectionType(value)
	return nil
}

// NewTimeseriesProjectionTypeFromValue returns a pointer to a valid TimeseriesProjectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesProjectionTypeFromValue(v string) (*TimeseriesProjectionType, error) {
	ev := TimeseriesProjectionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesProjectionType: valid values are %v", v, allowedTimeseriesProjectionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesProjectionType) IsValid() bool {
	for _, existing := range allowedTimeseriesProjectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesProjectionType value.
func (v TimeseriesProjectionType) Ptr() *TimeseriesProjectionType {
	return &v
}
