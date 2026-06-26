// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ControlNotificationTargetType The type of notification destination.
type ControlNotificationTargetType string

// List of ControlNotificationTargetType.
const (
	CONTROLNOTIFICATIONTARGETTYPE_EMAIL      ControlNotificationTargetType = "email"
	CONTROLNOTIFICATIONTARGETTYPE_SLACK      ControlNotificationTargetType = "slack"
	CONTROLNOTIFICATIONTARGETTYPE_AT_MENTION ControlNotificationTargetType = "at_mention"
	CONTROLNOTIFICATIONTARGETTYPE_CASE       ControlNotificationTargetType = "case"
)

var allowedControlNotificationTargetTypeEnumValues = []ControlNotificationTargetType{
	CONTROLNOTIFICATIONTARGETTYPE_EMAIL,
	CONTROLNOTIFICATIONTARGETTYPE_SLACK,
	CONTROLNOTIFICATIONTARGETTYPE_AT_MENTION,
	CONTROLNOTIFICATIONTARGETTYPE_CASE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ControlNotificationTargetType) GetAllowedValues() []ControlNotificationTargetType {
	return allowedControlNotificationTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ControlNotificationTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ControlNotificationTargetType(value)
	return nil
}

// NewControlNotificationTargetTypeFromValue returns a pointer to a valid ControlNotificationTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewControlNotificationTargetTypeFromValue(v string) (*ControlNotificationTargetType, error) {
	ev := ControlNotificationTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ControlNotificationTargetType: valid values are %v", v, allowedControlNotificationTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ControlNotificationTargetType) IsValid() bool {
	for _, existing := range allowedControlNotificationTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ControlNotificationTargetType value.
func (v ControlNotificationTargetType) Ptr() *ControlNotificationTargetType {
	return &v
}
