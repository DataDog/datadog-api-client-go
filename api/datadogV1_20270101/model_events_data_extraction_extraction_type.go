// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventsDataExtractionExtractionType Test-only generated schema for the dashboard versioning proof of concept.
type EventsDataExtractionExtractionType string

// List of EventsDataExtractionExtractionType.
const (
	EVENTSDATAEXTRACTIONEXTRACTIONTYPE_GROKEXTRACTION  EventsDataExtractionExtractionType = "grokExtraction"
	EVENTSDATAEXTRACTIONEXTRACTIONTYPE_REGEXEXTRACTION EventsDataExtractionExtractionType = "regexExtraction"
)

var allowedEventsDataExtractionExtractionTypeEnumValues = []EventsDataExtractionExtractionType{
	EVENTSDATAEXTRACTIONEXTRACTIONTYPE_GROKEXTRACTION,
	EVENTSDATAEXTRACTIONEXTRACTIONTYPE_REGEXEXTRACTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventsDataExtractionExtractionType) GetAllowedValues() []EventsDataExtractionExtractionType {
	return allowedEventsDataExtractionExtractionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventsDataExtractionExtractionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventsDataExtractionExtractionType(value)
	return nil
}

// NewEventsDataExtractionExtractionTypeFromValue returns a pointer to a valid EventsDataExtractionExtractionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventsDataExtractionExtractionTypeFromValue(v string) (*EventsDataExtractionExtractionType, error) {
	ev := EventsDataExtractionExtractionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventsDataExtractionExtractionType: valid values are %v", v, allowedEventsDataExtractionExtractionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventsDataExtractionExtractionType) IsValid() bool {
	for _, existing := range allowedEventsDataExtractionExtractionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsDataExtractionExtractionType value.
func (v EventsDataExtractionExtractionType) Ptr() *EventsDataExtractionExtractionType {
	return &v
}
