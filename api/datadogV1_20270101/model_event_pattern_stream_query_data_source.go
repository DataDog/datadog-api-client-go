// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPatternStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type EventPatternStreamQueryDataSource string

// List of EventPatternStreamQueryDataSource.
const (
	EVENTPATTERNSTREAMQUERYDATASOURCE_EVENT_PATTERN_STREAM EventPatternStreamQueryDataSource = "event_pattern_stream"
)

var allowedEventPatternStreamQueryDataSourceEnumValues = []EventPatternStreamQueryDataSource{
	EVENTPATTERNSTREAMQUERYDATASOURCE_EVENT_PATTERN_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventPatternStreamQueryDataSource) GetAllowedValues() []EventPatternStreamQueryDataSource {
	return allowedEventPatternStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventPatternStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventPatternStreamQueryDataSource(value)
	return nil
}

// NewEventPatternStreamQueryDataSourceFromValue returns a pointer to a valid EventPatternStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventPatternStreamQueryDataSourceFromValue(v string) (*EventPatternStreamQueryDataSource, error) {
	ev := EventPatternStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventPatternStreamQueryDataSource: valid values are %v", v, allowedEventPatternStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventPatternStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedEventPatternStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventPatternStreamQueryDataSource value.
func (v EventPatternStreamQueryDataSource) Ptr() *EventPatternStreamQueryDataSource {
	return &v
}
