// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailsType Event emails resource type.
type EventEmailsType string

// List of EventEmailsType.
const (
	EVENTEMAILSTYPE_EVENT_EMAILS EventEmailsType = "event_emails"
)

var allowedEventEmailsTypeEnumValues = []EventEmailsType{
	EVENTEMAILSTYPE_EVENT_EMAILS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventEmailsType) GetAllowedValues() []EventEmailsType {
	return allowedEventEmailsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventEmailsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventEmailsType(value)
	return nil
}

// NewEventEmailsTypeFromValue returns a pointer to a valid EventEmailsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventEmailsTypeFromValue(v string) (*EventEmailsType, error) {
	ev := EventEmailsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventEmailsType: valid values are %v", v, allowedEventEmailsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventEmailsType) IsValid() bool {
	for _, existing := range allowedEventEmailsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventEmailsType value.
func (v EventEmailsType) Ptr() *EventEmailsType {
	return &v
}
