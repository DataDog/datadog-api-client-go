// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringNotificationRuleType The type of the resource. The value should always be `notification_rules`.
type SecurityMonitoringNotificationRuleType string

// List of SecurityMonitoringNotificationRuleType.
const (
	SECURITYMONITORINGNOTIFICATIONRULETYPE_NOTIFICATION_RULES SecurityMonitoringNotificationRuleType = "notification_rules"
)

var allowedSecurityMonitoringNotificationRuleTypeEnumValues = []SecurityMonitoringNotificationRuleType{
	SECURITYMONITORINGNOTIFICATIONRULETYPE_NOTIFICATION_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringNotificationRuleType) GetAllowedValues() []SecurityMonitoringNotificationRuleType {
	return allowedSecurityMonitoringNotificationRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringNotificationRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringNotificationRuleType(value)
	return nil
}

// NewSecurityMonitoringNotificationRuleTypeFromValue returns a pointer to a valid SecurityMonitoringNotificationRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringNotificationRuleTypeFromValue(v string) (*SecurityMonitoringNotificationRuleType, error) {
	ev := SecurityMonitoringNotificationRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringNotificationRuleType: valid values are %v", v, allowedSecurityMonitoringNotificationRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringNotificationRuleType) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringNotificationRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringNotificationRuleType value.
func (v SecurityMonitoringNotificationRuleType) Ptr() *SecurityMonitoringNotificationRuleType {
	return &v
}
