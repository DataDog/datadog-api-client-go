// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloMeasureMinutes Test-only generated schema for the dashboard versioning proof of concept.
type SloMeasureMinutes string

// List of SloMeasureMinutes.
const (
	SLOMEASUREMINUTES_GOOD_MINUTES SloMeasureMinutes = "good_minutes"
	SLOMEASUREMINUTES_BAD_MINUTES  SloMeasureMinutes = "bad_minutes"
)

var allowedSloMeasureMinutesEnumValues = []SloMeasureMinutes{
	SLOMEASUREMINUTES_GOOD_MINUTES,
	SLOMEASUREMINUTES_BAD_MINUTES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloMeasureMinutes) GetAllowedValues() []SloMeasureMinutes {
	return allowedSloMeasureMinutesEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloMeasureMinutes) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloMeasureMinutes(value)
	return nil
}

// NewSloMeasureMinutesFromValue returns a pointer to a valid SloMeasureMinutes
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloMeasureMinutesFromValue(v string) (*SloMeasureMinutes, error) {
	ev := SloMeasureMinutes(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloMeasureMinutes: valid values are %v", v, allowedSloMeasureMinutesEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloMeasureMinutes) IsValid() bool {
	for _, existing := range allowedSloMeasureMinutesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloMeasureMinutes value.
func (v SloMeasureMinutes) Ptr() *SloMeasureMinutes {
	return &v
}
