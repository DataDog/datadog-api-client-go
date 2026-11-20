// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventStreamDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type EventStreamDefinitionType string

// List of EventStreamDefinitionType.
const (
	EVENTSTREAMDEFINITIONTYPE_EVENT_STREAM EventStreamDefinitionType = "event_stream"
)

var allowedEventStreamDefinitionTypeEnumValues = []EventStreamDefinitionType{
	EVENTSTREAMDEFINITIONTYPE_EVENT_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventStreamDefinitionType) GetAllowedValues() []EventStreamDefinitionType {
	return allowedEventStreamDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventStreamDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventStreamDefinitionType(value)
	return nil
}

// NewEventStreamDefinitionTypeFromValue returns a pointer to a valid EventStreamDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventStreamDefinitionTypeFromValue(v string) (*EventStreamDefinitionType, error) {
	ev := EventStreamDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventStreamDefinitionType: valid values are %v", v, allowedEventStreamDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventStreamDefinitionType) IsValid() bool {
	for _, existing := range allowedEventStreamDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventStreamDefinitionType value.
func (v EventStreamDefinitionType) Ptr() *EventStreamDefinitionType {
	return &v
}
