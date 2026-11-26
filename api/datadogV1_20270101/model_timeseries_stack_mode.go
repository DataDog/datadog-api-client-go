// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesStackMode Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesStackMode string

// List of TimeseriesStackMode.
const (
	TIMESERIESSTACKMODE_AUTO      TimeseriesStackMode = "auto"
	TIMESERIESSTACKMODE_NORMALIZE TimeseriesStackMode = "normalize"
)

var allowedTimeseriesStackModeEnumValues = []TimeseriesStackMode{
	TIMESERIESSTACKMODE_AUTO,
	TIMESERIESSTACKMODE_NORMALIZE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimeseriesStackMode) GetAllowedValues() []TimeseriesStackMode {
	return allowedTimeseriesStackModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimeseriesStackMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimeseriesStackMode(value)
	return nil
}

// NewTimeseriesStackModeFromValue returns a pointer to a valid TimeseriesStackMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimeseriesStackModeFromValue(v string) (*TimeseriesStackMode, error) {
	ev := TimeseriesStackMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimeseriesStackMode: valid values are %v", v, allowedTimeseriesStackModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimeseriesStackMode) IsValid() bool {
	for _, existing := range allowedTimeseriesStackModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesStackMode value.
func (v TimeseriesStackMode) Ptr() *TimeseriesStackMode {
	return &v
}
