// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceDetailsDataType Devices resource type.
type DeviceDetailsDataType string

// List of DeviceDetailsDataType.
const (
	DEVICEDETAILSDATATYPE_DEVICES DeviceDetailsDataType = "devices"
)

var allowedDeviceDetailsDataTypeEnumValues = []DeviceDetailsDataType{
	DEVICEDETAILSDATATYPE_DEVICES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DeviceDetailsDataType) GetAllowedValues() []DeviceDetailsDataType {
	return allowedDeviceDetailsDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DeviceDetailsDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DeviceDetailsDataType(value)
	return nil
}

// NewDeviceDetailsDataTypeFromValue returns a pointer to a valid DeviceDetailsDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDeviceDetailsDataTypeFromValue(v string) (*DeviceDetailsDataType, error) {
	ev := DeviceDetailsDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DeviceDetailsDataType: valid values are %v", v, allowedDeviceDetailsDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DeviceDetailsDataType) IsValid() bool {
	for _, existing := range allowedDeviceDetailsDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceDetailsDataType value.
func (v DeviceDetailsDataType) Ptr() *DeviceDetailsDataType {
	return &v
}
