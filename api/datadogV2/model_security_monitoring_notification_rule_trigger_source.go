// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringNotificationRuleTriggerSource Specifies the evaluation result (Signal or Finding).
type SecurityMonitoringNotificationRuleTriggerSource string

// List of SecurityMonitoringNotificationRuleTriggerSource.
const (
	SECURITYMONITORINGNOTIFICATIONRULETRIGGERSOURCE_SECURITY_SIGNALS  SecurityMonitoringNotificationRuleTriggerSource = "security_signals"
	SECURITYMONITORINGNOTIFICATIONRULETRIGGERSOURCE_SECURITY_FINDINGS SecurityMonitoringNotificationRuleTriggerSource = "security_findings"
)

var allowedSecurityMonitoringNotificationRuleTriggerSourceEnumValues = []SecurityMonitoringNotificationRuleTriggerSource{
	SECURITYMONITORINGNOTIFICATIONRULETRIGGERSOURCE_SECURITY_SIGNALS,
	SECURITYMONITORINGNOTIFICATIONRULETRIGGERSOURCE_SECURITY_FINDINGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringNotificationRuleTriggerSource) GetAllowedValues() []SecurityMonitoringNotificationRuleTriggerSource {
	return allowedSecurityMonitoringNotificationRuleTriggerSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringNotificationRuleTriggerSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringNotificationRuleTriggerSource(value)
	return nil
}

// NewSecurityMonitoringNotificationRuleTriggerSourceFromValue returns a pointer to a valid SecurityMonitoringNotificationRuleTriggerSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringNotificationRuleTriggerSourceFromValue(v string) (*SecurityMonitoringNotificationRuleTriggerSource, error) {
	ev := SecurityMonitoringNotificationRuleTriggerSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringNotificationRuleTriggerSource: valid values are %v", v, allowedSecurityMonitoringNotificationRuleTriggerSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringNotificationRuleTriggerSource) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringNotificationRuleTriggerSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringNotificationRuleTriggerSource value.
func (v SecurityMonitoringNotificationRuleTriggerSource) Ptr() *SecurityMonitoringNotificationRuleTriggerSource {
	return &v
}
