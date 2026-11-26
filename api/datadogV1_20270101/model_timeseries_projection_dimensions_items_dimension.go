// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesProjectionDimensionsItemsDimension Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesProjectionDimensionsItemsDimension string

// List of TimeseriesProjectionDimensionsItemsDimension.
const (
	TIMESERIESPROJECTIONDIMENSIONSITEMSDIMENSION_GROUP TimeseriesProjectionDimensionsItemsDimension = "group"
	TIMESERIESPROJECTIONDIMENSIONSITEMSDIMENSION_TIME  TimeseriesProjectionDimensionsItemsDimension = "time"
	TIMESERIESPROJECTIONDIMENSIONSITEMSDIMENSION_VALUE TimeseriesProjectionDimensionsItemsDimension = "value"
)

var allowedTimeseriesProjectionDimensionsItemsDimensionEnumValues = []TimeseriesProjectionDimensionsItemsDimension{
	TIMESERIESPROJECTIONDIMENSIONSITEMSDIMENSION_GROUP,
	TIMESERIESPROJECTIONDIMENSIONSITEMSDIMENSION_TIME,
	TIMESERIESPROJECTIONDIMENSIONSITEMSDIMENSION_VALUE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesProjectionDimensionsItemsDimension) GetAllowedValues() []TimeseriesProjectionDimensionsItemsDimension {
	return allowedTimeseriesProjectionDimensionsItemsDimensionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesProjectionDimensionsItemsDimension) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesProjectionDimensionsItemsDimension(value)
	return nil
}

// NewTimeseriesProjectionDimensionsItemsDimensionFromValue returns a pointer to a valid TimeseriesProjectionDimensionsItemsDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesProjectionDimensionsItemsDimensionFromValue(v string) (*TimeseriesProjectionDimensionsItemsDimension, error) {
	ev := TimeseriesProjectionDimensionsItemsDimension(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesProjectionDimensionsItemsDimension: valid values are %v", v, allowedTimeseriesProjectionDimensionsItemsDimensionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesProjectionDimensionsItemsDimension) IsValid() bool {
	for _, existing := range allowedTimeseriesProjectionDimensionsItemsDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesProjectionDimensionsItemsDimension value.
func (v TimeseriesProjectionDimensionsItemsDimension) Ptr() *TimeseriesProjectionDimensionsItemsDimension {
	return &v
}
