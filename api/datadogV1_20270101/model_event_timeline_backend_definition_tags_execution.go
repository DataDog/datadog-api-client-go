// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventTimelineBackendDefinitionTagsExecution Test-only generated schema for the dashboard versioning proof of concept.
type EventTimelineBackendDefinitionTagsExecution string

// List of EventTimelineBackendDefinitionTagsExecution.
const (
	EVENTTIMELINEBACKENDDEFINITIONTAGSEXECUTION_AND EventTimelineBackendDefinitionTagsExecution = "and"
	EVENTTIMELINEBACKENDDEFINITIONTAGSEXECUTION_OR  EventTimelineBackendDefinitionTagsExecution = "or"
)

var allowedEventTimelineBackendDefinitionTagsExecutionEnumValues = []EventTimelineBackendDefinitionTagsExecution{
	EVENTTIMELINEBACKENDDEFINITIONTAGSEXECUTION_AND,
	EVENTTIMELINEBACKENDDEFINITIONTAGSEXECUTION_OR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventTimelineBackendDefinitionTagsExecution) GetAllowedValues() []EventTimelineBackendDefinitionTagsExecution {
	return allowedEventTimelineBackendDefinitionTagsExecutionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventTimelineBackendDefinitionTagsExecution) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventTimelineBackendDefinitionTagsExecution(value)
	return nil
}

// NewEventTimelineBackendDefinitionTagsExecutionFromValue returns a pointer to a valid EventTimelineBackendDefinitionTagsExecution
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventTimelineBackendDefinitionTagsExecutionFromValue(v string) (*EventTimelineBackendDefinitionTagsExecution, error) {
	ev := EventTimelineBackendDefinitionTagsExecution(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventTimelineBackendDefinitionTagsExecution: valid values are %v", v, allowedEventTimelineBackendDefinitionTagsExecutionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventTimelineBackendDefinitionTagsExecution) IsValid() bool {
	for _, existing := range allowedEventTimelineBackendDefinitionTagsExecutionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventTimelineBackendDefinitionTagsExecution value.
func (v EventTimelineBackendDefinitionTagsExecution) Ptr() *EventTimelineBackendDefinitionTagsExecution {
	return &v
}
