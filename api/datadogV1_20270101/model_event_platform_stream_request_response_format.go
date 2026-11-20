// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventPlatformStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type EventPlatformStreamRequestResponseFormat string

// List of EventPlatformStreamRequestResponseFormat.
const (
	EVENTPLATFORMSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST EventPlatformStreamRequestResponseFormat = "event_list"
)

var allowedEventPlatformStreamRequestResponseFormatEnumValues = []EventPlatformStreamRequestResponseFormat{
	EVENTPLATFORMSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventPlatformStreamRequestResponseFormat) GetAllowedValues() []EventPlatformStreamRequestResponseFormat {
	return allowedEventPlatformStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventPlatformStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventPlatformStreamRequestResponseFormat(value)
	return nil
}

// NewEventPlatformStreamRequestResponseFormatFromValue returns a pointer to a valid EventPlatformStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventPlatformStreamRequestResponseFormatFromValue(v string) (*EventPlatformStreamRequestResponseFormat, error) {
	ev := EventPlatformStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventPlatformStreamRequestResponseFormat: valid values are %v", v, allowedEventPlatformStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventPlatformStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedEventPlatformStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventPlatformStreamRequestResponseFormat value.
func (v EventPlatformStreamRequestResponseFormat) Ptr() *EventPlatformStreamRequestResponseFormat {
	return &v
}
