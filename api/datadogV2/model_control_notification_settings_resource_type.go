// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ControlNotificationSettingsResourceType Control notification settings resource type.
type ControlNotificationSettingsResourceType string

// List of ControlNotificationSettingsResourceType.
const (
	CONTROLNOTIFICATIONSETTINGSRESOURCETYPE_CONTROL_NOTIFICATION_SETTINGS ControlNotificationSettingsResourceType = "control_notification_settings"
)

var allowedControlNotificationSettingsResourceTypeEnumValues = []ControlNotificationSettingsResourceType{
	CONTROLNOTIFICATIONSETTINGSRESOURCETYPE_CONTROL_NOTIFICATION_SETTINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ControlNotificationSettingsResourceType) GetAllowedValues() []ControlNotificationSettingsResourceType {
	return allowedControlNotificationSettingsResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ControlNotificationSettingsResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ControlNotificationSettingsResourceType(value)
	return nil
}

// NewControlNotificationSettingsResourceTypeFromValue returns a pointer to a valid ControlNotificationSettingsResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewControlNotificationSettingsResourceTypeFromValue(v string) (*ControlNotificationSettingsResourceType, error) {
	ev := ControlNotificationSettingsResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ControlNotificationSettingsResourceType: valid values are %v", v, allowedControlNotificationSettingsResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ControlNotificationSettingsResourceType) IsValid() bool {
	for _, existing := range allowedControlNotificationSettingsResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ControlNotificationSettingsResourceType value.
func (v ControlNotificationSettingsResourceType) Ptr() *ControlNotificationSettingsResourceType {
	return &v
}
