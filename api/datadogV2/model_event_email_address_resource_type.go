// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressResourceType The type of the resource. Must be `event_emails`.
type EventEmailAddressResourceType string

// List of EventEmailAddressResourceType.
const (
	EVENTEMAILADDRESSRESOURCETYPE_EVENT_EMAILS EventEmailAddressResourceType = "event_emails"
)

var allowedEventEmailAddressResourceTypeEnumValues = []EventEmailAddressResourceType{
	EVENTEMAILADDRESSRESOURCETYPE_EVENT_EMAILS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventEmailAddressResourceType) GetAllowedValues() []EventEmailAddressResourceType {
	return allowedEventEmailAddressResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventEmailAddressResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventEmailAddressResourceType(value)
	return nil
}

// NewEventEmailAddressResourceTypeFromValue returns a pointer to a valid EventEmailAddressResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventEmailAddressResourceTypeFromValue(v string) (*EventEmailAddressResourceType, error) {
	ev := EventEmailAddressResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventEmailAddressResourceType: valid values are %v", v, allowedEventEmailAddressResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventEmailAddressResourceType) IsValid() bool {
	for _, existing := range allowedEventEmailAddressResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventEmailAddressResourceType value.
func (v EventEmailAddressResourceType) Ptr() *EventEmailAddressResourceType {
	return &v
}
