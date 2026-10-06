// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorAlertTriggerType The type of monitor alert trigger.
type MonitorAlertTriggerType string

// List of MonitorAlertTriggerType.
const (
	MONITORALERTTRIGGERTYPE_MONITOR_ALERT_TRIGGER MonitorAlertTriggerType = "monitor_alert_trigger"
)

var allowedMonitorAlertTriggerTypeEnumValues = []MonitorAlertTriggerType{
	MONITORALERTTRIGGERTYPE_MONITOR_ALERT_TRIGGER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorAlertTriggerType) GetAllowedValues() []MonitorAlertTriggerType {
	return allowedMonitorAlertTriggerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorAlertTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorAlertTriggerType(value)
	return nil
}

// NewMonitorAlertTriggerTypeFromValue returns a pointer to a valid MonitorAlertTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorAlertTriggerTypeFromValue(v string) (*MonitorAlertTriggerType, error) {
	ev := MonitorAlertTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorAlertTriggerType: valid values are %v", v, allowedMonitorAlertTriggerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorAlertTriggerType) IsValid() bool {
	for _, existing := range allowedMonitorAlertTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorAlertTriggerType value.
func (v MonitorAlertTriggerType) Ptr() *MonitorAlertTriggerType {
	return &v
}
