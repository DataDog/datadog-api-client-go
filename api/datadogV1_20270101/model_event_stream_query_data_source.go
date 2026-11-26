// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type EventStreamQueryDataSource string

// List of EventStreamQueryDataSource.
const (
	EVENTSTREAMQUERYDATASOURCE_EVENT_STREAM EventStreamQueryDataSource = "event_stream"
)

var allowedEventStreamQueryDataSourceEnumValues = []EventStreamQueryDataSource{
	EVENTSTREAMQUERYDATASOURCE_EVENT_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventStreamQueryDataSource) GetAllowedValues() []EventStreamQueryDataSource {
	return allowedEventStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventStreamQueryDataSource(value)
	return nil
}

// NewEventStreamQueryDataSourceFromValue returns a pointer to a valid EventStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventStreamQueryDataSourceFromValue(v string) (*EventStreamQueryDataSource, error) {
	ev := EventStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventStreamQueryDataSource: valid values are %v", v, allowedEventStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedEventStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventStreamQueryDataSource value.
func (v EventStreamQueryDataSource) Ptr() *EventStreamQueryDataSource {
	return &v
}
