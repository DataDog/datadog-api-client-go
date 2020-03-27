/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// MonitorDeviceID ID of the device the Synthetics monitor is running on. Same as SyntheticsDeviceID.
type MonitorDeviceID string

// List of MonitorDeviceID
const (
	MONITORDEVICEID_LAPTOP_LARGE MonitorDeviceID = "laptop_large"
	MONITORDEVICEID_TABLET       MonitorDeviceID = "tablet"
	MONITORDEVICEID_MOBILE_SMALL MonitorDeviceID = "mobile_small"
)

// Ptr returns reference to MonitorDeviceID value
func (v MonitorDeviceID) Ptr() *MonitorDeviceID {
	return &v
}

type NullableMonitorDeviceID struct {
	value *MonitorDeviceID
	isSet bool
}

func (v NullableMonitorDeviceID) Get() *MonitorDeviceID {
	return v.value
}

func (v NullableMonitorDeviceID) Set(val *MonitorDeviceID) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorDeviceID) IsSet() bool {
	return v.isSet
}

func (v NullableMonitorDeviceID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorDeviceID(val *MonitorDeviceID) *NullableMonitorDeviceID {
	return &NullableMonitorDeviceID{value: val, isSet: true}
}

func (v NullableMonitorDeviceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorDeviceID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
