// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressAlertType The alert type of events generated from the email address.
type EventEmailAddressAlertType string

// List of EventEmailAddressAlertType.
const (
	EVENTEMAILADDRESSALERTTYPE_INFO    EventEmailAddressAlertType = "info"
	EVENTEMAILADDRESSALERTTYPE_WARN    EventEmailAddressAlertType = "warn"
	EVENTEMAILADDRESSALERTTYPE_ERROR   EventEmailAddressAlertType = "error"
	EVENTEMAILADDRESSALERTTYPE_SUCCESS EventEmailAddressAlertType = "success"
)

var allowedEventEmailAddressAlertTypeEnumValues = []EventEmailAddressAlertType{
	EVENTEMAILADDRESSALERTTYPE_INFO,
	EVENTEMAILADDRESSALERTTYPE_WARN,
	EVENTEMAILADDRESSALERTTYPE_ERROR,
	EVENTEMAILADDRESSALERTTYPE_SUCCESS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *EventEmailAddressAlertType) GetAllowedValues() []EventEmailAddressAlertType {
	return allowedEventEmailAddressAlertTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EventEmailAddressAlertType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EventEmailAddressAlertType(value)
	return nil
}

// NewEventEmailAddressAlertTypeFromValue returns a pointer to a valid EventEmailAddressAlertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEventEmailAddressAlertTypeFromValue(v string) (*EventEmailAddressAlertType, error) {
	ev := EventEmailAddressAlertType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EventEmailAddressAlertType: valid values are %v", v, allowedEventEmailAddressAlertTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EventEmailAddressAlertType) IsValid() bool {
	for _, existing := range allowedEventEmailAddressAlertTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventEmailAddressAlertType value.
func (v EventEmailAddressAlertType) Ptr() *EventEmailAddressAlertType {
	return &v
}
