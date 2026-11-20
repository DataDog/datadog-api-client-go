// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventTimelineBackendDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type EventTimelineBackendDefinitionType string

// List of EventTimelineBackendDefinitionType.
const (
	EVENTTIMELINEBACKENDDEFINITIONTYPE_EVENT_TIMELINE EventTimelineBackendDefinitionType = "event_timeline"
)

var allowedEventTimelineBackendDefinitionTypeEnumValues = []EventTimelineBackendDefinitionType{
	EVENTTIMELINEBACKENDDEFINITIONTYPE_EVENT_TIMELINE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventTimelineBackendDefinitionType) GetAllowedValues() []EventTimelineBackendDefinitionType {
	return allowedEventTimelineBackendDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventTimelineBackendDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventTimelineBackendDefinitionType(value)
	return nil
}

// NewEventTimelineBackendDefinitionTypeFromValue returns a pointer to a valid EventTimelineBackendDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventTimelineBackendDefinitionTypeFromValue(v string) (*EventTimelineBackendDefinitionType, error) {
	ev := EventTimelineBackendDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventTimelineBackendDefinitionType: valid values are %v", v, allowedEventTimelineBackendDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventTimelineBackendDefinitionType) IsValid() bool {
	for _, existing := range allowedEventTimelineBackendDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventTimelineBackendDefinitionType value.
func (v EventTimelineBackendDefinitionType) Ptr() *EventTimelineBackendDefinitionType {
	return &v
}
