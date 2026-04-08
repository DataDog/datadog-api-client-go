// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressFormat The format of events ingested through the email address.
type EventEmailAddressFormat string

// List of EventEmailAddressFormat.
const (
	EVENTEMAILADDRESSFORMAT_JSON       EventEmailAddressFormat = "json"
	EVENTEMAILADDRESSFORMAT_PLAIN_TEXT EventEmailAddressFormat = "plain-text"
)

var allowedEventEmailAddressFormatEnumValues = []EventEmailAddressFormat{
	EVENTEMAILADDRESSFORMAT_JSON,
	EVENTEMAILADDRESSFORMAT_PLAIN_TEXT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventEmailAddressFormat) GetAllowedValues() []EventEmailAddressFormat {
	return allowedEventEmailAddressFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventEmailAddressFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventEmailAddressFormat(value)
	return nil
}

// NewEventEmailAddressFormatFromValue returns a pointer to a valid EventEmailAddressFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventEmailAddressFormatFromValue(v string) (*EventEmailAddressFormat, error) {
	ev := EventEmailAddressFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventEmailAddressFormat: valid values are %v", v, allowedEventEmailAddressFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventEmailAddressFormat) IsValid() bool {
	for _, existing := range allowedEventEmailAddressFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventEmailAddressFormat value.
func (v EventEmailAddressFormat) Ptr() *EventEmailAddressFormat {
	return &v
}
