// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsHistogramRequestRequestType Test-only generated schema for the dashboard versioning proof of concept.
type EventsHistogramRequestRequestType string

// List of EventsHistogramRequestRequestType.
const (
	EVENTSHISTOGRAMREQUESTREQUESTTYPE_HISTOGRAM EventsHistogramRequestRequestType = "histogram"
)

var allowedEventsHistogramRequestRequestTypeEnumValues = []EventsHistogramRequestRequestType{
	EVENTSHISTOGRAMREQUESTREQUESTTYPE_HISTOGRAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventsHistogramRequestRequestType) GetAllowedValues() []EventsHistogramRequestRequestType {
	return allowedEventsHistogramRequestRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventsHistogramRequestRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsHistogramRequestRequestType(value)
	return nil
}

// NewEventsHistogramRequestRequestTypeFromValue returns a pointer to a valid EventsHistogramRequestRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsHistogramRequestRequestTypeFromValue(v string) (*EventsHistogramRequestRequestType, error) {
	ev := EventsHistogramRequestRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsHistogramRequestRequestType: valid values are %v", v, allowedEventsHistogramRequestRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsHistogramRequestRequestType) IsValid() bool {
	for _, existing := range allowedEventsHistogramRequestRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsHistogramRequestRequestType value.
func (v EventsHistogramRequestRequestType) Ptr() *EventsHistogramRequestRequestType {
	return &v
}
