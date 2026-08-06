// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRuleTargetType The type of notification target.
type NotificationRuleTargetType string

// List of NotificationRuleTargetType.
const (
	NOTIFICATIONRULETARGETTYPE_EMAIL             NotificationRuleTargetType = "EMAIL"
	NOTIFICATIONRULETARGETTYPE_SLACK_CHANNEL     NotificationRuleTargetType = "SLACK_CHANNEL"
	NOTIFICATIONRULETARGETTYPE_SLACK_USER        NotificationRuleTargetType = "SLACK_USER"
	NOTIFICATIONRULETARGETTYPE_WEBHOOK           NotificationRuleTargetType = "WEBHOOK"
	NOTIFICATIONRULETARGETTYPE_PAGERDUTY_SERVICE NotificationRuleTargetType = "PAGERDUTY_SERVICE"
	NOTIFICATIONRULETARGETTYPE_MS_TEAMS_CHANNEL  NotificationRuleTargetType = "MS_TEAMS_CHANNEL"
)

var allowedNotificationRuleTargetTypeEnumValues = []NotificationRuleTargetType{
	NOTIFICATIONRULETARGETTYPE_EMAIL,
	NOTIFICATIONRULETARGETTYPE_SLACK_CHANNEL,
	NOTIFICATIONRULETARGETTYPE_SLACK_USER,
	NOTIFICATIONRULETARGETTYPE_WEBHOOK,
	NOTIFICATIONRULETARGETTYPE_PAGERDUTY_SERVICE,
	NOTIFICATIONRULETARGETTYPE_MS_TEAMS_CHANNEL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *NotificationRuleTargetType) GetAllowedValues() []NotificationRuleTargetType {
	return allowedNotificationRuleTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *NotificationRuleTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = NotificationRuleTargetType(value)
	return nil
}

// NewNotificationRuleTargetTypeFromValue returns a pointer to a valid NotificationRuleTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewNotificationRuleTargetTypeFromValue(v string) (*NotificationRuleTargetType, error) {
	ev := NotificationRuleTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NotificationRuleTargetType: valid values are %v", v, allowedNotificationRuleTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v NotificationRuleTargetType) IsValid() bool {
	for _, existing := range allowedNotificationRuleTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationRuleTargetType value.
func (v NotificationRuleTargetType) Ptr() *NotificationRuleTargetType {
	return &v
}
