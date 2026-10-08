// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorAutomationType The resource type for monitor automation settings.
type MonitorAutomationType string

// List of MonitorAutomationType.
const (
	MONITORAUTOMATIONTYPE_MONITOR_AUTOMATION MonitorAutomationType = "monitor_automation"
)

var allowedMonitorAutomationTypeEnumValues = []MonitorAutomationType{
	MONITORAUTOMATIONTYPE_MONITOR_AUTOMATION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonitorAutomationType) GetAllowedValues() []MonitorAutomationType {
	return allowedMonitorAutomationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonitorAutomationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorAutomationType(value)
	return nil
}

// NewMonitorAutomationTypeFromValue returns a pointer to a valid MonitorAutomationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonitorAutomationTypeFromValue(v string) (*MonitorAutomationType, error) {
	ev := MonitorAutomationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonitorAutomationType: valid values are %v", v, allowedMonitorAutomationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonitorAutomationType) IsValid() bool {
	for _, existing := range allowedMonitorAutomationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorAutomationType value.
func (v MonitorAutomationType) Ptr() *MonitorAutomationType {
	return &v
}
