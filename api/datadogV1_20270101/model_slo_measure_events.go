// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SloMeasureEvents Test-only generated schema for the dashboard versioning proof of concept.
type SloMeasureEvents string

// List of SloMeasureEvents.
const (
	SLOMEASUREEVENTS_GOOD_EVENTS SloMeasureEvents = "good_events"
	SLOMEASUREEVENTS_BAD_EVENTS  SloMeasureEvents = "bad_events"
)

var allowedSloMeasureEventsEnumValues = []SloMeasureEvents{
	SLOMEASUREEVENTS_GOOD_EVENTS,
	SLOMEASUREEVENTS_BAD_EVENTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SloMeasureEvents) GetAllowedValues() []SloMeasureEvents {
	return allowedSloMeasureEventsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SloMeasureEvents) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SloMeasureEvents(value)
	return nil
}

// NewSloMeasureEventsFromValue returns a pointer to a valid SloMeasureEvents
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSloMeasureEventsFromValue(v string) (*SloMeasureEvents, error) {
	ev := SloMeasureEvents(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SloMeasureEvents: valid values are %v", v, allowedSloMeasureEventsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SloMeasureEvents) IsValid() bool {
	for _, existing := range allowedSloMeasureEventsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SloMeasureEvents value.
func (v SloMeasureEvents) Ptr() *SloMeasureEvents {
	return &v
}
