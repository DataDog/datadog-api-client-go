// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceStatus Health status of the device, computed from its issues and recent telemetry.
// `inactive` means the device stopped reporting. `unknown` means no telemetry is available yet.
type DeviceStatus string

// List of DeviceStatus.
const (
	DEVICESTATUS_HEALTHY  DeviceStatus = "healthy"
	DEVICESTATUS_DEGRADED DeviceStatus = "degraded"
	DEVICESTATUS_CRITICAL DeviceStatus = "critical"
	DEVICESTATUS_INACTIVE DeviceStatus = "inactive"
	DEVICESTATUS_UNKNOWN  DeviceStatus = "unknown"
)

var allowedDeviceStatusEnumValues = []DeviceStatus{
	DEVICESTATUS_HEALTHY,
	DEVICESTATUS_DEGRADED,
	DEVICESTATUS_CRITICAL,
	DEVICESTATUS_INACTIVE,
	DEVICESTATUS_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeviceStatus) GetAllowedValues() []DeviceStatus {
	return allowedDeviceStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeviceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeviceStatus(value)
	return nil
}

// NewDeviceStatusFromValue returns a pointer to a valid DeviceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeviceStatusFromValue(v string) (*DeviceStatus, error) {
	ev := DeviceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeviceStatus: valid values are %v", v, allowedDeviceStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeviceStatus) IsValid() bool {
	for _, existing := range allowedDeviceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceStatus value.
func (v DeviceStatus) Ptr() *DeviceStatus {
	return &v
}
