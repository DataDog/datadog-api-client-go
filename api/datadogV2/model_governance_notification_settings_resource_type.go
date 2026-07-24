// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceNotificationSettingsResourceType Governance notification settings resource type.
type GovernanceNotificationSettingsResourceType string

// List of GovernanceNotificationSettingsResourceType.
const (
	GOVERNANCENOTIFICATIONSETTINGSRESOURCETYPE_GOVERNANCE_NOTIFICATION_SETTINGS GovernanceNotificationSettingsResourceType = "governance_notification_settings"
)

var allowedGovernanceNotificationSettingsResourceTypeEnumValues = []GovernanceNotificationSettingsResourceType{
	GOVERNANCENOTIFICATIONSETTINGSRESOURCETYPE_GOVERNANCE_NOTIFICATION_SETTINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceNotificationSettingsResourceType) GetAllowedValues() []GovernanceNotificationSettingsResourceType {
	return allowedGovernanceNotificationSettingsResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceNotificationSettingsResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceNotificationSettingsResourceType(value)
	return nil
}

// NewGovernanceNotificationSettingsResourceTypeFromValue returns a pointer to a valid GovernanceNotificationSettingsResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceNotificationSettingsResourceTypeFromValue(v string) (*GovernanceNotificationSettingsResourceType, error) {
	ev := GovernanceNotificationSettingsResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceNotificationSettingsResourceType: valid values are %v", v, allowedGovernanceNotificationSettingsResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceNotificationSettingsResourceType) IsValid() bool {
	for _, existing := range allowedGovernanceNotificationSettingsResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceNotificationSettingsResourceType value.
func (v GovernanceNotificationSettingsResourceType) Ptr() *GovernanceNotificationSettingsResourceType {
	return &v
}
