// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPatternStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type EventPatternStreamRequestResponseFormat string

// List of EventPatternStreamRequestResponseFormat.
const (
	EVENTPATTERNSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST EventPatternStreamRequestResponseFormat = "event_list"
)

var allowedEventPatternStreamRequestResponseFormatEnumValues = []EventPatternStreamRequestResponseFormat{
	EVENTPATTERNSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventPatternStreamRequestResponseFormat) GetAllowedValues() []EventPatternStreamRequestResponseFormat {
	return allowedEventPatternStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventPatternStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventPatternStreamRequestResponseFormat(value)
	return nil
}

// NewEventPatternStreamRequestResponseFormatFromValue returns a pointer to a valid EventPatternStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventPatternStreamRequestResponseFormatFromValue(v string) (*EventPatternStreamRequestResponseFormat, error) {
	ev := EventPatternStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventPatternStreamRequestResponseFormat: valid values are %v", v, allowedEventPatternStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventPatternStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedEventPatternStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventPatternStreamRequestResponseFormat value.
func (v EventPatternStreamRequestResponseFormat) Ptr() *EventPatternStreamRequestResponseFormat {
	return &v
}
