// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type EventStreamRequestResponseFormat string

// List of EventStreamRequestResponseFormat.
const (
	EVENTSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST EventStreamRequestResponseFormat = "event_list"
)

var allowedEventStreamRequestResponseFormatEnumValues = []EventStreamRequestResponseFormat{
	EVENTSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventStreamRequestResponseFormat) GetAllowedValues() []EventStreamRequestResponseFormat {
	return allowedEventStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventStreamRequestResponseFormat(value)
	return nil
}

// NewEventStreamRequestResponseFormatFromValue returns a pointer to a valid EventStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventStreamRequestResponseFormatFromValue(v string) (*EventStreamRequestResponseFormat, error) {
	ev := EventStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventStreamRequestResponseFormat: valid values are %v", v, allowedEventStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedEventStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventStreamRequestResponseFormat value.
func (v EventStreamRequestResponseFormat) Ptr() *EventStreamRequestResponseFormat {
	return &v
}
